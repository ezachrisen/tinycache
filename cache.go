// Package tinycache is a minimal, thread-safe, expiring key/value store for strings.
// An expired item will be removed from the cache if someone attempts to read it, there is no periodic clearing of expired items.
package tinycache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	data map[string]entry
	mu   sync.Mutex
}

const (
	// Shortcut to set the TTL for non-expiring cache items
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
// Overwrites an existing key, without regard to the expiration time of the existing entry.
// In other words, if there's an existing entry that expires in 1 hour, and a new entry is
// set with an expiration in 1 second, the entry will expire in 1 second.
func (c *Cache) Set(key, value string, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = entry{
		payload:    value,
		expiration: time.Now().Add(ttl),
	}
}

// Get retrieves the value for the key. Returns false if either the key does not exist
// or the item has expired. Removes the key from the cache if it has expired.
func (c *Cache) Get(key string) (string, bool) {

	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.data[key]

	if !ok {
		return "", false
	}

	if entry.expiration.Before(time.Now()) {
		delete(c.data, key)
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
