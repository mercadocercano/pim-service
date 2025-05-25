package mapper

import (
	"pim/src/brand/application/request"
	"pim/src/brand/application/response"
	"pim/src/brand/domain/entity"
	"pim/src/brand/domain/value_object"
	"pim/src/shared/domain/criteria"
)

// BrandMapper maneja las conversiones entre entidades y DTOs
type BrandMapper struct{}

// NewBrandMapper crea una nueva instancia del mapper
func NewBrandMapper() *BrandMapper {
	return &BrandMapper{}
}

// ToResponse convierte una entidad Brand a BrandResponse
func (m *BrandMapper) ToResponse(brand *entity.Brand) *response.BrandResponse {
	if brand == nil {
		return nil
	}

	return &response.BrandResponse{
		ID:          brand.ID,
		Name:        brand.Name,
		Description: brand.Description,
		LogoURL:     brand.LogoURL,
		Website:     brand.Website,
		Status:      brand.Status.String(),
		CreatedAt:   brand.CreatedAt,
		UpdatedAt:   brand.UpdatedAt,
	}
}

// ToResponseList convierte una lista de entidades Brand a BrandListResponse
func (m *BrandMapper) ToResponseList(brands []*entity.Brand, totalCount int, crit criteria.Criteria) *response.BrandListResponse {
	items := make([]*response.BrandResponse, len(brands))
	for i, brand := range brands {
		items[i] = m.ToResponse(brand)
	}

	totalPages := (totalCount + crit.Pagination.PageSize - 1) / crit.Pagination.PageSize

	return &response.BrandListResponse{
		Items:      items,
		TotalCount: totalCount,
		Page:       crit.Pagination.Page,
		PageSize:   crit.Pagination.PageSize,
		TotalPages: totalPages,
	}
}

// ToReferenceResponse convierte una BrandReference a BrandReferenceResponse
func (m *BrandMapper) ToReferenceResponse(ref *value_object.BrandReference) *response.BrandReferenceResponse {
	if ref == nil {
		return nil
	}

	return &response.BrandReferenceResponse{
		ID:          ref.ID,
		Name:        ref.Name,
		Description: ref.Description,
	}
}

// FromCreateRequest convierte un CreateBrandRequest a entidad Brand
func (m *BrandMapper) FromCreateRequest(req *request.CreateBrandRequest, tenantID string) (*entity.Brand, error) {
	return entity.NewBrand(tenantID, req.Name, req.Description, req.LogoURL, req.Website)
}

// ApplyUpdateRequest aplica los cambios de un UpdateBrandRequest a una entidad Brand
func (m *BrandMapper) ApplyUpdateRequest(brand *entity.Brand, req *request.UpdateBrandRequest) error {
	return brand.Update(req.Name, req.Description, req.LogoURL, req.Website)
}
