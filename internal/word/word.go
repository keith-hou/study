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

func CamelCaseToUnderline(s string) string {
	var output [] rune
	for i, v := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(v))
		}
		if unicode.IsUpper(v) {
			output = append(output,'_')
		}
		output = append(output, unicode.ToLower(v))
	}

	return string(output)
}