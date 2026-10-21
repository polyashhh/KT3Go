package main
import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	mutex   sync.Mutex
	values  map[int]int
	waiting map[int]chan struct{}
}

func NewCache() *Cache {
	return &Cache{
		values:  make(map[int]int),
		waiting: make(map[int]chan struct{}),
	}
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func (c *Cache) Get(key int) int {
	c.mutex.Lock()
	if value, ok := c.values[key]; ok {
		c.mutex.Unlock()
		return value
	}
	if ch, ok := c.waiting[key]; ok {
		c.mutex.Unlock()
		<-ch
		c.mutex.Lock()
		value := c.values[key]
		c.mutex.Unlock()
		return value
	}
	ch := make(chan struct{})
	c.waiting[key] = ch
	c.mutex.Unlock()
	time.Sleep(time.Second)
	result := fibonacci(key)
	c.mutex.Lock()
	c.values[key] = result
	close(ch)
	delete(c.waiting, key)
	c.mutex.Unlock()
	return result
}

func main() {
	cache := NewCache()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println(cache.Get(10))
	}()
	go func() {
		defer wg.Done()
		fmt.Println(cache.Get(10))
	}()
	wg.Wait()
}