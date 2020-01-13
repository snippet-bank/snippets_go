package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "apples|oranges|pears"

	var parts []string
	parts = strings.Split(str, "|")

	fmt.Println(parts)
}

// result:
// [apples oranges pears]
