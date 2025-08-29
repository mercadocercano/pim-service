package usecase

import (
	"context"
	"github.com/stretchr/testify/assert"
	"saas-mt-pim-service/src/product/tenant/application/mapper"
	"saas-mt-pim-service/src/product/application/request"
	"saas-mt-pim-service/src/product/application/usecase"
	"testing"
)

func TestCreateProductUseCase_InvalidRequest(t *testing.T) {
	mapper := mapper.NewProductMapper()
	useCase := usecase.NewCreateProductUseCase(nil, nil, mapper)
	req := &request.CreateProductRequest{Name: ""}
	result, err := useCase.Execute(context.Background(), req, "tenant-123")
	assert.Error(t, err)
	assert.Nil(t, result)
}
