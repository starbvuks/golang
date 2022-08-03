package main

import "fmt"

func Loops() {
	// for loop
	i := 1
	for i <= 2 {
		fmt.Printf("in long: %v\n", i)
		i++
	}
	// or
	for i := 1; i <= 2; i++ {
		fmt.Printf("in short: %v\n", i)
	}
}
