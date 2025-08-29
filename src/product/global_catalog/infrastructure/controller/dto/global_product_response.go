package dto

import (
	"saas-mt-pim-service/src/product/global_catalog/domain/entity"
	"time"
)

// GlobalProductResponse representa la respuesta JSON de un producto global
type GlobalProductResponse struct {
	ID            string                 `json:"id"`
	EAN           *string                `json:"ean,omitempty"`
	Name          string                 `json:"name"`
	Description   *string                `json:"description"`
	Brand         *string                `json:"brand"`
	Category      *string                `json:"category"`
	Price         *float64               `json:"price"`
	ImageURL      *string                `json:"image_url"`
	ImageURLs     []string               `json:"image_urls"`
	Source        string                 `json:"source"`
	SourceURL     *string                `json:"source_url"`
	Reliability   float64                `json:"reliability"`
	QualityScore  int                    `json:"quality_score"`
	IsVerified    bool                   `json:"is_verified"`
	IsActive      bool                   `json:"is_active"`
	BusinessType  *string                `json:"business_type"`
	Tags          []string               `json:"tags"`
	Metadata      map[string]interface{} `json:"metadata"`
	CreatedAt     time.Time              `json:"created_at"`
	UpdatedAt     time.Time              `json:"updated_at"`
	LastScrapedAt *time.Time             `json:"last_scraped_at"`
}

// FromEntity convierte una entidad GlobalProduct a DTO
func (dto *GlobalProductResponse) FromEntity(product *entity.GlobalProduct) *GlobalProductResponse {
	if product == nil {
		return nil
	}

	var sourceURL *string
	var reliability float64
	if product.Source() != nil {
		sourceURL = product.Source().SourceURL()
		reliability = product.Source().Reliability()
	}

	var qualityScore int
	if product.QualityScore() != nil {
		qualityScore = product.QualityScore().Value()
	}

	var ean *string
	if product.EAN() != nil {
		eanValue := product.EAN().Value()
		ean = &eanValue
	}

	var source string
	if product.Source() != nil {
		source = product.Source().Source()
	}

	return &GlobalProductResponse{
		ID:            product.IDString(),
		EAN:           ean,
		Name:          product.Name(),
		Description:   product.Description(),
		Brand:         product.Brand(),
		Category:      product.Category(),
		Price:         product.Price(),
		ImageURL:      product.ImageURL(),
		ImageURLs:     product.ImageURLs(),
		Source:        source,
		SourceURL:     sourceURL,
		Reliability:   reliability,
		QualityScore:  qualityScore,
		IsVerified:    product.IsVerified(),
		IsActive:      product.IsActive(),
		BusinessType:  product.BusinessType(),
		Tags:          product.Tags(),
		Metadata:      product.Metadata(),
		CreatedAt:     product.CreatedAt(),
		UpdatedAt:     product.UpdatedAt(),
		LastScrapedAt: product.LastScrapedAt(),
	}
}

// NewGlobalProductResponse crea un nuevo DTO desde una entidad
func NewGlobalProductResponse(product *entity.GlobalProduct) *GlobalProductResponse {
	return new(GlobalProductResponse).FromEntity(product)
}
