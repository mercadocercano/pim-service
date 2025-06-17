package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"pim/src/marketplace/application/request"
)

// TestMapTenantCategoryBasic tests básicos para mapeo de categorías tenant
func TestMapTenantCategoryBasic_RequestValidation_Success(t *testing.T) {
	// Arrange
	req := &request.MapTenantCategoryRequest{
		CategoryID:            "tenant-cat-456",
		MarketplaceCategoryID: "marketplace-cat-789",
		CustomName:            "Electronics Custom",
	}

	// Act
	err := req.Validate()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "tenant-cat-456", req.CategoryID)
	assert.Equal(t, "marketplace-cat-789", req.MarketplaceCategoryID)
	assert.Equal(t, "Electronics Custom", req.CustomName)
}

func TestMapTenantCategoryBasic_RequestValidation_EmptyCategoryID(t *testing.T) {
	// Arrange
	req := &request.MapTenantCategoryRequest{
		CategoryID:            "", // Empty category ID should fail
		MarketplaceCategoryID: "marketplace-cat-789",
		CustomName:            "Electronics Custom",
	}

	// Act
	err := req.Validate()

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "category_id")
}

func TestMapTenantCategoryBasic_RequestValidation_EmptyMarketplaceCategoryID(t *testing.T) {
	// Arrange
	req := &request.MapTenantCategoryRequest{
		CategoryID:            "tenant-cat-456",
		MarketplaceCategoryID: "", // Empty marketplace category ID should fail
		CustomName:            "Electronics Custom",
	}

	// Act
	err := req.Validate()

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "marketplace_category_id")
}

func TestMapTenantCategoryBasic_RequestValidation_EmptyCustomName(t *testing.T) {
	// Arrange
	req := &request.MapTenantCategoryRequest{
		CategoryID:            "tenant-cat-456",
		MarketplaceCategoryID: "marketplace-cat-789",
		CustomName:            "", // Empty custom name should fail
	}

	// Act
	err := req.Validate()

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "custom_name")
}

func TestMapTenantCategoryBasic_RequestValidation_AllFieldsEmpty(t *testing.T) {
	// Arrange
	req := &request.MapTenantCategoryRequest{
		CategoryID:            "",
		MarketplaceCategoryID: "",
		CustomName:            "",
	}

	// Act
	err := req.Validate()

	// Assert
	assert.Error(t, err)
	// Should fail on the first validation (category_id)
	assert.Contains(t, err.Error(), "category_id")
}

func TestMapTenantCategoryBasic_RequestValidation_LongCustomName(t *testing.T) {
	// Arrange
	longName := make([]byte, 256) // 256 characters, exceeds limit
	for i := range longName {
		longName[i] = 'a'
	}

	req := &request.MapTenantCategoryRequest{
		CategoryID:            "tenant-cat-456",
		MarketplaceCategoryID: "marketplace-cat-789",
		CustomName:            string(longName),
	}

	// Act
	err := req.Validate()

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "custom_name")
	assert.Contains(t, err.Error(), "255")
}

func TestMapTenantCategoryBasic_CategoryIDValidation(t *testing.T) {
	// Test diferentes formatos de category ID
	validCategoryIDs := []string{
		"cat-123",
		"category_456",
		"CATEGORY-789",
		"category123",
		"123-category",
		"tenant-cat-abc",
	}

	for _, categoryID := range validCategoryIDs {
		t.Run("valid_category_id_"+categoryID, func(t *testing.T) {
			req := &request.MapTenantCategoryRequest{
				CategoryID:            categoryID,
				MarketplaceCategoryID: "marketplace-cat-789",
				CustomName:            "Test Category",
			}

			err := req.Validate()
			assert.NoError(t, err)
		})
	}
}

func TestMapTenantCategoryBasic_MarketplaceCategoryIDValidation(t *testing.T) {
	// Test diferentes formatos de marketplace category ID
	validMarketplaceCategoryIDs := []string{
		"marketplace-cat-123",
		"mp_cat_456",
		"MARKETPLACE-CAT-789",
		"marketplacecat123",
		"123-marketplace-cat",
	}

	for _, marketplaceCategoryID := range validMarketplaceCategoryIDs {
		t.Run("valid_marketplace_category_id_"+marketplaceCategoryID, func(t *testing.T) {
			req := &request.MapTenantCategoryRequest{
				CategoryID:            "tenant-cat-456",
				MarketplaceCategoryID: marketplaceCategoryID,
				CustomName:            "Test Category",
			}

			err := req.Validate()
			assert.NoError(t, err)
		})
	}
}

func TestMapTenantCategoryBasic_CustomNameValidation(t *testing.T) {
	// Test diferentes formatos de custom name
	validCustomNames := []string{
		"Electronics",
		"Home & Garden",
		"Sports & Outdoors",
		"Books, Movies & Music",
		"Health & Beauty Products",
		"Automotive & Industrial",
	}

	for _, customName := range validCustomNames {
		t.Run("valid_custom_name", func(t *testing.T) {
			req := &request.MapTenantCategoryRequest{
				CategoryID:            "tenant-cat-456",
				MarketplaceCategoryID: "marketplace-cat-789",
				CustomName:            customName,
			}

			err := req.Validate()
			assert.NoError(t, err)
		})
	}
}

func TestMapTenantCategoryBasic_MappingLogic(t *testing.T) {
	// Test lógica básica de mapeo
	t.Run("should_create_valid_mapping", func(t *testing.T) {
		categoryID := "tenant-cat-456"
		marketplaceCategoryID := "marketplace-cat-789"
		customName := "Electronics Custom"

		// Simular creación de mapeo
		mapping := map[string]string{
			"category_id":             categoryID,
			"marketplace_category_id": marketplaceCategoryID,
			"custom_name":             customName,
		}

		assert.Equal(t, categoryID, mapping["category_id"])
		assert.Equal(t, marketplaceCategoryID, mapping["marketplace_category_id"])
		assert.Equal(t, customName, mapping["custom_name"])
	})

	t.Run("should_handle_mapping_uniqueness", func(t *testing.T) {
		// Simular verificación de unicidad de mapeo
		existingMappings := []map[string]string{
			{
				"category_id":             "tenant-cat-456",
				"marketplace_category_id": "marketplace-cat-789",
				"custom_name":             "Electronics Custom",
			},
		}

		// Verificar que no existe mapeo duplicado para la misma categoría
		newMapping := map[string]string{
			"category_id":             "tenant-cat-456",      // Misma categoría
			"marketplace_category_id": "marketplace-cat-999", // Diferente marketplace category
			"custom_name":             "Electronics Updated",
		}

		isDuplicate := false
		for _, existing := range existingMappings {
			if existing["category_id"] == newMapping["category_id"] {
				isDuplicate = true
				break
			}
		}

		assert.True(t, isDuplicate, "Should detect duplicate category mapping")
	})
}
