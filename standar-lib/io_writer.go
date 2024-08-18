package main

import (
	"bufio"
	"os"
)

// https://pkg.go.dev/bufio
func main() {
	write := bufio.NewWriter(os.Stdout)
	_, _ = write.WriteString("Halo \n")
	_, _ = write.WriteString("Ini Apa ?")
	write.Flush()
}
