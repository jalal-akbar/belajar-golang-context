package belajar_golang_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func WithDeadline(ctx context.Context) chan int {
	channel := make(chan int)

	go func() {
		defer close(channel)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				channel <- counter
				counter++
				time.Sleep(1 * time.Second)
			}
		}
	}()
	return channel
}

func TestWithDeadline(t *testing.T) {
	fmt.Println("Total Goroutines Before", runtime.NumGoroutine())

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()

	channel := WithDeadline(ctx)
	fmt.Println("Total Goroutines on Process", runtime.NumGoroutine())
	for i := range channel {
		fmt.Println("counter", i)
	}
	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutines After", runtime.NumGoroutine())
}
