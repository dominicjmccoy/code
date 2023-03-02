package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	x := rand.Intn(10)
	fmt.Println("A random number: ", x)
	fmt.Printf("Now you have %g problems\n", math.Sqrt(7))
}
