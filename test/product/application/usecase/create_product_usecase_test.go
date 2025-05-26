package usecase

import (
	"context"
	"github.com/stretchr/testify/assert"
	"pim/src/product/application/mapper"
	"pim/src/product/application/request"
	"pim/src/product/application/usecase"
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
