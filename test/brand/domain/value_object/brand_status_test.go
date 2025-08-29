package value_object

import (
	"testing"

	"saas-mt-pim-service/src/brand/domain/value_object"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewBrandStatus(t *testing.T) {
	t.Run("should create active status", func(t *testing.T) {
		// Act
		status, err := value_object.NewBrandStatus("active")

		// Assert
		require.NoError(t, err)
		assert.True(t, status.IsActive())
		assert.False(t, status.IsDeleted())
		assert.Equal(t, "active", status.String())
	})

	t.Run("should create inactive status", func(t *testing.T) {
		// Act
		status, err := value_object.NewBrandStatus("inactive")

		// Assert
		require.NoError(t, err)
		assert.False(t, status.IsActive())
		assert.False(t, status.IsDeleted())
		assert.Equal(t, "inactive", status.String())
	})

	t.Run("should create deleted status", func(t *testing.T) {
		// Act
		status, err := value_object.NewBrandStatus("deleted")

		// Assert
		require.NoError(t, err)
		assert.False(t, status.IsActive())
		assert.True(t, status.IsDeleted())
		assert.Equal(t, "deleted", status.String())
	})

	t.Run("should fail with invalid status", func(t *testing.T) {
		// Act
		status, err := value_object.NewBrandStatus("invalid")

		// Assert
		assert.Error(t, err)
		assert.Equal(t, value_object.BrandStatus(""), status)
	})

	t.Run("should fail with empty status", func(t *testing.T) {
		// Act
		status, err := value_object.NewBrandStatus("")

		// Assert
		assert.Error(t, err)
		assert.Equal(t, value_object.BrandStatus(""), status)
	})
}

func TestBrandStatusConstants(t *testing.T) {
	t.Run("should have correct constant values", func(t *testing.T) {
		assert.Equal(t, "active", value_object.BrandStatusActive.String())
		assert.Equal(t, "inactive", value_object.BrandStatusInactive.String())
		assert.Equal(t, "deleted", value_object.BrandStatusDeleted.String())
	})

	t.Run("should validate constants correctly", func(t *testing.T) {
		assert.True(t, value_object.BrandStatusActive.IsActive())
		assert.False(t, value_object.BrandStatusInactive.IsActive())
		assert.True(t, value_object.BrandStatusDeleted.IsDeleted())
	})
}

func TestBrandStatus_IsActive(t *testing.T) {
	t.Run("should return true for active status", func(t *testing.T) {
		assert.True(t, value_object.BrandStatusActive.IsActive())
	})

	t.Run("should return false for inactive status", func(t *testing.T) {
		assert.False(t, value_object.BrandStatusInactive.IsActive())
	})

	t.Run("should return false for deleted status", func(t *testing.T) {
		assert.False(t, value_object.BrandStatusDeleted.IsActive())
	})
}

func TestBrandStatus_IsDeleted(t *testing.T) {
	t.Run("should return false for active status", func(t *testing.T) {
		assert.False(t, value_object.BrandStatusActive.IsDeleted())
	})

	t.Run("should return false for inactive status", func(t *testing.T) {
		assert.False(t, value_object.BrandStatusInactive.IsDeleted())
	})

	t.Run("should return true for deleted status", func(t *testing.T) {
		assert.True(t, value_object.BrandStatusDeleted.IsDeleted())
	})
}

func TestBrandStatus_IsValid(t *testing.T) {
	t.Run("should return true for valid statuses", func(t *testing.T) {
		assert.True(t, value_object.BrandStatusActive.IsValid())
		assert.True(t, value_object.BrandStatusInactive.IsValid())
		assert.True(t, value_object.BrandStatusDeleted.IsValid())
	})

	t.Run("should return false for invalid status", func(t *testing.T) {
		invalidStatus := value_object.BrandStatus("invalid")
		assert.False(t, invalidStatus.IsValid())
	})

	t.Run("should return false for empty status", func(t *testing.T) {
		emptyStatus := value_object.BrandStatus("")
		assert.False(t, emptyStatus.IsValid())
	})
}
