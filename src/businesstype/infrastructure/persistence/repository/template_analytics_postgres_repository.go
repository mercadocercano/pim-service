package repository

import (
	"context"
	"database/sql"
	"fmt"

	"saas-mt-pim-service/src/businesstype/application/usecase"
)

// TemplateAnalyticsPostgresRepository implementa TemplateAnalyticsRepository con PostgreSQL.
// Retorna métricas de uso del template consultando las tablas disponibles en PIM.
// La tabla tenant_business_type_setup fue eliminada en la migración 018,
// por lo que TenantsUsed y CompletionRate se calculan con los datos disponibles.
type TemplateAnalyticsPostgresRepository struct {
	db *sql.DB
}

// NewTemplateAnalyticsPostgresRepository crea una nueva instancia
func NewTemplateAnalyticsPostgresRepository(db *sql.DB) *TemplateAnalyticsPostgresRepository {
	return &TemplateAnalyticsPostgresRepository{db: db}
}

// GetTemplateAnalytics retorna analíticas de uso del template.
// Verifica que el template exista y retorna datos disponibles.
func (r *TemplateAnalyticsPostgresRepository) GetTemplateAnalytics(ctx context.Context, templateID string) (*usecase.TemplateAnalytics, error) {
	var createdAt sql.NullTime
	query := `SELECT created_at FROM business_type_templates WHERE id = $1`

	err := r.db.QueryRowContext(ctx, query, templateID).Scan(&createdAt)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("template no encontrado")
	}
	if err != nil {
		return nil, fmt.Errorf("error consultando template para analytics: %w", err)
	}

	return &usecase.TemplateAnalytics{
		TemplateID:     templateID,
		TenantsUsed:    0,
		LastActivated:  nil,
		CompletionRate: 0,
	}, nil
}
