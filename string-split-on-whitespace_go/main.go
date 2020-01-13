package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "   some     spaces    and\tsome \t  tabs too   "

	var fields []string
	fields = strings.Fields(str)

	fmt.Println(fields)
}

// result:
// [some spaces and some tabs too]
