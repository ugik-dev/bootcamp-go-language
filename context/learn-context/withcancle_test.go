package learncontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

// maksudnya leak=tidak berhenti atau sperti zombie ketika tidak digunakan lagi
func CreateConterLeak() chan int {
	destination := make(chan int)
	go func() {
		defer close(destination)
		counter := 1
		for {
			destination <- counter
			counter++
		}
	}()

	return destination
}

func TestContextCancle(t *testing.T) {
	// background := context.Background()
	fmt.Println("Runttime", runtime.NumGoroutine())
	destination := CreateConterLeak()

	for n := range destination {
		fmt.Println("conter ", n)
		if n == 10 {
			break
		}
	}
	fmt.Println("Runttime", runtime.NumGoroutine())
}

// penyelesaian masalah diatas
func CreateConterLeakSolve(ctx context.Context) chan int {
	destination := make(chan int)
	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
			}

		}
	}()

	return destination
}
func TestContextCancleSolve(t *testing.T) {
	fmt.Println("Runttime", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancle := context.WithCancel(parent)

	destination := CreateConterLeakSolve(ctx)

	for n := range destination {
		fmt.Println("conter ", n)
		if n == 10 {
			cancle()
			break
		}
	}
	time.Sleep(3 * time.Second)
	fmt.Println("Runttime", runtime.NumGoroutine())
}
