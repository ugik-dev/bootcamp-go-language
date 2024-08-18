package main

import (
	"fmt"
	"strconv"
)

func main() {
	// other https://pkg.go.dev/strconv
	intAge := 28
	stringAge := "-29"

	convStringAge, err := strconv.Atoi(stringAge)
	if err == nil {
		var result = intAge + convStringAge
		fmt.Println("hasil jumlah ", result)
	} else {
		fmt.Println("Error jumlah", err.Error())
	}


	fmt.Println("Umur saya " + strconv.Itoa(intAge))
}
