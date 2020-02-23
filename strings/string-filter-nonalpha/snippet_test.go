package snippets

import "testing"

func TestFilterNonAlpha(t *testing.T) {
	str := "why - you! that's it, we are *through*"
	want := "whyyouthatsitwearethrough"
	got := filterNonAlpha(str)
	if got != want {
		t.Errorf("Got %s, want %s", got, want)
	}
}
