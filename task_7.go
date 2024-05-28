package main

import (
	"fmt"
	"sync"
)

func task_7() {
	var wg sync.WaitGroup
	m := make(map[int]int)
	var mu sync.Mutex
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go writeToMap(i, i*i, &wg, &mu, m)
	}
	wg.Wait()
	for k, v := range m {
		fmt.Printf("key: %d, value: %d\n", k, v)
	}
}

func writeToMap(key, value int, wg *sync.WaitGroup, mu *sync.Mutex, m map[int]int) {
	defer wg.Done()
	mu.Lock()
	defer mu.Unlock()
	m[key] = value
}
