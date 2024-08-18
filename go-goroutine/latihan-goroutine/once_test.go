package latihangoroutine

import (
	"fmt"
	"sync"
	"testing"
)

/*
*
once hanya dijalankan 1 kali di dalam goroutine
*/
var counter int = 0
var looping int = 0

func TestOnce(t *testing.T) {
	group := &sync.WaitGroup{}
	once := sync.Once{}

	for i := 0; i <= 100; i++ {
		go runAsyncronusOnce(group, &once)
	}
	// untuk menunggu semua prosess go routine selesai
	// sehingga tidak perlu mengunakan timer untuk menebak prosess
	group.Wait()
	fmt.Println("Selesai Bro counter = ", counter, " | looping ", looping)
}

func runAsyncronusOnce(group *sync.WaitGroup, once *sync.Once) {
	// untuk antisipasi jika terjadi error
	defer group.Done()
	group.Add(1)
	looping++
	once.Do(counters)
}

func counters() {
	counter++
}
