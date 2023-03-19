package main

import "fmt"

func main() {
	i, j := 42, 2701
	fmt.Println("i: ", i, "j: ", j)
	p := &i // points to i
	fmt.Println("p*: ", *p)
	*p = 21
	fmt.Println("i is now: ", i)

	p = &j
	*p = *p / 37
	fmt.Println("j is now: ", j)
}
