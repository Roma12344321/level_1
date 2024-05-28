package main

import (
	"context"
	"fmt"
	"sync"
)

func task_9() {
	var wg sync.WaitGroup
	valChan := make(chan int)
	resChan := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go write(ctx, valChan, resChan)
	go printArr(ctx, &wg, resChan)
	arr := []int{2, 3, 3, 5, 2, 3, 5, 7, 9}
	for i := 0; i < len(arr); i++ {
		wg.Add(1)
		valChan <- arr[i]
	}
	close(valChan)
	wg.Wait()
}

func write(ctx context.Context, val chan int, res chan int) {
	for {
		select {
		case <-ctx.Done():
			close(res)
			return
		case value, ok := <-val:
			if !ok {
				close(res)
				return
			}
			res <- value * 2
		}
	}
}

func printArr(ctx context.Context, wg *sync.WaitGroup, res chan int) {
	for {
		select {
		case <-ctx.Done():
			return
		case value, ok := <-res:
			if !ok {
				return
			}
			fmt.Println(value)
			wg.Done()
		}
	}
}
