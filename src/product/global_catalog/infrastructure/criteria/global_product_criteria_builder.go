package criteria

import (
	cr "github.com/mercadocercano/criteria"

	"github.com/gin-gonic/gin"
)

// GlobalProductCriteriaBuilder construye criterios específicos para productos globales
type GlobalProductCriteriaBuilder struct {
	helper  *cr.EntityCriteriaHelper
	builder *cr.CriteriaBuilder
}

// NewGlobalProductCriteriaBuilder crea un nuevo builder para criterios de productos globales
func NewGlobalProductCriteriaBuilder() *GlobalProductCriteriaBuilder {
	return &GlobalProductCriteriaBuilder{
		helper: cr.NewEntityCriteriaHelper(),
	}
}

// FromContext construye criterios desde el contexto de Gin
func (b *GlobalProductCriteriaBuilder) FromContext(c *gin.Context) *GlobalProductCriteriaBuilder {
	b.builder = b.helper.BuildBaseFromContext(c)

	// Filtros específicos de productos globales
	if v := c.Query("source"); v != "" {
		b.builder.AddEqualFilter("source", v)
	}
	if v := c.Query("business_type"); v != "" {
		b.builder.AddEqualFilter("business_type", v)
	}
	if v := c.Query("search"); v != "" {
		b.builder.AddLikeFilter("name", v)
	}
	if v := c.Query("description"); v != "" {
		b.builder.AddLikeFilter("description", v)
	}

	// Filtros múltiples para marcas
	if brands := c.QueryArray("brand"); len(brands) > 0 {
		// Filtrar valores vacíos
		validBrands := make([]interface{}, 0)
		for _, brand := range brands {
			if brand != "" && brand != "all" {
				validBrands = append(validBrands, brand)
			}
		}
		if len(validBrands) > 0 {
			b.builder.AddInFilter("brand", validBrands)
		}
	} else if brand := c.Query("brand"); brand != "" && brand != "all" {
		// Compatibilidad con filtro único
		b.builder.AddLikeFilter("brand", brand)
	}

	// Filtros múltiples para categorías
	if categories := c.QueryArray("category"); len(categories) > 0 {
		// Filtrar valores vacíos
		validCategories := make([]interface{}, 0)
		for _, category := range categories {
			if category != "" && category != "all" {
				validCategories = append(validCategories, category)
			}
		}
		if len(validCategories) > 0 {
			b.builder.AddInFilter("category", validCategories)
		}
	} else if category := c.Query("category"); category != "" && category != "all" {
		// Compatibilidad con filtro único
		b.builder.AddLikeFilter("category", category)
	}

	// Filtros múltiples para fuentes
	if sources := c.QueryArray("source"); len(sources) > 0 {
		// Filtrar valores vacíos
		validSources := make([]interface{}, 0)
		for _, source := range sources {
			if source != "" && source != "all" {
				validSources = append(validSources, source)
			}
		}
		if len(validSources) > 0 {
			b.builder.AddInFilter("source", validSources)
		}
	}

	if v := c.Query("ean"); v != "" {
		b.builder.AddEqualFilter("ean", v)
	}

	// Filtros booleanos
	if isVerified := c.Query("is_verified"); isVerified != "" && isVerified != "all" {
		if isVerified == "true" {
			b.builder.AddEqualFilter("is_verified", true)
		} else if isVerified == "false" {
			b.builder.AddEqualFilter("is_verified", false)
		}
	}

	if hasImage := c.Query("has_image"); hasImage == "true" {
		b.builder.AddFilter("image_url", cr.OpIsNotNull, nil)
	}

	if isActive := c.Query("is_active"); isActive != "" && isActive != "all" {
		if isActive == "true" {
			b.builder.AddEqualFilter("is_active", true)
		} else if isActive == "false" {
			b.builder.AddEqualFilter("is_active", false)
		}
	}

	if isArgentine := c.Query("is_argentine"); isArgentine != "" && isArgentine != "all" {
		if isArgentine == "true" {
			b.builder.AddEqualFilter("is_argentine_product", true)
		} else if isArgentine == "false" {
			b.builder.AddEqualFilter("is_argentine_product", false)
		}
	}

	// Filtros numéricos
	if minQuality := c.Query("min_quality"); minQuality != "" {
		b.builder.AddGreaterThanOrEqualFilter("quality_score", minQuality)
	}

	if maxQuality := c.Query("max_quality"); maxQuality != "" {
		b.builder.AddLessThanOrEqualFilter("quality_score", maxQuality)
	}

	return b
}

// Build construye los criterios finales
func (b *GlobalProductCriteriaBuilder) Build() cr.Criteria {
	if b.builder == nil {
		// Si no se ha inicializado desde contexto, crear builder vacío
		b.builder = cr.NewCriteriaBuilder()
	}
	return b.builder.Build()
}

// GetAllowedFields retorna los campos permitidos para filtrado de productos globales
func (b *GlobalProductCriteriaBuilder) GetAllowedFields() []string {
	return []string{
		"id", "ean", "name", "description", "brand", "category",
		"price", "source", "quality_score", "is_verified", "is_active",
		"business_type", "created_at", "updated_at", "image_url",
	}
}

// BuildValidated construye criterios validados desde el contexto
func (b *GlobalProductCriteriaBuilder) BuildValidated(c *gin.Context) cr.Criteria {
	searchCriteria := b.FromContext(c).Build()
	return b.helper.ValidateAndSanitizeCriteria(searchCriteria, b.GetAllowedFields())
}
