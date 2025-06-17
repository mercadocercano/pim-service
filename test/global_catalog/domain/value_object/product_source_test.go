package value_object_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"pim/src/global_catalog/domain/value_object"
)

func TestNewProductSource(t *testing.T) {
	t.Run("should create valid product source", func(t *testing.T) {
		// Arrange
		source := "manual"
		sourceURL := "https://example.com"
		scrapedAt := time.Now()
		reliability := 0.8

		// Act
		ps, err := value_object.NewProductSource(source, &sourceURL, &scrapedAt, reliability)

		// Assert
		require.NoError(t, err)
		assert.NotNil(t, ps)
		assert.Equal(t, source, ps.Source())
		assert.Equal(t, &sourceURL, ps.SourceURL())
		assert.Equal(t, &scrapedAt, ps.ScrapedAt())
		assert.Equal(t, reliability, ps.Reliability())
	})

	t.Run("should create product source with nil optional fields", func(t *testing.T) {
		// Arrange
		source := "manual"
		reliability := 1.0

		// Act
		ps, err := value_object.NewProductSource(source, nil, nil, reliability)

		// Assert
		require.NoError(t, err)
		assert.NotNil(t, ps)
		assert.Equal(t, source, ps.Source())
		assert.Nil(t, ps.SourceURL())
		assert.Nil(t, ps.ScrapedAt())
		assert.Equal(t, reliability, ps.Reliability())
	})

	t.Run("should fail with empty source", func(t *testing.T) {
		// Act
		ps, err := value_object.NewProductSource("", nil, nil, 1.0)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, ps)
		assert.Contains(t, err.Error(), "fuente del producto es obligatoria")
	})

	t.Run("should fail with invalid source", func(t *testing.T) {
		// Act
		ps, err := value_object.NewProductSource("invalid_source", nil, nil, 1.0)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, ps)
		assert.Contains(t, err.Error(), "fuente de producto no válida")
	})

	t.Run("should fail with reliability below 0", func(t *testing.T) {
		// Act
		ps, err := value_object.NewProductSource("manual", nil, nil, -0.1)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, ps)
		assert.Contains(t, err.Error(), "confiabilidad debe estar entre 0 y 1")
	})

	t.Run("should fail with reliability above 1", func(t *testing.T) {
		// Act
		ps, err := value_object.NewProductSource("manual", nil, nil, 1.1)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, ps)
		assert.Contains(t, err.Error(), "confiabilidad debe estar entre 0 y 1")
	})

	t.Run("should accept all valid sources", func(t *testing.T) {
		validSources := []string{
			"manual", "disco", "carrefour", "fravega", "coto", "jumbo", "api",
		}

		for _, source := range validSources {
			t.Run("Valid source: "+source, func(t *testing.T) {
				// Act
				ps, err := value_object.NewProductSource(source, nil, nil, 0.8)

				// Assert
				require.NoError(t, err)
				assert.Equal(t, source, ps.Source())
			})
		}
	})
}

func TestNewManualSource(t *testing.T) {
	t.Run("should create manual source with high reliability", func(t *testing.T) {
		// Act
		ps, err := value_object.NewManualSource()

		// Assert
		require.NoError(t, err)
		assert.Equal(t, "manual", ps.Source())
		assert.Nil(t, ps.SourceURL())
		assert.Nil(t, ps.ScrapedAt())
		assert.Equal(t, 1.0, ps.Reliability())
	})
}

func TestNewScrapingSource(t *testing.T) {
	t.Run("should create scraping source with current timestamp", func(t *testing.T) {
		// Arrange
		source := "disco"
		sourceURL := "https://disco.com.ar/product/123"
		reliability := 0.8
		beforeTime := time.Now()

		// Act
		ps, err := value_object.NewScrapingSource(source, sourceURL, reliability)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, source, ps.Source())
		assert.Equal(t, &sourceURL, ps.SourceURL())
		assert.NotNil(t, ps.ScrapedAt())
		assert.True(t, ps.ScrapedAt().After(beforeTime) || ps.ScrapedAt().Equal(beforeTime))
		assert.Equal(t, reliability, ps.Reliability())
	})
}

