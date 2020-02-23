package snippets

import "testing"

func TestSumToDigit(t *testing.T) {
	got := sumToDigit(3)
	want := 6
	if got != want {
		t.Errorf("Got %d, expected %d", got, want)
	}
}
