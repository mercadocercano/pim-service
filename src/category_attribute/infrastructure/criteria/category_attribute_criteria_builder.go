package criteria

import (
	"net/url"

	cr "github.com/hornosg/go-shared/criteria"

	"github.com/gin-gonic/gin"
)

// CategoryAttributeCriteriaBuilder construye criterios específicos para atributos de categoría
type CategoryAttributeCriteriaBuilder struct {
	*cr.CriteriaBuilder
	helper *cr.EntityCriteriaHelper
}

// NewCategoryAttributeCriteriaBuilder crea un nuevo builder para criterios de atributos de categoría
func NewCategoryAttributeCriteriaBuilder() *CategoryAttributeCriteriaBuilder {
	return &CategoryAttributeCriteriaBuilder{
		CriteriaBuilder: cr.NewCriteriaBuilder(),
		helper:          cr.NewEntityCriteriaHelper(),
	}
}

// BuildFromContext construye criterios desde el contexto de Gin con filtros específicos de atributos de categoría
func (b *CategoryAttributeCriteriaBuilder) BuildFromContext(c *gin.Context) *CategoryAttributeCriteriaBuilder {
	// Construir criterios base desde query parameters
	b.CriteriaBuilder = b.helper.BuildBaseFromContext(c)

	// Agregar filtros específicos de atributos de categoría
	b.addCategoryAttributeFilters(c.Request.URL.Query())

	return b
}

// BuildValidated construye y valida criterios desde el contexto
func (b *CategoryAttributeCriteriaBuilder) BuildValidated(c *gin.Context) cr.Criteria {
	criteria := b.BuildFromContext(c).Build()
	return b.helper.ValidateAndSanitizeCriteria(criteria, b.GetAllowedFields())
}

// addCategoryAttributeFilters agrega filtros específicos de atributos de categoría
func (b *CategoryAttributeCriteriaBuilder) addCategoryAttributeFilters(values url.Values) {
	// Filtro por tenant_id (obligatorio)
	if tenantID := values.Get("tenant_id"); tenantID != "" {
		b.AddUUIDFilter("tenant_id", tenantID)
	}

	// Filtro por category_id
	if categoryID := values.Get("category_id"); categoryID != "" {
		b.AddUUIDFilter("category_id", categoryID)
	}

	// Filtro por attribute_id
	if attributeID := values.Get("attribute_id"); attributeID != "" {
		b.AddUUIDFilter("attribute_id", attributeID)
	}

	// Filtro por status
	if status := values.Get("status"); status != "" {
		b.AddEqualFilter("status", status)
	}

	// Filtros especiales
	if active := values.Get("active"); active == "true" {
		b.AddEqualFilter("status", "active")
	}

	// Filtro por valores permitidos (si contiene un valor específico)
	if allowedValue := values.Get("allowed_value"); allowedValue != "" {
		// Usar el nuevo operador ARRAY_CONTAINS para PostgreSQL
		b.AddArrayContainsFilter("allowed_values", allowedValue)
	}
}

// GetAllowedFields retorna los campos permitidos para filtrado y ordenamiento
func (b *CategoryAttributeCriteriaBuilder) GetAllowedFields() []string {
	return []string{
		"id",
		"tenant_id",
		"category_id",
		"attribute_id",
		"allowed_values",
		"status",
		"created_at",
		"updated_at",
	}
}
