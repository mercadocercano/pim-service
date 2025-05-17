package criteria

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

// Order representa el criterio de ordenación
type Order struct {
	Field     string
	Direction string
}

// NewOrder crea un nuevo criterio de ordenación
func NewOrder(field string, direction string) Order {
	return Order{
		Field:     field,
		Direction: direction,
	}
}

// Pagination representa los criterios de paginación
type Pagination struct {
	Limit  int
	Offset int
}

// NewPagination crea un nuevo criterio de paginación
func NewPagination(limit int, offset int) Pagination {
	return Pagination{
		Limit:  limit,
		Offset: offset,
	}
}
