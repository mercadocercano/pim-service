package mapper

import (
	"pim/src/category_attribute/domain/entity"
	"pim/src/category_attribute/infrastructure/persistence/model"
)

// CategoryAttributeMapper maneja las conversiones entre entidad y modelo
type CategoryAttributeMapper struct{}

// NewCategoryAttributeMapper crea una nueva instancia del mapper
func NewCategoryAttributeMapper() *CategoryAttributeMapper {
	return &CategoryAttributeMapper{}
}

// ToEntity convierte un modelo de base de datos a una entidad de dominio
func (m *CategoryAttributeMapper) ToEntity(model *model.CategoryAttribute) *entity.CategoryAttribute {
	if model == nil {
		return nil
	}

	return &entity.CategoryAttribute{
		ID:            model.ID,
		TenantID:      model.TenantID,
		CategoryID:    model.CategoryID,
		AttributeID:   model.AttributeID,
		AllowedValues: []string(model.AllowedValues),
		Status:        model.Status,
		CreatedAt:     model.CreatedAt,
		UpdatedAt:     model.UpdatedAt,
	}
}

// ToModel convierte una entidad de dominio a un modelo de base de datos
func (m *CategoryAttributeMapper) ToModel(entity *entity.CategoryAttribute) *model.CategoryAttribute {
	if entity == nil {
		return nil
	}

	return &model.CategoryAttribute{
		ID:            entity.ID,
		TenantID:      entity.TenantID,
		CategoryID:    entity.CategoryID,
		AttributeID:   entity.AttributeID,
		AllowedValues: entity.AllowedValues,
		Status:        entity.Status,
		CreatedAt:     entity.CreatedAt,
		UpdatedAt:     entity.UpdatedAt,
	}
}

// ToEntityList convierte una lista de modelos a una lista de entidades
func (m *CategoryAttributeMapper) ToEntityList(models []*model.CategoryAttribute) []*entity.CategoryAttribute {
	if models == nil {
		return nil
	}

	entities := make([]*entity.CategoryAttribute, len(models))
	for i, model := range models {
		entities[i] = m.ToEntity(model)
	}
	return entities
}

// ToModelList convierte una lista de entidades a una lista de modelos
func (m *CategoryAttributeMapper) ToModelList(entities []*entity.CategoryAttribute) []*model.CategoryAttribute {
	if entities == nil {
		return nil
	}

	models := make([]*model.CategoryAttribute, len(entities))
	for i, entity := range entities {
		models[i] = m.ToModel(entity)
	}
	return models
}
