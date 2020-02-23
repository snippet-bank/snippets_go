package snippets

import "testing"

func TestGetFirstUnicodeChar(t *testing.T) {
	arg := "日本語"
	want := "日"
	got := getFirstUnicodeChar(arg)
	if got != want {
		t.Errorf("Got %s but want %s", got, want)
	}
}
