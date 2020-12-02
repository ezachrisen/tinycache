package tinycache_test

import (
	"testing"
	"time"

	cache "github.com/ezachrisen/tinycache"
	"github.com/matryer/is"
)

func TestCache(t *testing.T) {

	is := is.New(t)

	c := cache.NewCache()

	c.Set("donald", "duck", 1000*time.Millisecond)

	d, ok := c.Get("donald")
	is.True(ok)
	is.Equal(d, "duck")

	time.Sleep(time.Duration(time.Millisecond) * 1000)

	_, ok = c.Get("donald")
	is.True(!ok)

	c.Set("mickey", "mouse", cache.NoExpiration)
	time.Sleep(time.Duration(time.Millisecond) * 1000)
	m, ok := c.Get("mickey")
	is.True(ok)
	is.Equal(m, "mouse")

}
