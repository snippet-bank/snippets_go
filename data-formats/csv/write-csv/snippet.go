package snippets

import (
	"encoding/csv"
	"log"
	"os"
)

// SaveCSV writes a csv file to the supplied path.
func SaveCSV(path string, data [][]string) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal("Cannot create output file", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		err := writer.Write(value)
		if err != nil {
			log.Fatal("Cannot write to file", err)
		}
	}
}
