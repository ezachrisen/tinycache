// Package tinycache provides a minimal, thread-safe, expiring key/value store for strings.

package tinycache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	data map[string]entry
	mu   sync.RWMutex
}

const (
	// Pass to Set for non-expiring cache items
	NoExpiration time.Duration = 1<<63 - 62135596801
)

type entry struct {
	payload    string
	expiration time.Time
}

// Initialize a new cache. Must be called before using.
func NewCache() *Cache {
	return &Cache{
		data: make(map[string]entry),
	}

}

// Set stores a value with a key, expiring the item after the time to live (ttl)
func (c *Cache) Set(key, value string, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = entry{
		payload:    value,
		expiration: time.Now().Add(ttl),
	}
}

// Get retrieves the value for the key. Returns false if either the key does not exist
// or the item has expired.
func (c *Cache) Get(key string) (string, bool) {

	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, ok := c.data[key]

	if !ok {
		return "", false
	}

	if entry.expiration.Before(time.Now()) {
		return "", false
	}

	return entry.payload, true
}

// Function to print the cache contents; useful for debugging
func (c *Cache) Dump() {
	for k, v := range c.data {
		fmt.Printf("%s:'%s' (%s)\n", k, v.payload, v.expiration)
	}
}
