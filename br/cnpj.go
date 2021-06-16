package brlocalflavor

import (
	"regexp"
)

// ValidateCNPJ validates CNPJ document.
func ValidateCNPJ(document string) error {
	pattern := regexp.MustCompile(`^\d{2}\.?\d{3}\.?\d{3}\/?(:?\d{3}[1-9]|\d{2}[1-9]\d|\d[1-9]\d{2}|[1-9]\d{3})-?\d{2}$`)
	if !pattern.MatchString(document) {
		return ErrInvalidDocumentPattern
	}

	digits := stringDigitsToIntSlice(document)

	firstDigit := calculateDocumentDigit(5, digits[:12])
	secondDigit := calculateDocumentDigit(6, digits[:13])

	if digits[12] != firstDigit || digits[13] != secondDigit {
		return ErrInvalidDocumentNumber
	}

	return nil
}
