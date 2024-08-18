package main

import "fmt"

func main() {
	breaking()
	continueLop()
	// loopingMap()
	// loopingFor()
	//
}

func breaking() {
	fmt.Println(" ---- break ----")
	for i := 1; i < 10; i++ {
		if i == 4 {
			break
		}
		fmt.Println("m2 looping ke ", i)
	}
}

func continueLop() {
	fmt.Println(" ---- continue ----")
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("m2 looping ke ", i)
	}
}
func loopingMap() {
	fruits := []string{"Jeruk", "Mangga", "Apel", "Manggis"}

	for index, fruit := range fruits {
		fmt.Println("key ke = ", index, " -> ", fruit)
	}
	for _, fruit := range fruits {
		fmt.Println("m2 value = ", fruit)
	}
}
func loopingFor() {
	conter := 5

	for conter <= 10 {
		fmt.Println("looping ke ", conter)
		conter++
	}

	for i := 1; i < 10; i++ {
		fmt.Println("m2 looping ke ", i)
	}
}
