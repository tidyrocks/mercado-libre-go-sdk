package resilience

import (
	"context"
	"errors"
	"sync"
	"time"
)

type State int

const (
	Closed State = iota
	Open
	HalfOpen
)

type CircuitBreaker struct {
	mu                  sync.RWMutex
	state               State
	failureCount        int
	successCount        int
	lastFailureTime     time.Time
	maxFailures         int
	timeout             time.Duration
	maxRequestsHalfOpen int
}

type CircuitBreakerConfig struct {
	MaxFailures         int           `yaml:"max_failures" default:"5"`
	Timeout             time.Duration `yaml:"timeout" default:"60s"`
	MaxRequestsHalfOpen int           `yaml:"max_requests_half_open" default:"3"`
}

// NewCircuitBreaker crea una nueva instancia de circuit breaker.
func NewCircuitBreaker(config CircuitBreakerConfig) *CircuitBreaker {
	return &CircuitBreaker{
		maxFailures:         config.MaxFailures,
		timeout:             config.Timeout,
		maxRequestsHalfOpen: config.MaxRequestsHalfOpen,
		state:               Closed,
	}
}

var (
	ErrCircuitBreakerOpen = errors.New("circuit breaker is open")
	ErrTooManyRequests    = errors.New("too many requests in half-open state")
)

// Execute evalúa estado antes del request y registra resultado después.
func (cb *CircuitBreaker) Execute(ctx context.Context, fn func(ctx context.Context) error) error {
	if err := cb.beforeRequest(); err != nil {
		return err
	}

	err := fn(ctx)
	cb.afterRequest(err == nil)

	return err
}

// beforeRequest verifica si permite el request según el estado actual.
func (cb *CircuitBreaker) beforeRequest() error {
	cb.mu.RLock()
	state := cb.state
	failureCount := cb.failureCount
	successCount := cb.successCount
	lastFailureTime := cb.lastFailureTime
	cb.mu.RUnlock()

	switch state {
	case Open:
		// Verificar si es tiempo de pasar a half-open
		if time.Since(lastFailureTime) > cb.timeout {
			cb.setState(HalfOpen)
			return nil
		}
		return ErrCircuitBreakerOpen

	case HalfOpen:
		// Limitar requests en half-open state
		if successCount+failureCount >= cb.maxRequestsHalfOpen {
			return ErrTooManyRequests
		}
		return nil

	case Closed:
		return nil

	default:
		return nil
	}
}

// afterRequest actualiza contadores y cambia estado según el resultado.
func (cb *CircuitBreaker) afterRequest(success bool) {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	switch cb.state {
	case Closed:
		if success {
			cb.failureCount = 0
		} else {
			cb.failureCount++
			cb.lastFailureTime = time.Now()

			if cb.failureCount >= cb.maxFailures {
				cb.setState(Open)
			}
		}

	case HalfOpen:
		if success {
			cb.successCount++

			// Si tenemos suficientes éxitos, volver a closed
			if cb.successCount >= cb.maxRequestsHalfOpen/2 {
				cb.setState(Closed)
			}
		} else {
			cb.failureCount++
			cb.lastFailureTime = time.Now()
			cb.setState(Open)
		}

	case Open:
		// No hacer nada, ya estamos open
	}
}

// setState cambia estado y resetea contadores apropiados.
func (cb *CircuitBreaker) setState(state State) {
	cb.state = state

	switch state {
	case Closed:
		cb.failureCount = 0
		cb.successCount = 0
	case Open:
		cb.successCount = 0
	case HalfOpen:
		cb.failureCount = 0
		cb.successCount = 0
	}
}

// GetState retorna el estado actual del circuit breaker.
func (cb *CircuitBreaker) GetState() State {
	cb.mu.RLock()
	defer cb.mu.RUnlock()
	return cb.state
}
