package mapper

import (
	"strings"
	
	"saas-mt-pim-service/src/category/domain/entity"
	"saas-mt-pim-service/src/category/infrastructure/persistence/model"
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

	// Generar slug desde el nombre si no existe
	slug := generateSlug(entity.Name)

	return &model.Category{
		ID:          entity.ID,
		TenantID:    entity.TenantID,
		Name:        entity.Name,
		Slug:        slug,
		Description: entity.Description,
		ParentID:    entity.ParentID,
		Status:      entity.Status,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}

// generateSlug genera un slug desde un nombre
func generateSlug(name string) string {
	// Simplificado: lowercase y reemplazar espacios por guiones
	slug := name
	slug = strings.ToLower(slug)
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.ReplaceAll(slug, "á", "a")
	slug = strings.ReplaceAll(slug, "é", "e")
	slug = strings.ReplaceAll(slug, "í", "i")
	slug = strings.ReplaceAll(slug, "ó", "o")
	slug = strings.ReplaceAll(slug, "ú", "u")
	slug = strings.ReplaceAll(slug, "ñ", "n")
	return slug
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
