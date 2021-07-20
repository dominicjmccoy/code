package main

import (
	"fmt"
	"strconv"
)

// Have to use defined variables or get runtime error
// can declare several at once
var (
	x int     = 56
	y float32 = 78
)

// if declaring variables at package level, can't use :=
var l float32 = 70

func main() {
	// ways of declaring variables
	var i int
	i = 42

	var j int = 43

	k := 44

	// converting variable types
	var a float32
	a = float32(i)

	var myString string
	myString = strconv.Itoa(i)

	fmt.Printf("%v , %T\n", i, i)
	fmt.Printf("%v , %T\n", j, j)
	fmt.Printf("%v , %T\n", k, k)
	fmt.Printf("%v , %T\n", l, l)
	fmt.Printf("%v , %T\n", x, x)
	fmt.Printf("%v , %T\n", y, y)
	fmt.Printf("%v , %T\n", a, a)
	fmt.Printf("%v , %T\n", myString, myString)

}
