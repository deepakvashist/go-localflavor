package brlocalflavor

import (
	"regexp"
)

// ValidateCPF validates CPF document.
func ValidateCPF(document string) error {
	pattern := regexp.MustCompile(`^\d{3}\.?\d{3}\.?\d{3}-?\d{2}$`)
	if !pattern.MatchString(document) {
		return ErrInvalidDocumentPattern
	}

	digits := stringDigitsToIntSlice(document)

	firstDigit := calculateDocumentDigit(10, digits[:9])
	secondDigit := calculateDocumentDigit(11, digits[:10])

	if digits[9] != firstDigit || digits[10] != secondDigit {
		return ErrInvalidDocumentNumber
	}

	return nil
}
