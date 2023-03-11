package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// declaring a constant
const Pi = 3.14

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

// takes two numbers and returns sum
func add(x, y int) int {
	return x + y
}

func needInt(x int) int           { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

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

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

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

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var c int
	var d float64
	var e bool
	var f string
	fmt.Printf("%v %v %v %q\n", c, d, e, f)

	var m, n int = 3, 4
	var o = math.Sqrt(float64(m*m + n*n))
	var p = uint(o)
	fmt.Println(m, n, p)

	v := 3.14
	fmt.Printf("v is of type: %T\n", v)

	fmt.Println("Pi:", Pi)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
