package main

import "fmt"

func main() {
	type NoKTP string
	var ktpCustomer NoKTP = "19010103c"
	const ktpOwner NoKTP = "19010102o"

	var temp1 string = "1234"
	fmt.Println(temp1)

	var temp2 NoKTP = NoKTP(ktpCustomer)
	fmt.Println(temp2)

	var temp3 NoKTP = NoKTP("tmp3")
	fmt.Println(temp3)
}
