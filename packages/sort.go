package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

// sort.Sort takes a sort.Interface and sorts it.
// sort.Interface requires three methods: Len, Less, Swap
// So here we are trying to sort our own data structure called Person and that too "ByName"
// Hence to confirm ByName to the sort.Interface we are defining 3 methods Len, Less and Swap on it
type ByName []Person

func (ps ByName) Len() int {
	return len(ps)
}

func (ps ByName) Less(i, j int) bool {
	return ps[i].Name < ps[j].Name
}

func (ps ByName) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

// Sorting by Age
type ByAge []Person

func (this ByAge) Len() int {
	return len(this)
}

func (this ByAge) Less(i, j int) bool {
	return this[i].Age < this[j].Age
}

func (this ByAge) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func main() {
	kids := []Person{
		{"Jill", 9},
		{"Jack", 10},
	}
	sort.Sort(ByName(kids))
	fmt.Print("Sorting By Name...")
	fmt.Println(kids)
	fmt.Print("Sorting By Age...")
	sort.Sort(ByAge(kids))
	fmt.Println(kids)
}
