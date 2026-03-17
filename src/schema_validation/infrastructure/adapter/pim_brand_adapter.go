package adapter

import (
	"context"
	"database/sql"
)

// PimBrandAdapter retrieves tenant brand names from the PIM database.
type PimBrandAdapter struct {
	db *sql.DB
}

func NewPimBrandAdapter(db *sql.DB) *PimBrandAdapter {
	return &PimBrandAdapter{db: db}
}

func (a *PimBrandAdapter) GetBrandNames(ctx context.Context, tenantID string) ([]string, error) {
	query := `SELECT name FROM brands WHERE tenant_id = $1 AND status = 'active' ORDER BY name`
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
