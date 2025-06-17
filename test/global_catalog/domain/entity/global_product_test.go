package entity_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"pim/src/global_catalog/domain/entity"
	"pim/src/global_catalog/domain/value_object"
)

func TestNewGlobalProduct(t *testing.T) {
	t.Run("should create global product with valid data", func(t *testing.T) {
		// Arrange
		ean, _ := value_object.NewEAN13("7790001234568")
		source, _ := value_object.NewManualSource()
		name := "Producto Test"
		description := "Descripción del producto"
		brand := "Marca Test"
		category := "Categoría Test"
		price := 100.50
		imageURL := "https://example.com/image.jpg"

		// Act
		product, err := entity.NewGlobalProduct(
			ean,
			name,
			&description,
			&brand,
			&category,
			&price,
			&imageURL,
			source,
		)

		// Assert
		require.NoError(t, err)
		assert.NotNil(t, product)
		assert.NotEqual(t, uuid.Nil, product.ID())
		assert.Equal(t, ean, product.EAN())
		assert.Equal(t, name, product.Name())
		assert.Equal(t, &description, product.Description())
		assert.Equal(t, &brand, product.Brand())
		assert.Equal(t, &category, product.Category())
		assert.Equal(t, &price, product.Price())
		assert.Equal(t, &imageURL, product.ImageURL())
		assert.Equal(t, source, product.Source())
		assert.NotNil(t, product.QualityScore())
		assert.False(t, product.IsVerified())
		assert.True(t, product.IsActive())
		assert.Empty(t, product.Tags())
		assert.NotNil(t, product.Metadata())
		assert.WithinDuration(t, time.Now(), product.CreatedAt(), time.Second)
		assert.WithinDuration(t, time.Now(), product.UpdatedAt(), time.Second)
	})

	t.Run("should create global product with minimal data", func(t *testing.T) {
		// Arrange
		ean, _ := value_object.NewEAN13("7790001234568")
		source, _ := value_object.NewManualSource()
		name := "Producto Mínimo"

		// Act
		product, err := entity.NewGlobalProduct(
			ean,
			name,
			nil, // description
			nil, // brand
			nil, // category
			nil, // price
			nil, // imageURL
			source,
		)

		// Assert
		require.NoError(t, err)
		assert.NotNil(t, product)
		assert.Equal(t, name, product.Name())
		assert.Nil(t, product.Description())
		assert.Nil(t, product.Brand())
		assert.Nil(t, product.Category())
		assert.Nil(t, product.Price())
		assert.Nil(t, product.ImageURL())
		assert.Empty(t, product.ImageURLs())
	})

	t.Run("should fail when EAN is nil", func(t *testing.T) {
		// Arrange
		source, _ := value_object.NewManualSource()
		name := "Producto Test"

		// Act
		product, err := entity.NewGlobalProduct(
			nil, // EAN nil
			name,
			nil,
			nil,
			nil,
			nil,
			nil,
			source,
		)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, product)
		assert.Contains(t, err.Error(), "EAN-13 es obligatorio")
	})

	t.Run("should fail when name is empty", func(t *testing.T) {
		// Arrange
		ean, _ := value_object.NewEAN13("7790001234568")
		source, _ := value_object.NewManualSource()

		// Act
		product, err := entity.NewGlobalProduct(
			ean,
			"", // nombre vacío
			nil,
			nil,
			nil,
			nil,
			nil,
			source,
		)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, product)
		assert.Contains(t, err.Error(), "nombre del producto es obligatorio")
	})

	t.Run("should fail when source is nil", func(t *testing.T) {
		// Arrange
		ean, _ := value_object.NewEAN13("7790001234568")
		name := "Producto Test"

		// Act
		product, err := entity.NewGlobalProduct(
			ean,
			name,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil, // source nil
		)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, product)
		assert.Contains(t, err.Error(), "fuente del producto es obligatoria")
	})

	t.Run("should trim whitespace from name", func(t *testing.T) {
		// Arrange
		ean, _ := value_object.NewEAN13("7790001234568")
		source, _ := value_object.NewManualSource()
		name := "  Producto con espacios  "

		// Act
		product, err := entity.NewGlobalProduct(
			ean,
			name,
			nil,
			nil,
			nil,
			nil,
			nil,
			source,
		)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, "Producto con espacios", product.Name())
	})

	t.Run("should add main image to image URLs list", func(t *testing.T) {
		// Arrange
		ean, _ := value_object.NewEAN13("7790001234568")
		source, _ := value_object.NewManualSource()
		name := "Producto Test"
		imageURL := "https://example.com/image.jpg"

		// Act
		product, err := entity.NewGlobalProduct(
			ean,
			name,
			nil,
			nil,
			nil,
			nil,
			&imageURL,
			source,
		)

		// Assert
		require.NoError(t, err)
		assert.Len(t, product.ImageURLs(), 1)
		assert.Contains(t, product.ImageURLs(), imageURL)
	})
}

