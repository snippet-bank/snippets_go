package snippets

import (
	"os"
	"testing"
)

func TestGetEnvVariable(t *testing.T) {
	os.Setenv("FOO", "BAR")

	v := getEnvVariable("FOO")
	if v != "BAR" {
		t.Errorf("Got %s, expected BAR", v)
	}
}
