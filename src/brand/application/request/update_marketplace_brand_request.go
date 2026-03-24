package request

// UpdateMarketplaceBrandRequest representa la petición para actualizar una marca marketplace global
type UpdateMarketplaceBrandRequest struct {
	ID           string   `json:"id" validate:"required"`
	Name         string   `json:"name" validate:"required,min=2,max=100"`
	Description  string   `json:"description"`
	LogoURL      string   `json:"logo_url" validate:"omitempty,url"`
	Website      string   `json:"website" validate:"omitempty,url"`
	Aliases      []string `json:"aliases"`
	CategoryTags []string `json:"category_tags"`
	Sources      []string `json:"sources"`
	QualityScore float64  `json:"quality_score"`
	IsActive     bool     `json:"is_active"`
}
