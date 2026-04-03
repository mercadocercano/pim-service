package usecase_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"saas-mt-pim-service/src/quickstart/application/usecase"
)

func TestSetupTenantUseCase_Execute_Success(t *testing.T) {
	// Arrange
	uc := usecase.NewSetupTenantUseCase()
	setupData := map[string]interface{}{
		"business_type_id": "ferreteria",
		"template_id":      "tpl-1",
	}

	// Act
	result, err := uc.Execute(context.Background(), "tenant-123", setupData)

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "tenant-123", result.TenantID)
	assert.NotEmpty(t, result.SetupData)
}

func TestSetupTenantUseCase_Execute_EmptySetupData(t *testing.T) {
	// Arrange
	uc := usecase.NewSetupTenantUseCase()

	// Act
	result, err := uc.Execute(context.Background(), "tenant-123", map[string]interface{}{})

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "tenant-123", result.TenantID)
}

func TestSetupTenantUseCase_Execute_WithMultipleFields(t *testing.T) {
	// Arrange
	uc := usecase.NewSetupTenantUseCase()
	setupData := map[string]interface{}{
		"business_type_id": "kiosco",
		"categories":       []string{"golosinas", "bebidas"},
		"brands":           []string{"Coca-Cola", "Arcor"},
	}

	// Act
	result, err := uc.Execute(context.Background(), "tenant-456", setupData)

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "tenant-456", result.TenantID)
}
