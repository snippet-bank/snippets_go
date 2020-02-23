package snippets

import (
	"reflect"
	"testing"
)

func TestSplitOnWhitespace(t *testing.T) {
	str := "   some     spaces    and\tsome \t  tabs too   "
	want := []string{"some", "spaces", "and", "some", "tabs", "too"}
	got := splitOnWhitespace(str)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v, want %v", got, want)
	}
}
