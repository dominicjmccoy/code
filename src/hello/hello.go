package main

import (
	"fmt"
	"time"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
	fmt.Println("The time is now", time.Now())
}
