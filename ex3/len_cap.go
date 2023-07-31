package main

import "fmt"

func LenCap() {
	slice := make([]int, 3, 5)
	fmt.Println("Length of the slice:", len(slice))   // Output: 3
	fmt.Println("Capacity of the slice:", cap(slice)) // Output: 5

	slice = append(slice, 1)
	fmt.Println(slice)
	fmt.Println("Length of the slice:", len(slice))   // Output: 4
	fmt.Println("Capacity of the slice:", cap(slice)) // Output: 5

	slice = append(slice, 2, 3)
	fmt.Println(slice)
	fmt.Println("Length of the slice:", len(slice))   // Output: 4
	fmt.Println("Capacity of the slice:", cap(slice)) // Output: 5

	// Pre-allocating a slice to minimize reallocations when appending elements.
	initCap := 500 // 10, 100, 200
	slice2 := make([]int, 0, initCap)
	currCap := cap(slice2)
	for i := 0; i < 1000; i++ {
		slice2 = append(slice2, i)
		if currCap != cap(slice2) {
			fmt.Println("Reallocate new slice with capacity: ", cap(slice2))
			currCap = cap(slice2)
		}
	}
}
