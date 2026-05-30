package response

import "time"

// BrandResponse representa la respuesta de una marca
type BrandResponse struct {
	ID          string    `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Name        string    `json:"name" example:"Nike"`
	Description string    `json:"description" example:"Marca deportiva internacional"`
	LogoURL     *string   `json:"logo_url,omitempty" example:"https://example.com/logo.png"`
	Website     *string   `json:"website,omitempty" example:"https://nike.com"`
	Color       string    `json:"color" example:"#FF5733"`
	Status      string    `json:"status" example:"active"`
	CreatedAt   time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt   time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z"`
}

// BrandListResponse representa la respuesta de una lista de marcas
type BrandListResponse struct {
	Items      []*BrandResponse `json:"items"`
	TotalCount int              `json:"total_count" example:"100"`
	Page       int              `json:"page" example:"1"`
	PageSize   int              `json:"page_size" example:"10"`
	TotalPages int              `json:"total_pages" example:"10"`
}

// BrandReferenceResponse representa una referencia simplificada de marca
type BrandReferenceResponse struct {
	ID          string `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Name        string `json:"name" example:"Nike"`
	Description string `json:"description" example:"Marca deportiva internacional"`
}
