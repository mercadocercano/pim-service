package port

import (
	"context"
	"time"

	cr "github.com/hornosg/go-shared/criteria"
	"saas-mt-pim-service/src/category_attribute/domain/entity"
)

// CategoryAttributeRepository define el contrato para acceder a los datos de atributos de categoría
type CategoryAttributeRepository interface {
	// Create guarda un nuevo atributo de categoría
	Create(ctx context.Context, categoryAttribute *entity.CategoryAttribute) error

	// FindByID busca un atributo de categoría por su ID y tenantID
	FindByID(ctx context.Context, id string, tenantID string) (*entity.CategoryAttribute, error)

	// FindByCategoryAndTenant busca atributos de categoría por categoryID y tenantID
	FindByCategoryAndTenant(ctx context.Context, categoryID string, tenantID string) ([]*entity.CategoryAttribute, error)

	// FindDetailedByCategoryAndTenant busca atributos completos con JOIN para una categoría
	FindDetailedByCategoryAndTenant(ctx context.Context, categoryID string, tenantID string) ([]*DetailedCategoryAttribute, error)

	// FindByTenant recupera todos los atributos de categoría de un tenant
	FindByTenant(ctx context.Context, tenantID string) ([]*entity.CategoryAttribute, error)

	// Update actualiza un atributo de categoría existente
	Update(ctx context.Context, categoryAttribute *entity.CategoryAttribute) error

	// Delete elimina un atributo de categoría por su ID y tenantID
	Delete(ctx context.Context, id string, tenantID string) error

	// FindByAttributeAndCategory busca una relación específica por attributeID, categoryID y tenantID
	FindByAttributeAndCategory(ctx context.Context, attributeID, categoryID, tenantID string) (*entity.CategoryAttribute, error)
}

// CategoryAttributeCriteriaRepository extiende CategoryAttributeRepository con soporte para criteria
type CategoryAttributeCriteriaRepository interface {
	CategoryAttributeRepository
	cr.CriteriaRepository[entity.CategoryAttribute]
	cr.ListRepository[entity.CategoryAttribute]
}

// DetailedCategoryAttribute representa un atributo de categoría con datos completos del atributo
type DetailedCategoryAttribute struct {
	ID               string
	CategoryID       string
	AttributeID      string
	AttributeName    string
	AttributeType    string
	Description      string
	Required         bool
	AttributeOptions []string
	AllowedValues    []string
	Active           bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
