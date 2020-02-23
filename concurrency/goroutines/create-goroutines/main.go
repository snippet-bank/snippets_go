package main

import (
	"fmt"
	"time"
)

func f(num int) {
	time.Sleep(time.Second)
	fmt.Printf("goroutine number %d exits\n", num)
}

func main() {
	for i := 0; i < 3; i++ {
		go f(i) // it doesn't wait for the goroutine to finish execution
	}

	// if there was no sleeping here, you wouldn't see anything printed in the terminal
	// as the goroutines would be killed when main exits
	time.Sleep(3 * time.Second)
	fmt.Println("main exits")
}
