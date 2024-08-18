package latihangoroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

/**
atomic merupakan librart baru
untuk data primitif seperti number
https://pkg.go.dev/sync/atomic

alternatif untuk mutex dan lock

sampel code di ambil dari rase
*/

func TestAtomic(t *testing.T) {
	var kena_racecondition int64 = 0
	var bebas_racecondition int64 = 0
	group := &sync.WaitGroup{}

	for i := 1; i <= 100; i++ {
		go func() {
			// defer group.Done()
			group.Add(1)
			for j := 1; j <= 100; j++ {
				kena_racecondition++
				atomic.AddInt64(&bebas_racecondition, 1)
			}
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Race Condition x =", kena_racecondition)
	fmt.Println("Bebas Race Condition x =", bebas_racecondition)
}
