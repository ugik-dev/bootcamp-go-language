package main

import (
	"fmt"
	"path"
	"path/filepath"
)

func main() {
	//untuk url
	fmt.Println(path.Dir("hellow/world.go"))
	fmt.Println(path.Base("hellow/world.go"))
	fmt.Println(path.Ext("hellow/world.go"))
	fmt.Println(path.Join("hellow", "word", "main.go"))

	//untuk filebase OS, karna diwindows tidak suport backslise
	fmt.Println(filepath.Dir("hellow/world.go"))
	fmt.Println(filepath.Base("hellow/world.go"))
	fmt.Println(filepath.Ext("hellow/world.go"))
	fmt.Println(filepath.IsAbs("hellow/world.go"))
	fmt.Println(filepath.IsLocal("hellow/world.go"))
	fmt.Println(filepath.Join("hellow", "word", "main.go"))
}
