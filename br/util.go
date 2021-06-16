package brlocalflavor

import "unicode"

func stringDigitsToIntSlice(value string) []int {
	var digits []int
	for _, char := range value {
		if unicode.IsDigit(char) {
			digits = append(digits, int(char-'0'))
		}
	}

	return digits
}

func calculateDocumentDigit(multiplierCounter int, digits []int) int {
	var total int

	for _, digit := range digits {
		total += digit * multiplierCounter
		multiplierCounter--

		if multiplierCounter < 2 {
			multiplierCounter = 9
		}
	}

	mod := total % 11
	if mod < 2 {
		return 0
	}

	return 11 - mod
}
