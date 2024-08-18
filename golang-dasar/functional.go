package main

import "fmt"

// alias parameter function
type Categgory func(string) string

// bisa mengunakan alias atau identifikasi secara lengkap
// func buah(name string, filter func(string) string) {
func buah(name string, filter Categgory) {
	cat := categgory(name)
	res := name + " Merupakan Buah " + cat
	fmt.Println(res)
}

func categgory(name string) string {
	switch name {
	case "Jeruk":
		return "Asem"
	case "Durian":
		return "Gak Enak"
	default:
		return "Kira-kira gmna ya?"
	}
}

// input seperti array
func variadicFunc(action string, numbers ...int) int {
	total := 0
	for _, number := range numbers {
		if action == "jumlah" {
			total += number
		} else if action == "perkalian" {
			total *= number
		} else {
			total -= number
		}
	}
	return total
}
func main() {

	// res_variadic := variadicFunc("jumlah", 10, 10, 5)
	// fmt.Println(res_variadic)

	// numbering := []int{2, 3, 4}
	//konversi slice menjadi variable argumen
	// fmt.Println(variadicFunc("perkalian", numbering...))

	// anonymus function
	// anonymFunc := func(name string) string {
	// 	return "Hay, " + name
	// }
	// fmt.Println(anonymFunc("Ugik"))

	// function as parameter
	// res := categgory("Durian")
	// fmt.Println(res)
	// buah("Jeruks", categgory)

}

func recrusifFunc() {

}
