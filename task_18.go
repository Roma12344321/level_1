package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.RWMutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.value
}

func task_18() {
	var wg sync.WaitGroup
	counter := new(Counter)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}
	wg.Wait()
	fmt.Println(counter.Value())
}
