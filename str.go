package utils

import (
	"bytes"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
	"unicode"
	"unicode/utf16"
)

func ShortStr4Web(str string, max int) string {
	str = strings.ReplaceAll(str, "\n", "\\n")
	str = strings.ReplaceAll(str, "\r", "\\r")
	str = strings.ReplaceAll(str, "\t", "\\t")

	newStr := ShortStr(str, max)
	if len(newStr) != len(str) {
		newStr += "..."
	}
	return newStr
}

func ShortStr(str string, max int) string {
	if max < 0 {
		return str
	}
	if len(str) > max {
		return str[:max]
	}
	return str
}

func Utf16KLen(str string) int {
	return len(utf16.Encode([]rune(str)))
}

func CamelToSnake(s string) string {
	var b bytes.Buffer
	for i, c := range s {
		if unicode.IsUpper(c) {
			if i > 0 {
				b.WriteString("_")
			}
			b.WriteRune(unicode.ToLower(c))
		} else {
			b.WriteRune(c)
		}
	}
	return b.String()
}

func SnakeToCamel(s string) string {
	s = strings.ReplaceAll(s, "_", " ")
	s = cases.Title(language.English).String(s)
	return strings.ReplaceAll(s, " ", "")
}
