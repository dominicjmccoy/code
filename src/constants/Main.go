package main

import "fmt"

func main() {
	// Constants have to be ready at compile time
	const myConst int = 42
	// e.g  const myConst float64 = sin(1.57) will not work as sin() will only
	// be available at run time
	fmt.Printf("%v, %T\n", myConst, myConst)
}
