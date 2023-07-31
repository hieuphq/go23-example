/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"flag"
	"fmt"
)

func main() {
	// Flags
	flag.Parse()
	args := flag.Args()

	fmt.Println("Output: ", args)
	fmt.Println("Args: ", flag.Args())

}
