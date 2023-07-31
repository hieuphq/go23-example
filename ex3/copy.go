package main

import "fmt"

func Copy() {

	// Creating the original slice
	originalSlice := []int{1, 2, 3, 4, 5}

	// Creating a new slice with the same length as the original slice
	copiedSlice := make([]int, len(originalSlice))

	// Using the copy() function to copy elements from the original slice to the new slice
	copy(copiedSlice, originalSlice)

	// Printing the original slice
	fmt.Println("Original Slice:", originalSlice)
	// Output: [1 2 3 4 5]

	// Printing the copied slice
	fmt.Println("Copied Slice:", copiedSlice)
	// Output: [1 2 3 4 5]

	// Modifying the first element of the copiedSlice
	copiedSlice[0] = 100

	// Printing the copied slice
	fmt.Println("Copied Slice after changed:", copiedSlice)
	// Output: [100 2 3 4 5]
}
