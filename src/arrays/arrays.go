package main

import "fmt"

func main() {
	// create an array of 2 elements
	// the length of the array is part of the declaration
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// create an array of prime numbers
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// slices allow us to take a section of an array
	var slice []int = primes[1:4]
	fmt.Println(slice)

	// slices do not store data, just describe a section of the underlying array
	//  chages to a slice will happen to  the array
	names := [4]string{"John", "Paul", "Ringo", "George"}
	fmt.Println(names)

	c := names[0:2]
	d := names[1:3]
	fmt.Println(c, d)
	d[0] = "xxxx"
	fmt.Println(c, d)
	fmt.Println(names)

	// Slice literals are like array literals without the length
	arrayLit := [3]bool{true, true, true}
	fmt.Println(arrayLit)

	// no length needed for this slice. creates the array then builds a slive as reference to it
	sliceLit := []bool{true, true, true}
	fmt.Println(sliceLit)

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	// slice defaults
	t := []int{2, 3, 5, 7, 11, 13}
	u := t[1:4]
	fmt.Println(u)

	v := t[:2]
	fmt.Println(v)

	g := t[1:]
	fmt.Println(g)
	fmt.Println(t)

}
