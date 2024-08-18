package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	/**
	  argumen adalah seperti url terakhir setelah go run os.go sugi pramana
	  maka sugi pramana merupakan args / argumen
	  gunakan kutip untuk menyatukan arg jadi satu gunakan petik
	*/

	for index, arg := range args {
		fmt.Printf("argumen ke %d = %s \n", index, arg)
	}

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("error", err)
	}
}
