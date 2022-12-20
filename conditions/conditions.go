package main

import (
	"fmt"
)

func main() {
	x := 5
	y := 6
	if x%2 == 0 {
		fmt.Println("Even number ")
	} else {
		fmt.Println("Odd number")
	}

	switch {
	case y%2 == 0:

		fmt.Printf("%v Even number", y)
		break

	default:
		fmt.Printf("%v Odd number", y)
	}
}
