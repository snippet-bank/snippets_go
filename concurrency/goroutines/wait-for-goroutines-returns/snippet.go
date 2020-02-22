package snippets

import (
	"time"
)

func f(ch chan string) {
	time.Sleep(time.Second * 1)
	ch <- "msg"
}

func getGoroutinesResults() []string {
	ch := make(chan string)

	// creates three goroutines
	for i := 0; i < 3; i++ {
		go f(ch)
	}

	// waits for three messages from the channel
	var results []string
	for i := 0; i < 3; i++ {
		results = append(results, <-ch)
	}
	return results
}
