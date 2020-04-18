package main

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestCreateLRU(t *testing.T) {
	cache, _ := NewLru(3)
	assert.Equal(t, cache.Length(), 0, "Cache should be empty while created")
}

func TestShouldOnlyHaveThreeValues(t *testing.T) {
	cache, _ := NewLru(3)
	cache.Add("v1", "somevale")
	cache.Add("v2", "extraval")
	cache.Add("v3", "testval")
	cache.Add("v4", "updateval")

	assert.Equal(t, 3, cache.Length(), "They should be equal")
}

func TestCacheShouldHaveLatestEntry(t *testing.T) {
	cache, _ := NewLru(3)

	cache.Add("v1", "<h1>Hello v1</h1>")
	cache.Add("v2", "extraval")
	cache.Add("v3", "testval")
	cache.Add("v4", "<h1>Hello v4</h1>")

	assert.Equal(t, nil, cache.Get("v1"), "Should not be in cache")
	assert.Equal(t, "<h1>Hello v4</h1>", cache.Get("v4"), "Should not be in cache")
}

func TestCacheShouldNotStoreDuplicates(t *testing.T) {
	cache, _ := NewLru(3)

	cache.Add("v1", "<h1>Hello v1</h1>")
	cache.Add("v1", "<h1>Hello v1</h1>")
	cache.Add("v1", "<h1>Hello v1</h1>")

	assert.Equal(t, 1, cache.Length(), "Should not store duplicates")
}

func BenchmarkCache(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cache, _ := NewLru(3)

		cache.Add("v1", "<h1>Hello v1</h1>")
		cache.Add("v2", "extraval")
	    cache.Add("v3", "testval")
	    cache.Add("v4", "<h1>Hello v4</h1>")
		cache.Add("v1", "<h1>Hello v1</h1>")
		cache.Add("v1", "<h1>Hello v1</h1>")
		cache.Add("v2", "extraval")
		cache.Add("v3", "<h1>Hello v3</h1>")
		cache.Add("v2", "<h1>Hello v2</h1>")
	    cache.Add("v3", "testval")
	}
}
