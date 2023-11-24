package main

import "fmt"

func second() {
	print("Second...\n")
}

func first() {
	print("First...\n")
}

func main() {
	// defer moves the "defered" function to be called at the end of the parent function
	// here the function second will be called at the end of the main function
	// usually used when we need to free resources in some way...like making sure closing a file, we defer the
	// closing operation whenever a file is opened
	// defered functions are run even when there is a runtime panic
	defer second()
	first()
	// panic function causes runtime error, which mainly happens due to programming errors
	// now handling a runtime panic can be done via the in-built recover function that stops the panic function
	// and returns the value that was passed to the panic function during the runtime error
	// the recover function needs to be paired with the defer keyword so as to run it even after the runtime panic
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC!!!")
}
