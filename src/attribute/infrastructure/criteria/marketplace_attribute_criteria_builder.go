package criteria

import (
	"pim/src/shared/domain/criteria"
	sharedCriteria "pim/src/shared/infrastructure/criteria"

	"github.com/gin-gonic/gin"
)

// MarketplaceAttributeCriteriaBuilder construye criterios específicos para MarketplaceAttribute
type MarketplaceAttributeCriteriaBuilder struct {
	helper  *sharedCriteria.EntityCriteriaHelper
	builder *criteria.CriteriaBuilder
}

// NewMarketplaceAttributeCriteriaBuilder crea una nueva instancia del builder
func NewMarketplaceAttributeCriteriaBuilder() *MarketplaceAttributeCriteriaBuilder {
	return &MarketplaceAttributeCriteriaBuilder{
		helper: sharedCriteria.NewEntityCriteriaHelper(),
	}
}

// FromContext construye criterios desde el contexto de Gin
func (b *MarketplaceAttributeCriteriaBuilder) FromContext(c *gin.Context) *MarketplaceAttributeCriteriaBuilder {
	// Construir criterios base (paginación, ordenamiento)
	b.builder = b.helper.BuildBaseFromContext(c)

	// Filtros específicos de MarketplaceAttribute
	b.addMarketplaceAttributeFilters(c)

	return b
}

// BuildValidated construye y valida los criterios
func (b *MarketplaceAttributeCriteriaBuilder) BuildValidated(c *gin.Context) criteria.Criteria {
	searchCriteria := b.FromContext(c).Build()
	return b.helper.ValidateAndSanitizeCriteria(searchCriteria, b.GetAllowedFields())
}

// Build construye los criterios sin validación
func (b *MarketplaceAttributeCriteriaBuilder) Build() criteria.Criteria {
	return b.builder.Build()
}

// addMarketplaceAttributeFilters agrega filtros específicos de MarketplaceAttribute
func (b *MarketplaceAttributeCriteriaBuilder) addMarketplaceAttributeFilters(c *gin.Context) {
	// Filtro por tipo de atributo
	if attrType := c.Query("type"); attrType != "" {
		b.builder.AddEqualFilter("type", attrType)
	}

	// Filtro por nombre (búsqueda LIKE)
	if name := c.Query("name"); name != "" {
		b.builder.AddLikeFilter("name", name)
	}

	// Filtro por estado activo/inactivo
	if isActive := c.Query("is_active"); isActive != "" {
		if isActive == "true" {
			b.builder.AddEqualFilter("is_active", true)
		} else if isActive == "false" {
			b.builder.AddEqualFilter("is_active", false)
		}
	}

	// Filtro por atributos requeridos
	if isRequired := c.Query("is_required"); isRequired != "" {
		if isRequired == "true" {
			b.builder.AddEqualFilter("is_required_for_listing", true)
		} else if isRequired == "false" {
			b.builder.AddEqualFilter("is_required_for_listing", false)
		}
	}

	// Filtro por atributos filtrables
	if isFilterable := c.Query("is_filterable"); isFilterable != "" {
		if isFilterable == "true" {
			b.builder.AddEqualFilter("is_filterable", true)
		} else if isFilterable == "false" {
			b.builder.AddEqualFilter("is_filterable", false)
		}
	}

	// Filtro por atributos buscables
	if isSearchable := c.Query("is_searchable"); isSearchable != "" {
		if isSearchable == "true" {
			b.builder.AddEqualFilter("is_searchable", true)
		} else if isSearchable == "false" {
			b.builder.AddEqualFilter("is_searchable", false)
		}
	}

	// Filtro por descripción (búsqueda LIKE)
	if description := c.Query("description"); description != "" {
		b.builder.AddLikeFilter("description", description)
	}

	// Filtro por group_name si existe en la tabla
	if groupName := c.Query("group_name"); groupName != "" {
		b.builder.AddEqualFilter("group_name", groupName)
	}

	// Por defecto, solo mostrar atributos activos (a menos que se especifique lo contrario)
	if includeInactive := c.Query("include_inactive"); includeInactive != "true" {
		b.builder.AddEqualFilter("is_active", true)
	}
}

// GetAllowedFields retorna los campos permitidos para filtrado y ordenamiento
func (b *MarketplaceAttributeCriteriaBuilder) GetAllowedFields() []string {
	return []string{
		"id",
		"name",
		"type",
		"description",
		"is_required_for_listing",
		"is_searchable",
		"is_filterable",
		"is_active",
		"group_name",
		"sort_order",
		"created_at",
		"updated_at",
	}
}

// GetDefaultOrderField retorna el campo de ordenamiento por defecto
func (b *MarketplaceAttributeCriteriaBuilder) GetDefaultOrderField() string {
	return "sort_order"
}

// GetDefaultOrderDirection retorna la dirección de ordenamiento por defecto
func (b *MarketplaceAttributeCriteriaBuilder) GetDefaultOrderDirection() string {
	return "ASC"
}
