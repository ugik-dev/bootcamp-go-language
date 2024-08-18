package latihangoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
*
guna mutex adalah untuk mengatasi race condition
dimana saat goroutine mengeksekusi 1 variable  dlm waktu yg sama
sehingga bentrok dan hanya 1 kali dieksekusi dimana seharusnya ada 2

coba hapus mutex untuk perbandingan

RWMutex adalah process untuk menghandle saat membaca sehingga tidak rebutan antara membaca(Read) dan menulis (Write)
*/
func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x++
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Result x =", x)
}

// ----------------- RWMutex

type BankAccount struct {
	RWMutex sync.RWMutex
	Saldo   int
}

// method
func (account *BankAccount) AddSaldo(amount int) {
	account.RWMutex.Lock()
	account.Saldo = account.Saldo + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetSaldo() int {
	account.RWMutex.RLock()
	saldo := account.Saldo
	account.RWMutex.RUnlock()
	return saldo
}
func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				account.AddSaldo(1)
				fmt.Println("Result x =", account.GetSaldo())
			}
		}()
	}
	time.Sleep(15 * time.Second)
	fmt.Println("Last Result x =", account.GetSaldo())
}
