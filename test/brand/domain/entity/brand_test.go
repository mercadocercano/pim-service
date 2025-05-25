package entity

import (
	"testing"

	"pim/src/brand/domain/entity"
	"pim/src/brand/domain/exception"
	"pim/src/brand/domain/value_object"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewBrand(t *testing.T) {
	t.Run("should create brand with valid data", func(t *testing.T) {
		// Arrange
		tenantID := "tenant-123"
		name := "Nike"
		description := "Marca deportiva"
		logoURL := stringPtr("https://nike.com/logo.png")
		website := stringPtr("https://nike.com")

		// Act
		brand, err := entity.NewBrand(tenantID, name, description, logoURL, website)

		// Assert
		require.NoError(t, err)
		assert.NotEmpty(t, brand.ID)
		assert.Equal(t, tenantID, brand.TenantID)
		assert.Equal(t, name, brand.Name)
		assert.Equal(t, description, brand.Description)
		assert.Equal(t, logoURL, brand.LogoURL)
		assert.Equal(t, website, brand.Website)
		assert.True(t, brand.Status.IsActive())
		assert.NotZero(t, brand.CreatedAt)
		assert.NotZero(t, brand.UpdatedAt)
	})

	t.Run("should create brand without optional fields", func(t *testing.T) {
		// Arrange
		tenantID := "tenant-123"
		name := "Generic Brand"
		description := "Generic description"

		// Act
		brand, err := entity.NewBrand(tenantID, name, description, nil, nil)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, name, brand.Name)
		assert.Nil(t, brand.LogoURL)
		assert.Nil(t, brand.Website)
	})

	t.Run("should fail when tenant ID is empty", func(t *testing.T) {
		// Arrange
		tenantID := ""
		name := "Nike"
		description := "Marca deportiva"

		// Act
		brand, err := entity.NewBrand(tenantID, name, description, nil, nil)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, exception.ErrBrandTenantRequired, err)
		assert.Nil(t, brand)
	})

	t.Run("should fail when name is empty", func(t *testing.T) {
		// Arrange
		tenantID := "tenant-123"
		name := ""
		description := "Marca deportiva"

		// Act
		brand, err := entity.NewBrand(tenantID, name, description, nil, nil)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, exception.ErrBrandNameRequired, err)
		assert.Nil(t, brand)
	})
}

func TestBrand_Update(t *testing.T) {
	t.Run("should update brand with valid data", func(t *testing.T) {
		// Arrange
		brand := Create().WithDefaults()
		originalUpdatedAt := brand.UpdatedAt

		newName := "Updated Brand"
		newDescription := "Updated description"
		newLogoURL := stringPtr("https://updated.com/logo.png")
		newWebsite := stringPtr("https://updated.com")

		// Act
		err := brand.Update(newName, newDescription, newLogoURL, newWebsite)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, newName, brand.Name)
		assert.Equal(t, newDescription, brand.Description)
		assert.Equal(t, newLogoURL, brand.LogoURL)
		assert.Equal(t, newWebsite, brand.Website)
		assert.True(t, brand.UpdatedAt.After(originalUpdatedAt))
	})

	t.Run("should fail when name is empty", func(t *testing.T) {
		// Arrange
		brand := Create().WithDefaults()
		originalName := brand.Name

		// Act
		err := brand.Update("", "New description", nil, nil)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, exception.ErrBrandNameRequired, err)
		assert.Equal(t, originalName, brand.Name) // No debe cambiar
	})
}

func TestBrand_Activate(t *testing.T) {
	t.Run("should activate inactive brand", func(t *testing.T) {
		// Arrange
		brand := Create().Inactive()
		originalUpdatedAt := brand.UpdatedAt

		// Act
		brand.Activate()

		// Assert
		assert.True(t, brand.IsActive())
		assert.True(t, brand.UpdatedAt.After(originalUpdatedAt))
	})
}

func TestBrand_Deactivate(t *testing.T) {
	t.Run("should deactivate active brand", func(t *testing.T) {
		// Arrange
		brand := Create().WithDefaults()
		originalUpdatedAt := brand.UpdatedAt

		// Act
		brand.Deactivate()

		// Assert
		assert.False(t, brand.IsActive())
		assert.Equal(t, value_object.BrandStatusInactive, brand.Status)
		assert.True(t, brand.UpdatedAt.After(originalUpdatedAt))
	})
}

