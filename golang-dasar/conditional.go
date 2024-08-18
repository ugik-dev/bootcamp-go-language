package main

import "fmt"

func main() {
	status := true
	password := "myPass"

	if status && password == "myPass" {
		fmt.Println("masuk")
	} else if password != "myPass" {
		fmt.Println("wrong Password")
	} else {
		fmt.Println("status bukan true")
	}
	hari := "Senin"
	switch hari {
	case "Sabtu":
		fmt.Println("ini hari sabtu libur")
	case "Minggu":
		fmt.Println("ini hari minggu libur")
	default:
		fmt.Println("working day!!")
	}

}
