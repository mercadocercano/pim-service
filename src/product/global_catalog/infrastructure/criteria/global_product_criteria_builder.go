package criteria

import (
	"github.com/gin-gonic/gin"

	"pim/src/shared/domain/criteria"
	sharedCriteria "pim/src/shared/infrastructure/criteria"
)

// GlobalProductCriteriaBuilder construye criterios específicos para productos globales
type GlobalProductCriteriaBuilder struct {
	helper  *sharedCriteria.EntityCriteriaHelper
	builder *criteria.CriteriaBuilder
}

// NewGlobalProductCriteriaBuilder crea un nuevo builder para criterios de productos globales
func NewGlobalProductCriteriaBuilder() *GlobalProductCriteriaBuilder {
	return &GlobalProductCriteriaBuilder{
		helper: sharedCriteria.NewEntityCriteriaHelper(),
	}
}

// FromContext construye criterios desde el contexto de Gin
func (b *GlobalProductCriteriaBuilder) FromContext(c *gin.Context) *GlobalProductCriteriaBuilder {
	b.builder = b.helper.BuildBaseFromContext(c)

	// Filtros específicos de productos globales
	b.builder.AddEqualFilter("source", c.Query("source"))
	b.builder.AddEqualFilter("business_type", c.Query("business_type"))
	b.builder.AddLikeFilter("name", c.Query("search_name"))
	b.builder.AddLikeFilter("brand", c.Query("search_brand"))
	b.builder.AddLikeFilter("category", c.Query("search_category"))
	b.builder.AddEqualFilter("ean", c.Query("ean"))
	
	// Filtros booleanos
	if c.Query("is_active") == "true" {
		b.builder.AddEqualFilter("is_active", "true")
	}
	if c.Query("is_verified") == "true" {
		b.builder.AddEqualFilter("is_verified", "true")
	}
	
	// Filtros de calidad
	if minQuality := c.Query("min_quality"); minQuality != "" {
		b.builder.AddFilter("quality_score", criteria.OpGreaterThanOrEqual, minQuality)
	}
	if maxQuality := c.Query("max_quality"); maxQuality != "" {
		b.builder.AddFilter("quality_score", criteria.OpLessThanOrEqual, maxQuality)
	}
	
	// Filtro de productos argentinos (si existe columna is_argentine)
	if c.Query("is_argentine") == "true" {
		// Esto podría requerir un filtro especial basado en el campo metadata
		// o una columna calculada, dependiendo de la implementación
		b.builder.AddEqualFilter("metadata->>'is_argentine'", "true")
	}

	return b
}

// Build construye los criterios finales
func (b *GlobalProductCriteriaBuilder) Build() criteria.Criteria {
	if b.builder == nil {
		// Si no se ha inicializado desde contexto, crear builder vacío
		b.builder = criteria.NewCriteriaBuilder()
	}
	return b.builder.Build()
}

// GetAllowedFields retorna los campos permitidos para filtrado de productos globales
func (b *GlobalProductCriteriaBuilder) GetAllowedFields() []string {
	return []string{
		"id", "ean", "name", "description", "brand", "category", 
		"price", "source", "quality_score", "is_verified", "is_active",
		"business_type", "created_at", "updated_at",
	}
}

// BuildValidated construye criterios validados desde el contexto
func (b *GlobalProductCriteriaBuilder) BuildValidated(c *gin.Context) criteria.Criteria {
	searchCriteria := b.FromContext(c).Build()
	return b.helper.ValidateAndSanitizeCriteria(searchCriteria, b.GetAllowedFields())
}