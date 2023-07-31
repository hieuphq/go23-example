package main

import "fmt"

func handlePanic() {
	if r := recover(); r != nil {
		// r contains the value passed to panic()
		fmt.Println("Recovered from panic:", r)
	}
}

func ExampleFunction() {
	defer handlePanic()

	// Some code that may panic
	panic("something went wrong!")
}
