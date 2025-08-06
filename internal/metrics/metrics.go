package metrics

import (
	"context"
	"sync"
	"time"
)

// MetricsCollector interface para diferentes backends (Prometheus, StatsD, etc.)
type MetricsCollector interface {
	RecordHTTPRequest(method, endpoint string, statusCode int, duration time.Duration)
	RecordCacheHit(cacheName string)
	RecordCacheMiss(cacheName string)
	RecordRetryAttempt(operation string, attempt int)
	RecordCircuitBreakerState(name string, state string)
	IncrementCounter(name string, tags map[string]string)
	RecordHistogram(name string, value float64, tags map[string]string)
	RecordGauge(name string, value float64, tags map[string]string)
}

// InMemoryMetrics implementación en memoria para desarrollo/testing
type InMemoryMetrics struct {
	mu       sync.RWMutex
	counters map[string]int64
	gauges   map[string]float64
	histos   map[string][]float64
}

// NewInMemoryMetrics crea una nueva instancia de métricas en memoria.
func NewInMemoryMetrics() *InMemoryMetrics {
	return &InMemoryMetrics{
		counters: make(map[string]int64),
		gauges:   make(map[string]float64),
		histos:   make(map[string][]float64),
	}
}

func (m *InMemoryMetrics) RecordHTTPRequest(method, endpoint string, statusCode int, duration time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()

	key := "http_requests_total"
	m.counters[key]++

	durKey := "http_request_duration_ms"
	m.histos[durKey] = append(m.histos[durKey], float64(duration.Milliseconds()))
}

func (m *InMemoryMetrics) RecordCacheHit(cacheName string) {
	m.IncrementCounter("cache_hits_total", map[string]string{"cache": cacheName})
}

func (m *InMemoryMetrics) RecordCacheMiss(cacheName string) {
	m.IncrementCounter("cache_misses_total", map[string]string{"cache": cacheName})
}

func (m *InMemoryMetrics) RecordRetryAttempt(operation string, attempt int) {
	m.IncrementCounter("retry_attempts_total", map[string]string{"operation": operation})
}

func (m *InMemoryMetrics) RecordCircuitBreakerState(name string, state string) {
	m.RecordGauge("circuit_breaker_state", 1.0, map[string]string{"name": name, "state": state})
}

func (m *InMemoryMetrics) IncrementCounter(name string, tags map[string]string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	key := name
	if tags != nil {
		for k, v := range tags {
			key += "," + k + "=" + v
		}
	}

	m.counters[key]++
}

func (m *InMemoryMetrics) RecordHistogram(name string, value float64, tags map[string]string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	key := name
	if tags != nil {
		for k, v := range tags {
			key += "," + k + "=" + v
		}
	}

	m.histos[key] = append(m.histos[key], value)
}

func (m *InMemoryMetrics) RecordGauge(name string, value float64, tags map[string]string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	key := name
	if tags != nil {
		for k, v := range tags {
			key += "," + k + "=" + v
		}
	}

	m.gauges[key] = value
}

// GetStats retorna todas las métricas para debugging.
func (m *InMemoryMetrics) GetStats() map[string]interface{} {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return map[string]interface{}{
		"counters":   m.counters,
		"gauges":     m.gauges,
		"histograms": m.histos,
	}
}

// Global metrics collector
var (
	globalCollector MetricsCollector
	once            sync.Once
)

// GetCollector usa singleton para garantizar una sola instancia global.
func GetCollector() MetricsCollector {
	once.Do(func() {
		// En producción esto podría ser Prometheus, StatsD, etc.
		globalCollector = NewInMemoryMetrics()
	})
	return globalCollector
}

// RecordHTTPCall es un helper para registrar llamadas HTTP.
func RecordHTTPCall(ctx context.Context, method, endpoint string, statusCode int, duration time.Duration) {
	GetCollector().RecordHTTPRequest(method, endpoint, statusCode, duration)
}
