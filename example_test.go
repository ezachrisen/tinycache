package tinycache_test

import (
	"fmt"
	"time"

	cache "github.com/ezachrisen/tinycache"
)

// Example of simple usage
func ExampleSuccess() {
	c := cache.NewCache()
	c.Set("donald", "duck", 1000*time.Millisecond)

	if d, ok := c.Get("donald"); ok {
		fmt.Println("donald", d)
	}
	// Output: donald duck
}

// Example of a cache entry expiring
func ExampleExpired() {
	c := cache.NewCache()
	c.Set("donald", "duck", 1000*time.Millisecond)
	time.Sleep(time.Duration(time.Millisecond) * 1000)

	if d, ok := c.Get("donald"); ok {
		fmt.Println("donald", d)
	} else {
		fmt.Println("donald expired")
	}
	// Output: donald expired
}

// Example of a cache entry that never expires
func ExampleNoExpiration() {
	c := cache.NewCache()
	c.Set("donald", "duck", cache.NoExpiration)
	time.Sleep(time.Duration(time.Millisecond) * 1000)

	if d, ok := c.Get("donald"); ok {
		fmt.Println("donald", d)
	}
	// Output: donald duck
}
