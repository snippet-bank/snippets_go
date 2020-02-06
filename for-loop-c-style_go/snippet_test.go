package snippets

import "testing"

func TestSumToDigit(t *testing.T) {
	sum := sumToDigit(3)
	if sum != 6 {
		t.Errorf("Got %d, expected 6", sum)
	}
}
