package entity

import (
	"github.com/stretchr/testify/assert"
	"saas-mt-pim-service/src/product/tenant/domain/entity"
	"testing"
)

func TestProduct_Simple(t *testing.T) {
	product, err := entity.NewProduct("tenant-123", "Test Product", nil, nil, nil, nil)
	assert.NoError(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, "Test Product", product.Name())
	assert.Equal(t, "tenant-123", product.TenantID())
}

func TestProduct_Update(t *testing.T) {
	product := createTestProduct()
	newName := "Updated Product"
	err := product.Update(newName, nil, nil, nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, newName, product.Name())
}

func TestProduct_StatusTransitions(t *testing.T) {
	product := createTestProduct()
	assert.True(t, product.IsActive())

	product.Deactivate()
	assert.False(t, product.IsActive())

	product.Activate()
	assert.True(t, product.IsActive())
}

func createTestProduct() *entity.Product {
	product, _ := entity.NewProduct("tenant-123", "Test Product", nil, nil, nil, nil)
	return product
}
