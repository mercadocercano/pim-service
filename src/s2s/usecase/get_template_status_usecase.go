package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

// ErrBusinessTypeNotFound se retorna cuando el slug no existe en business_types.
var ErrBusinessTypeNotFound = errors.New("business type not found")

// TemplateStatusResult describe el estado del template computado vs editorial para un rubro.
type TemplateStatusResult struct {
	BusinessTypeSlug string     `json:"business_type_slug"`
	Source           string     `json:"source"` // "computed" | "editorial"
	ComputedCount    int        `json:"computed_count"`
	EditorialCount   int        `json:"editorial_count"`
	LastRefresh      *time.Time `json:"last_refresh"` // nil si no hay template computado
}

// GetTemplateStatusUseCase consulta el estado del template de quickstart para un rubro.
type GetTemplateStatusUseCase struct {
	db *sql.DB
}

const templateStatusQuery = `
	SELECT
		COALESCE(jsonb_array_length(btpt.suggested_products), 0) AS computed_count,
		COALESCE(jsonb_array_length(btt.products), 0)            AS editorial_count,
		btpt.updated_at                                          AS last_refresh
	FROM business_type_templates btt
	JOIN business_types bt ON bt.id = btt.business_type_id
	LEFT JOIN business_type_product_templates btpt ON btpt.business_type_template_id = btt.id
	WHERE bt.code = $1
	  AND btt.is_active = true
	  AND btt.is_default = true
	LIMIT 1
`

func NewGetTemplateStatusUseCase(db *sql.DB) *GetTemplateStatusUseCase {
	return &GetTemplateStatusUseCase{db: db}
}

func (uc *GetTemplateStatusUseCase) Execute(ctx context.Context, businessTypeSlug string) (*TemplateStatusResult, error) {
	var computedCount, editorialCount int
	var lastRefresh sql.NullTime

	err := uc.db.QueryRowContext(ctx, templateStatusQuery, businessTypeSlug).Scan(&computedCount, &editorialCount, &lastRefresh)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("%w: %s", ErrBusinessTypeNotFound, businessTypeSlug)
	}
	if err != nil {
		return nil, fmt.Errorf("get template status %q: %w", businessTypeSlug, err)
	}

	return &TemplateStatusResult{
		BusinessTypeSlug: businessTypeSlug,
		Source:           resolveSource(computedCount),
		ComputedCount:    computedCount,
		EditorialCount:   editorialCount,
		LastRefresh:      nullTimePtr(lastRefresh),
	}, nil
}

func resolveSource(computedCount int) string {
	if computedCount > 0 {
		return "computed"
	}
	return "editorial"
}

func nullTimePtr(t sql.NullTime) *time.Time {
	if t.Valid {
		return &t.Time
	}
	return nil
}
