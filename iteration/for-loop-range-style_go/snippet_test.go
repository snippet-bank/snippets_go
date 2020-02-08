package snippets

import "testing"

func TestSum(t *testing.T) {
	nums := []float64{5, 7, 12}
	var expected float64 = 24
	total := sum(nums)
	if total != expected {
		t.Errorf("Got %f, expected %f", total, expected)
	}
}
