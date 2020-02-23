package snippets

import "testing"

func TestSum(t *testing.T) {
	nums := []float64{5, 7, 12}
	var want float64 = 24
	got := sum(nums)
	if got != want {
		t.Errorf("Got %f, want %f", got, want)
	}
}
