package main

import "fmt"

func main() {
	// a function like below with the nonlocal variable it references is called a closure
	add := func(args ...int) int {
		total := 0
		for _, v := range args {
			total += v
		}
		return total
	}
	fmt.Println(add(1, 2, 3, 4, 5))
	fmt.Println("Generating Even Numbers...")
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
}

// closures can be used to write functions that return another functions like...
// below function can be interpreted as a function that returns another function
// which in turn returns unsigned integers, in this case generates the next even number
func makeEvenGenerator() func() uint {
	i := uint(0)
	// here the function func references a nonlocal variable `i` initialized to 0 of the parent function `makeEvenGenerator`
	// increments it by 2 and returns the same
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}
