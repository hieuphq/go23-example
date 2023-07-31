package main

func main() {
	// arraySlide()
	ClonePointers()
	// ClonePointers()
	// InterfaceAssert()

	// v := Vertex{X: 1, Y: 1}

	// updatedV := v.UpdateX(2)

	// fmt.Println(v)
	// fmt.Println(updatedV)
}

type Vertex struct {
	X int
	Y int
}

func (v Vertex) UpdateX(x int) Vertex {
	v.X = x

	return v
}
