package main

import "fmt"

func Pointers() {
	a := 5
	b := &a
	// *b == *&a
	fmt.Printf("%d, %d\n", a, *b)

	*b = 10
	fmt.Printf("%d, %d\n", a, b)
}
