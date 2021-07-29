package main

import (
	"fmt"
)

// constants can be shadowed, inner version of a takes precedence
const a int32 = 50

// enumerated constants
const (
	x = iota
	y
	z
)

func main() {
	// Constants have to be ready at compile time
	const a int = 42
	// e.g  const myConst float64 = sin(1.57) will not work as sin() will only
	// be available at run time
	fmt.Printf("%v, %T\n", a, a)
	// perform operations on constants
	var b int = 34
	fmt.Printf("%v, %T\n", a+b, a+b)

	// compiler replaces constant with its value, so type can be left out
	// and the it will be implied
	const c = 43
	var d int16 = 56
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", c+d, c+d)
	// the above works despite differences in type

	// enumerated constants
	fmt.Print("Enumerated Constants\n")
	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v, %T\n", y, y)
	fmt.Printf("%v, %T\n", z, z)

}
