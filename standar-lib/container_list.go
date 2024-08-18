package main

import (
	"container/list"
	"fmt"
)

func main() {
	// data := list.New()
	/**
	Blok container ring juga bisa
	https://pkg.go.dev/container/ring
	https://pkg.go.dev/container/list
	*/
	var data *list.List = list.New()

	data.PushBack("A")
	data.PushBack("B")
	data.PushBack("C")
	data.PushFront("D")

	var firstData *list.Element = data.Front()
	fmt.Println(firstData.Value)

	for i := data.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
