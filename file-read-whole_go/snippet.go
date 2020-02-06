package snippets

import (
	"io/ioutil"
	"log"
)

// readWholeFile reads an entire file into memory at once.
func readWholeFile(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}
