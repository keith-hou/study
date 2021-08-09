package word

import "strings"
import "unicode"

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

func UnderlineToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.ReplaceAll(s, " ", "")
}

func UnderlineToLowerCamelCase(s string) string {
	s = UnderlineToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[:1]
}