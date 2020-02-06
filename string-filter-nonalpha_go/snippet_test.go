package snippets

import "testing"

func TestFilterNonAlpha(t *testing.T) {
	str := "why - you! that's it, we are *through*"
	expected := "whyyouthatsitwearethrough"
	filtered := filterNonAlpha(str)
	if filtered != expected {
		t.Errorf("Got %s, expected %s", filtered, expected)
	}
}
