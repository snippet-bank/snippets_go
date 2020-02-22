package main

import (
	"fmt"
	"sync"
	"time"
)

// global declaration of a WaitGroup
var wg sync.WaitGroup

func f(num int) {
	// decrements the value of wg by one, thus informing that this goroutine is done executing
	// this must happen at the end of the execution, no matter if it ends naturally or is panicking
	// defer is the best way to achieve that
	defer wg.Done()

	time.Sleep(time.Second)
	fmt.Printf("goroutine number %d exits\n", num)
}

func main() {
	n := 10

	// adds n to the initial value of wg (that is 0)
	wg.Add(n)

	for i := 0; i < n; i++ {
		go f(i)
	}

	// waits for the value to drop to 0
	wg.Wait()
}
