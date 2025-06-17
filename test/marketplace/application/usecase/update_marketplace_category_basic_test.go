package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"pim/src/marketplace/application/request"
)

// TestUpdateMarketplaceCategoryBasic tests básicos para actualización de categorías marketplace
func TestUpdateMarketplaceCategoryBasic_RequestValidation_Success(t *testing.T) {
	// Arrange
	name := "Updated Electronics"
	slug := "updated-electronics"
	description := "Updated electronic products category"
	isActive := true
	sortOrder := 5

	req := &request.UpdateMarketplaceCategoryRequest{
		Name:        &name,
		Slug:        &slug,
		Description: &description,
		IsActive:    &isActive,
		SortOrder:   &sortOrder,
	}

	// Act
	err := req.Validate()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "Updated Electronics", *req.Name)
	assert.Equal(t, "updated-electronics", *req.Slug)
	assert.Equal(t, "Updated electronic products category", *req.Description)
	assert.True(t, *req.IsActive)
	assert.Equal(t, 5, *req.SortOrder)
}

func TestUpdateMarketplaceCategoryBasic_RequestValidation_PartialUpdate(t *testing.T) {
	// Arrange - solo actualizar nombre
	name := "New Category Name"
	req := &request.UpdateMarketplaceCategoryRequest{
		Name: &name,
	}

	// Act
	err := req.Validate()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "New Category Name", *req.Name)
	assert.Nil(t, req.Slug)
	assert.Nil(t, req.Description)
	assert.Nil(t, req.IsActive)
	assert.Nil(t, req.SortOrder)
}

func TestUpdateMarketplaceCategoryBasic_RequestValidation_EmptyFields(t *testing.T) {
	// Test que campos vacíos causan errores de validación
	t.Run("empty_name_should_fail", func(t *testing.T) {
		name := ""
		req := &request.UpdateMarketplaceCategoryRequest{
			Name: &name,
		}

		err := req.Validate()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "name")
	})

	t.Run("empty_slug_should_fail", func(t *testing.T) {
		slug := ""
		req := &request.UpdateMarketplaceCategoryRequest{
			Slug: &slug,
		}

		err := req.Validate()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "slug")
	})

	t.Run("negative_sort_order_should_fail", func(t *testing.T) {
		sortOrder := -1
		req := &request.UpdateMarketplaceCategoryRequest{
			SortOrder: &sortOrder,
		}

		err := req.Validate()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "sort_order")
	})
}

func TestUpdateMarketplaceCategoryBasic_RequestValidation_FieldLimits(t *testing.T) {
	t.Run("long_name_should_fail", func(t *testing.T) {
		// Crear nombre muy largo (más de 255 caracteres)
		longName := make([]byte, 256)
		for i := range longName {
			longName[i] = 'a'
		}
		name := string(longName)

		req := &request.UpdateMarketplaceCategoryRequest{
			Name: &name,
		}

		err := req.Validate()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "name")
		assert.Contains(t, err.Error(), "255")
	})

	t.Run("long_slug_should_fail", func(t *testing.T) {
		// Crear slug muy largo (más de 255 caracteres)
		longSlug := make([]byte, 256)
		for i := range longSlug {
			longSlug[i] = 'a'
		}
		slug := string(longSlug)

		req := &request.UpdateMarketplaceCategoryRequest{
			Slug: &slug,
		}

		err := req.Validate()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "slug")
		assert.Contains(t, err.Error(), "255")
	})

	t.Run("long_description_should_fail", func(t *testing.T) {
		// Crear descripción muy larga (más de 1000 caracteres)
		longDescription := make([]byte, 1001)
		for i := range longDescription {
			longDescription[i] = 'a'
		}
		description := string(longDescription)

		req := &request.UpdateMarketplaceCategoryRequest{
			Description: &description,
		}

		err := req.Validate()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "description")
		assert.Contains(t, err.Error(), "1000")
	})
}

func TestUpdateMarketplaceCategoryBasic_RequestValidation_SlugFormat(t *testing.T) {
	invalidSlugs := []string{
		"Invalid Slug", // espacios y mayúsculas
		"invalid_slug", // guiones bajos
		"invalid&slug", // caracteres especiales
		"INVALID-SLUG", // mayúsculas
	}

	for _, invalidSlug := range invalidSlugs {
		t.Run("invalid_slug_"+invalidSlug, func(t *testing.T) {
			req := &request.UpdateMarketplaceCategoryRequest{
				Slug: &invalidSlug,
			}

			err := req.Validate()
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "slug")
		})
	}

	validSlugs := []string{
		"valid-slug",
		"electronics",
		"home-garden",
		"sports-outdoors",
	}

	for _, validSlug := range validSlugs {
		t.Run("valid_slug_"+validSlug, func(t *testing.T) {
			req := &request.UpdateMarketplaceCategoryRequest{
				Slug: &validSlug,
			}

			err := req.Validate()
			assert.NoError(t, err)
		})
	}
}
