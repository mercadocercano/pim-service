package criteria

import (
	"math"
)

// Criteria representa un conjunto de criterios de filtrado, ordenación y paginación
type Criteria struct {
	Filters    Filters
	Order      Order
	Pagination Pagination
}

// NewCriteria crea una nueva instancia de Criteria
func NewCriteria(filters Filters, order Order, pagination Pagination) Criteria {
	return Criteria{
		Filters:    filters,
		Order:      order,
		Pagination: pagination,
	}
}

// IsEmpty verifica si el criteria está vacío
func (c Criteria) IsEmpty() bool {
	return len(c.Filters.Items) == 0 && c.Order.Field == "" && c.Pagination.Limit == 0
}

// Filter representa un filtro individual
type Filter struct {
	Field    string
	Operator string
	Value    interface{}
}

// NewFilter crea un nuevo filtro
func NewFilter(field string, operator string, value interface{}) Filter {
	return Filter{
		Field:    field,
		Operator: operator,
		Value:    value,
	}
}

// Filters representa una colección de filtros
type Filters struct {
	Items []Filter
}

// NewFilters crea una nueva colección de filtros
func NewFilters(items ...Filter) Filters {
	return Filters{
		Items: items,
	}
}

// Add agrega un filtro a la colección
func (f *Filters) Add(filter Filter) {
	f.Items = append(f.Items, filter)
}

// Count retorna el número de filtros
func (f Filters) Count() int {
	return len(f.Items)
}

// IsEmpty verifica si no hay filtros
func (f Filters) IsEmpty() bool {
	return len(f.Items) == 0
}

// Order representa el criterio de ordenación
type Order struct {
	Field     string
	Direction string
}

// NewOrder crea un nuevo criterio de ordenación
func NewOrder(field string, direction string) Order {
	if direction != "ASC" && direction != "DESC" {
		direction = "DESC" // Default
	}
	return Order{
		Field:     field,
		Direction: direction,
	}
}

// IsEmpty verifica si el orden está vacío
func (o Order) IsEmpty() bool {
	return o.Field == ""
}

// Pagination representa los criterios de paginación
type Pagination struct {
	Page     int
	PageSize int
	Limit    int
	Offset   int
}

// NewPagination crea un nuevo criterio de paginación
func NewPagination(page, pageSize int) Pagination {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}

	offset := (page - 1) * pageSize

	return Pagination{
		Page:     page,
		PageSize: pageSize,
		Limit:    pageSize,
		Offset:   offset,
	}
}

// IsEmpty verifica si la paginación está vacía
func (p Pagination) IsEmpty() bool {
	return p.Limit == 0
}

// GetTotalPages calcula el número total de páginas
func (p Pagination) GetTotalPages(totalCount int) int {
	if p.PageSize == 0 {
		return 0
	}
	return int(math.Ceil(float64(totalCount) / float64(p.PageSize)))
}

// ListResponse representa una respuesta de listado genérica
type ListResponse[T any] struct {
	Items      []*T `json:"items"`
	TotalCount int  `json:"total_count"`
	Page       int  `json:"page"`
	PageSize   int  `json:"page_size"`
	TotalPages int  `json:"total_pages"`
}

// NewListResponse crea una nueva respuesta de listado
func NewListResponse[T any](items []*T, totalCount int, criteria Criteria) *ListResponse[T] {
	totalPages := criteria.Pagination.GetTotalPages(totalCount)

	return &ListResponse[T]{
		Items:      items,
		TotalCount: totalCount,
		Page:       criteria.Pagination.Page,
		PageSize:   criteria.Pagination.PageSize,
		TotalPages: totalPages,
	}
}
