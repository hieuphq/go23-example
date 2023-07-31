package main

import "fmt"

func memory() {
	// Create a slice with length 3 and capacity 5
	slice1 := make([]int, 3, 5)
	slice1[0] = 10
	slice1[1] = 20
	slice1[2] = 30

	// Append two elements to the slice
	slice1 = append(slice1, 40, 50)

	// Display the contents of slice1
	fmt.Println("Slice 1:", slice1) // Output: Slice 1: [10 20 30 40 50]

	// Create a new slice that is a subset of slice1
	slice2 := slice1[1:4]
	fmt.Println("Slice 2:", slice2) // Output: Slice 2: [20 30 40]

	// Modify an element in slice2
	slice2[0] = 100

	// Display the contents of slice1 after modifying slice2
	fmt.Println("Slice 1 after modifying Slice 2:", slice1) // Output: Slice 1 after modifying Slice 2: [10 100 30 40 50]

}
