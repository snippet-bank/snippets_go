package snippets

import "testing"

func TestGetFirstUnicodeChar(t *testing.T) {
	str := "日本語"
	expected := "日"
	firstChar := getFirstUnicodeChar(str)
	if firstChar != expected {
		t.Errorf("Got %s but expected %s", firstChar, expected)
	}
}
