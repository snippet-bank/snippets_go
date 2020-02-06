package snippets

import (
	"reflect"
	"testing"
)

func TestSplitOnWhitespace(t *testing.T) {
	str := "   some     spaces    and\tsome \t  tabs too   "
	expected := []string{"some", "spaces", "and", "some", "tabs", "too"}
	split := splitOnWhitespace(str)

	if !reflect.DeepEqual(split, expected) {
		t.Errorf("Got %v, expected %v", split, expected)
	}
}
