package main

import (
	"fmt"
	"reflect"
)

/*
*
Masih sangat banyak fungsi lain dari reflect
https://pkg.go.dev/reflect
*/
type Inputan struct { //valueType.Name()
	Name    string `required:"true" max:"10" min:"1"`
	Address string `required:"true" max:"100"`
}

func scanInputan(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)
	fmt.Println(valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		var structField reflect.StructField = valueType.Field(i)
		fmt.Println(structField.Name, "dengan Type ", structField.Type)
		structField.Tag.Get("required")
	}
}

func IsValid(value any) bool {
	v := reflect.TypeOf(value)
	fmt.Println("ini merukan struct dar ", v.Name())
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			if data == "" {
				return false
			}
		}
	}

	return true
}
func main() {

	inputan := Inputan{"", "portugal"}
	scanInputan(inputan)
	res := IsValid(inputan)
	if res {
		fmt.Println("Data Valid")
	} else {
		fmt.Println("Data Tidak Valid")
	}
}
