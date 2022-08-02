package main

import (
	"fmt"
)

func main() {
	var confName = "Go Conference"
	const confTickets = 50
	var remTickets = 50

	fmt.Printf("Welcome to %v\n", confName)
	fmt.Println("We have a total of", confTickets, "tickets with", remTickets, "left")

	// must explicity define type on initialization
	var userName string
	userName = "Me"
	fmt.Printf("%v\n", userName)
}
