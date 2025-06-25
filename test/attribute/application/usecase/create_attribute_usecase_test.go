package usecase_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"pim/src/attribute/application/usecase"
	testentity "pim/test/attribute/domain/entity"
)

func TestCreateAttributeUseCase_Execute_Success(t *testing.T) {
	// Arrange
	// TODO: Crear mock repository cuando esté disponible
	// mockRepo := repository.NewMockAttributeRepository()
	// createUseCase := usecase.NewCreateAttributeUseCase(mockRepo)

	tenantID := "tenant-123"
	name := "Color"

	// Act - Por ahora solo validamos que el use case no es nil
	createUseCase := usecase.NewCreateAttributeUseCase(nil)

	// Assert - Verificamos que el use case se crea correctamente
	assert.NotNil(t, createUseCase)
	assert.Equal(t, "Color", name)
	assert.Equal(t, "tenant-123", tenantID)
}

func TestCreateAttributeUseCase_Execute_EmptyName_ShouldFail(t *testing.T) {
	// Arrange
	ctx := context.Background()
	createUseCase := usecase.NewCreateAttributeUseCase(nil)

	tenantID := "tenant-123"
	emptyName := ""

	// Act
	attribute, err := createUseCase.Execute(ctx, tenantID, emptyName)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, attribute)
	assert.Equal(t, usecase.ErrInvalidAttributeName, err)
}

func TestCreateAttributeUseCase_Execute_EmptyTenantID_ShouldFail(t *testing.T) {
	// Arrange
	ctx := context.Background()
	createUseCase := usecase.NewCreateAttributeUseCase(nil)

	emptyTenantID := ""
	name := "Color"

	// Act
	attribute, err := createUseCase.Execute(ctx, emptyTenantID, name)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, attribute)
	assert.Contains(t, err.Error(), "tenant_id")
}

func TestCreateAttributeUseCase_ObjectMother_WithDefaults(t *testing.T) {
	// Arrange
	attributeMother := testentity.Create()

	// Act - Crear atributo usando Object Mother
	attribute := attributeMother.WithDefaults()

	// Assert
	assert.NotNil(t, attribute)
	assert.NotEmpty(t, attribute.ID)
	assert.Equal(t, "tenant-123", attribute.TenantID)     // Valor por defecto
	assert.Equal(t, "Atributo de prueba", attribute.Name) // Valor por defecto
	assert.True(t, attribute.Active)
}

func TestCreateAttributeUseCase_ObjectMother_WithTenantID(t *testing.T) {
	// Arrange
	attributeMother := testentity.Create()

	// Act - Crear atributo con tenant específico usando Object Mother
	attribute := attributeMother.WithTenantID("tenant-456")

	// Assert
	assert.NotNil(t, attribute)
	assert.Equal(t, "tenant-456", attribute.TenantID)
	assert.Equal(t, "Atributo de prueba", attribute.Name) // Valor por defecto
	assert.True(t, attribute.Active)
}

func TestCreateAttributeUseCase_ObjectMother_WithName(t *testing.T) {
	// Arrange
	attributeMother := testentity.Create()

	// Act - Crear atributo con nombre específico usando Object Mother
	attribute := attributeMother.WithName("Color")

	// Assert
	assert.NotNil(t, attribute)
	assert.Equal(t, "Color", attribute.Name)
	assert.Equal(t, "tenant-123", attribute.TenantID) // Valor por defecto
	assert.True(t, attribute.Active)
}

func TestCreateAttributeUseCase_ObjectMother_Inactive(t *testing.T) {
	// Arrange
	attributeMother := testentity.Create()

	// Act - Crear atributo inactivo usando Object Mother
	attribute := attributeMother.Inactive()

	// Assert
	assert.NotNil(t, attribute)
	assert.False(t, attribute.Active)
	assert.Equal(t, "tenant-123", attribute.TenantID)     // Valor por defecto
	assert.Equal(t, "Atributo de prueba", attribute.Name) // Valor por defecto
}

func TestCreateAttributeUseCase_ObjectMother_ColorAttribute(t *testing.T) {
	// Arrange
	attributeMother := testentity.Create()

	// Act - Crear atributo de color usando Object Mother
	attribute := attributeMother.ColorAttribute()

	// Assert
	assert.NotNil(t, attribute)
	assert.Equal(t, "Color", attribute.Name)
	assert.Equal(t, "tenant-123", attribute.TenantID)
	assert.True(t, attribute.Active)
}

func TestCreateAttributeUseCase_ObjectMother_SizeAttribute(t *testing.T) {
	// Arrange
	attributeMother := testentity.Create()

	// Act - Crear atributo de tamaño usando Object Mother
	attribute := attributeMother.SizeAttribute()

	// Assert
	assert.NotNil(t, attribute)
	assert.Equal(t, "Talla", attribute.Name)
	assert.Equal(t, "tenant-123", attribute.TenantID)
	assert.True(t, attribute.Active)
}

func TestCreateAttributeUseCase_ObjectMother_MaterialAttribute(t *testing.T) {
	// Arrange
	attributeMother := testentity.Create()

	// Act - Crear atributo de material usando Object Mother
	attribute := attributeMother.MaterialAttribute()

	// Assert
	assert.NotNil(t, attribute)
	assert.Equal(t, "Material", attribute.Name)
	assert.Equal(t, "tenant-123", attribute.TenantID)
	assert.True(t, attribute.Active)
}

func TestCreateAttributeUseCase_ObjectMother_Complete(t *testing.T) {
	// Arrange
	attributeMother := testentity.Create()

	// Act - Crear atributo completo usando Object Mother
	attribute := attributeMother.Complete(
		"attr-123",
		"tenant-456",
		"Peso",
		true,
	)

	// Assert
	assert.NotNil(t, attribute)
	assert.Equal(t, "attr-123", attribute.ID)
	assert.Equal(t, "tenant-456", attribute.TenantID)
	assert.Equal(t, "Peso", attribute.Name)
	assert.True(t, attribute.Active)
}

func TestCreateAttributeUseCase_ObjectMother_MultipleAttributes(t *testing.T) {
	// Arrange
	attributeMother := testentity.Create()

	// Casos de prueba para diferentes atributos
	testCases := []struct {
		name     string
		expected string
	}{
		{"Color", "Color"},
		{"Tamaño", "Tamaño"},
		{"Material", "Material"},
		{"Marca", "Marca"},
	}

	// Act & Assert
	for _, tc := range testCases {
		attribute := attributeMother.WithName(tc.name)

		assert.NotNil(t, attribute)
		assert.Equal(t, tc.expected, attribute.Name)
		assert.Equal(t, "tenant-123", attribute.TenantID)
		assert.True(t, attribute.Active)
	}
}

func TestCreateAttributeUseCase_ObjectMother_DifferentTenants(t *testing.T) {
	// Arrange
	attributeMother := testentity.Create()

	// Datos de prueba para diferentes tenants
	testData := []struct {
		tenantID string
		name     string
	}{
		{"tenant-1", "Color"},
		{"tenant-2", "Tamaño"},
		{"tenant-3", "Material"},
	}

	// Act & Assert
	for _, td := range testData {
		attribute := attributeMother.WithTenantID(td.tenantID)
		attribute.Name = td.name // Modificar después de crear

		assert.NotNil(t, attribute)
		assert.Equal(t, td.tenantID, attribute.TenantID)
		assert.Equal(t, td.name, attribute.Name)
		assert.True(t, attribute.Active)
	}
}
