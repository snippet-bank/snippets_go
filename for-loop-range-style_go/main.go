package main

import "fmt"

func main() {
	fruits := []string{"apple", "orange", "lemon"}

	for i, fruit := range fruits {
		fmt.Printf("array index %d: %s \n", i, fruit)
	}
}
