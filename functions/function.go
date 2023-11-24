package main

import "fmt"

// return types in go can also be named as below
func average(values []float64) (avg float64) {
	avg = 0.0
	for _, value := range values {
		avg += value
	}
	return avg / float64(len(values))
}

// variadic functions - where the last parameter can be of special form as below
func add(args ...int) (t int) {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func main() {
	values := []float64{1, 2, 3, 43, 5}
	fmt.Println(average(values))
	fmt.Println(add(1, 2, 23))
}
