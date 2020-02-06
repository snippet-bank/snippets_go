package snippets

import (
	"strings"
	"testing"
)

func TestReadByLines(t *testing.T) {
	var sb strings.Builder
	callback := func(s string) {
		sb.WriteString(s)
	}

	readByLines("shakespeare_sonnet30.txt", callback)
	fileContents := sb.String()

	if !strings.HasPrefix(fileContents, "When to the sessions of sweet silent thought") {
		t.Errorf("File did not contain correct first line")
	}

	if !strings.HasSuffix(fileContents, "All losses are restored and sorrows end.") {
		t.Errorf("File did not contain correct last line")
	}
}
