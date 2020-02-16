package snippets

import (
	"encoding/csv"
	"os"
)

// ReadCSV reads a CSV file as strings.
func ReadCSV(path string, skipHeaderRow bool) ([][]string, error) {
	// Open CSV file
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.LazyQuotes = true
	r.TrimLeadingSpace = true
	lines, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	if skipHeaderRow {
		return lines[1:], nil
	}
	return lines, nil
}
