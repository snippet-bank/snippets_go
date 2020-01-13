package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "    some stuff      "
	trimmedStr := strings.TrimSpace(str)
	fmt.Println(trimmedStr)
}

// result:
// "some stuff"
