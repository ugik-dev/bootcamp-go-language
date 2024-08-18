package main

import "fmt"

func main() {
	var arr [3]string
	fmt.Println(arr)
	arr[0] = "Fazea"
	arr[1] = "Audrilia"
	arr[2] = "Pramita"
	fmt.Println(arr)

	var arr2 = [3]string{"Fazea", "Audrilia", "Pramita"}

	fmt.Println(arr2[1])
	fmt.Println("jumlah array", len(arr2))

	var arr3 = [...]string{"Fazea", "Audrilia", "Pramita", "Binti", "Sugi", "Pramana"}

	fmt.Println(arr3)
	fmt.Println("jumlah array", len(arr3))
	arr3[3] = ""
	fmt.Println(arr3)

	// Slice

	arr4 := [...]string{"Fazea", "Audrilia", "Pramita", "Binti", "Sugi", "Pramana"}
	slice := arr4[2:4]
	fmt.Println(slice)
	// slice2 := arr4[:5]
	fmt.Println(arr4[:5])
	fmt.Println(arr4[2:])
	fmt.Println(arr4[:])

	var slicePure []string 
	slicePure[0] = "satu"
	fmt.Println(slicePure)
}
