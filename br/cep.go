package brlocalflavor

import "regexp"

// ValidateCEP validates CEP.
func ValidateCEP(cep string) error {
	pattern := regexp.MustCompile(`^\d{5}-?\d{3}$`)
	if !pattern.MatchString(cep) {
		return ErrInvalidCEPPattern
	}

	return nil
}
