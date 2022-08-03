package main

import "fmt"

func Conditionals() {

	// if else
	x := 5
	y := 10

	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	// if elseif
	color := "purple"

	if color == "red" {
		fmt.Printf("color is red \n")
	} else if color == "green" {
		fmt.Printf("color is green \n")
	} else {
		fmt.Printf("color is purple \n")
	}

	//switch
	switch color {
	case "red":
		fmt.Printf("color is red \n")
	case "blue":
		fmt.Printf("color is blue \n")
	default:
		fmt.Printf("color is purple \n")
	}
}
