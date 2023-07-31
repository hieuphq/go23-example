package main

import "fmt"

func arraySlide() {
	coral := [4]string{
		"blue coral",
		"staghorn coral",
		"pillar coral",
		"elkhorn coral",
	}

	coral[1] = "foliose coral"

	fmt.Println(coral[0])
	fmt.Println(coral[1])
	fmt.Println(coral[2])
	fmt.Println(coral[3])

	// Compile error
	// fmt.Println(coral[4])

	a := coral[0:2]
	b := coral[1:3]

	fmt.Println(a, b)

	b[0] = "Sth else"

	fmt.Println(a, b)
	fmt.Println(coral)
}
