package string

import "unicode"

func CapitalizeStringLetterByIndex(input string, index int) string {
	nameRune := []rune(input)
	nameRune[index] = unicode.ToUpper(nameRune[index])

	return string(nameRune)
}
