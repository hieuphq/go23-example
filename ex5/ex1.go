package main

import (
	"fmt"
)

func hello() {
	fmt.Println("Hello world goroutine")
}
func ex1() {
	go hello()

	// FIX: wait for 1 second
	// time.Sleep(1 * time.Second)

	fmt.Println("main function")
}
