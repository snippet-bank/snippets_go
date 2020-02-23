package snippets

import "testing"

func TestTrim(t *testing.T) {
	arg := "    some stuff      "
	got := trim(arg)
	want := "some stuff"

	if got != want {
		t.Errorf("Got %s, want %s", got, want)
	}
}
