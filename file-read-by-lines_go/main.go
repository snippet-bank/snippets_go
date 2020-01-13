package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	readByLines()
}

// Minimal, *unsafe* version without error handling or closing file handles.
// This is just to highlight the core functionality in the fewest lines.
// See also "safe.go" for the version with error handling.
func readByLines() {
	file, _ := os.Open("shakespeare_sonnet30.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}
