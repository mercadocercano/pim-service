package usecase

import (
	"context"
	"fmt"

	"pim/src/attribute/domain/port"
)

// DeleteTenantCustomAttributeUseCase maneja la eliminación de atributos personalizados de un tenant
type DeleteTenantCustomAttributeUseCase struct {
	customAttributeRepo port.TenantCustomAttributeRepository
}

// NewDeleteTenantCustomAttributeUseCase crea una nueva instancia del caso de uso
func NewDeleteTenantCustomAttributeUseCase(
	customAttributeRepo port.TenantCustomAttributeRepository,
) *DeleteTenantCustomAttributeUseCase {
	return &DeleteTenantCustomAttributeUseCase{
		customAttributeRepo: customAttributeRepo,
	}
}

// Execute ejecuta el caso de uso para eliminar un atributo personalizado
func (uc *DeleteTenantCustomAttributeUseCase) Execute(
	ctx context.Context,
	tenantID string,
	attributeID string,
) error {
	// TODO: Verificar que el atributo pertenece al tenant
	// Esta validación se puede hacer en el controlador

	// Eliminar el atributo (soft delete)
	if err := uc.customAttributeRepo.Delete(ctx, attributeID); err != nil {
		return fmt.Errorf("failed to delete attribute: %w", err)
	}

	return nil
}
