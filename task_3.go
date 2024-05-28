package main

import (
	"fmt"
	"sync"
)

func task_3() {
	numbers := []int{2, 4, 6, 8, 10}
	ch := make(chan int, len(numbers))
	var wg sync.WaitGroup
	for _, i := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			ch <- n * n
		}(i)
	}
	wg.Wait()
	close(ch)
	res := 0
	for i := range ch {
		res += i
	}
	fmt.Println(res)
}
