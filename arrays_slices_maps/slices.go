package main

import "fmt"

func main() {
	// slices in go are more predominantly used in comparision to arrays
	// as we don't have to provide the length here unlike arrays
	// a slice of float64 type with length 5
	x := make([]float64, 5)
	fmt.Println(x)
	// a slice y with length 5 but capactity 10 of the underlying array
	y := make([]float64, 5, 10)
	fmt.Println(y)

	// using [low:high] approach to create a slice
	arr := [5]float64{1, 2, 3, 4, 5}
	arrSlice := arr[0:5]
	fmt.Println(arrSlice)
}
