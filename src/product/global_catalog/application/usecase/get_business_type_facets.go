package usecase

import (
	"fmt"
	"saas-mt-pim-service/src/product/global_catalog/domain/port"
)

type BusinessTypeFacets struct {
	BusinessType string   `json:"business_type"`
	Brands       []string `json:"brands"`
	Categories   []string `json:"categories"`
}

type GetBusinessTypeFacets struct {
	repo port.GlobalProductRepository
}

func NewGetBusinessTypeFacets(repo port.GlobalProductRepository) *GetBusinessTypeFacets {
	return &GetBusinessTypeFacets{repo: repo}
}

func (uc *GetBusinessTypeFacets) Execute(businessType string) (*BusinessTypeFacets, error) {
	if businessType == "" {
		return nil, fmt.Errorf("business_type es requerido")
	}

	brands, err := uc.repo.FindDistinctBrandsByBusinessType(businessType)
	if err != nil {
		return nil, fmt.Errorf("buscando marcas: %w", err)
	}

	categories, err := uc.repo.FindDistinctCategoriesByBusinessType(businessType)
	if err != nil {
		return nil, fmt.Errorf("buscando categorías: %w", err)
	}

	return &BusinessTypeFacets{
		BusinessType: businessType,
		Brands:       brands,
		Categories:   categories,
	}, nil
}
