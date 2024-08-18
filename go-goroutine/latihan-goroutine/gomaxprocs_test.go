package latihangoroutine

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestGomaxprocs(t *testing.T) {

	// defaul goroutine yg berjalan 2 bawaan golang

	for i := 0; i <= 100; i++ {
		go func() {
			time.Sleep(3 * time.Second)
		}()
	}
	totalMyPcCPU := runtime.NumCPU()
	fmt.Println("Total CPU Device Saya =", totalMyPcCPU)

	// untuk merubah treads menjadi 20 tread menggunakan runtime.GOMAXPROCS(20)
	// disarankan jangan karna dari golang sudah optimal pengaturannya

	// runtime.GOMAXPROCS(20)
	totalMyPcTread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Treads Device Saya =", totalMyPcTread)

	currentRunningGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine yang sedang berjalan", currentRunningGoroutine)
}
