package belajar_golang_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

// With Cancel
// Cth Leak
func CreateCounterLeak() chan int {
	dest := make(chan int)
	go func() {
		defer close(dest)
		count := 1
		for {
			dest <- count
			count++
		}
	}()
	return dest
}
func TestCreateCounterLeak(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine()) // Total 2

	leak := CreateCounterLeak()
	go func() {
		for i := range leak {
			fmt.Println("counter", i)
			if i == 10 {
				break
			}
		}
	}()

	fmt.Println("Total Goroutine", runtime.NumGoroutine()) //  Total 4
}

// Cth With Cancel
func CreateCounter(ctx context.Context) chan int {
	channel := make(chan int)

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
			}
		}
	}()

	return channel
}

func TestWithCancel(t *testing.T) {
	fmt.Println("Total Goroutine Before", runtime.NumGoroutine())
	// ctx Cancel
	ctx, cancel := context.WithCancel(context.Background())
	channel := CreateCounter(ctx)

	fmt.Println("Total Goroutine on Process", runtime.NumGoroutine())
	for i := range channel {
		fmt.Println("Counter:", i)
		if i == 10 {
			break
		}
	}
	cancel() //  Mengirim Sinyal Cancel

	time.Sleep(2 * time.Second) // Wait cause async

	fmt.Println("Total Goroutine After", runtime.NumGoroutine())
}
