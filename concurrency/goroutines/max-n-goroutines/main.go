package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	semaphoreSize = 4
	wg            sync.WaitGroup
)

func f(sem chan bool, num int) {
	defer wg.Done()
	defer func() {
		// "release semaphore" - release one place in the channel buffer, thus allowing other goroutines to start
		<-sem
	}()

	time.Sleep(time.Second)
	fmt.Printf("goroutine number %d exits\n", num)
}

func main() {
	n := 20
	wg.Add(n)

	sem := make(chan bool, semaphoreSize)

	for i := 0; i < n; i++ {
		// "acquire semaphore" - take one place in the channel buffer
		// if there are no places left, block until there is
		sem <- true
		go f(sem, i)
	}

	wg.Wait()
}
