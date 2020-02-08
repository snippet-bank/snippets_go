package snippets

import (
	"strings"
	"testing"
)

func TestReadWholeFile(t *testing.T) {
	fileContents := readWholeFile("shakespeare_sonnet30.txt")

	if !strings.HasPrefix(fileContents, "When to the sessions of sweet silent thought") {
		t.Errorf("File did not contain correct first line")
	}

	if !strings.HasSuffix(fileContents, "All losses are restored and sorrows end.") {
		t.Errorf("File did not contain correct last line")
	}
}
