package adapter

import (
	"context"
	"database/sql"
)

// PimCategoryAdapter retrieves tenant category names from the PIM database.
type PimCategoryAdapter struct {
	db *sql.DB
}

func NewPimCategoryAdapter(db *sql.DB) *PimCategoryAdapter {
	return &PimCategoryAdapter{db: db}
}

func (a *PimCategoryAdapter) GetCategoryNames(ctx context.Context, tenantID string) ([]string, error) {
	query := `SELECT name FROM categories WHERE tenant_id = $1 ORDER BY name`
	rows, err := a.db.QueryContext(ctx, query, tenantID)
	if err != nil {
		return []string{}, nil
	}
	defer rows.Close()

	var names []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			continue
		}
		names = append(names, name)
	}
	return names, nil
}