func TestGlobalProduct_UpdateFromScraping(t *testing.T) {
	t.Run("should update product with new scraping data", func(t *testing.T) {
		// Arrange
		ean, _ := value_object.NewEAN13("7790001234568")
		originalSource, _ := value_object.NewManualSource()
		product, _ := entity.NewGlobalProduct(
			ean,
			"Producto Original",
			nil,
			nil,
			nil,
			nil,
			nil,
			originalSource,
		)

		newSource, _ := value_object.NewScrapingSource("disco", "https://disco.com.ar/product/123", 0.8)
		newName := "Producto Actualizado"
		newDescription := "Nueva descripción"
		newBrand := "Nueva Marca"
		newCategory := "Nueva Categoría"
		newPrice := 200.75
		newImageURL := "https://example.com/new-image.jpg"

		// Act
		err := product.UpdateFromScraping(
			newName,
			&newDescription,
			&newBrand,
			&newCategory,
			&newPrice,
			&newImageURL,
			newSource,
		)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, newName, product.Name())
		assert.Equal(t, &newDescription, product.Description())
		assert.Equal(t, &newBrand, product.Brand())
		assert.Equal(t, &newCategory, product.Category())
		assert.Equal(t, &newPrice, product.Price())
		assert.Equal(t, &newImageURL, product.ImageURL())
		assert.Equal(t, newSource, product.Source())
		assert.NotNil(t, product.LastScrapedAt())
	})

	t.Run("should fail when name is empty", func(t *testing.T) {
		// Arrange
		ean, _ := value_object.NewEAN13("7790001234568")
		source, _ := value_object.NewManualSource()
		product, _ := entity.NewGlobalProduct(ean, "Producto Original", nil, nil, nil, nil, nil, source)

		newSource, _ := value_object.NewScrapingSource("disco", "https://disco.com.ar/product/123", 0.8)

		// Act
		err := product.UpdateFromScraping(
			"", // nombre vacío
			nil,
			nil,
			nil,
			nil,
			nil,
			newSource,
		)

		// Assert
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "nombre del producto es obligatorio")
	})

	t.Run("should fail when source is nil", func(t *testing.T) {
		// Arrange
		ean, _ := value_object.NewEAN13("7790001234568")
		source, _ := value_object.NewManualSource()
		product, _ := entity.NewGlobalProduct(ean, "Producto Original", nil, nil, nil, nil, nil, source)

		// Act
		err := product.UpdateFromScraping(
			"Producto Actualizado",
			nil,
			nil,
			nil,
			nil,
			nil,
			nil, // source nil
		)

		// Assert
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "fuente del producto es obligatoria")
	})
}

func TestGlobalProduct_Verify(t *testing.T) {
	t.Run("should verify product", func(t *testing.T) {
		// Arrange
		ean, _ := value_object.NewEAN13("7790001234568")
		source, _ := value_object.NewManualSource()
		product, _ := entity.NewGlobalProduct(ean, "Producto Test", nil, nil, nil, nil, nil, source)

		// Act
		product.Verify()

		// Assert
		assert.True(t, product.IsVerified())
	})

	t.Run("should unverify product", func(t *testing.T) {
		// Arrange
		ean, _ := value_object.NewEAN13("7790001234568")
		source, _ := value_object.NewManualSource()
		product, _ := entity.NewGlobalProduct(ean, "Producto Test", nil, nil, nil, nil, nil, source)
		product.Verify()

		// Act
		product.Unverify()

		// Assert
		assert.False(t, product.IsVerified())
	})
}

