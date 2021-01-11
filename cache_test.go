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

	c.Set("donald", "duck", 900*time.Millisecond)

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

func TestConcurrencyNoExpirations(t *testing.T) {

	var wg sync.WaitGroup

	c := cache.NewCache()

	for i := 1; i <= 100000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			rand.Seed(time.Now().UnixNano())
			min := 1
			max := 30
			time.Sleep(time.Duration(time.Millisecond * time.Duration((rand.Intn(max-min+1) + min))))
			c.Set("mickey", "mouse", 1000*time.Second)
		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			rand.Seed(time.Now().UnixNano())
			min := 1
			max := 30
			time.Sleep(time.Duration(time.Millisecond * time.Duration((rand.Intn(max-min+1) + min))))
			if _, ok := c.Get("mickey"); !ok {
				t.Fatal("Couldn't get mickey")
			}
		}()
	}

	wg.Wait()
}

func TestConcurrencyWithExpirations(t *testing.T) {

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

			ttl := time.Duration(time.Millisecond * time.Duration((rand.Intn(max-min+1) + min)))
			c.Set("minnie", "mouse", ttl)
		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			rand.Seed(time.Now().UnixNano())
			min := 1
			max := 30
			time.Sleep(time.Duration(time.Millisecond * time.Duration((rand.Intn(max-min+1) + min))))
			// We don't check the return value here. It is expected that
			// for some invokations, minnie will not be found.
			c.Get("minnie")
		}()
	}

	wg.Wait()
}