func TestBrand_Delete(t *testing.T) {
	t.Run("should delete active brand", func(t *testing.T) {
		// Arrange
		brand := Create().WithDefaults()
		originalUpdatedAt := brand.UpdatedAt

		// Act
		err := brand.Delete()

		// Assert
		require.NoError(t, err)
		assert.True(t, brand.IsDeleted())
		assert.Equal(t, value_object.BrandStatusDeleted, brand.Status)
		assert.True(t, brand.UpdatedAt.After(originalUpdatedAt))
	})

	t.Run("should fail to delete already deleted brand", func(t *testing.T) {
		// Arrange
		brand := Create().Deleted()

		// Act
		err := brand.Delete()

		// Assert
		assert.Error(t, err)
		assert.Equal(t, exception.ErrBrandCannotDelete, err)
	})
}

func TestBrand_IsActive(t *testing.T) {
	t.Run("should return true for active brand", func(t *testing.T) {
		// Arrange
		brand := Create().WithDefaults()

		// Act & Assert
		assert.True(t, brand.IsActive())
	})

	t.Run("should return false for inactive brand", func(t *testing.T) {
		// Arrange
		brand := Create().Inactive()

		// Act & Assert
		assert.False(t, brand.IsActive())
	})

	t.Run("should return false for deleted brand", func(t *testing.T) {
		// Arrange
		brand := Create().Deleted()

		// Act & Assert
		assert.False(t, brand.IsActive())
	})
}

func TestBrand_IsDeleted(t *testing.T) {
	t.Run("should return true for deleted brand", func(t *testing.T) {
		// Arrange
		brand := Create().Deleted()

		// Act & Assert
		assert.True(t, brand.IsDeleted())
	})

	t.Run("should return false for active brand", func(t *testing.T) {
		// Arrange
		brand := Create().WithDefaults()

		// Act & Assert
		assert.False(t, brand.IsDeleted())
	})
}

func TestBrand_CanBeDeleted(t *testing.T) {
	t.Run("should return true for active brand", func(t *testing.T) {
		// Arrange
		brand := Create().WithDefaults()

		// Act & Assert
		assert.True(t, brand.CanBeDeleted())
	})

	t.Run("should return true for inactive brand", func(t *testing.T) {
		// Arrange
		brand := Create().Inactive()

		// Act & Assert
		assert.True(t, brand.CanBeDeleted())
	})

	t.Run("should return false for deleted brand", func(t *testing.T) {
		// Arrange
		brand := Create().Deleted()

		// Act & Assert
		assert.False(t, brand.CanBeDeleted())
	})
}

func TestBrand_ToReference(t *testing.T) {
	t.Run("should create brand reference", func(t *testing.T) {
		// Arrange
		brand := Create().Nike()

		// Act
		ref := brand.ToReference()

		// Assert
		assert.NotNil(t, ref)
		assert.Equal(t, brand.ID, ref.ID)
		assert.Equal(t, brand.Name, ref.Name)
		assert.Equal(t, brand.Description, ref.Description)
	})
}

func TestBrand_HasLogo(t *testing.T) {
	t.Run("should return true when brand has logo", func(t *testing.T) {
		// Arrange
		brand := Create().WithLogoURL("https://example.com/logo.png")

		// Act & Assert
		assert.True(t, brand.HasLogo())
	})

	t.Run("should return false when brand has no logo", func(t *testing.T) {
		// Arrange
		brand := Create().WithoutLogo()

		// Act & Assert
		assert.False(t, brand.HasLogo())
	})

	t.Run("should return false when logo URL is empty", func(t *testing.T) {
		// Arrange
		brand := Create().WithLogoURL("")

		// Act & Assert
		assert.False(t, brand.HasLogo())
	})
}

func TestBrand_HasWebsite(t *testing.T) {
	t.Run("should return true when brand has website", func(t *testing.T) {
		// Arrange
		brand := Create().WithWebsite("https://example.com")

		// Act & Assert
		assert.True(t, brand.HasWebsite())
	})

	t.Run("should return false when brand has no website", func(t *testing.T) {
		// Arrange
		brand := Create().WithoutWebsite()

		// Act & Assert
		assert.False(t, brand.HasWebsite())
	})

	t.Run("should return false when website is empty", func(t *testing.T) {
		// Arrange
		brand := Create().WithWebsite("")

		// Act & Assert
		assert.False(t, brand.HasWebsite())
	})
}
