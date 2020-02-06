package snippets

import (
	"regexp"
)

// filterNonAlpha removes all non-alphabetic characters.
func filterNonAlpha(str string) string {
	var re = regexp.MustCompile("[^a-zA-Z]+")
	return re.ReplaceAllString(str, "")
}
