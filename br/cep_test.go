package brlocalflavor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCEP(t *testing.T) {
	t.Run("Valid CEP", func(t *testing.T) {
		cep := "04538-133"

		err := ValidateCEP(cep)

		assert.Nil(t, err)
	})

	t.Run("Invalid CEP", func(t *testing.T) {
		cep := "38-133"

		err := ValidateCEP(cep)

		assert.Equal(t, ErrInvalidCEPPattern, err)
	})
}
