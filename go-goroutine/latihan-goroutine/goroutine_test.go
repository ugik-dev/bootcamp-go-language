package latihangoroutine

/**
Go routine seperti asyncronus
tapi tidak bisa menangkap return atau value kembalian

chanel bisa menghandle tangkapan data dari goroutine seperti async await
konsep menunggu parameter di ambil, bukan menunggu tangkapan data


*/
import (
	"fmt"
	"testing"
	"time"
)

func HelloWord() {
	fmt.Println("it is function HelloWord")
}

func TestCreateGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go HelloWord()
	}
	fmt.Println("Ups")

	// agar go HelloWord tetap berjalan atau menunggu sedkit agar goroutine selesai di jalankn
	time.Sleep(1 * time.Second)
}

func displayNumber(i int) {
	fmt.Printf("Looping %d\n", i)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100; i++ {
		go displayNumber(i)
	}

	// tanpa go 7.06s
	// dengan go routine
	// time.Sleep(5 * time.Second)
}
