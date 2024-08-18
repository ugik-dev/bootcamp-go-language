package main

// import "fmt"
import (
	"errors"
	"fmt"
	"golang-dasar/config"
	_ "golang-dasar/nouse"
	"golang-dasar/utils"
)

func TestError() error {
	return &utils.ErrorHandler{"ER"}
}

func main() {
	fmt.Println("Happy Coding!!, First time Go-Lang!!")
	dariUtil := utils.SayHello("Ugik")
	db := config.GetDatabase()
	fmt.Println(db)
	fmt.Println(dariUtil)
	hasil, err := perbandingan(3, 2)
	if err != nil {
		fmt.Println("Something wrong ", err)
	} else {
		fmt.Println("Success pembagian", hasil)
	}
	err2 := TestError()

	if err2 != nil {
		fmt.Println("Something wrong in Error Handler 1", err2.Error())
		fmt.Println("Something wrong in Error Handler 2", err2)
	}

	// returning
	// ress := callPerson("Ugik")
	// fmt.Println(ress)

	// ress1, ress2 := returnMultiple("Ugik")
	// fmt.Println(ress1, ress2)
	// ress3, _ := returnMultiple("Ugik")
	// fmt.Println("ignore return ", ress3)

	// return with class
	a, b, _ := getCat()
	fmt.Println(a, b)
	fmt.Println(getCat())
	// konstanta()
	// learnStringAndDeclareVariable()
	// learnBoolean()
	// learnNumber()
}

// jika tipe semua sama cukup type var di ahir func getCat() (jenis , jenis_kelamin string)
func getCat() (jenis string, jenis_kelamin string, umur int) {
	jenis = "Kucing"
	jenis_kelamin = "Laki-laki"
	umur = 4
	return jenis, jenis_kelamin, umur
}
func callPerson(firstName string) string {
	return "Hello " + firstName
}
func returnMultiple(firstName string) (string, string) {
	return "Hello " + firstName, "Jendral Ayani"
}
func perbandingan(a int, b int) (int, error) {
	if a > b {
		return 0, errors.New("nilai a lebih besar")
	} else {
		fmt.Println("nilai b lebih besar ")
	}

	return a / b, nil
	// var name1 = "Zea"
	// var name2 = "Zea"
	// var result bool = name1 == name2
	// fmt.Println("result string = ", result)

	// var var1 int = 3
	// var var2 int = 3
	// var result2 bool = var1 === var2
	// fmt.Println("result var = ", result2)
}

func konstanta() {
	const appName string = "First Time Golang - "
	const developerName = "Ugik DEV, "
	const (
		address   = "JL Totokarto, "
		kabupaten = "Pringsewu"
	)
	fmt.Println(appName, developerName, address, kabupaten)
}

func learnStringAndDeclareVariable() {
	var name string
	name = "Fazea"
	var secondName = "Audrilia"
	lastName := "Pramita"
	printName(name)
	printName(secondName)
	printName(lastName)

	var (
		firstName  = "Yuvita"
		middleName = "Tri"
		// lastName   = "Rejeki"
	)
	printName(firstName)
	printName(middleName)

}

func printName(name string) {
	fmt.Println("Data = ", name)
	// fmt.Println("Jumlah Karakter = ", len(name))
	fmt.Println("String Ke 1 (byte) = ", name[0])
	fmt.Println("String Ke 1 = ", string(name[0]))
}

func learnBoolean() {
	fmt.Println("ini true = ", true)
	fmt.Println("ini false = ", false)
}
func learnNumber() {
	fmt.Println("() init Integer Umur : ", 1)
	fmt.Println("(float) /  : ", 1.3)
}
