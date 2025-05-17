package mapper

import (
	"pim/src/category/domain/entity"
	"pim/src/category/infrastructure/persistence/model"
)

// CategoryMapper maneja las conversiones entre entidad y modelo
type CategoryMapper struct{}

// NewCategoryMapper crea una nueva instancia del mapper
func NewCategoryMapper() *CategoryMapper {
	return &CategoryMapper{}
}

// ToEntity convierte un modelo de base de datos a una entidad de dominio
func (m *CategoryMapper) ToEntity(model *model.Category) *entity.Category {
	if model == nil {
		return nil
	}

	return &entity.Category{
		ID:          model.ID,
		TenantID:    model.TenantID,
		Name:        model.Name,
		Description: model.Description,
		ParentID:    model.ParentID,
		Status:      model.Status,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
}

// ToModel convierte una entidad de dominio a un modelo de base de datos
func (m *CategoryMapper) ToModel(entity *entity.Category) *model.Category {
	if entity == nil {
		return nil
	}

	return &model.Category{
		ID:          entity.ID,
		TenantID:    entity.TenantID,
		Name:        entity.Name,
		Description: entity.Description,
		ParentID:    entity.ParentID,
		Status:      entity.Status,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}

// ToEntityList convierte una lista de modelos a una lista de entidades
func (m *CategoryMapper) ToEntityList(models []*model.Category) []*entity.Category {
	if models == nil {
		return nil
	}

	entities := make([]*entity.Category, len(models))
	for i, model := range models {
		entities[i] = m.ToEntity(model)
	}
	return entities
}

// ToModelList convierte una lista de entidades a una lista de modelos
func (m *CategoryMapper) ToModelList(entities []*entity.Category) []*model.Category {
	if entities == nil {
		return nil
	}

	models := make([]*model.Category, len(entities))
	for i, entity := range entities {
		models[i] = m.ToModel(entity)
	}
	return models
}
