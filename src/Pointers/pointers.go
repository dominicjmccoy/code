package main

import "fmt"

func main() {
	i, j := 42, 2701
	p := &i // points to i
	fmt.Println(*p)
	q := &j
	fmt.Println(*q)
}
