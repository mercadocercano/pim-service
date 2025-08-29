package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"saas-mt-pim-service/src/category/application/request"
)

// TestMarketplaceCategoryValidation tests para validaciones de categorías marketplace
func TestMarketplaceCategoryValidation_CreateRequest_Success(t *testing.T) {
	// Arrange
	req := &request.CreateMarketplaceCategoryRequest{
		Name:        "Electronics",
		Slug:        "electronics",
		Description: stringPtr("Electronic products category"),
		SortOrder:   1,
	}

	// Act
	err := req.Validate()

	// Assert
	assert.NoError(t, err)
}

func TestMarketplaceCategoryValidation_CreateRequest_EmptyName(t *testing.T) {
	// Arrange
	req := &request.CreateMarketplaceCategoryRequest{
		Name: "",
		Slug: "electronics",
	}

	// Act
	err := req.Validate()

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "name")
}

func TestMarketplaceCategoryValidation_CreateRequest_EmptySlug(t *testing.T) {
	// Arrange
	req := &request.CreateMarketplaceCategoryRequest{
		Name: "Electronics",
		Slug: "",
	}

	// Act
	err := req.Validate()

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "slug")
}

func TestMarketplaceCategoryValidation_CreateRequest_InvalidSlugFormat(t *testing.T) {
	// Arrange
	invalidSlugs := []string{
		"Electronics",     // uppercase
		"home garden",     // spaces
		"sports&outdoors", // special characters
		"home_garden",     // underscores
	}

	for _, slug := range invalidSlugs {
		t.Run("invalid_slug_"+slug, func(t *testing.T) {
			req := &request.CreateMarketplaceCategoryRequest{
				Name: "Test Category",
				Slug: slug,
			}

			// Act
			err := req.Validate()

			// Assert
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "slug")
		})
	}
}

func TestMarketplaceCategoryValidation_CreateRequest_ValidSlugFormats(t *testing.T) {
	// Arrange
	validSlugs := []string{
		"electronics",
		"home-garden",
		"sports-outdoors",
		"books-magazines",
		"health-beauty",
	}

	for _, slug := range validSlugs {
		t.Run("valid_slug_"+slug, func(t *testing.T) {
			req := &request.CreateMarketplaceCategoryRequest{
				Name: "Test Category",
				Slug: slug,
			}

			// Act
			err := req.Validate()

			// Assert
			assert.NoError(t, err)
		})
	}
}

func TestMarketplaceCategoryValidation_CreateRequest_LongName(t *testing.T) {
	// Arrange
	longName := make([]byte, 256) // 256 characters, exceeds limit
	for i := range longName {
		longName[i] = 'a'
	}

	req := &request.CreateMarketplaceCategoryRequest{
		Name: string(longName),
		Slug: "test-category",
	}

	// Act
	err := req.Validate()

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "name")
	assert.Contains(t, err.Error(), "255")
}

func TestMarketplaceCategoryValidation_CreateRequest_LongSlug(t *testing.T) {
	// Arrange
	longSlug := make([]byte, 256) // 256 characters, exceeds limit
	for i := range longSlug {
		longSlug[i] = 'a'
	}

	req := &request.CreateMarketplaceCategoryRequest{
		Name: "Test Category",
		Slug: string(longSlug),
	}

	// Act
	err := req.Validate()

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "slug")
	assert.Contains(t, err.Error(), "255")
}

func TestMarketplaceCategoryValidation_CreateRequest_LongDescription(t *testing.T) {
	// Arrange
	longDescription := make([]byte, 1001) // 1001 characters, exceeds limit
	for i := range longDescription {
		longDescription[i] = 'a'
	}

	req := &request.CreateMarketplaceCategoryRequest{
		Name:        "Test Category",
		Slug:        "test-category",
		Description: stringPtr(string(longDescription)),
	}

	// Act
	err := req.Validate()

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "description")
	assert.Contains(t, err.Error(), "1000")
}

func TestMarketplaceCategoryValidation_CreateRequest_NegativeSortOrder(t *testing.T) {
	// Arrange
	req := &request.CreateMarketplaceCategoryRequest{
		Name:      "Test Category",
		Slug:      "test-category",
		SortOrder: -1,
	}

	// Act
	err := req.Validate()

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "sort_order")
	assert.Contains(t, err.Error(), "negative")
}

func TestMarketplaceCategoryValidation_UpdateRequest_Success(t *testing.T) {
	// Arrange
	name := "Updated Electronics"
	slug := "updated-electronics"
	description := "Updated electronic products category"
	sortOrder := 5

	req := &request.UpdateMarketplaceCategoryRequest{
		Name:        &name,
		Slug:        &slug,
		Description: &description,
		SortOrder:   &sortOrder,
	}

	// Act
	err := req.Validate()

	// Assert
	assert.NoError(t, err)
}

func TestMarketplaceCategoryValidation_UpdateRequest_EmptyName(t *testing.T) {
	// Arrange
	name := ""
	req := &request.UpdateMarketplaceCategoryRequest{
		Name: &name,
	}

	// Act
	err := req.Validate()

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "name")
}

func TestMarketplaceCategoryValidation_UpdateRequest_EmptySlug(t *testing.T) {
	// Arrange
	slug := ""
	req := &request.UpdateMarketplaceCategoryRequest{
		Slug: &slug,
	}

	// Act
	err := req.Validate()

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "slug")
}

func TestMarketplaceCategoryValidation_UpdateRequest_NegativeSortOrder(t *testing.T) {
	// Arrange
	sortOrder := -1
	req := &request.UpdateMarketplaceCategoryRequest{
		SortOrder: &sortOrder,
	}

	// Act
	err := req.Validate()

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "sort_order")
}

// Helper function
func stringPtr(s string) *string {
	return &s
}
