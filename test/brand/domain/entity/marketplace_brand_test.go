package entity

import (
	"strings"
	"testing"

	"saas-mt-pim-service/src/brand/domain/entity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// --- helpers ---

func newTestMarketplacebrand(t *testing.T) *entity.Marketplacebrand {
	t.Helper()
	brand, err := entity.NewMarketplacebrand("TestBrand")
	require.NoError(t, err)
	return brand
}

// --- tests ---

func TestMarketplacebrand_WithValidHexColors_ReturnsSuccess(t *testing.T) {
	// Arrange
	brand := newTestMarketplacebrand(t)
	params := entity.VisualIdentityParams{
		BackgroundColor: "#FF5733",
		TextColor:       "#FFFFFF",
		Typography:      "Oswald",
	}

	// Act
	err := brand.SetVisualIdentity(params)

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "#FF5733", brand.BackgroundColor)
	assert.Equal(t, "#FFFFFF", brand.TextColor)
	assert.Equal(t, "Oswald", brand.Typography)
}

func TestMarketplacebrand_WithLowercaseHex_ReturnsSuccess(t *testing.T) {
	// Arrange
	brand := newTestMarketplacebrand(t)
	params := entity.VisualIdentityParams{
		BackgroundColor: "#ff5733",
		TextColor:       "#ffffff",
		Typography:      "Roboto",
	}

	// Act
	err := brand.SetVisualIdentity(params)

	// Assert
	require.NoError(t, err, "hex en minúsculas debe ser válido")
	assert.Equal(t, "#ff5733", brand.BackgroundColor)
}

func TestMarketplacebrand_WithInvalidHex_ReturnsValidationError(t *testing.T) {
	invalidHexValues := []struct {
		name  string
		value string
	}{
		{"solo hash sin dígitos", "#ZZZ"},
		{"nombre de color en inglés", "red"},
		{"hex corto sin hash", "#FFF"},
		{"hex sin hash", "FFFFFF"},
		{"hex con 8 dígitos", "#FFFFFFFF"},
		{"hex con 5 dígitos", "#FFFFF"},
		{"hex con caracteres inválidos", "#GG1122"},
	}

	for _, tc := range invalidHexValues {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			brand := newTestMarketplacebrand(t)
			params := entity.VisualIdentityParams{BackgroundColor: tc.value}

			// Act
			err := brand.SetVisualIdentity(params)

			// Assert
			assert.Error(t, err, "valor %q debe ser inválido", tc.value)
			assert.Contains(t, err.Error(), "background_color inválido")
		})
	}
}

func TestMarketplacebrand_WithInvalidTextColor_ReturnsValidationError(t *testing.T) {
	// Arrange
	brand := newTestMarketplacebrand(t)
	params := entity.VisualIdentityParams{
		BackgroundColor: "#FF5733",
		TextColor:       "red", // inválido
	}

	// Act
	err := brand.SetVisualIdentity(params)

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "text_color inválido")
}

func TestMarketplacebrand_WithEmptyVisualFields_IsValid(t *testing.T) {
	// Arrange
	brand := newTestMarketplacebrand(t)
	params := entity.VisualIdentityParams{
		BackgroundColor: "",
		TextColor:       "",
		Typography:      "",
	}

	// Act
	err := brand.SetVisualIdentity(params)

	// Assert
	require.NoError(t, err, "campos vacíos son válidos — significan fallback al design system")
	assert.Equal(t, "", brand.BackgroundColor)
	assert.Equal(t, "", brand.TextColor)
	assert.Equal(t, "", brand.Typography)
}

func TestMarketplacebrand_WithTypographyTooLong_ReturnsValidationError(t *testing.T) {
	// Arrange
	brand := newTestMarketplacebrand(t)
	longTypography := strings.Repeat("A", 101) // 101 caracteres — excede límite
	params := entity.VisualIdentityParams{
		Typography: longTypography,
	}

	// Act
	err := brand.SetVisualIdentity(params)

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "typography inválido")
}

func TestMarketplacebrand_WithTypographyExactlyAtLimit_IsValid(t *testing.T) {
	// Arrange
	brand := newTestMarketplacebrand(t)
	exactTypography := strings.Repeat("A", 100) // exactamente 100 chars — límite inclusivo
	params := entity.VisualIdentityParams{
		Typography: exactTypography,
	}

	// Act
	err := brand.SetVisualIdentity(params)

	// Assert
	require.NoError(t, err, "typography de exactamente 100 caracteres debe ser válida")
	assert.Equal(t, exactTypography, brand.Typography)
}

func TestNewMarketplacebrand_WithEmptyName_ReturnsError(t *testing.T) {
	// Act
	brand, err := entity.NewMarketplacebrand("")

	// Assert
	assert.Error(t, err)
	assert.Nil(t, brand)
	assert.Contains(t, err.Error(), "name es requerido")
}

func TestNewMarketplacebrand_WithValidName_HasEmptyVisualFields(t *testing.T) {
	// Act
	brand, err := entity.NewMarketplacebrand("Coca-Cola")

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "", brand.BackgroundColor, "nuevo brand debe tener BackgroundColor vacío")
	assert.Equal(t, "", brand.TextColor, "nuevo brand debe tener TextColor vacío")
	assert.Equal(t, "", brand.Typography, "nuevo brand debe tener Typography vacío")
}
