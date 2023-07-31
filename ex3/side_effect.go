package main

import "fmt"

// Function that sums all elements in a slice without modifying the original slice
func sumSliceElements(slice []int) int {
	sum := 0
	for _, num := range slice {
		sum += num
	}
	return sum
}

// Function that creates a new slice and reverses its elements without modifying the original slice
func reverseSlice(slice []int) []int {
	reversed := make([]int, len(slice))
	for i := 0; i < len(slice); i++ {
		reversed[i] = slice[len(slice)-1-i]
	}
	return reversed
}

// Function that reverses its elements modifying the original slice
func reverseOriginalSlice(slice []int) []int {
	for i := 0; i < len(slice)/2; i++ {
		tmp := slice[i]
		slice[i] = slice[len(slice)-1-i]
		slice[len(slice)-1-i] = tmp
	}

	return slice
}

func SideEffectSlice() {
	originalSlice := []int{1, 2, 3, 4, 5}

	// Avoid unintended side effects by passing a copy of the slice
	sum := sumSliceElements(append([]int{}, originalSlice...))
	fmt.Println("Sum of elements in the original slice:", sum)

	// Avoid unintended side effects by passing a copy of the slice
	reversedSlice := reverseSlice(append([]int{}, originalSlice...))
	fmt.Println("Reversed slice:", reversedSlice)

	// The original slice remains unchanged
	fmt.Println("Original slice:", originalSlice)

	// Avoid unintended side effects by passing a copy of the slice
	// reverseOriginalSlice(originalSlice)
	reverseOriginalSlice(append([]int{}, originalSlice...))

	// The original slice remains changed
	fmt.Println("Original slice changed:", originalSlice)
}
