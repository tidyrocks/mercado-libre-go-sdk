package httpx

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/tidyrocks/mercado-libre-go-sdk/internal/cache"
	"github.com/tidyrocks/mercado-libre-go-sdk/internal/logger"
	"github.com/tidyrocks/mercado-libre-go-sdk/internal/metrics"
	"github.com/tidyrocks/mercado-libre-go-sdk/internal/resilience"
	"github.com/tidyrocks/mercado-libre-go-sdk/internal/validation"
)

// ProfessionalClient cliente HTTP con todas las funcionalidades empresariales
type ProfessionalClient struct {
	httpClient     *http.Client
	retryConfig    *resilience.RetryConfig
	circuitBreaker *resilience.CircuitBreaker
	cache          cache.Cache
	metricsEnabled bool

	// Rate limiting
	rateLimiter chan struct{}
}

type ClientConfig struct {
	Timeout        time.Duration                   `yaml:"timeout" default:"30s"`
	RetryConfig    *resilience.RetryConfig         `yaml:"retry"`
	CircuitBreaker resilience.CircuitBreakerConfig `yaml:"circuit_breaker"`
	CacheEnabled   bool                            `yaml:"cache_enabled" default:"true"`
	CacheTTL       time.Duration                   `yaml:"cache_ttl" default:"5m"`
	MetricsEnabled bool                            `yaml:"metrics_enabled" default:"true"`
	RateLimit      int                             `yaml:"rate_limit" default:"100"` // requests per second
}

// NewProfessionalClient inicia goroutine de rate limiting y configura circuit breaker.
func NewProfessionalClient(config ClientConfig) *ProfessionalClient {
	client := &ProfessionalClient{
		httpClient: &http.Client{
			Timeout: config.Timeout,
		},
		retryConfig:    config.RetryConfig,
		circuitBreaker: resilience.NewCircuitBreaker(config.CircuitBreaker),
		metricsEnabled: config.MetricsEnabled,
		rateLimiter:    make(chan struct{}, config.RateLimit),
	}

	if config.CacheEnabled {
		client.cache = cache.GetCache()
	}

	if config.RetryConfig == nil {
		client.retryConfig = resilience.DefaultRetryConfig()
	}

	// Fill rate limiter
	for i := 0; i < config.RateLimit; i++ {
		client.rateLimiter <- struct{}{}
	}

	// Refill rate limiter every second
	go client.refillRateLimiter(config.RateLimit)

	return client
}

// refillRateLimiter corre en goroutine separada rellenando tokens cada segundo.
func (c *ProfessionalClient) refillRateLimiter(limit int) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for range ticker.C {
		// Refill up to limit
		for len(c.rateLimiter) < limit {
			select {
			case c.rateLimiter <- struct{}{}:
			default:
				break
			}
		}
	}
}

