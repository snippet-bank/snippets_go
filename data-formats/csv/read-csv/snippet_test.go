package snippets

import (
	"testing"
)

type CSVRow struct {
	Letter     string
	Frequency  string
	Percentage string
}

var tests = []struct {
	path              string
	skipFirstLine     bool
	expectedLineCount int
	expectedFirstRow  CSVRow
	expectedSecondRow CSVRow
}{
	{
		"fixtures/sample.csv",
		true,
		26,
		CSVRow{"A", "24373121", "8.1"},
		CSVRow{"B", "4762938", "1.6"},
	},
	{
		"fixtures/sample.csv",
		false,
		27,
		CSVRow{"Letter", "Frequency", "Percentage"},
		CSVRow{"A", "24373121", "8.1"},
	},
}

func TestReadCSV(t *testing.T) {
	for _, tt := range tests {
		lines, err := ReadCSV(tt.path, tt.skipFirstLine)
		if err != nil {
			t.Error("Error processing CSV", err)
		}

		if len(lines) != tt.expectedLineCount {
			t.Errorf("Not enough data read. Got %d lines, expected %d lines.", len(lines), tt.expectedLineCount)
		}

		got := ToCSVRow(lines[0])
		if got != tt.expectedFirstRow {
			t.Errorf("Got %v, expected %v", lines[0], tt.expectedFirstRow)
		}

		got = ToCSVRow(lines[1])
		if got != tt.expectedSecondRow {
			t.Errorf("Got %v, expected %v", lines[0], tt.expectedSecondRow)
		}

	}
}

func ToCSVRow(rawLine []string) CSVRow {
	return CSVRow{
		Letter:     rawLine[0],
		Frequency:  rawLine[1],
		Percentage: rawLine[2],
	}
}
