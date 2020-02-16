package snippets

import (
	"bufio"
	"log"
	"os"
)

// readByLines reads a file into memory line by line.
func readByLines(path string, callback func(str string)) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		callback(line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
