package snippets

// IndexOf returns the index of the first instance of the target element, or -1 if not found.
func IndexOf(vals []string, target string) int {
	for i, v := range vals {
		if v == target {
			return i
		}
	}
	return -1
}
