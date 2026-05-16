package usecase

import (
	"saas-mt-pim-service/src/product/global_catalog/domain/port"
)

type GetDistinctBusinessTypes struct {
	repo port.GlobalProductRepository
}

func NewGetDistinctBusinessTypes(repo port.GlobalProductRepository) *GetDistinctBusinessTypes {
	return &GetDistinctBusinessTypes{repo: repo}
}

func (uc *GetDistinctBusinessTypes) Execute() ([]string, error) {
	return uc.repo.FindDistinctBusinessTypes()
}
