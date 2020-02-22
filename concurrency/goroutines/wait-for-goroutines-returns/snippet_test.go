package snippets

import "testing"

func TestGetGoroutinesResults(t *testing.T) {
	results := getGoroutinesResults()

	if len(results) != 3 {
		t.Errorf("Got array of length %d, expected 3", len(results))
	}

	for i := 0; i < 3; i++ {
		if results[i] != "msg" {
			t.Errorf("Got msg \"%s\" at index %d, expected \"msg\"", results[i], i)
		}
	}
}
