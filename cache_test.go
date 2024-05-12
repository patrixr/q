package q

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCacheGetSet(t *testing.T) {
	var cache Cache[int] = NewInMemoryCache[int](time.Minute)

	cache.Set("a", 1)
	cache.Set("b", 2)
	cache.Set("c", 3)

	a, ok := cache.Get("a")
	assert.True(t, ok)
	assert.Equal(t, 1, a)

	b, ok := cache.Get("b")
	assert.True(t, ok)
	assert.Equal(t, 2, b)

	c, ok := cache.Get("c")
	assert.True(t, ok)
	assert.Equal(t, 3, c)

	_, ok = cache.Get("d")
	assert.False(t, ok)
}

func TestCacheGetSetWithTTL(t *testing.T) {
	var cache Cache[int] = NewInMemoryCache[int](time.Millisecond)

	cache.Set("a", 1)
	cache.Set("b", 2)
	cache.Set("c", 3)

	time.Sleep(time.Millisecond * 2)

	_, ok := cache.Get("a")
	assert.False(t, ok)

	_, ok = cache.Get("b")
	assert.False(t, ok)

	_, ok = cache.Get("c")
	assert.False(t, ok)
}

func TestCleanup(t *testing.T) {
	var cache Cache[int] = NewInMemoryCache[int](time.Millisecond)

	cache.Set("a", 1)
	cache.Set("b", 2)
	cache.Set("c", 3)

	time.Sleep(time.Millisecond * 2)

	assert.Equal(t, 3, cache.Len())

	cache.Cleanup()

	assert.Equal(t, 0, cache.Len())

	_, ok := cache.Get("a")
	assert.False(t, ok)

	_, ok = cache.Get("b")
	assert.False(t, ok)

	_, ok = cache.Get("c")
	assert.False(t, ok)
}

func TestInvalidate(t *testing.T) {
	var cache Cache[int] = NewInMemoryCache[int](time.Minute)

	cache.Set("a", 1)
	cache.Set("b", 2)
	cache.Set("c", 3)

	cache.Invalidate("a")

	_, ok := cache.Get("a")
	assert.False(t, ok)

	b, ok := cache.Get("b")
	assert.True(t, ok)
	assert.Equal(t, 2, b)

	c, ok := cache.Get("c")
	assert.True(t, ok)
	assert.Equal(t, 3, c)
}

func TestClear(t *testing.T) {
	var cache Cache[int] = NewInMemoryCache[int](time.Minute)

	cache.Set("a", 1)
	cache.Set("b", 2)
	cache.Set("c", 3)

	cache.Clear()

	_, ok := cache.Get("a")
	assert.False(t, ok)

	_, ok = cache.Get("b")
	assert.False(t, ok)

	_, ok = cache.Get("c")
	assert.False(t, ok)
}
