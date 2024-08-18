package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("Its Validation Error")
	NotFoundError   = errors.New("Its Not Found Error")
)

func lgoin() error {
	username := "ugikdev2"
	password := "123"

	if username == "" || password == "" {
		return ValidationError
	} else if username != "ugikdev" {
		return NotFoundError
	}

	return nil
}

func main() {
	err := lgoin()
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("Error Validation")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("Error Not Found", err)
		} else {
			fmt.Println("Error Unkown")

		}
	}
}
