package criteria

import (
	cr "github.com/mercadocercano/criteria"

	"github.com/gin-gonic/gin"
)

// BusinessTypeTemplateCriteriaBuilder construye criterios específicos para BusinessTypeTemplate
type BusinessTypeTemplateCriteriaBuilder struct {
	helper  *cr.EntityCriteriaHelper
	builder *cr.CriteriaBuilder
}

// NewBusinessTypeTemplateCriteriaBuilder crea una nueva instancia del builder
func NewBusinessTypeTemplateCriteriaBuilder() *BusinessTypeTemplateCriteriaBuilder {
	return &BusinessTypeTemplateCriteriaBuilder{
		helper: cr.NewEntityCriteriaHelper(),
	}
}

// FromContext construye criterios desde el contexto de Gin
func (b *BusinessTypeTemplateCriteriaBuilder) FromContext(c *gin.Context) *BusinessTypeTemplateCriteriaBuilder {
	// Construir criterios base (paginación, ordenamiento)
	b.builder = b.helper.BuildBaseFromContext(c)

	// Filtros específicos de BusinessTypeTemplate
	b.addTemplateFilters(c)

	return b
}

// BuildValidated construye y valida los criterios
func (b *BusinessTypeTemplateCriteriaBuilder) BuildValidated(c *gin.Context) cr.Criteria {
	searchCriteria := b.FromContext(c).Build()
	return b.helper.ValidateAndSanitizeCriteria(searchCriteria, b.GetAllowedFields())
}

// Build construye los criterios sin validación
func (b *BusinessTypeTemplateCriteriaBuilder) Build() cr.Criteria {
	return b.builder.Build()
}

// addTemplateFilters agrega filtros específicos de BusinessTypeTemplate
func (b *BusinessTypeTemplateCriteriaBuilder) addTemplateFilters(c *gin.Context) {
	// Filtro por business_type_id
	if businessTypeID := c.Query("business_type_id"); businessTypeID != "" {
		b.builder.AddUUIDFilter("business_type_id", businessTypeID)
	}

	// Filtro por región
	if region := c.Query("region"); region != "" {
		b.builder.AddEqualFilter("region", region)
	}

	// Filtro por nombre (búsqueda LIKE)
	if name := c.Query("name"); name != "" {
		b.builder.AddLikeFilter("name", name)
	}

	// Filtro por descripción (búsqueda LIKE)
	if description := c.Query("description"); description != "" {
		b.builder.AddLikeFilter("description", description)
	}

	// Filtro para búsqueda general (nombre o descripción)
	if search := c.Query("search"); search != "" {
		// Para búsqueda general, usaremos un filtro personalizado
		// que busque en múltiples campos
		b.builder.AddLikeFilter("name", search)
		b.builder.AddLikeFilter("description", search)
	}

	// Filtro por estado activo/inactivo
	if isActive := c.Query("is_active"); isActive != "" {
		b.builder.AddBoolFilter("is_active", isActive)
	}

	// Filtro por template por defecto
	if isDefault := c.Query("is_default"); isDefault != "" {
		b.builder.AddBoolFilter("is_default", isDefault)
	}

	// Filtro para incluir solo templates activos (por defecto) - solo si no se especifica is_active explícitamente
	if isActive := c.Query("is_active"); isActive == "" {
		if includeInactive := c.Query("include_inactive"); includeInactive != "true" {
			b.builder.AddEqualFilter("is_active", true)
		}
	}

	// Filtro por versión
	if version := c.Query("version"); version != "" {
		b.builder.AddEqualFilter("version", version)
	}
}

// GetAllowedFields retorna los campos permitidos para filtrado y ordenamiento
func (b *BusinessTypeTemplateCriteriaBuilder) GetAllowedFields() []string {
	return []string{
		"id",
		"business_type_id",
		"name",
		"description",
		"version",
		"region",
		"is_active",
		"is_default",
		"created_at",
		"updated_at",
	}
}

// GetDefaultOrderField retorna el campo de ordenamiento por defecto
func (b *BusinessTypeTemplateCriteriaBuilder) GetDefaultOrderField() string {
	return "is_default"
}

// GetDefaultOrderDirection retorna la dirección de ordenamiento por defecto
func (b *BusinessTypeTemplateCriteriaBuilder) GetDefaultOrderDirection() string {
	return "DESC" // Templates por defecto primero
}
