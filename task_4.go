package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func task_4() {
	ctx, cancel := context.WithCancel(context.Background())
	var workNum int
	if _, err := fmt.Scan(&workNum); err != nil {
		log.Fatalf("scan error")
	}
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	ch := make(chan int, workNum)
	var wg sync.WaitGroup
	wg.Add(workNum)
	for i := 0; i < workNum; i++ {
		go worker(ctx, &wg, ch)
	}
	go func() {
		i := 0
		for {
			select {
			case <-ctx.Done():
				close(ch)
				return
			default:
				ch <- i
				i++
			}
		}
	}()
	sig := <-sigChan
	cancel()
	wg.Wait()
	fmt.Println("Stopped by signal:", sig)
}

func worker(ctx context.Context, wg *sync.WaitGroup, ch <-chan int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case val, ok := <-ch:
			if !ok {
				return
			}
			fmt.Println(val)
		}
	}
}
