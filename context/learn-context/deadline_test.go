package learncontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func CreateDeadline(ctx context.Context) chan int {
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
				time.Sleep(1 * time.Second) //simulasi slow response
			}
		}
	}()
	return destination
}

func TestDeadline(t *testing.T) {
	fmt.Println("Runttime", runtime.NumGoroutine())
	parent := context.Background()
	waktu_deadline := time.Now().Add(7 * time.Second)
	fmt.Println(waktu_deadline)
	ctx, cancle := context.WithDeadline(parent, waktu_deadline)

	/**
	hampir sama dengan timeout
	disini misalnya mau jam 12 siang close
	*/
	defer cancle()
	destination := CreateTimeout(ctx)
	for n := range destination {
		fmt.Println("conter ", n, "Runttime", runtime.NumGoroutine())
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Runttime", runtime.NumGoroutine())
}
