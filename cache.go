package q

import (
	"sync"
	"time"
)

type Cache[T any] interface {
	Set(key string, value T) error
	Get(key string) (T, bool)
	Invalidate(key string)
	Cleanup()
	Clear()
	Len() int
}

type InMemoryCacheItem[T any] struct {
	Data      T
	Timestamp time.Time
}

type InMemoryCache[T any] struct {
	ttl   time.Duration
	data  map[string]InMemoryCacheItem[T]
	mutex sync.RWMutex
}

// NewInMemoryCache creates a new instance of InMemoryCache with the specified time-to-live (TTL) for cache items.
// It initializes an empty cache with the given TTL value.
//
// Parameters:
// - ttl: A time.Duration value that specifies the time-to-live for each cache item.
//
// Returns:
// - A pointer to the newly created InMemoryCache instance.
func NewInMemoryCache[T any](ttl time.Duration) *InMemoryCache[T] {
	return &InMemoryCache[T]{
		ttl:  ttl,
		data: make(map[string]InMemoryCacheItem[T]),
	}
}

func (c *InMemoryCache[T]) Set(key string, value T) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.data[key] = InMemoryCacheItem[T]{
		Data:      value,
		Timestamp: time.Now(),
	}

	return nil
}

func (c *InMemoryCache[T]) Get(key string) (T, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	item, found := c.data[key]
	if !found {
		var zero T // Create a zero value for the type T
		return zero, false
	}

	// Check if the item has expired
	if time.Since(item.Timestamp) > c.ttl {
		c.mutex.RUnlock()
		c.mutex.Lock()
		delete(c.data, key) // Remove the expired item
		c.mutex.Unlock()
		c.mutex.RLock()
		var zero T // Create a zero value for the type T
		return zero, false
	}

	return item.Data, true
}

func (c *InMemoryCache[T]) Invalidate(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.data, key)
}

func (c *InMemoryCache[T]) Cleanup() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	now := time.Now()
	for key, item := range c.data {
		if now.Sub(item.Timestamp) > c.ttl {
			delete(c.data, key)
		}
	}
}

func (c *InMemoryCache[T]) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for key := range c.data {
		delete(c.data, key)
	}
}

func (c *InMemoryCache[T]) Len() int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return len(c.data)
}
