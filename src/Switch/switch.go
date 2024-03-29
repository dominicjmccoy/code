package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Print("Switch 1\n")
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s\n", os)

	}

	fmt.Print("Switch 2\n")
	fmt.Println("When is Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")

	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In 2 days")
	default:
		fmt.Println("Not for ages")
	}
	fmt.Println("Switch 3")
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good Evening")
	}
}
