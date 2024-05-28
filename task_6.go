package main

import (
	"context"
	"time"
)

func withChannel() {
	stopChan := make(chan bool)
	go func() {
		for {
			select {
			case <-stopChan:
				return
			default:
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
	time.Sleep(2 * time.Second)
	stopChan <- true
	time.Sleep(1 * time.Second)
}

func withContextCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctx)
	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

func withTimeOut() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctx)
	time.Sleep(3 * time.Second)
}

func withTimeAfter() {
	stopChan := time.After(2 * time.Second)
	go func() {
		for {
			select {
			case <-stopChan:
				return
			default:
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
	time.Sleep(3 * time.Second)
}