func TestGlobalProduct_Activate(t *testing.T) {
	t.Run("should activate product", func(t *testing.T) {
		// Arrange
		ean, _ := value_object.NewEAN13("7790001234568")
		source, _ := value_object.NewManualSource()
		product, _ := entity.NewGlobalProduct(ean, "Producto Test", nil, nil, nil, nil, nil, source)
		product.Deactivate()

		// Act
		product.Activate()

		// Assert
		assert.True(t, product.IsActive())
	})

	t.Run("should deactivate product", func(t *testing.T) {
		// Arrange
		ean, _ := value_object.NewEAN13("7790001234568")
		source, _ := value_object.NewManualSource()
		product, _ := entity.NewGlobalProduct(ean, "Producto Test", nil, nil, nil, nil, nil, source)

		// Act
		product.Deactivate()

		// Assert
		assert.False(t, product.IsActive())
	})
}

func TestGlobalProduct_BusinessType(t *testing.T) {
	t.Run("should set business type", func(t *testing.T) {
		// Arrange
		ean, _ := value_object.NewEAN13("7790001234568")
		source, _ := value_object.NewManualSource()
		product, _ := entity.NewGlobalProduct(ean, "Producto Test", nil, nil, nil, nil, nil, source)
		businessType := "supermercado"

		// Act
		product.SetBusinessType(businessType)

		// Assert
		assert.Equal(t, &businessType, product.BusinessType())
	})
}

func TestGlobalProduct_Tags(t *testing.T) {
	t.Run("should add tag", func(t *testing.T) {
		// Arrange
		ean, _ := value_object.NewEAN13("7790001234568")
		source, _ := value_object.NewManualSource()
		product, _ := entity.NewGlobalProduct(ean, "Producto Test", nil, nil, nil, nil, nil, source)
		tag := "promoción"

		// Act
		product.AddTag(tag)

		// Assert
		assert.Contains(t, product.Tags(), tag)
	})

	t.Run("should not add duplicate tag", func(t *testing.T) {
		// Arrange
		ean, _ := value_object.NewEAN13("7790001234568")
		source, _ := value_object.NewManualSource()
		product, _ := entity.NewGlobalProduct(ean, "Producto Test", nil, nil, nil, nil, nil, source)
		tag := "promoción"

		// Act
		product.AddTag(tag)
		product.AddTag(tag) // agregar duplicado

		// Assert
		count := 0
		for _, t := range product.Tags() {
			if t == tag {
				count++
			}
		}
		assert.Equal(t, 1, count)
	})

	t.Run("should remove tag", func(t *testing.T) {
		// Arrange
		ean, _ := value_object.NewEAN13("7790001234568")
		source, _ := value_object.NewManualSource()
		product, _ := entity.NewGlobalProduct(ean, "Producto Test", nil, nil, nil, nil, nil, source)
		tag := "promoción"
		product.AddTag(tag)

		// Act
		product.RemoveTag(tag)

		// Assert
		assert.NotContains(t, product.Tags(), tag)
	})
}

func TestGlobalProduct_Images(t *testing.T) {
	t.Run("should add image", func(t *testing.T) {
		// Arrange
		ean, _ := value_object.NewEAN13("7790001234568")
		source, _ := value_object.NewManualSource()
		product, _ := entity.NewGlobalProduct(ean, "Producto Test", nil, nil, nil, nil, nil, source)
		imageURL := "https://example.com/new-image.jpg"

		// Act
		product.AddImage(imageURL)

		// Assert
		assert.Contains(t, product.ImageURLs(), imageURL)
	})

	t.Run("should not add duplicate image", func(t *testing.T) {
		// Arrange
		ean, _ := value_object.NewEAN13("7790001234568")
		source, _ := value_object.NewManualSource()
		imageURL := "https://example.com/image.jpg"
		product, _ := entity.NewGlobalProduct(ean, "Producto Test", nil, nil, nil, nil, &imageURL, source)

		// Act
		product.AddImage(imageURL) // agregar duplicado

		// Assert
		count := 0
		for _, img := range product.ImageURLs() {
			if img == imageURL {
				count++
			}
		}
		assert.Equal(t, 1, count)
	})
}

