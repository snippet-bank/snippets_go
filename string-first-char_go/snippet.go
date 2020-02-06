package snippets

func getFirstUnicodeChar(str string) string {
	runes := []rune(str)
	return string(runes[0])
}
