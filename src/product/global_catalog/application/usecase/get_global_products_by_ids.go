package usecase

import (
	"context"
	"errors"

	"saas-mt-pim-service/src/product/global_catalog/domain/entity"
	"saas-mt-pim-service/src/product/global_catalog/domain/exception"
	"saas-mt-pim-service/src/product/global_catalog/domain/port"
)

const maxIDsPerRequest = 100

// GetGlobalProductsByIDsRequest parámetros del endpoint on-demand por IDs.
type GetGlobalProductsByIDsRequest struct {
	IDs []string `json:"ids"`
}

// GetGlobalProductsByIDsResponse respuesta compatible con NeedsEnrichmentResult de webdata.
type GetGlobalProductsByIDsResponse struct {
	Products []EnrichmentItemDTO `json:"products"`
	Total    int                 `json:"total"`
}

// EnrichmentItemDTO representa un producto en el formato que espera webdata-service.
type EnrichmentItemDTO struct {
	ID           string `json:"id"`
	EAN          string `json:"ean"`
	Name         string `json:"name"`
	Brand        string `json:"brand"`
	Category     string `json:"category"`
	BusinessType string `json:"business_type"`
	QualityScore int    `json:"quality_score"`
}

// GetGlobalProductsByIDs devuelve productos por lista de IDs para on-demand enrichment.
type GetGlobalProductsByIDs struct {
	repo port.GlobalProductRepository
}

// NewGetGlobalProductsByIDs crea el use case.
func NewGetGlobalProductsByIDs(repo port.GlobalProductRepository) *GetGlobalProductsByIDs {
	return &GetGlobalProductsByIDs{repo: repo}
}

// Execute retorna los productos correspondientes a los IDs recibidos.
func (uc *GetGlobalProductsByIDs) Execute(ctx context.Context, req GetGlobalProductsByIDsRequest) (*GetGlobalProductsByIDsResponse, error) {
	if len(req.IDs) == 0 {
		return nil, exception.NewValidationError("la lista de IDs no puede estar vacía", errors.New("ids requeridos"))
	}

	if len(req.IDs) > maxIDsPerRequest {
		return nil, exception.NewValidationError("se permiten máximo 100 IDs por request", errors.New("límite excedido"))
	}

	products, err := uc.repo.FindByIDs(ctx, req.IDs)
	if err != nil {
		return nil, exception.NewInternalError("error buscando productos por IDs", err)
	}

	items := make([]EnrichmentItemDTO, len(products))
	for i, p := range products {
		items[i] = productToEnrichmentDTO(p)
	}

	return &GetGlobalProductsByIDsResponse{
		Products: items,
		Total:    len(items),
	}, nil
}

func productToEnrichmentDTO(p *entity.GlobalProduct) EnrichmentItemDTO {
	dto := EnrichmentItemDTO{
		ID:           p.IDString(),
		EAN:          p.EANString(),
		Name:         p.Name(),
		QualityScore: p.QualityScore().Value(),
	}
	if b := p.Brand(); b != nil {
		dto.Brand = *b
	}
	if c := p.Category(); c != nil {
		dto.Category = *c
	}
	if bt := p.BusinessType(); bt != nil {
		dto.BusinessType = *bt
	}
	return dto
}
