package criteria

import (
	"fmt"
	"strings"

	domainCriteria "pim/src/shared/domain/criteria"
)

// SQLCriteriaConverter convierte un objeto Criteria en una consulta SQL
type SQLCriteriaConverter struct{}

// NewSQLCriteriaConverter crea una nueva instancia del conversor
func NewSQLCriteriaConverter() *SQLCriteriaConverter {
	return &SQLCriteriaConverter{}
}

// ToSQL convierte un criteria a una consulta SQL con sus parámetros
func (s *SQLCriteriaConverter) ToSQL(criteria domainCriteria.Criteria) (string, []interface{}) {
	var conditions []string
	var params []interface{}

	// Procesar los filtros
	for _, filter := range criteria.Filters.Items {
		condition, value := s.processFilter(filter)
		conditions = append(conditions, condition)
		params = append(params, value)
	}

	// Construir la cláusula WHERE
	var whereClause string
	if len(conditions) > 0 {
		whereClause = fmt.Sprintf("WHERE %s", strings.Join(conditions, " AND "))
	}

	// Construir la cláusula ORDER BY
	var orderByClause string
	if criteria.Order.Field != "" {
		orderByClause = fmt.Sprintf("ORDER BY %s %s", criteria.Order.Field, criteria.Order.Direction)
	}

	// Construir la cláusula LIMIT y OFFSET
	var limitOffsetClause string
	if criteria.Pagination.Limit > 0 {
		limitOffsetClause = fmt.Sprintf("LIMIT %d OFFSET %d", criteria.Pagination.Limit, criteria.Pagination.Offset)
	}

	// Combinar las cláusulas
	clauses := []string{whereClause, orderByClause, limitOffsetClause}
	var filteredClauses []string
	for _, clause := range clauses {
		if clause != "" {
			filteredClauses = append(filteredClauses, clause)
		}
	}

	return strings.Join(filteredClauses, " "), params
}

// processFilter convierte un filtro en una condición SQL
func (s *SQLCriteriaConverter) processFilter(filter domainCriteria.Filter) (string, interface{}) {
	var condition string

	switch filter.Operator {
	case "=", "!=", ">", ">=", "<", "<=":
		condition = fmt.Sprintf("%s %s $?", filter.Field, filter.Operator)
	case "LIKE":
		condition = fmt.Sprintf("%s LIKE $?", filter.Field)
		// Asegurar que el valor sea compatible con LIKE
		if str, ok := filter.Value.(string); ok {
			if !strings.Contains(str, "%") {
				filter.Value = "%" + str + "%"
			}
		}
	case "IN":
		// Manejar arrays para cláusulas IN
		condition = fmt.Sprintf("%s IN ($?)", filter.Field)
	case "NULL":
		condition = fmt.Sprintf("%s IS NULL", filter.Field)
		return condition, nil
	case "NOT NULL":
		condition = fmt.Sprintf("%s IS NOT NULL", filter.Field)
		return condition, nil
	default:
		condition = fmt.Sprintf("%s = $?", filter.Field)
	}

	return condition, filter.Value
}
