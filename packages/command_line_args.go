package main

import (
	"flag"
	"fmt"
	"math/rand"
)

// Command line program to generate a number bet 0 and 6.
// We can change the max value by sending a flag (-max=100)
func main() {
	// Define flags
	maxp := flag.Int("max", 6, "the max value")
	// Parse
	flag.Parse()
	// Generate a number between 0 and max
	fmt.Println(rand.Intn(*maxp))
}
