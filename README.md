# Tinycache
Tinycache is a minimal, concurrency-safe, expiring key/value store for strings

# Usage
``` go 
import (
	"fmt"
	"time"
	cache "github.com/ezachrisen/tinycache"
)

func ExampleCache() {
	c := cache.NewCache()
	c.Set("donald", "duck", 1000*time.Millisecond)

	if d, ok := c.Get("donald"); ok {
		fmt.Println("donald", d)
	}
	// Output: donald duck
}

```