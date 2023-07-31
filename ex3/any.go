package main

import "fmt"

// Function that takes an interface{} parameter and uses type switches
// to handle different data types
func process(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Received an int:", v)
	case string:
		fmt.Println("Received a string:", v)
	case bool:
		fmt.Println("Received a bool:", v)
	default:
		fmt.Println("Received an unknown type")
	}
}

func InterfaceAssert() {
	var x interface{}

	x = 42
	process(x) // Output: Received an int: 42

	x = "Hello"
	process(x) // Output: Received a string: Hello

	x = true
	process(x) // Output: Received a bool: true

	x = 3.14
	process(x) // Output: Received an unknown type
}
