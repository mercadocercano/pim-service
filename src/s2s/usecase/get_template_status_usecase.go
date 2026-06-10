package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"saas-mt-pim-service/src/s2s/domain/port"
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
	repo port.TemplateRepository
}

func NewGetTemplateStatusUseCase(repo port.TemplateRepository) *GetTemplateStatusUseCase {
	return &GetTemplateStatusUseCase{repo: repo}
}

func (uc *GetTemplateStatusUseCase) Execute(ctx context.Context, businessTypeSlug string) (*TemplateStatusResult, error) {
	row, err := uc.repo.GetTemplateStatus(ctx, businessTypeSlug)
	if err != nil {
		return nil, fmt.Errorf("get template status %q: %w", businessTypeSlug, err)
	}
	if row == nil {
		return nil, fmt.Errorf("%w: %s", ErrBusinessTypeNotFound, businessTypeSlug)
	}

	return &TemplateStatusResult{
		BusinessTypeSlug: businessTypeSlug,
		Source:           resolveSource(row.ComputedCount),
		ComputedCount:    row.ComputedCount,
		EditorialCount:   row.EditorialCount,
		LastRefresh:      row.LastRefresh,
	}, nil
}

func resolveSource(computedCount int) string {
	if computedCount > 0 {
		return "computed"
	}
	return "editorial"
}
