package main

import (
	"fmt"
	"math"
)

// Exercis to return square root of argument
func Sqrt(x float64) float64 {
	z := 1.0
	z -= (z*z - x) / (2 * z)
	fmt.Println("Square Root: ", z)
	return 0.0
}

// if statement can have logic before it
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

// if statement can have logic before it
func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {

	// standard for loop
	sum := 0
	for i := 0; i < 10; i++ {
		sum += 1
	}
	fmt.Println(sum)

	// init and post statments are optional
	sum1 := 1
	for sum1 < 1000 {
		sum1 += sum1
	}
	fmt.Println(sum1)

	// infinite loop
	// for {}

	//
	fmt.Println(sqrt(2), sqrt(-4))

	// if statement can have logic before it
	fmt.Println("Power 1: ",
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	fmt.Println("Power 2: ",
		pow2(3, 2, 10),
		pow2(3, 3, 20),
	)

	// Exercise find Square Root
	fmt.Println(Sqrt(2))
}
