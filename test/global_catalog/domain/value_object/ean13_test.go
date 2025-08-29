package value_object_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"saas-mt-pim-service/src/product/global_catalog/domain/value_object"
)

func TestNewEAN13(t *testing.T) {
	t.Run("should create valid EAN13", func(t *testing.T) {
		// Arrange
		validEAN := "7790001234568"

		// Act
		ean, err := value_object.NewEAN13(validEAN)

		// Assert
		require.NoError(t, err)
		assert.NotNil(t, ean)
		assert.Equal(t, validEAN, ean.Value())
	})

	t.Run("should create valid EAN13 with different valid codes", func(t *testing.T) {
		validEANs := []string{
			"7790001234568", // Argentina
			"1234567890128", // Genérico válido
			"0123456789012", // Con cero inicial
		}

		for _, validEAN := range validEANs {
			t.Run("EAN: "+validEAN, func(t *testing.T) {
				// Act
				ean, err := value_object.NewEAN13(validEAN)

				// Assert
				require.NoError(t, err)
				assert.Equal(t, validEAN, ean.Value())
			})
		}
	})

	t.Run("should fail with empty EAN", func(t *testing.T) {
		// Act
		ean, err := value_object.NewEAN13("")

		// Assert
		assert.Error(t, err)
		assert.Nil(t, ean)
		assert.Contains(t, err.Error(), "EAN-13 es obligatorio")
	})

	t.Run("should fail with invalid length", func(t *testing.T) {
		invalidEANs := []string{
			"123456789012",   // 12 dígitos
			"12345678901234", // 14 dígitos
			"123",            // muy corto
		}

		for _, invalidEAN := range invalidEANs {
			t.Run("Invalid EAN: "+invalidEAN, func(t *testing.T) {
				// Act
				ean, err := value_object.NewEAN13(invalidEAN)

				// Assert
				assert.Error(t, err)
				assert.Nil(t, ean)
				assert.Contains(t, err.Error(), "13 dígitos")
			})
		}
	})

	t.Run("should fail with non-numeric characters", func(t *testing.T) {
		invalidEANs := []string{
			"779000123456A", // letra al final
			"A790001234568", // letra al inicio
			"7790001-34568", // guión
			"7790001 34568", // espacio
		}

		for _, invalidEAN := range invalidEANs {
			t.Run("Non-numeric EAN: "+invalidEAN, func(t *testing.T) {
				// Act
				ean, err := value_object.NewEAN13(invalidEAN)

				// Assert
				assert.Error(t, err)
				assert.Nil(t, ean)
				assert.Contains(t, err.Error(), "13 dígitos")
			})
		}
	})

	t.Run("should fail with invalid checksum", func(t *testing.T) {
		// Arrange - EAN con checksum incorrecto
		invalidEAN := "7790001234567" // último dígito incorrecto

		// Act
		ean, err := value_object.NewEAN13(invalidEAN)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, ean)
		assert.Contains(t, err.Error(), "checksum inválido")
	})
}

func TestEAN13_IsArgentineProduct(t *testing.T) {
	t.Run("should return true for Argentine EAN", func(t *testing.T) {
		// Arrange
		argentineEANs := []string{
			"7790001234568", // 779
			"7800001234564", // 780 válido
			"7990001234562", // 799 válido
		}

		for _, eanCode := range argentineEANs {
			t.Run("Argentine EAN: "+eanCode, func(t *testing.T) {
				// Act
				ean, _ := value_object.NewEAN13(eanCode)

				// Assert
				assert.True(t, ean.IsArgentineProduct())
			})
		}
	})

	t.Run("should return false for non-Argentine EAN", func(t *testing.T) {
		// Arrange
		nonArgentineEANs := []string{
			"1234567890128", // 123 válido
			"0123456789012", // 012
			"7780001234569", // 778 válido (no es rango argentino)
		}

		for _, eanCode := range nonArgentineEANs {
			t.Run("Non-Argentine EAN: "+eanCode, func(t *testing.T) {
				// Act
				ean, _ := value_object.NewEAN13(eanCode)

				// Assert
				assert.False(t, ean.IsArgentineProduct())
			})
		}
	})
}

func TestEAN13_CountryCode(t *testing.T) {
	t.Run("should return correct country code", func(t *testing.T) {
		testCases := []struct {
			ean         string
			countryCode string
		}{
			{"7790001234568", "779"}, // Argentina
			{"1234567890128", "123"}, // Genérico válido
			{"0123456789012", "012"}, // Con cero
		}

		for _, tc := range testCases {
			t.Run("EAN: "+tc.ean, func(t *testing.T) {
				// Act
				ean, _ := value_object.NewEAN13(tc.ean)

				// Assert
				assert.Equal(t, tc.countryCode, ean.CountryCode())
			})
		}
	})
}

func TestEAN13_Formatted(t *testing.T) {
	t.Run("should return formatted EAN", func(t *testing.T) {
		// Arrange
		ean, _ := value_object.NewEAN13("7790001234568")

		// Act
		formatted := ean.Formatted()

		// Assert
		assert.Equal(t, "779-0-00123-456-8", formatted)
	})
}

func TestEAN13_Equals(t *testing.T) {
	t.Run("should return true for equal EANs", func(t *testing.T) {
		// Arrange
		ean1, _ := value_object.NewEAN13("7790001234568")
		ean2, _ := value_object.NewEAN13("7790001234568")

		// Act & Assert
		assert.True(t, ean1.Equals(ean2))
		assert.True(t, ean2.Equals(ean1))
	})

	t.Run("should return false for different EANs", func(t *testing.T) {
		// Arrange
		ean1, _ := value_object.NewEAN13("7790001234568")
		ean2, _ := value_object.NewEAN13("1234567890128")

		// Act & Assert
		assert.False(t, ean1.Equals(ean2))
		assert.False(t, ean2.Equals(ean1))
	})

	t.Run("should return false when comparing with nil", func(t *testing.T) {
		// Arrange
		ean, _ := value_object.NewEAN13("7790001234568")

		// Act & Assert
		assert.False(t, ean.Equals(nil))
	})
}

func TestEAN13_String(t *testing.T) {
	t.Run("should return EAN value as string", func(t *testing.T) {
		// Arrange
		eanValue := "7790001234568"
		ean, _ := value_object.NewEAN13(eanValue)

		// Act & Assert
		assert.Equal(t, eanValue, ean.String())
	})
}
