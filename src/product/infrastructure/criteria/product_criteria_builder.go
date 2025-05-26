package criteria

import (
	"github.com/gin-gonic/gin"

	"pim/src/shared/domain/criteria"
	sharedCriteria "pim/src/shared/infrastructure/criteria"
)

// ProductCriteriaBuilder construye criterios específicos para productos
type ProductCriteriaBuilder struct {
	helper  *sharedCriteria.EntityCriteriaHelper
	builder *criteria.CriteriaBuilder
}

// NewProductCriteriaBuilder crea una nueva instancia del builder
func NewProductCriteriaBuilder() *ProductCriteriaBuilder {
	return &ProductCriteriaBuilder{
		helper: sharedCriteria.NewEntityCriteriaHelper(),
	}
}

// FromContext construye criterios desde el contexto de Gin
func (b *ProductCriteriaBuilder) FromContext(c *gin.Context) *ProductCriteriaBuilder {
	b.builder = b.helper.BuildBaseFromContext(c)

	// Filtros específicos de productos
	b.builder.AddLikeFilter("name", c.Query("name"))
	b.builder.AddLikeFilter("description", c.Query("description"))
	b.builder.AddLikeFilter("sku", c.Query("sku"))
	b.builder.AddEqualFilter("status", c.Query("status"))
	b.builder.AddUUIDFilter("category_id", c.Query("category_id"))
	b.builder.AddUUIDFilter("brand_id", c.Query("brand_id"))
	b.builder.AddLikeFilter("category_name", c.Query("category_name"))
	b.builder.AddLikeFilter("brand_name", c.Query("brand_name"))

	// Filtro para excluir productos eliminados por defecto
	if c.Query("include_deleted") != "true" {
		b.builder.AddNotEqualFilter("status", "deleted")
	}

	return b
}

// BuildValidated construye y valida los criterios
func (b *ProductCriteriaBuilder) BuildValidated(c *gin.Context) criteria.Criteria {
	searchCriteria := b.FromContext(c).Build()
	return b.helper.ValidateAndSanitizeCriteria(searchCriteria, b.GetAllowedFields())
}

// Build construye los criterios sin validación adicional
func (b *ProductCriteriaBuilder) Build() criteria.Criteria {
	return b.builder.Build()
}

// GetAllowedFields retorna los campos permitidos para filtrado y ordenamiento
func (b *ProductCriteriaBuilder) GetAllowedFields() []string {
	return []string{
		"id",
		"tenant_id",
		"name",
		"description",
		"sku",
		"status",
		"category_id",
		"category_name",
		"brand_id",
		"brand_name",
		"created_at",
		"updated_at",
	}
}

// GetDefaultSortField retorna el campo de ordenamiento por defecto
func (b *ProductCriteriaBuilder) GetDefaultSortField() string {
	return "created_at"
}

// GetDefaultSortDirection retorna la dirección de ordenamiento por defecto
func (b *ProductCriteriaBuilder) GetDefaultSortDirection() string {
	return "desc"
}
