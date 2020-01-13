package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// Load the sample file.
	var data []byte
	var _ error
	data, _ = ioutil.ReadFile("letters.txt")

	// Note the return type of "ReadFile()"
	// is a byte array.

	// Important point:
	// All files are just byte arrays.

	// Now print the bytes we get, in hexadecimal
	for i := 0; i < len(data); i++ {
		fmt.Printf("%x \n", data[i])
	}

	// Result:
	// 41
	// 42
	// 43
	// 44
	// 45

	// Very Important Point:
	// Each byte is just a *number*.
	// from 00000000 to 11111111, in binary
	// from 00 to FF, in hexadecimal
	// from 0 to 255, in decimal

	// What each number means is completely dependent
	// on context.

	// So what do the numbers in *this* particular file mean?

	// Well we *know* (because we wrote this program) that
	// these numbers represent Unicode.

	// Each byte is a number mapping to a "code point"
	// in Unicode (roughly, a "letter")
	// see http://www.unicode.org/charts/PDF/U0000.pdf

	// hexadecimal 41 is "A"
	// hexadecimal 42 is "B"
	// etc.

	// So we choose to *interpret* those numbers
	// as a Unicode string:
	var str string
	str = string(data)
	fmt.Println(str)

	// Result:
	// ABCDE
}
