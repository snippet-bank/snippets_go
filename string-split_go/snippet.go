package snippets

import "strings"

func splitOnPipes(str string) []string {
	return strings.Split(str, "|")
}
