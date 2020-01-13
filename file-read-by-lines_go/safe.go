package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readByLinesSafe() {
	file, err := os.Open("shakespeare_sonnet30.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
