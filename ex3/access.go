package main

import "fmt"

func ClonePointers() {
	s := []int{10, 20, 30, 40}
	s2 := make([]*int, 0, len(s))

	for i, val := range s {
		fmt.Println(val)
		s2 = append(s2, &s[i])
	}

	for _, val := range s2 {
		fmt.Print(*val)
		fmt.Print(" ")
	}

	fmt.Print("\n")
	fmt.Println(s2)
}
