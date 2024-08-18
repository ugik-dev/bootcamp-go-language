package main

import (
	"fmt"
	"sort"
)

type Person2 struct {
	name string
	age  int
}
type PersonSlice []Person2

func (p PersonSlice) Len() int {
	return len(p)
}
func (p PersonSlice) Less(i, j int) bool {
	return p[i].age < p[j].age // by age
	// return p[i].name < p[j].name //byname
}

func (p PersonSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func main() {

	// ugik := Person{"Ugik", "Pringsewu", 29}
	persons := []Person2{
		{"Ugik", 29},
		{"Vita", 28},
		{"Zea", 3},
		{"Eyang", 50},
		{"Kakak Anum", 10},
	}
	fmt.Println(persons)

	sort.Sort(PersonSlice(persons))
	fmt.Println(persons)

}
