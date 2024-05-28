package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

func task_5() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		log.Fatalf("fail")
	}
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(2)
	go send(ch, ctx, &wg)
	go receive(ch, ctx, &wg)
	time.Sleep(time.Duration(n) * time.Second)
	cancel()
	wg.Wait()
	close(ch)
}

func send(ch chan int, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	i := 0
	for {
		select {
		case <-ctx.Done():
			return
		case ch <- i:
			i++
		}
	}
}

func receive(ch chan int, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case val := <-ch:
			fmt.Println(val)
		}
	}
}
