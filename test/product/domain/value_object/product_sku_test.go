package value_object

import (
	"testing"
	"pim/src/product/domain/value_object"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewProductSKU(t *testing.T) {
	sku, err := value_object.NewProductSKU("PROD-001")
	require.NoError(t, err)
	assert.NotNil(t, sku)
	assert.Equal(t, "PROD-001", sku.Value())
}

func TestNewProductSKU_EmptyValue(t *testing.T) {
	sku, err := value_object.NewProductSKU("")
	assert.Error(t, err)
	assert.Nil(t, sku)
}

func TestProductSKU_Equals(t *testing.T) {
	sku1, _ := value_object.NewProductSKU("PROD-001")
	sku2, _ := value_object.NewProductSKU("PROD-001")
	sku3, _ := value_object.NewProductSKU("PROD-002")
	
	assert.True(t, sku1.Equals(sku2))
	assert.False(t, sku1.Equals(sku3))
	assert.False(t, sku1.Equals(nil))
}

func TestProductSKU_String(t *testing.T) {
	sku, _ := value_object.NewProductSKU("PROD-001")
	assert.Equal(t, "PROD-001", sku.String())
}