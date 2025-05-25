package criteria

import (
	"net/url"

	domainCriteria "pim/src/shared/domain/criteria"
	sharedCriteria "pim/src/shared/infrastructure/criteria"

	"github.com/gin-gonic/gin"
)

// CategoryCriteriaBuilder construye criterios específicos para categorías
type CategoryCriteriaBuilder struct {
	*domainCriteria.CriteriaBuilder
	helper *sharedCriteria.EntityCriteriaHelper
}

// NewCategoryCriteriaBuilder crea un nuevo builder para criterios de categorías
func NewCategoryCriteriaBuilder() *CategoryCriteriaBuilder {
	return &CategoryCriteriaBuilder{
		CriteriaBuilder: domainCriteria.NewCriteriaBuilder(),
		helper:          sharedCriteria.NewEntityCriteriaHelper(),
	}
}

// BuildFromContext construye criterios desde el contexto de Gin con filtros específicos de categorías
func (b *CategoryCriteriaBuilder) BuildFromContext(c *gin.Context) *CategoryCriteriaBuilder {
	// Construir criterios base desde query parameters
	b.CriteriaBuilder = b.helper.BuildBaseFromContext(c)

	// Agregar filtros específicos de categorías
	b.addCategoryFilters(c.Request.URL.Query())

	return b
}

// BuildValidated construye y valida criterios desde el contexto
func (b *CategoryCriteriaBuilder) BuildValidated(c *gin.Context) domainCriteria.Criteria {
	criteria := b.BuildFromContext(c).Build()
	return b.helper.ValidateAndSanitizeCriteria(criteria, b.GetAllowedFields())
}

// addCategoryFilters agrega filtros específicos de categorías
func (b *CategoryCriteriaBuilder) addCategoryFilters(values url.Values) {
	// Filtro por tenant_id (obligatorio)
	if tenantID := values.Get("tenant_id"); tenantID != "" {
		b.AddUUIDFilter("tenant_id", tenantID)
	}

	// Filtro por status
	if status := values.Get("status"); status != "" {
		b.AddEqualFilter("status", status)
	}

	// Filtro por parent_id
	if parentID := values.Get("parent_id"); parentID != "" {
		if parentID == "null" || parentID == "NULL" {
			b.AddFilter("parent_id", domainCriteria.OpIsNull, nil)
		} else {
			b.AddUUIDFilter("parent_id", parentID)
		}
	}

	// Filtro por name (búsqueda LIKE)
	if name := values.Get("name"); name != "" {
		b.AddLikeFilter("name", name)
	}

	// Filtro por description (búsqueda LIKE)
	if description := values.Get("description"); description != "" {
		b.AddLikeFilter("description", description)
	}

	// Filtros especiales
	if active := values.Get("active"); active == "true" {
		b.AddEqualFilter("status", "active")
	}

	if rootOnly := values.Get("root_only"); rootOnly == "true" {
		b.AddFilter("parent_id", domainCriteria.OpIsNull, nil)
	}
}

// GetAllowedFields retorna los campos permitidos para filtrado y ordenamiento
func (b *CategoryCriteriaBuilder) GetAllowedFields() []string {
	return []string{
		"id",
		"tenant_id",
		"name",
		"description",
		"parent_id",
		"status",
		"created_at",
		"updated_at",
	}
}
