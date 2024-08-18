package latihangoroutine

import (
	"fmt"
	"sync"
	"testing"
)

/*
*
pool ibarat antrian yang bisa diisi, jadi ketia put data diisi dan di get diambil dan data yang diput hilang
lalu diput lagi atau dikembalikan

jadi ketika di put akan hilang
*/
func TestPool(t *testing.T) {
	// data kosong akan dianggap nil
	// pool := sync.Pool{}

	// untuk definisi data kosong
	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}

	pool.Put("Fazea")
	pool.Put("Pramita")
	pool.Put("Audrilia")

	// for i := 0; i <= 10; i++ {
	// 	go func() {
	data := pool.Get()
	fmt.Println("d1", data)
	data2 := pool.Get()
	fmt.Println("d2", data2)
	data3 := pool.Get()
	fmt.Println("d3", data3)
	pool.Put(data3) //contoh data dikembalikan lagi
	data3 = pool.Get()
	fmt.Println("d4", data3)
	data3 = pool.Get()
	fmt.Println("d5", data3)
	data3 = pool.Get()
	fmt.Println("d6", data3)
	// }()
	// }
}
