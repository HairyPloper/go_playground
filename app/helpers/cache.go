// Package helpers provides utility functions and structures to assist in various tasks.

package helpers

import (
	"sync"
)

// Cache represents a simple in-memory cache for storing serialized JSON objects.
type Cache struct {
	data map[string]string // Internal data store for the cache
	mu   sync.RWMutex      // Mutex for concurrent access
}

// NewCache creates a new instance of the Cache.
// It initializes the internal data store.
//
// Returns:
//   - *Cache: A pointer to the newly created Cache instance.
func NewCache() *Cache {
	return &Cache{
		data: make(map[string]string),
	}
}

// Set adds or updates a serialized JSON object in the cache based on the provided key.
// It uses a mutex to handle concurrent access safely.
//
// Parameters:
//   - key: The key associated with the serialized JSON object.
//   - serializedValue: The serialized JSON object to be stored in the cache.
func (c *Cache) Set(key string, serializedValue string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = serializedValue
}

// Get retrieves a serialized JSON object from the cache based on the provided key.
// It returns the serialized JSON object and a boolean indicating whether the key exists in the cache.
// It uses a read-lock to allow multiple concurrent reads.
//
// Parameters:
//   - key: The key associated with the desired serialized JSON object.
//
// Returns:
//   - string: The serialized JSON object.
//   - bool: A boolean indicating whether the key exists in the cache.
func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	data, exists := c.data[key]
	return data, exists
}
