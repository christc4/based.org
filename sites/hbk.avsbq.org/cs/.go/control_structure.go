package main

import "fmt"

func main() {
	x := 10

	if x > 0 {
		fmt.Println("x is positive")
	} else {
		fmt.Println("x is negative")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	switch x {
	case 10:
		fmt.Println("x is ten")
	default:
		fmt.Println("x isn't ten")
	}
}

