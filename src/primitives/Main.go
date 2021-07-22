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

	print("\n\n")
	// Bit shifting
	v := 8              // 2^3
	fmt.Println(v << 3) // 2^3 * 2^3 = 2^6 = 64
	fmt.Println(v >> 3) // 2^3 / 2^3 = 2^0 = 1
	println("\n\n")
	// floating point numbers
	// float32 float64
	// module and bitwise operator not available
	var h float64 = 3.14
	j := 13.7e72
	k := 2.1e14
	fmt.Printf("%v , %T\n", h, h)
	fmt.Printf("%v , %T\n", j, j)
	fmt.Printf("%v , %T\n", k, k)
	fmt.Println(h + j)
	fmt.Println(h - j)
	fmt.Println(h * j)
	fmt.Println(h / j)

	println("\n\n")

	// Complex Types
	// complex64 and complex128
	var f complex128 = 1 + 2i
	g := 2 + 5.2i
	fmt.Printf("%v , %T\n", f, f)

	fmt.Println(f + g)
	fmt.Println(f - g)
	fmt.Println(f * g)
	fmt.Println(f / g)

	println("\n\n")

	// text type
	// string = any UTF8 character
	s := "this is a string"
	fmt.Printf("%v , %T\n", s, s)
	// This will not print "i", it will print "105 uint8"
	// this is because strings are aliases for bytes
	fmt.Printf("%v , %T\n", s[2], s[2])
	// this will work
	fmt.Printf("%v , %T\n", string(s[2]), s[2])

	// stirng are not immutable, indivudal letter can't be changed
	// eg. s[3] = "g" will not work

	// operators on strings
	s2 := "this is also a string"
	fmt.Printf("%v , %T\n", s+s2, s+s2)

	// can covert to collection of bytes
	t := []byte(s)
	fmt.Printf("%v , %T\n", t, t)

	// Rune represents UTF32
	// declared using single quotes
	r := 'a'
	var z rune = 'a'
	fmt.Printf("%v , %T\n", r, r)
	fmt.Printf("%v , %T\n", z, z)

}