func TestNewAPISource(t *testing.T) {
	t.Run("should create API source", func(t *testing.T) {
		// Arrange
		sourceURL := "https://api.example.com/products"
		reliability := 0.9

		// Act
		ps, err := value_object.NewAPISource(sourceURL, reliability)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, "api", ps.Source())
		assert.Equal(t, &sourceURL, ps.SourceURL())
		assert.Nil(t, ps.ScrapedAt())
		assert.Equal(t, reliability, ps.Reliability())
	})
}

func TestProductSource_IsManual(t *testing.T) {
	t.Run("should return true for manual source", func(t *testing.T) {
		// Arrange
		ps, _ := value_object.NewManualSource()

		// Act & Assert
		assert.True(t, ps.IsManual())
	})

	t.Run("should return false for non-manual source", func(t *testing.T) {
		// Arrange
		ps, _ := value_object.NewScrapingSource("disco", "https://disco.com.ar", 0.8)

		// Act & Assert
		assert.False(t, ps.IsManual())
	})
}

func TestProductSource_IsScraped(t *testing.T) {
	t.Run("should return true for scraped sources", func(t *testing.T) {
		scrapedSources := []string{"disco", "carrefour", "fravega", "coto", "jumbo"}

		for _, source := range scrapedSources {
			t.Run("Scraped source: "+source, func(t *testing.T) {
				// Arrange
				ps, _ := value_object.NewScrapingSource(source, "https://example.com", 0.8)

				// Act & Assert
				assert.True(t, ps.IsScraped())
			})
		}
	})

	t.Run("should return false for non-scraped sources", func(t *testing.T) {
		// Arrange
		manualPS, _ := value_object.NewManualSource()
		apiPS, _ := value_object.NewAPISource("https://api.example.com", 0.9)

		// Act & Assert
		assert.False(t, manualPS.IsScraped())
		assert.False(t, apiPS.IsScraped())
	})
}

func TestProductSource_IsFromAPI(t *testing.T) {
	t.Run("should return true for API source", func(t *testing.T) {
		// Arrange
		ps, _ := value_object.NewAPISource("https://api.example.com", 0.9)

		// Act & Assert
		assert.True(t, ps.IsFromAPI())
	})

	t.Run("should return false for non-API source", func(t *testing.T) {
		// Arrange
		manualPS, _ := value_object.NewManualSource()
		scrapedPS, _ := value_object.NewScrapingSource("disco", "https://disco.com.ar", 0.8)

		// Act & Assert
		assert.False(t, manualPS.IsFromAPI())
		assert.False(t, scrapedPS.IsFromAPI())
	})
}

func TestProductSource_IsHighReliability(t *testing.T) {
	t.Run("should return true for high reliability (>= 0.7)", func(t *testing.T) {
		reliabilities := []float64{0.7, 0.8, 0.9, 1.0}

		for _, reliability := range reliabilities {
			t.Run("Reliability: "+fmt.Sprintf("%.1f", reliability), func(t *testing.T) {
				// Arrange
				ps, _ := value_object.NewProductSource("manual", nil, nil, reliability)

				// Act & Assert
				assert.True(t, ps.IsHighReliability())
			})
		}
	})

	t.Run("should return false for low reliability (< 0.7)", func(t *testing.T) {
		reliabilities := []float64{0.0, 0.3, 0.5, 0.69}

		for _, reliability := range reliabilities {
			t.Run("Reliability: "+fmt.Sprintf("%.2f", reliability), func(t *testing.T) {
				// Arrange
				ps, _ := value_object.NewProductSource("manual", nil, nil, reliability)

				// Act & Assert
				assert.False(t, ps.IsHighReliability())
			})
		}
	})
}

