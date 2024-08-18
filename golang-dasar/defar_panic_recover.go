package main

import "fmt"

func log() {
	fmt.Println("Finising")
	message := recover()
	if message != nil {
		fmt.Println("With message error from panic : ", message)
	}
}

func runApp() {
	/**
	KOMENTAR
	defer akan dieksekusi walaupun terjadi error
	*/
	// diakhir fungsi maka log() akan berjalan walaupun di atas posisinya
	defer log()
	fmt.Println("App Runing")
	fmt.Println("Ex1")
	panic("something wrong guys")
	fmt.Println("Ex2")
	fmt.Println("Ex2")
}

func main() {
	runApp()
}
