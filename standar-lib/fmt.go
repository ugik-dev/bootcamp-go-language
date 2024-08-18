package main

import "fmt"

type Person struct {
	name    string
	address string
	age     int
}

func main() {
	formatingprint()
}

func formatingprint() {
	fmt.Println("Starting Go-Lang Standar Library")
	ugik := Person{"Ugik", "Pringsewu", 29}
	fmt.Println(ugik.name)
	fmt.Printf("Hallo %s, umur anda %d tahun\nAlamat anda di %s", ugik.name, ugik.age, ugik.address)
}
