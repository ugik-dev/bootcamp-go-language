package main

import (
	"flag"
	"fmt"
)

func main() {

	//  go run flag.go -username=ugikdev -host=192.168.1.1
	username := flag.String("username", "root", "DB username")
	password := flag.String("password", "", "DB password")
	host := flag.String("host", "localhost", "DB host")

	flag.Parse()

	fmt.Println("Username", *username)
	fmt.Println("Pasword", *password)
	fmt.Println("Host", *host)
}
