package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"saas-mt-pim-service/src/schema_validation/domain/service"
)

func TestCategoryDeductionService_DeduceCategory_MatchesKeywordDictionary(t *testing.T) {
	// Arrange
	svc := service.NewCategoryDeductionService()

	tests := []struct {
		productName      string
		expectedCategory string
	}{
		{"Coca Cola 2.25L", "Bebidas"},
		{"Alfajor Triple Havanna", "Golosinas"},
		{"Galletitas Oreo x 118g", "Galletitas"},
		{"Doritos Nachos 150g", "Snacks"},
		{"Leche Entera Serenisima 1L", "Lácteos"},
		{"Detergente Magistral 750ml", "Limpieza"},
		{"Shampoo Head & Shoulders 400ml", "Higiene"},
		{"Arroz Largo Fino Gallo 1kg", "Alimentos"},
		{"Tornillo autoperforante 6x1", "Ferretería"},
		{"Martillo Carpintero 500g", "Herramientas Manuales"},
		{"Taladro Percutor 13mm", "Herramientas Eléctricas"},
		{"Cable unipolar 2.5mm", "Electricidad"},
		{"Pintura Latex Interior 4L", "Pinturas"},
	}

	for _, tt := range tests {
		t.Run(tt.productName, func(t *testing.T) {
			// Act
			result := svc.DeduceCategory(tt.productName, nil)

			// Assert
			assert.Equal(t, tt.expectedCategory, result)
		})
	}
}

func TestCategoryDeductionService_DeduceCategory_PrioritizesTenantCategories(t *testing.T) {
	// Arrange
	svc := service.NewCategoryDeductionService()
	tenantCategories := []string{"Bebidas Frías", "Galletitas Dulces"}

	// "Galletitas Oreo" contiene "Galletitas" que existe en tenant categories como substring
	// pero "Galletitas Dulces" no matchea porque no es substring de "Galletitas Oreo"
	result := svc.DeduceCategory("Bebidas Frías Coca Cola", tenantCategories)

	// Assert - should match tenant category first
	assert.Equal(t, "Bebidas Frías", result)
}

func TestCategoryDeductionService_DeduceCategory_UnknownProduct(t *testing.T) {
	// Arrange
	svc := service.NewCategoryDeductionService()

	// Act
	result := svc.DeduceCategory("XYZXYZ Producto Inexistente", nil)

	// Assert
	assert.Empty(t, result)
}

func TestCategoryDeductionService_DeduceCategory_CaseInsensitive(t *testing.T) {
	// Arrange
	svc := service.NewCategoryDeductionService()

	// Act
	result := svc.DeduceCategory("COCA COLA 2.25L", nil)

	// Assert
	assert.Equal(t, "Bebidas", result)
}

func TestCategoryDeductionService_DeduceCategory_EmptyName(t *testing.T) {
	// Arrange
	svc := service.NewCategoryDeductionService()

	// Act
	result := svc.DeduceCategory("", nil)

	// Assert
	assert.Empty(t, result)
}

func TestCategoryDeductionService_DeduceCategory_EmptyTenantCategories(t *testing.T) {
	// Arrange
	svc := service.NewCategoryDeductionService()

	// Act
	result := svc.DeduceCategory("Coca Cola Zero", []string{})

	// Assert
	assert.Equal(t, "Bebidas", result)
}
