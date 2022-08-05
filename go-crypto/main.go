package main

import (
	"fmt"
	"unicode"
)

// Cipher interface encrypts and decrypts a string
type Cipher interface {
	Encryption(string) string
	Decryption(string) string
}

// holds cipher key for enc and dec
type cipherKey []int

// cipherAlg encodes a letter based on shiftFunc
// here l = letters
func (c cipherKey) cipherAlg(l string, shift func(int, int) int) string {
	shiftedTxt := ""
	for _, letter := range l {
		if !unicode.IsLetter(letter) {
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
func (c cipherKey) Encryption(plainText string) string {
	return c.cipherAlg(plainText, func(a, b int) int { return a + b })
}

func (c cipherKey) Decryption(plainText string) string {
	return c.cipherAlg(plainText, func(a, b int) int { return a - b })
}

func NewCaesar(key int) Cipher {
	return NewShift(key)
}

// NewShift creates a new Shift cipher.
func NewShift(shift int) Cipher {
	if shift < -25 || 25 < shift || shift == 0 {
		return nil
	}
	c := cipherKey([]int{shift})
	return &c
}

func main() {
	c := NewCaesar(2)
	test := c.Encryption("sarvag")
	fmt.Println(test)

	huh := c.Decryption("uctxci")
	fmt.Println(huh)
}
