package brlocalflavor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCNPJ(t *testing.T) {
	t.Run("Valid CNPJ", func(t *testing.T) {
		document := "14.543.367/0001-66"

		err := ValidateCNPJ(document)

		assert.Nil(t, err)
	})

	t.Run("Invalid CNPJ pattern", func(t *testing.T) {
		document := "14.543.ABC/0001-66"

		err := ValidateCNPJ(document)

		assert.Equal(t, ErrInvalidDocumentPattern, err)
	})

	t.Run("Invalid CNPJ number", func(t *testing.T) {
		document := "14.540.367/0001-66"

		err := ValidateCNPJ(document)

		assert.Equal(t, ErrInvalidDocumentNumber, err)
	})
}
