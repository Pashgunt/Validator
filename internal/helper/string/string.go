package string

import (
	"math"
	"unicode"
)

const (
	WeightValueLowerAlphabet   = 26
	WeightValueUpperAlphabet   = 26
	WeightValueDigitAlphabet   = 10
	WeightValueSpecialAlphabet = 32
)

func CapitalizeStringLetterByIndex(input string, index int) string {
	nameRune := []rune(input)
	nameRune[index] = unicode.ToUpper(nameRune[index])

	return string(nameRune)
}

func getAlphabetSize(data string) int {
	hasLower := false
	hasUpper := false
	hasDigit := false
	hasSpecial := false

	for _, ch := range data {
		switch {
		case unicode.IsLower(ch):
			hasLower = true
		case unicode.IsUpper(ch):
			hasUpper = true
		case unicode.IsDigit(ch):
			hasDigit = true
		default:
			hasSpecial = true
		}
	}

	size := 0

	if hasLower {
		size += WeightValueLowerAlphabet
	}

	if hasUpper {
		size += WeightValueUpperAlphabet
	}

	if hasDigit {
		size += WeightValueDigitAlphabet
	}

	if hasSpecial {
		size += WeightValueSpecialAlphabet
	}

	return size
}

func CalculateEntropy(data string) float64 {
	alphabetSize := getAlphabetSize(data)

	if alphabetSize == 0 || len(data) == 0 {
		return 0
	}

	return float64(len(data)) * math.Log2(float64(alphabetSize))
}
