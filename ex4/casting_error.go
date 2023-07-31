package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	Msg string
}

func (e *MyError) Error() string {
	return e.Msg
}

func foo() error {
	return fmt.Errorf("foo: %w", &MyError{Msg: "custom error"})
}

func Casting() {
	err := foo()

	// Check if the error contains MyError in the chain
	if errors.Is(err, &MyError{}) {
		fmt.Println("MyError found in the error chain.")
	}

	// Extract the MyError from the error chain
	var myErr *MyError
	if errors.As(err, &myErr) {
		fmt.Println("Extracted MyError:", myErr)
	}

	// Unwrap the error and continue checking
	for err != nil {
		fmt.Println("Error:", err)
		err = errors.Unwrap(err)
	}
}
