package brlocalflavor

import "errors"

var (
	// ErrInvalidDocumentPattern is returned when the document pattern is invalid.
	ErrInvalidDocumentPattern = errors.New("invalid document patter")
	// ErrInvalidDocumentPattern is returned when the document number is invalid.
	ErrInvalidDocumentNumber = errors.New("invalid document number")
)