package value_object

import (
	"github.com/stretchr/testify/assert"
	"pim/src/product/domain/value_object"
	"testing"
)

func TestProductStatus_Factories(t *testing.T) {
	activeStatus := value_object.ProductStatusActive()
	assert.True(t, activeStatus.IsActive())
	assert.False(t, activeStatus.IsDeleted())
	assert.Equal(t, "active", activeStatus.Value())

	inactiveStatus := value_object.ProductStatusInactive()
	assert.False(t, inactiveStatus.IsActive())
	assert.True(t, inactiveStatus.IsInactive())
	assert.Equal(t, "inactive", inactiveStatus.Value())

	deletedStatus := value_object.ProductStatusDeleted()
	assert.True(t, deletedStatus.IsDeleted())
	assert.Equal(t, "deleted", deletedStatus.Value())
}

func TestProductStatus_Validation(t *testing.T) {
	status := value_object.ProductStatusActive()
	assert.True(t, status.IsValid())

	// Test NewProductStatus with valid value
	status2, err := value_object.NewProductStatus("active")
	assert.NoError(t, err)
	assert.True(t, status2.IsActive())

	// Test NewProductStatus with invalid value
	_, err = value_object.NewProductStatus("invalid")
	assert.Error(t, err)
}
