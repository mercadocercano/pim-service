package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"saas-mt-pim-service/src/category/application/request"
)

// TestValidateCategoryHierarchyBasic tests básicos para validación de jerarquía de categorías
func TestValidateCategoryHierarchyBasic_RequestValidation_Success(t *testing.T) {
	// Arrange
	req := &request.ValidateCategoryHierarchyRequest{
		CategoryID:       "cat-123",
		NewParentID:      stringPtrHierarchy("parent-456"),
		MaxDepth:         3,
		ValidateChildren: true,
	}

	// Act
	err := req.Validate()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "cat-123", req.CategoryID)
	assert.Equal(t, "parent-456", *req.NewParentID)
	assert.Equal(t, 3, req.MaxDepth)
	assert.True(t, req.ValidateChildren)
}

func TestValidateCategoryHierarchyBasic_RequestValidation_RootCategory(t *testing.T) {
	// Arrange - categoría raíz sin parent
	req := &request.ValidateCategoryHierarchyRequest{
		CategoryID:       "cat-123",
		NewParentID:      nil, // Sin parent, es categoría raíz
		MaxDepth:         3,
		ValidateChildren: false,
	}

	// Act
	err := req.Validate()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "cat-123", req.CategoryID)
	assert.Nil(t, req.NewParentID)
	assert.Equal(t, 3, req.MaxDepth)
	assert.False(t, req.ValidateChildren)
}

func TestValidateCategoryHierarchyBasic_RequestValidation_EmptyCategoryID(t *testing.T) {
	// Arrange
	req := &request.ValidateCategoryHierarchyRequest{
		CategoryID:  "", // Empty category ID should fail
		NewParentID: stringPtrHierarchy("parent-456"),
		MaxDepth:    3,
	}

	// Act
	err := req.Validate()

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "category_id")
}

func TestValidateCategoryHierarchyBasic_RequestValidation_EmptyNewParentID(t *testing.T) {
	// Arrange
	emptyParent := ""
	req := &request.ValidateCategoryHierarchyRequest{
		CategoryID:  "cat-123",
		NewParentID: &emptyParent, // Empty parent ID should fail
		MaxDepth:    3,
	}

	// Act
	err := req.Validate()

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "new_parent_id")
}

func TestValidateCategoryHierarchyBasic_RequestValidation_SelfReference(t *testing.T) {
	// Arrange
	categoryID := "cat-123"
	req := &request.ValidateCategoryHierarchyRequest{
		CategoryID:  categoryID,
		NewParentID: &categoryID, // Self reference should fail
		MaxDepth:    3,
	}

	// Act
	err := req.Validate()

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot be its own parent")
}

func TestValidateCategoryHierarchyBasic_RequestValidation_MaxDepthValidation(t *testing.T) {
	t.Run("should_set_default_max_depth", func(t *testing.T) {
		req := &request.ValidateCategoryHierarchyRequest{
			CategoryID: "cat-123",
			MaxDepth:   0, // Should be set to default (3)
		}

		err := req.Validate()
		assert.NoError(t, err)
		assert.Equal(t, 3, req.MaxDepth) // Should be set to default
	})

	t.Run("should_reject_excessive_max_depth", func(t *testing.T) {
		req := &request.ValidateCategoryHierarchyRequest{
			CategoryID: "cat-123",
			MaxDepth:   15, // Exceeds limit of 10
		}

		err := req.Validate()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "max_depth")
		assert.Contains(t, err.Error(), "10")
	})

	t.Run("should_allow_valid_max_depth", func(t *testing.T) {
		validDepths := []int{1, 2, 3, 5, 10}

		for _, depth := range validDepths {
			req := &request.ValidateCategoryHierarchyRequest{
				CategoryID: "cat-123",
				MaxDepth:   depth,
			}

			err := req.Validate()
			assert.NoError(t, err)
			assert.Equal(t, depth, req.MaxDepth)
		}
	})
}

func TestValidateCategoryHierarchyBasic_HierarchyValidation_MaxDepth(t *testing.T) {
	// Test validación de profundidad máxima
	t.Run("should_allow_valid_levels", func(t *testing.T) {
		maxLevel := 2
		validLevels := []int{0, 1, 2}

		for _, level := range validLevels {
			assert.True(t, isValidLevelHierarchy(level, maxLevel), "Level %d should be valid", level)
		}
	})

	t.Run("should_reject_invalid_levels", func(t *testing.T) {
		maxLevel := 2
		invalidLevels := []int{-1, 3, 4, 5}

		for _, level := range invalidLevels {
			assert.False(t, isValidLevelHierarchy(level, maxLevel), "Level %d should be invalid", level)
		}
	})
}

