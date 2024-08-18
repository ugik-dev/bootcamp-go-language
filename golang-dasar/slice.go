package main

import (
	"fmt"
)

func main() {
	freshSlice()
	// fromArray()
}

func mapping() {
	// slice := []string{"Januari", "Februari", "Maret", "April", "Mei", "Juni"}
}

func freshSlice() {
	slices := make([]string, 2, 5)
	slices[0] = "Kucing"
	slices[1] = "Anjing"
	// slices[2] = "Domba" // tidak bisa ditambah karna panjang hanya 2 harus mengunakan append

	fmt.Println(slices)
	fmt.Println("panjang array = ", len(slices), " | kapasitas array", cap(slices))

	slices2 := append(slices, "Singa")
	fmt.Println(slices2)
	fmt.Println("panjang array = ", len(slices2), " | kapasitas array", cap(slices2))

	copySlice := slices2
	fmt.Println("copy 1", copySlice)
	copySlice2 := make([]string, len(slices2), cap(slices2))
	copy(copySlice2, slices2)
	fmt.Println("copy 2", copySlice2)

	slices2[0] = "Harimau"
	fmt.Println("copy 1", copySlice)
	fmt.Println("copy 2", copySlice2)
	// saat mengcopy maka jika terjadi perubahan parent tidak ikut berubah

}

func fromArray() {
	fruits := [...]string{"Mangga", "Jeruk", "Apel"}
	sliceFruits := fruits[2:]
	fmt.Println(sliceFruits)
	sliceFruits[0] = "Kepala Baru"
	fmt.Println("di slice berubah", sliceFruits)
	fmt.Println("di array ikut berubah", fruits)

	sliceFruits2 := append(sliceFruits, "Durian Append")
	fmt.Println("sliceFruits", sliceFruits)
	fmt.Println("sliceFruits2", sliceFruits2)
	fmt.Println("fruits", fruits)
}
