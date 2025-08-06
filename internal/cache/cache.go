package cache

import (
	"context"
	"encoding/json"
	"sync"
	"time"
)

type Cache interface {
	Get(ctx context.Context, key string, dest interface{}) (bool, error)
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	Delete(ctx context.Context, key string) error
	Clear(ctx context.Context) error
	Stats() CacheStats
}

type CacheStats struct {
	Hits      int64   `json:"hits"`
	Misses    int64   `json:"misses"`
	Sets      int64   `json:"sets"`
	Deletes   int64   `json:"deletes"`
	HitRatio  float64 `json:"hit_ratio"`
	ItemCount int     `json:"item_count"`
}

type cacheItem struct {
	data      []byte
	expiresAt time.Time
}

type InMemoryCache struct {
	mu    sync.RWMutex
	items map[string]*cacheItem
	stats CacheStats

	// Cleanup goroutine
	stopCh   chan struct{}
	stopOnce sync.Once
}

// NewInMemoryCache crea una nueva instancia de cache en memoria e inicia el cleanup.
func NewInMemoryCache() *InMemoryCache {
	c := &InMemoryCache{
		items:  make(map[string]*cacheItem),
		stopCh: make(chan struct{}),
	}

	// Start background cleanup goroutine
	go c.cleanupExpired()

	return c
}

// Get verifica expiración antes de retornar y actualiza hit ratio.
func (c *InMemoryCache) Get(ctx context.Context, key string, dest interface{}) (bool, error) {
	c.mu.RLock()
	item, exists := c.items[key]
	c.mu.RUnlock()

	if !exists {
		c.mu.Lock()
		c.stats.Misses++
		c.mu.Unlock()
		return false, nil
	}

	// Check if expired
	if time.Now().After(item.expiresAt) {
		c.Delete(ctx, key)
		c.mu.Lock()
		c.stats.Misses++
		c.mu.Unlock()
		return false, nil
	}

	// Deserialize
	if err := json.Unmarshal(item.data, dest); err != nil {
		return false, err
	}

	c.mu.Lock()
	c.stats.Hits++
	c.stats.HitRatio = float64(c.stats.Hits) / float64(c.stats.Hits+c.stats.Misses)
	c.mu.Unlock()

	return true, nil
}

// Set almacena un valor en cache con tiempo de expiración.
func (c *InMemoryCache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	item := &cacheItem{
		data:      data,
		expiresAt: time.Now().Add(ttl),
	}

	c.mu.Lock()
	c.items[key] = item
	c.stats.Sets++
	c.mu.Unlock()

	return nil
}

// Delete elimina una clave del cache.
func (c *InMemoryCache) Delete(ctx context.Context, key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, exists := c.items[key]; exists {
		delete(c.items, key)
		c.stats.Deletes++
	}

	return nil
}

// Clear elimina todos los elementos del cache.
func (c *InMemoryCache) Clear(ctx context.Context) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items = make(map[string]*cacheItem)
	c.stats = CacheStats{} // Reset stats

	return nil
}

// Stats retorna estadísticas actuales del cache.
func (c *InMemoryCache) Stats() CacheStats {
	c.mu.RLock()
	defer c.mu.RUnlock()

	stats := c.stats
	stats.ItemCount = len(c.items)
	return stats
}

func (c *InMemoryCache) Close() {
	c.stopOnce.Do(func() {
		close(c.stopCh)
	})
}

// cleanupExpired corre en goroutine cada 5 minutos para limpiar elementos vencidos.
func (c *InMemoryCache) cleanupExpired() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.performCleanup()
		case <-c.stopCh:
			return
		}
	}
}

func (c *InMemoryCache) performCleanup() {
	now := time.Now()

	c.mu.Lock()
	defer c.mu.Unlock()

	for key, item := range c.items {
		if now.After(item.expiresAt) {
			delete(c.items, key)
		}
	}
}

// CacheWrapper wraps any function with caching
type CacheWrapper struct {
	cache Cache
}

func NewCacheWrapper(cache Cache) *CacheWrapper {
	return &CacheWrapper{cache: cache}
}

// WrapFunc envuelve cualquier función con cache transparente.
func (cw *CacheWrapper) WrapFunc(key string, ttl time.Duration, fn func() (interface{}, error)) func() (interface{}, error) {
	return func() (interface{}, error) {
		ctx := context.Background()

		// Try to get from cache first
		var result interface{}
		if hit, err := cw.cache.Get(ctx, key, &result); err == nil && hit {
			return result, nil
		}

		// Cache miss, execute function
		result, err := fn()
		if err != nil {
			return nil, err
		}

		// Store in cache (ignore errors)
		_ = cw.cache.Set(ctx, key, result, ttl)

		return result, nil
	}
}

// Global cache instance
var (
	defaultCache Cache
	cacheOnce    sync.Once
)

// GetCache usa singleton para compartir cache entre todos los módulos.
func GetCache() Cache {
	cacheOnce.Do(func() {
		defaultCache = NewInMemoryCache()
	})
	return defaultCache
}
