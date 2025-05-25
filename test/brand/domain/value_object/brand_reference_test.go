package value_object

import (
	"testing"

	"pim/src/brand/domain/value_object"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewBrandReference(t *testing.T) {
	t.Run("should create brand reference with valid data", func(t *testing.T) {
		// Arrange
		id := "brand-123"
		name := "Nike"
		description := "Marca deportiva"

		// Act
		ref, err := value_object.NewBrandReference(id, name, description)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, id, ref.ID)
		assert.Equal(t, name, ref.Name)
		assert.Equal(t, description, ref.Description)
	})

	t.Run("should create brand reference with empty description", func(t *testing.T) {
		// Arrange
		id := "brand-123"
		name := "Nike"
		description := ""

		// Act
		ref, err := value_object.NewBrandReference(id, name, description)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, id, ref.ID)
		assert.Equal(t, name, ref.Name)
		assert.Equal(t, description, ref.Description)
	})

	t.Run("should fail when ID is empty", func(t *testing.T) {
		// Arrange
		id := ""
		name := "Nike"
		description := "Marca deportiva"

		// Act
		ref, err := value_object.NewBrandReference(id, name, description)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, ref)
		assert.Contains(t, err.Error(), "ID")
	})

	t.Run("should fail when name is empty", func(t *testing.T) {
		// Arrange
		id := "brand-123"
		name := ""
		description := "Marca deportiva"

		// Act
		ref, err := value_object.NewBrandReference(id, name, description)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, ref)
		assert.Contains(t, err.Error(), "nombre")
	})
}

func TestBrandReference_IsEmpty(t *testing.T) {
	t.Run("should return false for valid reference", func(t *testing.T) {
		// Arrange
		ref, _ := value_object.NewBrandReference("brand-123", "Nike", "Marca deportiva")

		// Act & Assert
		assert.False(t, ref.IsEmpty())
	})

	t.Run("should return true for nil reference", func(t *testing.T) {
		// Arrange
		var ref *value_object.BrandReference

		// Act & Assert
		assert.True(t, ref.IsEmpty())
	})

	t.Run("should return true for reference with empty ID", func(t *testing.T) {
		// Arrange
		ref := &value_object.BrandReference{
			ID:          "",
			Name:        "Nike",
			Description: "Marca deportiva",
		}

		// Act & Assert
		assert.True(t, ref.IsEmpty())
	})
}

func TestBrandReference_Equals(t *testing.T) {
	t.Run("should return true for same references", func(t *testing.T) {
		// Arrange
		ref1, _ := value_object.NewBrandReference("brand-123", "Nike", "Marca deportiva")
		ref2, _ := value_object.NewBrandReference("brand-123", "Nike Different", "Different description")

		// Act & Assert
		assert.True(t, ref1.Equals(ref2))
	})

	t.Run("should return false for different references", func(t *testing.T) {
		// Arrange
		ref1, _ := value_object.NewBrandReference("brand-123", "Nike", "Marca deportiva")
		ref2, _ := value_object.NewBrandReference("brand-456", "Adidas", "Otra marca")

		// Act & Assert
		assert.False(t, ref1.Equals(ref2))
	})

	t.Run("should return true for both nil references", func(t *testing.T) {
		// Arrange
		var ref1, ref2 *value_object.BrandReference

		// Act & Assert
		assert.True(t, ref1.Equals(ref2))
	})

	t.Run("should return false when one reference is nil", func(t *testing.T) {
		// Arrange
		ref1, _ := value_object.NewBrandReference("brand-123", "Nike", "Marca deportiva")
		var ref2 *value_object.BrandReference

		// Act & Assert
		assert.False(t, ref1.Equals(ref2))
		assert.False(t, ref2.Equals(ref1))
	})
}
