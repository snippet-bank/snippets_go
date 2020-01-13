package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	readWholeFile()
}

func readWholeFile() {
	data, err := ioutil.ReadFile("shakespeare_sonnet30.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(data))
}