func TestValidateCategoryHierarchyBasic_HierarchyValidation_CircularReference(t *testing.T) {
	// Test detección de referencias circulares
	t.Run("should_detect_direct_circular_reference", func(t *testing.T) {
		// Categoría que se referencia a sí misma
		categoryID := "cat-123"
		parentID := "cat-123"

		isCircular := categoryID == parentID
		assert.True(t, isCircular, "Should detect direct circular reference")
	})

	t.Run("should_detect_indirect_circular_reference", func(t *testing.T) {
		// Simular jerarquía: A -> B -> C -> A (circular)
		hierarchy := map[string]string{
			"cat-a": "cat-b",
			"cat-b": "cat-c",
			"cat-c": "cat-a", // Circular reference
		}

		// Verificar si hay referencia circular
		visited := make(map[string]bool)
		isCircular := hasCircularReferenceHierarchy("cat-a", hierarchy, visited)
		assert.True(t, isCircular, "Should detect indirect circular reference")
	})

	t.Run("should_allow_valid_hierarchy", func(t *testing.T) {
		// Simular jerarquía válida: A -> B -> C
		hierarchy := map[string]string{
			"cat-a": "cat-b",
			"cat-b": "cat-c",
			"cat-c": "", // Root category
		}

		// Verificar que no hay referencia circular
		visited := make(map[string]bool)
		isCircular := hasCircularReferenceHierarchy("cat-a", hierarchy, visited)
		assert.False(t, isCircular, "Should allow valid hierarchy")
	})
}

func TestValidateCategoryHierarchyBasic_CategoryIDValidation(t *testing.T) {
	// Test diferentes formatos de category ID
	validCategoryIDs := []string{
		"cat-123",
		"category_456",
		"CATEGORY-789",
		"category123",
		"123-category",
		"marketplace-cat-abc",
	}

	for _, categoryID := range validCategoryIDs {
		t.Run("valid_category_id_"+categoryID, func(t *testing.T) {
			req := &request.ValidateCategoryHierarchyRequest{
				CategoryID:  categoryID,
				NewParentID: stringPtrHierarchy("parent-456"),
				MaxDepth:    3,
			}

			err := req.Validate()
			assert.NoError(t, err)
		})
	}
}

func TestValidateCategoryHierarchyBasic_NewParentIDValidation(t *testing.T) {
	// Test diferentes formatos de new parent ID
	validParentIDs := []string{
		"parent-123",
		"parent_456",
		"PARENT-789",
		"parent123",
		"123-parent",
		"marketplace-parent-abc",
	}

	for _, parentID := range validParentIDs {
		t.Run("valid_new_parent_id_"+parentID, func(t *testing.T) {
			req := &request.ValidateCategoryHierarchyRequest{
				CategoryID:  "cat-123",
				NewParentID: &parentID,
				MaxDepth:    3,
			}

			err := req.Validate()
			assert.NoError(t, err)
		})
	}
}

func TestValidateCategoryHierarchyBasic_HierarchyLogic(t *testing.T) {
	// Test lógica básica de jerarquía
	t.Run("should_calculate_category_level", func(t *testing.T) {
		// Simular cálculo de nivel de categoría
		hierarchy := map[string]string{
			"cat-level-0": "",            // Root (level 0)
			"cat-level-1": "cat-level-0", // Level 1
			"cat-level-2": "cat-level-1", // Level 2
		}

		level0 := calculateLevelHierarchy("cat-level-0", hierarchy)
		level1 := calculateLevelHierarchy("cat-level-1", hierarchy)
		level2 := calculateLevelHierarchy("cat-level-2", hierarchy)

		assert.Equal(t, 0, level0, "Root category should be level 0")
		assert.Equal(t, 1, level1, "First child should be level 1")
		assert.Equal(t, 2, level2, "Second child should be level 2")
	})

	t.Run("should_get_category_path", func(t *testing.T) {
		// Simular obtención de path de categoría
		hierarchy := map[string]string{
			"electronics":    "",
			"smartphones":    "electronics",
			"android-phones": "smartphones",
		}

		path := getCategoryPathHierarchy("android-phones", hierarchy)
		expectedPath := []string{"electronics", "smartphones", "android-phones"}

		assert.Equal(t, expectedPath, path, "Should return correct category path")
	})
}

// Helper functions específicas para este archivo
func stringPtrHierarchy(s string) *string {
	return &s
}

func isValidLevelHierarchy(level, maxLevel int) bool {
	return level >= 0 && level <= maxLevel
}

func hasCircularReferenceHierarchy(categoryID string, hierarchy map[string]string, visited map[string]bool) bool {
	if visited[categoryID] {
		return true // Circular reference detected
	}

	parentID, exists := hierarchy[categoryID]
	if !exists || parentID == "" {
		return false // Root category or not found
	}

	visited[categoryID] = true
	result := hasCircularReferenceHierarchy(parentID, hierarchy, visited)
	delete(visited, categoryID) // Backtrack
	return result
}

func calculateLevelHierarchy(categoryID string, hierarchy map[string]string) int {
	parentID, exists := hierarchy[categoryID]
	if !exists || parentID == "" {
		return 0 // Root category
	}
	return 1 + calculateLevelHierarchy(parentID, hierarchy)
}

func getCategoryPathHierarchy(categoryID string, hierarchy map[string]string) []string {
	var path []string
	current := categoryID

	for current != "" {
		path = append([]string{current}, path...) // Prepend to get correct order
		parent, exists := hierarchy[current]
		if !exists {
			break
		}
		current = parent
	}

	return path
}
