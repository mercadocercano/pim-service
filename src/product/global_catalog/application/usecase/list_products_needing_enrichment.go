package usecase

import (
	"saas-mt-pim-service/src/product/global_catalog/domain/exception"
	"saas-mt-pim-service/src/product/global_catalog/domain/port"
)

// ListProductsNeedingEnrichmentRequest parámetros de la cola de enrichment
type ListProductsNeedingEnrichmentRequest struct {
	BusinessType *string `json:"business_type,omitempty"`
	Limit        int     `json:"limit,omitempty"`
	Offset       int     `json:"offset,omitempty"`
}

// ListProductsNeedingEnrichmentResponse respuesta paginada de productos a enriquecer
type ListProductsNeedingEnrichmentResponse struct {
	Products   []*GlobalProductSummary `json:"products"`
	Pagination *PaginationInfo         `json:"pagination"`
}

// ListProductsNeedingEnrichment lista productos con datos incompletos para que
// webdata-service los encole para scraping.
type ListProductsNeedingEnrichment struct {
	globalProductRepository port.GlobalProductRepository
}

// NewListProductsNeedingEnrichment crea el use case
func NewListProductsNeedingEnrichment(repo port.GlobalProductRepository) *ListProductsNeedingEnrichment {
	return &ListProductsNeedingEnrichment{globalProductRepository: repo}
}

// Execute retorna productos con quality_score < 70 o campos nulos (price, image_url, brand).
func (uc *ListProductsNeedingEnrichment) Execute(req ListProductsNeedingEnrichmentRequest) (*ListProductsNeedingEnrichmentResponse, error) {
	limit := req.Limit
	if limit <= 0 || limit > 500 {
		limit = 100
	}

	offset := req.Offset
	if offset < 0 {
		offset = 0
	}

	products, err := uc.globalProductRepository.FindNeedingEnrichment(req.BusinessType, limit, offset)
	if err != nil {
		return nil, exception.NewInternalError("error listando productos para enrichment", err)
	}

	total, err := uc.globalProductRepository.CountNeedingEnrichment(req.BusinessType)
	if err != nil {
		return nil, exception.NewInternalError("error contando productos para enrichment", err)
	}

	summaries := make([]*GlobalProductSummary, len(products))
	for i, p := range products {
		summaries[i] = &GlobalProductSummary{
			ID:                p.IDString(),
			EAN:               p.EANString(),
			Name:              p.Name(),
			Brand:             p.Brand(),
			Category:          p.Category(),
			Price:             p.Price(),
			ImageURL:          p.ImageURL(),
			QualityScore:      p.QualityScore().Value(),
			QualityLevel:      p.QualityScore().Level(),
			Source:            p.Source().Source(),
			SourceDisplayName: p.Source().GetSourceDisplayName(),
			IsVerified:        p.IsVerified(),
			IsActive:          p.IsActive(),
			BusinessType:      p.BusinessType(),
			Tags:              p.Tags(),
			CreatedAt:         p.CreatedAt().Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt:         p.UpdatedAt().Format("2006-01-02T15:04:05Z07:00"),
		}
	}

	pagination := &PaginationInfo{
		Offset:     offset,
		Limit:      limit,
		Total:      total,
		HasNext:    offset+limit < total,
		HasPrev:    offset > 0,
		TotalPages: (total + limit - 1) / limit,
	}

	return &ListProductsNeedingEnrichmentResponse{
		Products:   summaries,
		Pagination: pagination,
	}, nil
}
