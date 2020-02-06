package snippets

// sumToDigit sums all the numbers from 0 to n.
// For example, given n=3, 0+1+2+3 is 6.
func sumToDigit(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	return sum
}
