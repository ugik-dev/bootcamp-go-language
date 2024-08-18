/**
Lebih ke pendekatan ke model
*/

package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// start interface
type HasName interface {
	GetName() string
}

func (person Customer) GetName() string {
	return person.Name
}

//interface kosong

func Ups() any {
	return "Ups"
}

// end interface

func SayHello(param HasName) {
	fmt.Println("Hello My Name is,", param.GetName())
}

// Pembuatan method, fungsi dalam strutc
func (person Customer) callMe() {
	fmt.Println("Hallo my name is, ", person.Name)
}
func (person Customer) myAddress(streetNumber string) string {
	personAge := person.Age.(string)
	return "Hallo my address in, " + person.Address + ", Number " + streetNumber + " My Age is " + personAge
}

func main() {
	var ugik Customer
	fmt.Println(ugik)
	ugik.Name = "Sugi Pramana"
	ugik.Age = 29
	ugik.Address = "Jl Totokarto"
	fmt.Println(ugik)
	fmt.Println(ugik.Address)
	var address = ugik.myAddress("30A")
	fmt.Println(address)
	ugik.callMe()

	yuvita := Customer{
		Name:    "Yuvita Tri Rejeki",
		Address: "Jl Bandung Baru",
		Age:     28,
	}
	fmt.Println(yuvita)

	zea := Customer{"Fazea Audrilia Pramita", "Pringsewu", 3}
	fmt.Println(zea)

	// Interface
	fmt.Println("Output dari interface = ", zea.GetName())

	// interface kosong
	var kosong any = Ups()
	fmt.Println(kosong)

}
