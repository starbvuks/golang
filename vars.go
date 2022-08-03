package main

import (
	"fmt"
)

var (
	movieName   string = "Stalker"
	director    string = "Tarkofsky"
	movieNumber int    = 3
	released    int    = 1976
)

var (
	counter int = 0
)

func Vars() {
	fmt.Printf("%s, %T", movieName, released)
}