// DoGetJSONWithCache aplica todas las funcionalidades: cache, rate limit, retry, circuit breaker y métricas.
func (c *ProfessionalClient) DoGetJSONWithCache(ctx context.Context, url, accessToken string, cacheTTL time.Duration, dest interface{}) error {
	// Validaciones
	if err := validation.ValidateAccessToken(accessToken); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	// Cache key
	cacheKey := fmt.Sprintf("GET:%s", url)

	// Try cache first
	if c.cache != nil {
		if hit, err := c.cache.Get(ctx, cacheKey, dest); err == nil && hit {
			if c.metricsEnabled {
				metrics.GetCollector().RecordCacheHit("http_get")
			}
			return nil
		}
		if c.metricsEnabled {
			metrics.GetCollector().RecordCacheMiss("http_get")
		}
	}

	// Rate limiting
	select {
	case <-c.rateLimiter:
		defer func() {
			// Return token after 1 second
			time.AfterFunc(time.Second, func() {
				select {
				case c.rateLimiter <- struct{}{}:
				default:
				}
			})
		}()
	case <-ctx.Done():
		return ctx.Err()
	}

	start := time.Now()
	var response *http.Response
	var requestErr error

	// Execute with circuit breaker and retry
	err := c.circuitBreaker.Execute(ctx, func(ctx context.Context) error {
		response, requestErr = resilience.WithExponentialBackoff(ctx, c.retryConfig, func(ctx context.Context) (*http.Response, error) {
			req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
			if err != nil {
				return nil, err
			}

			req.Header.Set("Authorization", "Bearer "+accessToken)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("User-Agent", "mercado-libre-go-sdk/1.0")

			return c.httpClient.Do(req)
		})

		return requestErr
	})

	duration := time.Since(start)

	// Metrics and logging
	statusCode := 0
	if response != nil {
		statusCode = response.StatusCode
		defer response.Body.Close()
	}

	if c.metricsEnabled {
		metrics.RecordHTTPCall(ctx, "GET", url, statusCode, duration)
	}

	logger.LogHTTPRequest(ctx, "GET", url, duration, statusCode, err)

	if err != nil {
		return fmt.Errorf("HTTP request failed after retries and circuit breaker: %w", err)
	}

	if statusCode < 200 || statusCode >= 300 {
		return fmt.Errorf("HTTP %d: request failed", statusCode)
	}

	// Parse response
	if err := json.NewDecoder(response.Body).Decode(dest); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	// Store in cache
	if c.cache != nil && cacheTTL > 0 {
		_ = c.cache.Set(ctx, cacheKey, dest, cacheTTL)
	}

	return nil
}

// DoPostJSONWithResilience aplica rate limit, retry y circuit breaker pero sin cache por ser POST.
func (c *ProfessionalClient) DoPostJSONWithResilience(ctx context.Context, url, accessToken string, payload, dest interface{}) error {
	// Validations
	if err := validation.ValidateAccessToken(accessToken); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	// Rate limiting
	select {
	case <-c.rateLimiter:
	case <-ctx.Done():
		return ctx.Err()
	}

	start := time.Now()
	var response *http.Response
	var requestErr error

	// Execute with circuit breaker and retry
	err := c.circuitBreaker.Execute(ctx, func(ctx context.Context) error {
		response, requestErr = resilience.WithExponentialBackoff(ctx, c.retryConfig, func(ctx context.Context) (*http.Response, error) {
			jsonData, err := json.Marshal(payload)
			if err != nil {
				return nil, err
			}

			req, err := http.NewRequestWithContext(ctx, "POST", url, strings.NewReader(string(jsonData)))
			if err != nil {
				return nil, err
			}

			req.Header.Set("Authorization", "Bearer "+accessToken)
			req.Header.Set("Content-Type", "application/json")

			return c.httpClient.Do(req)
		})

		return requestErr
	})

	duration := time.Since(start)
	statusCode := 0
	if response != nil {
		statusCode = response.StatusCode
		defer response.Body.Close()
	}

	// Metrics and logging
	if c.metricsEnabled {
		metrics.RecordHTTPCall(ctx, "POST", url, statusCode, duration)
	}

	logger.LogHTTPRequest(ctx, "POST", url, duration, statusCode, err)

	if err != nil {
		return err
	}

	if statusCode < 200 || statusCode >= 300 {
		return fmt.Errorf("HTTP %d: request failed", statusCode)
	}

	if dest != nil {
		return json.NewDecoder(response.Body).Decode(dest)
	}

	return nil
}

// GetCircuitBreakerState retorna el estado del circuit breaker para monitoring.
func (c *ProfessionalClient) GetCircuitBreakerState() resilience.State {
	return c.circuitBreaker.GetState()
}

// GetCacheStats retorna estadísticas del cache.
func (c *ProfessionalClient) GetCacheStats() cache.CacheStats {
	if c.cache != nil {
		return c.cache.Stats()
	}
	return cache.CacheStats{}
}
