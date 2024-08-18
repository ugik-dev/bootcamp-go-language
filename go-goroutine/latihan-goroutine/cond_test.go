package latihangoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var group = &sync.WaitGroup{}
var cond = sync.NewCond(&sync.Mutex{})

func waitCondition(value int) {
	defer group.Done()

	cond.L.Lock()
	cond.Wait()
	fmt.Println("Done ", value)
	cond.L.Unlock()
}
func TestCond(t *testing.T) {
	for i := 1; i < 10; i++ {
		group.Add(1)
		go waitCondition(i)
	}

	// cond.Signal mengirimkan signal 1 per 1 bahwa go cond.wait selanjutnya sudah bisa jalan
	go func() {
		for i := 1; i < 100; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal()
		}
	}()

	// memberitahu bahwa semua cond sudah bisa jalan yang sudah di lock, semua
	// go func() {
	// 	for i := 1; i < 100; i++ {
	// 		time.Sleep(1 * time.Second)
	// 		cond.Broadcast()
	// 	}
	// }()
	group.Wait()
}
