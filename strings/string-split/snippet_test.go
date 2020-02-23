package snippets

import (
	"reflect"
	"testing"
)

func TestSplitOnPipes(t *testing.T) {
	str := "apples|oranges|pears"
	got := splitOnPipes(str)
	want := []string{"apples", "oranges", "pears"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v, want %v", got, want)
	}
}
