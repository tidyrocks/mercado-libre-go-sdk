package resilience

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"time"
)

type RetryConfig struct {
	MaxAttempts  int           `yaml:"max_attempts" default:"3"`
	InitialDelay time.Duration `yaml:"initial_delay" default:"100ms"`
	MaxDelay     time.Duration `yaml:"max_delay" default:"30s"`
	Multiplier   float64       `yaml:"multiplier" default:"2.0"`
	Jitter       bool          `yaml:"jitter" default:"true"`
	RetryOn      []int         `yaml:"retry_on" default:"[429,500,502,503,504]"`
}

// DefaultRetryConfig retorna configuración por defecto para retry.
func DefaultRetryConfig() *RetryConfig {
	return &RetryConfig{
		MaxAttempts:  3,
		InitialDelay: 100 * time.Millisecond,
		MaxDelay:     30 * time.Second,
		Multiplier:   2.0,
		Jitter:       true,
		RetryOn:      []int{429, 500, 502, 503, 504},
	}
}

type RetryableFunc func(ctx context.Context) (*http.Response, error)

// WithExponentialBackoff aplica backoff exponencial con jitter para evitar thundering herd.
func WithExponentialBackoff(ctx context.Context, config *RetryConfig, fn RetryableFunc) (*http.Response, error) {
	var lastErr error

	for attempt := 1; attempt <= config.MaxAttempts; attempt++ {
		resp, err := fn(ctx)

		// Si no hay error, retornar exitosamente
		if err == nil && resp != nil && !shouldRetry(resp.StatusCode, config.RetryOn) {
			return resp, nil
		}

		lastErr = err

		// Si es el último intento, no esperar
		if attempt == config.MaxAttempts {
			break
		}

		// Calcular delay con exponential backoff
		delay := calculateDelay(attempt, config)

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(delay):
			// Continuar con el siguiente intento
		}
	}

	return nil, fmt.Errorf("max retry attempts (%d) exceeded: %w", config.MaxAttempts, lastErr)
}

func shouldRetry(statusCode int, retryOn []int) bool {
	for _, code := range retryOn {
		if statusCode == code {
			return true
		}
	}
	return false
}

// calculateDelay aplica fórmula exponencial con jitter aleatorio del 10%.
func calculateDelay(attempt int, config *RetryConfig) time.Duration {
	// Exponential backoff: initialDelay * (multiplier ^ (attempt - 1))
	delay := float64(config.InitialDelay) * math.Pow(config.Multiplier, float64(attempt-1))

	// Aplicar max delay
	if delay > float64(config.MaxDelay) {
		delay = float64(config.MaxDelay)
	}

	// Aplicar jitter para evitar thundering herd
	if config.Jitter {
		jitter := rand.Float64() * 0.1 // ±10%
		delay = delay * (1.0 + jitter)
	}

	return time.Duration(delay)
}
