package value_object

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"saas-mt-pim-service/src/product/tenant/domain/value_object"
	"testing"
)

func TestNewBrandReference(t *testing.T) {
	t.Run("should create brand reference with valid data", func(t *testing.T) {
		id := uuid.New().String()
		name := "Marca Test"
		brandRef, err := value_object.NewBrandReference(id, name)
		require.NoError(t, err)
		assert.NotNil(t, brandRef)
		assert.Equal(t, id, brandRef.ID())
		assert.Equal(t, name, brandRef.Name())
	})

	t.Run("should fail when ID is empty", func(t *testing.T) {
		brandRef, err := value_object.NewBrandReference("", "Marca Test")
		assert.Error(t, err)
		assert.Nil(t, brandRef)
		assert.Contains(t, err.Error(), "el ID de marca es obligatorio")
	})

	t.Run("should fail when name is empty", func(t *testing.T) {
		id := uuid.New().String()
		brandRef, err := value_object.NewBrandReference(id, "")
		assert.Error(t, err)
		assert.Nil(t, brandRef)
		assert.Contains(t, err.Error(), "el nombre de marca es obligatorio")
	})

	t.Run("should fail when ID is not UUID", func(t *testing.T) {
		brandRef, err := value_object.NewBrandReference("invalid-uuid", "Marca Test")
		assert.Error(t, err)
		assert.Nil(t, brandRef)
		assert.Contains(t, err.Error(), "el ID de marca debe ser un UUID válido")
	})
}

func TestBrandReference_Methods(t *testing.T) {
	id := uuid.New().String()
	brandRef, _ := value_object.NewBrandReference(id, "Marca Test")

	assert.Equal(t, id, brandRef.ID())
	assert.Equal(t, "Marca Test", brandRef.Name())
	assert.Equal(t, "Marca Test", brandRef.String())
	assert.False(t, brandRef.IsEmpty())
	assert.False(t, brandRef.Equals(nil))
}
