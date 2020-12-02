package tinycache_test

import (
	"math/rand"
	"sync"
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

func TestConcurrency(t *testing.T) {

	var wg sync.WaitGroup

	c := cache.NewCache()

	for i := 1; i <= 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			rand.Seed(time.Now().UnixNano())
			min := 1
			max := 30
			time.Sleep(time.Duration(time.Millisecond * time.Duration((rand.Intn(max-min+1) + min))))
			c.Set("donald", "duck", 50000*time.Millisecond)
		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			rand.Seed(time.Now().UnixNano())
			min := 1
			max := 30
			time.Sleep(time.Duration(time.Millisecond * time.Duration((rand.Intn(max-min+1) + min))))
			if _, ok := c.Get("donald"); !ok {
				t.Fatal("Couldn't get donald")
			}
		}()
	}

	wg.Wait()
}
