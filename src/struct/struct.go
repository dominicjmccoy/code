package main

// a struct is a collection of fields

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	fmt.Println("Printing a struct: ", Vertex{1, 2})
	// create a new vertex
	v := Vertex{3, 4}
	// set a new value or access a field using a dot.
	v.X = 5
	fmt.Println(v)

	z := Vertex{7, 8}
	p := &z
	p.X = 1e9
	fmt.Println("z:", z)

	// declare some vertices
	v1 := Vertex{1, 2}

	v2 := Vertex{X: 1}
	v3 := Vertex{}
	p = &Vertex{1, 2}

	fmt.Println("v1: ", v1, " p: ", p, "v2: ", v2, " v3:", v3)

}
