package usecase

import (
	"context"
	"encoding/json"
	"saas-mt-pim-service/src/quickstart/domain/entity"
)

// SetupTenantUseCase configura un tenant con los datos del quickstart
type SetupTenantUseCase struct {
	// Aquí irían las dependencias de servicios para crear categorías, productos, etc.
}

// NewSetupTenantUseCase crea una nueva instancia
func NewSetupTenantUseCase() *SetupTenantUseCase {
	return &SetupTenantUseCase{}
}

// Execute ejecuta el caso de uso
func (uc *SetupTenantUseCase) Execute(ctx context.Context, tenantID string, setupData map[string]interface{}) (*entity.TenantQuickstartHistory, error) {
	// Por ahora solo crea un registro de historia mock

	dataJSON, _ := json.Marshal(setupData)

	history := &entity.TenantQuickstartHistory{
		ID:        "mock-history-id",
		TenantID:  tenantID,
		SetupData: dataJSON,
	}

	return history, nil
}
