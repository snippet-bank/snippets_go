package snippets

import "testing"

func TestTrim(t *testing.T) {
	str := "    some stuff      "
	trimmed := trim(str)
	expected := "some stuff"

	if trimmed != expected {
		t.Errorf("Got %s, expected %s", trimmed, expected)
	}
}
