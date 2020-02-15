package snippets

import "testing"

var tests = []struct {
	values   []string
	target   string
	expected int
}{
	{[]string{"a", "b", "c"}, "a", 0},
	{[]string{"a", "b", "c"}, "b", 1},
	{[]string{"a", "b", "c"}, "c", 2},
	{[]string{"a", "b", "c"}, "d", -1},
	{[]string{}, "a", -1},
}

func TestIndexOf(t *testing.T) {
	for _, tt := range tests {
		result := IndexOf(tt.values, tt.target)
		if result != tt.expected {
			t.Errorf("For %v and target %s, got %v, expected %v", tt.values, tt.target, result, tt.expected)
		}
	}
}
