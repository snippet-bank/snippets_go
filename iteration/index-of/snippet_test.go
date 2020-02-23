package snippets

import "testing"

var tests = []struct {
	values []string
	target string
	want   int
}{
	{[]string{"a", "b", "c"}, "a", 0},
	{[]string{"a", "b", "c"}, "b", 1},
	{[]string{"a", "b", "c"}, "c", 2},
	{[]string{"a", "b", "c"}, "d", -1},
	{[]string{}, "a", -1},
}

func TestIndexOf(t *testing.T) {
	for _, tt := range tests {
		got := IndexOf(tt.values, tt.target)
		if got != tt.want {
			t.Errorf("For %v and target %s, got %v, want %v", tt.values, tt.target, got, tt.want)
		}
	}
}
