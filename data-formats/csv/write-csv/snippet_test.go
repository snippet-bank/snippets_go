package snippets

import (
	"bufio"
	"os"
	"testing"
)

var data = [][]string{{"Letter", "Frequency", "Percentage"},
	{"A", "24373121", "8.1"},
	{"B", "4762938", "1.6"},
	{"C", "8982417", "3.0"},
	{"D", "10805580", "3.6"},
	{"E", "37907119", "12.6"},
	{"F", "7486889", "2.5"},
	{"G", "5143059", "1.7"},
	{"H", "18058207", "6.0"},
	{"I", "21820970", "7.3"},
	{"J", "474021", "0.2"},
	{"K", "1720909", "0.6"},
	{"L", "11730498", "3.9"},
	{"M", "7391366", "2.5"},
	{"N", "21402466", "7.1"},
	{"O", "23215532", "7.7"},
	{"P", "5719422", "1.9"},
	{"Q", "297237", "0.1"},
	{"R", "17897352", "5.9"},
	{"S", "19059775", "6.3"},
	{"T", "28691274", "9.5"},
	{"U", "8022379", "2.7"},
	{"V", "2835696", "0.9"},
	{"W", "6505294", "2.2"},
	{"X", "562732", "0.2"},
	{"Y", "5910495", "2.0"},
	{"Z", "93172", "0.0"}}

func TestSaveCSV(t *testing.T) {
	path := "test-output/letter-frequency.csv"
	SaveCSV(path, data)

	file, err := os.Open(path)
	if err != nil {
		t.Errorf("Got error reading output file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if lineNum == 0 {
			expected := "Letter,Frequency,Percentage"
			if line != expected {
				t.Errorf("First line, got %s, expected %s", line, expected)
			}
		}
		if lineNum == 26 {
			expected := "Z,93172,0.0"
			if line != expected {
				t.Errorf("Last line, got %s, expected %s", line, expected)
			}
		}
		lineNum++
	}
	if err := scanner.Err(); err != nil {
		t.Errorf("Got error reading output file: %v", err)
	}
}
