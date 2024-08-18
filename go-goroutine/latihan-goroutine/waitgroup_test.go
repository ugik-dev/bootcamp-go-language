package latihangoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i <= 100; i++ {
		go runAsyncronus(group)
	}
	// untuk menunggu semua prosess go routine selesai
	// sehingga tidak perlu mengunakan timer untuk menebak
	group.Wait()
	fmt.Println("Selesai Bro")
}

func runAsyncronus(group *sync.WaitGroup) {
	// untuk antisipasi jika terjadi error
	defer group.Done()
	group.Add(1)
	fmt.Println("RunAsyncronus")
	time.Sleep(1 * time.Second)
}
