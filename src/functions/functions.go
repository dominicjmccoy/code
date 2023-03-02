package main

import "fmt"

// takes two numbers and returns sum
func add(x, y int) int {
	return x + y
}

// swaps two strings
func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// a list of boolean variables, defaults to false
var c, python, java bool

// more variables
var x, y int = 1, 2

func main() {
	// Add two numbers
	fmt.Println(add(42, 13))

	// swap two strings
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	// split
	fmt.Println(split(14))

	var i int
	fmt.Println(i, c, python, java)

	var english, french, german = true, true, "no!"
	fmt.Println(x, y, english, french, german)

	// no need to use var inside a function, can use :=
	z := 3
	fmt.Println(z)
}
