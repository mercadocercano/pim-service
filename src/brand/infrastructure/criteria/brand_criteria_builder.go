package criteria

import (
	cr "github.com/hornosg/go-shared/criteria"

	"github.com/gin-gonic/gin"
)

// BrandCriteriaBuilder construye criterios específicos para Brand
type BrandCriteriaBuilder struct {
	helper  *cr.EntityCriteriaHelper
	builder *cr.CriteriaBuilder
}

// NewBrandCriteriaBuilder crea una nueva instancia del builder
func NewBrandCriteriaBuilder() *BrandCriteriaBuilder {
	return &BrandCriteriaBuilder{
		helper: cr.NewEntityCriteriaHelper(),
	}
}

// FromContext construye criterios desde el contexto de Gin
func (b *BrandCriteriaBuilder) FromContext(c *gin.Context) *BrandCriteriaBuilder {
	// Construir criterios base (paginación, ordenamiento)
	b.builder = b.helper.BuildBaseFromContext(c)

	// Filtros específicos de Brand
	b.addBrandFilters(c)

	return b
}

// BuildValidated construye y valida los criterios
func (b *BrandCriteriaBuilder) BuildValidated(c *gin.Context) cr.Criteria {
	searchCriteria := b.FromContext(c).Build()
	return b.helper.ValidateAndSanitizeCriteria(searchCriteria, b.GetAllowedFields())
}

// Build construye los criterios sin validación
func (b *BrandCriteriaBuilder) Build() cr.Criteria {
	return b.builder.Build()
}

// addBrandFilters agrega filtros específicos de Brand
func (b *BrandCriteriaBuilder) addBrandFilters(c *gin.Context) {
	// Filtro por tenant_id (obligatorio)
	if tenantID := c.GetHeader("X-Tenant-ID"); tenantID != "" {
		b.builder.AddUUIDFilter("tenant_id", tenantID)
	}

	// Filtro por status
	if status := c.Query("status"); status != "" {
		b.builder.AddEqualFilter("status", status)
	}

	// Filtro por nombre (búsqueda LIKE)
	if name := c.Query("name"); name != "" {
		b.builder.AddLikeFilter("name", name)
	}

	// Filtro por descripción (búsqueda LIKE)
	if description := c.Query("description"); description != "" {
		b.builder.AddLikeFilter("description", description)
	}

	// Filtro para marcas activas solamente
	if active := c.Query("active"); active == "true" {
		b.builder.AddEqualFilter("status", "active")
	}

	// Filtro para excluir marcas eliminadas (por defecto)
	if includeDeleted := c.Query("include_deleted"); includeDeleted != "true" {
		b.builder.AddNotEqualFilter("status", "deleted")
	}

	// Filtro por website (si tiene sitio web)
	if hasWebsite := c.Query("has_website"); hasWebsite == "true" {
		b.builder.AddFilter("website", cr.OpIsNotNull, nil)
	} else if hasWebsite == "false" {
		b.builder.AddFilter("website", cr.OpIsNull, nil)
	}

	// Filtro por logo (si tiene logo)
	if hasLogo := c.Query("has_logo"); hasLogo == "true" {
		b.builder.AddFilter("logo_url", cr.OpIsNotNull, nil)
	} else if hasLogo == "false" {
		b.builder.AddFilter("logo_url", cr.OpIsNull, nil)
	}
}

// GetAllowedFields retorna los campos permitidos para filtrado y ordenamiento
func (b *BrandCriteriaBuilder) GetAllowedFields() []string {
	return []string{
		"id",
		"tenant_id",
		"name",
		"description",
		"logo_url",
		"website",
		"status",
		"created_at",
		"updated_at",
	}
}

// GetDefaultOrderField retorna el campo de ordenamiento por defecto
func (b *BrandCriteriaBuilder) GetDefaultOrderField() string {
	return "name"
}

// GetDefaultOrderDirection retorna la dirección de ordenamiento por defecto
func (b *BrandCriteriaBuilder) GetDefaultOrderDirection() string {
	return "ASC"
}
