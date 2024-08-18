package latihangoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type UserBalance struct {
	sync.Mutex
	Name  string
	Saldo int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Change(amount int) {
	user.Saldo = user.Saldo + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock user1 ", user1.Name)
	user1.Change(-amount)
	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user2 ", user2.Name)
	user2.Change(amount)
	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}
func userInfo(user *UserBalance) {
	fmt.Println("Name ", user.Name, " | Saldo =", user.Saldo)
}
func TestCaseDeadlock(t *testing.T) {
	//disini hanya sampai Lock user1 dan tidak pernah Lock user2 karna saling lock
	user1 := UserBalance{
		Name:  "Ugik",
		Saldo: 1000000,
	}

	user2 := UserBalance{
		Name:  "Vita",
		Saldo: 1000000,
	}

	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 200000)
	time.Sleep(4 * time.Second)
	// fmt.Println("Name ", user1.Name, " | Saldo =", user1.Saldo)
	// fmt.Println("Name ", user2.Name, "Saldo =", user2.Saldo)

	userInfo(&user1)
	userInfo(&user2)
}
