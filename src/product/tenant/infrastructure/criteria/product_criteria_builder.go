package criteria

import (
	"net/url"

	cr "github.com/hornosg/go-shared/criteria"

	"github.com/gin-gonic/gin"
)

// ProductCriteriaBuilder construye criterios específicos para productos
type ProductCriteriaBuilder struct {
	*cr.CriteriaBuilder
	helper *cr.EntityCriteriaHelper
}

// NewProductCriteriaBuilder crea un nuevo builder para criterios de productos
func NewProductCriteriaBuilder() *ProductCriteriaBuilder {
	return &ProductCriteriaBuilder{
		CriteriaBuilder: cr.NewCriteriaBuilder(),
		helper:          cr.NewEntityCriteriaHelper(),
	}
}

// BuildFromContext construye criterios desde el contexto de Gin con filtros específicos de productos
func (b *ProductCriteriaBuilder) BuildFromContext(c *gin.Context) *ProductCriteriaBuilder {
	// Construir criterios base desde query parameters
	b.CriteriaBuilder = b.helper.BuildBaseFromContext(c)

	// Agregar filtros específicos de productos
	b.addProductFilters(c.Request.URL.Query())

	return b
}

// BuildValidated construye y valida criterios desde el contexto
func (b *ProductCriteriaBuilder) BuildValidated(c *gin.Context) cr.Criteria {
	criteria := b.BuildFromContext(c).Build()
	return b.helper.ValidateAndSanitizeCriteria(criteria, b.GetAllowedFields())
}

// addProductFilters agrega filtros específicos de productos
func (b *ProductCriteriaBuilder) addProductFilters(values url.Values) {
	// Filtro por tenant_id (obligatorio)
	if tenantID := values.Get("tenant_id"); tenantID != "" {
		b.AddUUIDFilter("tenant_id", tenantID)
	}

	// Filtros de búsqueda por texto
	if name := values.Get("name"); name != "" {
		b.AddLikeFilter("name", name)
	}

	if description := values.Get("description"); description != "" {
		b.AddLikeFilter("description", description)
	}

	if sku := values.Get("sku"); sku != "" {
		b.AddLikeFilter("sku", sku)
	}

	// Filtros exactos
	if status := values.Get("status"); status != "" {
		b.AddEqualFilter("status", status)
	}

	if categoryID := values.Get("category_id"); categoryID != "" {
		b.AddUUIDFilter("category_id", categoryID)
	}

	if brandID := values.Get("brand_id"); brandID != "" {
		b.AddUUIDFilter("brand_id", brandID)
	}

	// Filtros por nombres relacionados
	if categoryName := values.Get("category_name"); categoryName != "" {
		b.AddLikeFilter("category_name", categoryName)
	}

	if brandName := values.Get("brand_name"); brandName != "" {
		b.AddLikeFilter("brand_name", brandName)
	}

	// Filtros especiales
	if active := values.Get("active"); active == "true" {
		b.AddEqualFilter("status", "active")
	}

	// Filtro para excluir productos eliminados por defecto
	if includeDeleted := values.Get("include_deleted"); includeDeleted != "true" {
		b.AddFilter("status", cr.OpNotEqual, "deleted")
	}

	// Filtros de disponibilidad
	if inStock := values.Get("in_stock"); inStock == "true" {
		b.AddFilter("stock_quantity", cr.OpGreaterThan, 0)
	} else if inStock == "false" {
		b.AddFilter("stock_quantity", cr.OpLessThanOrEqual, 0)
	}

	// Filtros de precio
	if minPrice := values.Get("min_price"); minPrice != "" {
		b.AddFilter("price", cr.OpGreaterThanOrEqual, minPrice)
	}

	if maxPrice := values.Get("max_price"); maxPrice != "" {
		b.AddFilter("price", cr.OpLessThanOrEqual, maxPrice)
	}
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
		"price",
		"stock_quantity",
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
