package brlocalflavor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCPF(t *testing.T) {
	t.Run("Valid CPF", func(t *testing.T) {
		document := "245.815.230-95"

		err := ValidateCPF(document)

		assert.Nil(t, err)
	})

	t.Run("Invalid CPF pattern", func(t *testing.T) {
		document := "ABC.815.230-95"

		err := ValidateCPF(document)

		assert.Equal(t, ErrInvalidDocumentPattern, err)
	})

	t.Run("Invalid CPF number", func(t *testing.T) {
		document := "245.915.230-95"

		err := ValidateCPF(document)

		assert.Equal(t, ErrInvalidDocumentNumber, err)
	})
}
