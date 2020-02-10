package snippets

import (
	"bufio"
	"os"
)

// readByLines reads a file into memory line by line.
func readByLines(path string, callback func(str string)) {
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		callback(line)
	}
}

// This is a minimal, *unsafe* version without error handling or closing file handles.
// This is just to highlight the core functionality in the fewest lines.
// See also "snippet_safe.go" for the version with error handling.
