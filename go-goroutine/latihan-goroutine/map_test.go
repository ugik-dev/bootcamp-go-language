package latihangoroutine

import (
	"fmt"
	"sync"
	"testing"
)

func addMap(data *sync.Map, group *sync.WaitGroup, value int) {
	defer group.Done()
	group.Add(1)
	data.Store(value, value)
}

func TestMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i <= 100; i++ {
		go addMap(data, group, i)
	}

	group.Wait()

	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, "value=", value)
		return true
	})
}
