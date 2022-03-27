package main

import (
	"fmt"
	"sync"
)

type CustomCounter struct {
	sync.RWMutex
	value int
}

func (c *CustomCounter) Increment() {
	c.Lock()
	defer c.Unlock()
	c.value++
}

func (c *CustomCounter) Get() int {
	c.RLock()
	defer c.RUnlock()
	return c.value
}

func main() {
	counter := &CustomCounter{}

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(c *CustomCounter) {
			for i := 0; i < 100; i++ {
				c.Increment()
			}
			wg.Done()
		}(counter)
	}

	wg.Wait()

	fmt.Println(counter.Get())
}
