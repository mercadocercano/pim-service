package response

import (
	"saas-mt-pim-service/src/brand/domain/entity"
	"time"
)

// MarketplaceBrandResponse representa la respuesta de una marca marketplace global
type MarketplaceBrandResponse struct {
	ID                 string   `json:"id"`
	Name               string   `json:"name"`
	Slug               string   `json:"slug"`
	NormalizedName     string   `json:"normalized_name"`
	Description        string   `json:"description"`
	LogoURL            string   `json:"logo_url"`
	Website            string   `json:"website"`
	CategoryTags       []string `json:"category_tags"`
	VerificationStatus string   `json:"verification_status"`
	QualityScore       float64  `json:"quality_score"`
	ProductCount       int      `json:"product_count"`
	Aliases            []string `json:"aliases"`
	Sources            []string `json:"sources"`
	IsActive           bool     `json:"is_active"`
	BackgroundColor    string   `json:"background_color"`
	TextColor          string   `json:"text_color"`
	Typography         string   `json:"typography"`
	CreatedAt          string   `json:"created_at"`
	UpdatedAt          string   `json:"updated_at"`
}

// NewMarketplaceBrandResponse crea una respuesta desde una entidad
func NewMarketplaceBrandResponse(brand *entity.Marketplacebrand) *MarketplaceBrandResponse {
	return &MarketplaceBrandResponse{
		ID:                 brand.ID,
		Name:               brand.Name,
		Slug:               generateSlug(brand.Name),
		NormalizedName:     generateNormalizedName(brand.Name),
		Description:        brand.Description,
		LogoURL:            brand.LogoURL,
		Website:            brand.Website,
		CategoryTags:       brand.CategoryTags,
		VerificationStatus: brand.VerificationStatus,
		QualityScore:       brand.QualityScore,
		ProductCount:       brand.ProductCount,
		Aliases:            brand.Aliases,
		Sources:            brand.Sources,
		IsActive:           brand.IsActive,
		BackgroundColor:    brand.BackgroundColor,
		TextColor:          brand.TextColor,
		Typography:         brand.Typography,
		CreatedAt:          brand.CreatedAt.Format(time.RFC3339),
		UpdatedAt:          brand.UpdatedAt.Format(time.RFC3339),
	}
}

// generateSlug genera un slug desde el nombre
func generateSlug(name string) string {
	// Implementación simple - en producción usar una librería más robusta
	return name // TODO: implementar generación de slug real
}

// generateNormalizedName genera un nombre normalizado
func generateNormalizedName(name string) string {
	// Implementación simple - en producción usar normalización más robusta
	return name // TODO: implementar normalización real
}
