package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// https://pkg.go.dev/bufio
func main() {
	input := strings.NewReader("this is long\nnewline\n")
	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()
		// line, prefix, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(string(line))
	}
}
