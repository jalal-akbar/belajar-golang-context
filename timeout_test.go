package belajar_golang_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func WithTimeOut(ctx context.Context) chan int {
	channel := make(chan int) // Channel
	go func() {
		defer close(channel)
		count := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				channel <- count
				count++
				time.Sleep(1 * time.Second) // simulate slow process
			}
		}
	}()
	return channel
}

func TestContextWithTimeOut(t *testing.T) {
	fmt.Println("Total Goroutine Before", runtime.NumGoroutine())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	channel := WithTimeOut(ctx)

	for i := range channel {
		fmt.Println("counter :", i)
	}

	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine After", runtime.NumGoroutine())
}
