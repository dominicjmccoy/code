package main

import "fmt"

func main() {
	// Boolean type
	var n bool = true
	fmt.Printf("%v, %T\n", n, n)

	b := 1 == 1
	fmt.Printf("%v, %T\n", b, b)

	var m bool
	fmt.Printf("%v, %T\n", m, m)

	// Numeric Types
	// Int - unspecified size, 32 bits minimum
	// int8 -128 - 127
	// int16 -32,768 - 33,2767
	// int32 -2,147,483,648 - 2,147,483,647
	// int64 .......
	x := 42
	fmt.Printf("%v, %T\n", x, x)

	// Can also have unsigned int
	var y uint16 = 42
	fmt.Printf("%v, %T\n\n", y, y)

	// uint goes from 0 - 255 etc

	// aritmethic operations
	p := 10 // 1010
	q := 3  // 0011
	fmt.Println(p + q)
	fmt.Println(p - q)
	fmt.Println(p * q)
	fmt.Println(p / q)
	fmt.Println(p % q)
	fmt.Print("\n\n")

	// bit operators
	fmt.Println(p & q)  // AND  0010
	fmt.Println(p | q)  // OR   1011
	fmt.Println(p ^ q)  // EXCLUSIVE OR 1001
	fmt.Println(p &^ q) // AND NOT 0100

	fmt.Print("\n\n")
	// Bit shifting
	v := 8              // 2^3
	fmt.Println(v << 3) // 2^3 * 2^3 = 2^6 = 64
	fmt.Println(v >> 3) // 2^3 / 2^3 = 2^0 = 1
}
