package latihangoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())
	time := <-timer.C
	fmt.Println(time)
}

func TestTimerAfter(t *testing.T) {
	channel := time.After(5 * time.Second)
	fmt.Println(time.Now())
	time := <-channel
	fmt.Println(time)
}

func TestTimerAfterFunction(t *testing.T) {
	group := &sync.WaitGroup{}
	group.Add(1)
	time.AfterFunc(5*time.Second, func() {
		fmt.Println("in func", time.Now())
		group.Done()
	})
	fmt.Println("out func", time.Now())
	group.Wait()
}

func TestTimerTicker(t *testing.T) {
	// sama seperti after dikembalikan mengunakan chanel
	// bisa di stop
	group := &sync.WaitGroup{}
	ticker := time.NewTicker(1 * time.Second)
	fmt.Println("start test")
	group.Add(1)
	go func() {
		fmt.Println("start go")
		time.Sleep(5 * time.Second)
		ticker.Stop()
		group.Done()
		fmt.Println("stoped ticker")
	}()

	for data := range ticker.C {
		fmt.Println("cuurent", data)
	}

	group.Wait()
	fmt.Println("end test")

	// fmt.Println(time)
}

func TestTimerTick(t *testing.T) {
	// sama seperti after dikembalikan mengunakan chanel
	// terus menerus diulang
	channel := time.Tick(1 * time.Second)

	go func() {
		time.Sleep(5 * time.Second)
	}()

	for data := range channel {
		fmt.Println(data)
	}
	// fmt.Println(time)
}
