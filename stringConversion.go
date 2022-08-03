package main

import (
	"fmt"
	"strconv"
)

/*
Var declaration:
var foo int = *num*
foo := *num*

Cant redeclare variables, but can re-assign

Visibility:
lower case first letter - package scope
upper case first letter - export
no private scope
*/

func StrConv() {
	var i int = 34
	fmt.Printf("%v, %T\n", i, i)

	var j float32
	j = float32(i)
	fmt.Printf("%v, %T\n", j, j)

	var k string
	k = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", k, k)
}