func TestProductSource_NeedsUpdate(t *testing.T) {
	t.Run("should return true when scraped source is old", func(t *testing.T) {
		// Arrange
		oldTime := time.Now().Add(-25 * time.Hour)
		ps, _ := value_object.NewProductSource("disco", nil, &oldTime, 0.8)

		// Act & Assert
		assert.True(t, ps.NeedsUpdate(24*time.Hour))
	})

	t.Run("should return false when scraped source is recent", func(t *testing.T) {
		// Arrange
		recentTime := time.Now().Add(-1 * time.Hour)
		ps, _ := value_object.NewProductSource("disco", nil, &recentTime, 0.8)

		// Act & Assert
		assert.False(t, ps.NeedsUpdate(24*time.Hour))
	})

	t.Run("should return false for non-scraped sources", func(t *testing.T) {
		// Arrange
		manualPS, _ := value_object.NewManualSource()
		apiPS, _ := value_object.NewAPISource("https://api.example.com", 0.9)

		// Act & Assert
		assert.False(t, manualPS.NeedsUpdate(24*time.Hour))
		assert.False(t, apiPS.NeedsUpdate(24*time.Hour))
	})

	t.Run("should return false when scraped at is nil", func(t *testing.T) {
		// Arrange
		ps, _ := value_object.NewProductSource("disco", nil, nil, 0.8)

		// Act & Assert
		assert.False(t, ps.NeedsUpdate(24*time.Hour))
	})
}

func TestProductSource_GetSourceDisplayName(t *testing.T) {
	t.Run("should return correct display names", func(t *testing.T) {
		testCases := []struct {
			source      string
			displayName string
		}{
			{"manual", "Manual"},
			{"disco", "Disco Argentina"},
			{"carrefour", "Carrefour Argentina"},
			{"fravega", "Fravega"},
			{"coto", "Coto Digital"},
			{"jumbo", "Jumbo Argentina"},
			{"api", "API Externa"},
		}

		for _, tc := range testCases {
			t.Run("Source: "+tc.source, func(t *testing.T) {
				// Arrange
				ps, _ := value_object.NewProductSource(tc.source, nil, nil, 0.8)

				// Act & Assert
				assert.Equal(t, tc.displayName, ps.GetSourceDisplayName())
			})
		}
	})
}

func TestProductSource_Equals(t *testing.T) {
	t.Run("should return true for equal sources", func(t *testing.T) {
		// Arrange
		sourceURL := "https://example.com"
		ps1, _ := value_object.NewProductSource("manual", &sourceURL, nil, 0.8)
		ps2, _ := value_object.NewProductSource("manual", &sourceURL, nil, 0.9) // reliability diferente

		// Act & Assert
		assert.True(t, ps1.Equals(ps2))
	})

	t.Run("should return false for different sources", func(t *testing.T) {
		// Arrange
		ps1, _ := value_object.NewManualSource()
		ps2, _ := value_object.NewScrapingSource("disco", "https://disco.com.ar", 0.8)

		// Act & Assert
		assert.False(t, ps1.Equals(ps2))
	})

	t.Run("should return false when comparing with nil", func(t *testing.T) {
		// Arrange
		ps, _ := value_object.NewManualSource()

		// Act & Assert
		assert.False(t, ps.Equals(nil))
	})

	t.Run("should return true for same source with different URLs", func(t *testing.T) {
		// Arrange
		url1 := "https://example1.com"
		url2 := "https://example2.com"
		ps1, _ := value_object.NewProductSource("api", &url1, nil, 0.8)
		ps2, _ := value_object.NewProductSource("api", &url2, nil, 0.8)

		// Act & Assert
		assert.False(t, ps1.Equals(ps2)) // URLs diferentes
	})
}

func TestProductSource_String(t *testing.T) {
	t.Run("should return display name as string", func(t *testing.T) {
		// Arrange
		ps, _ := value_object.NewManualSource()

		// Act & Assert
		assert.Equal(t, "Manual", ps.String())
	})
}
