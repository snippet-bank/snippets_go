package snippets

import (
	"reflect"
	"testing"
)

func TestSplitOnPipes(t *testing.T) {
	str := "apples|oranges|pears"
	split := splitOnPipes(str)
	expected := []string{"apples", "oranges", "pears"}
	if !reflect.DeepEqual(split, expected) {
		t.Errorf("Got %v, expected %v", split, expected)
	}
}
