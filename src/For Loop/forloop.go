package main

import (
	"fmt"
	"math"
)

// if statement can have logic before it
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
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
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

}
