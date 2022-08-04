package main

import (
	"fmt"
	"unicode"
)

// Cipher interface encrypts and decrypts a string
type cipher interface {
	enc(string) string
	dec(string) string
}

// holds cipher key for enc and dec
type cipherKey []int

// cipherAlg encodes a letter based on shiftFunc
// here l = letters
func (c cipherKey) cipherAlg(l string, shift func(int, int) int) string {
	shiftedTxt := ""
	for _, letter := range l {
		if !unicode.IsLetter(l) {
			continue
		}
		// shift distance is equivalent to remainder of
		// shifttxt length by cipherkey length
		shiftDist := c[len(shiftedTxt)%len(c)]
		s := shift(int(unicode.ToLower(letter)), shiftDist)
		switch {
		// case for if s < 97
		case s < 'a':
			s += 'z' - 'a' + 1
		case 'z' < s:
			s -= 'z' - 'a' + 1
		}
		shiftedTxt += string(s)
	}
	return shiftedTxt
}

// cipherAlg("hello", func(a, b, int) int {return a+b})

func main() {
	huh := "hello"
	test := cipherKey.cipherAlg(huh, func(a, b int) int { return a + b })
	fmt.Println(test)
}
