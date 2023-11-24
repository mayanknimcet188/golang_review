package main

import "fmt"

// embedded type basically supports "is-a" kind of relationship in contrast to the "has-a"
// relationship normally between the structs and its fields
// In other words to put it in a dumb way it is similar to inheritance concept
type Person struct {
	// here type Persons "has-a" name
	name string
}

func (p *Person) talk() {
	fmt.Println("This person can talk")
}

type Android struct {
	// here instead of saying Android "has-a" person, we would like to say Android "is-a" person
	Person
	Model string
}

func main() {
	a := new(Android)
	// since the Android is a person so it can talk as well
	a.talk()
}
