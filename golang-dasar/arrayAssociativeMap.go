package main

import "fmt"

func main() {

	//Struct.go untuk model lainya
	var person map[string]string = map[string]string{}
	person["name"] = "Fazea"
	person["address"] = "Pringsewu"

	person2 := map[string]string{
		"name":    "Ugikdev",
		"address": "Bangka Island",
		"wife":    "Yuvita",
	}

	fmt.Println(person)
	fmt.Println(person["address"])
	person2["second adress"] = "Totokarto"
	fmt.Println(person2)
	delete(person2, "second adress")
	fmt.Println(person2)

}