func TestGlobalProduct_Metadata(t *testing.T) {
	t.Run("should set and get metadata", func(t *testing.T) {
		// Arrange
		ean, _ := value_object.NewEAN13("7790001234568")
		source, _ := value_object.NewManualSource()
		product, _ := entity.NewGlobalProduct(ean, "Producto Test", nil, nil, nil, nil, nil, source)
		key := "test_key"
		value := "test_value"

		// Act
		product.SetMetadata(key, value)
		retrievedValue := product.GetMetadata(key)

		// Assert
		assert.Equal(t, value, retrievedValue)
	})

	t.Run("should return nil for non-existent metadata key", func(t *testing.T) {
		// Arrange
		ean, _ := value_object.NewEAN13("7790001234568")
		source, _ := value_object.NewManualSource()
		product, _ := entity.NewGlobalProduct(ean, "Producto Test", nil, nil, nil, nil, nil, source)

		// Act
		value := product.GetMetadata("non_existent_key")

		// Assert
		assert.Nil(t, value)
	})
}

func TestGlobalProduct_IsArgentineProduct(t *testing.T) {
	t.Run("should return true for Argentine EAN", func(t *testing.T) {
		// Arrange
		ean, _ := value_object.NewEAN13("7790001234568") // EAN argentino
		source, _ := value_object.NewManualSource()
		product, _ := entity.NewGlobalProduct(ean, "Producto Test", nil, nil, nil, nil, nil, source)

		// Act & Assert
		assert.True(t, product.IsArgentineProduct())
	})

	t.Run("should return false for non-Argentine EAN", func(t *testing.T) {
		// Arrange
		ean, _ := value_object.NewEAN13("1234567890128") // EAN no argentino válido
		source, _ := value_object.NewManualSource()
		product, _ := entity.NewGlobalProduct(ean, "Producto Test", nil, nil, nil, nil, nil, source)

		// Act & Assert
		assert.False(t, product.IsArgentineProduct())
	})
}

func TestGlobalProduct_HasCompleteData(t *testing.T) {
	t.Run("should return true for complete data", func(t *testing.T) {
		// Arrange
		ean, _ := value_object.NewEAN13("7790001234568")
		source, _ := value_object.NewManualSource()
		description := "Descripción completa"
		brand := "Marca"
		category := "Categoría"
		price := 100.0
		imageURL := "https://example.com/image.jpg"

		product, _ := entity.NewGlobalProduct(
			ean,
			"Producto Completo",
			&description,
			&brand,
			&category,
			&price,
			&imageURL,
			source,
		)

		// Act & Assert
		assert.True(t, product.HasCompleteData())
	})

	t.Run("should return false for incomplete data", func(t *testing.T) {
		// Arrange
		ean, _ := value_object.NewEAN13("7790001234568")
		source, _ := value_object.NewManualSource()
		product, _ := entity.NewGlobalProduct(ean, "Producto Incompleto", nil, nil, nil, nil, nil, source)

		// Act & Assert
		assert.False(t, product.HasCompleteData())
	})
}

func TestGlobalProduct_NeedsUpdate(t *testing.T) {
	t.Run("should return true when product needs update", func(t *testing.T) {
		// Arrange
		ean, _ := value_object.NewEAN13("7790001234568")
		oldTime := time.Now().Add(-25 * time.Hour) // hace 25 horas
		source, _ := value_object.NewProductSource("disco", nil, &oldTime, 0.8)
		product, _ := entity.NewGlobalProduct(ean, "Producto Test", nil, nil, nil, nil, nil, source)

		// Act & Assert
		assert.True(t, product.NeedsUpdate(24*time.Hour)) // máximo 24 horas
	})

	t.Run("should return false when product is recent", func(t *testing.T) {
		// Arrange
		ean, _ := value_object.NewEAN13("7790001234568")
		recentTime := time.Now().Add(-1 * time.Hour) // hace 1 hora
		source, _ := value_object.NewProductSource("disco", nil, &recentTime, 0.8)
		product, _ := entity.NewGlobalProduct(ean, "Producto Test", nil, nil, nil, nil, nil, source)

		// Act & Assert
		assert.False(t, product.NeedsUpdate(24*time.Hour)) // máximo 24 horas
	})
}
