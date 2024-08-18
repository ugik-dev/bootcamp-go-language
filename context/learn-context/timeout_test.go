package learncontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func CreateTimeout(ctx context.Context) chan int {
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

func TestTimeout(t *testing.T) {
	fmt.Println("Runttime", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancle := context.WithTimeout(parent, 5*time.Second)

	/**
	saat timeout tetap menjalankan cancle karna jika sukses pun tetap harus ditutup
	timeout hanya menghandle saat waktu melebihi parameter yg ditentukan
	perhatikan disini for di crete timeout harusnya 10 tapi di time 5
	dan tidak perlu lagi break seperti di withcancle
	*/
	defer cancle()
	destination := CreateTimeout(ctx)
	for n := range destination {
		fmt.Println("conter ", n, "Runttime", runtime.NumGoroutine())
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Runttime", runtime.NumGoroutine())
}
