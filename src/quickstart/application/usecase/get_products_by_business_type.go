package usecase

import (
	"context"

	"saas-mt-pim-service/src/quickstart/domain/port"
)

// GetProductsByBusinessTypeUseCase obtiene productos sugeridos por tipo de negocio
type GetProductsByBusinessTypeUseCase struct {
	repo port.GetProductsByBusinessTypeRepository
}

// NewGetProductsByBusinessTypeUseCase crea una nueva instancia
func NewGetProductsByBusinessTypeUseCase(repo port.GetProductsByBusinessTypeRepository) *GetProductsByBusinessTypeUseCase {
	return &GetProductsByBusinessTypeUseCase{repo: repo}
}

// Execute ejecuta el caso de uso
func (uc *GetProductsByBusinessTypeUseCase) Execute(ctx context.Context, businessType string) (interface{}, error) {
	products, err := uc.repo.GetProductsByBusinessType(ctx, businessType)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"products": products,
		"total":    len(products),
	}, nil
}
