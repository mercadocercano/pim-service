package response

import (
	"time"
)

// ProductResponse representa la respuesta de un producto
type ProductResponse struct {
	ID          string                     `json:"id"`
	Name        string                     `json:"name"`
	Description *string                    `json:"description,omitempty"`
	ImageURL    *string                    `json:"image_url,omitempty"`
	SKU         *string                    `json:"sku,omitempty"`
	Category    *CategoryReferenceResponse `json:"category,omitempty"`
	Brand       *BrandReferenceResponse    `json:"brand,omitempty"`
	Status      string                     `json:"status"`
	Price       *float64                   `json:"price,omitempty"`
	Stock       *int                       `json:"stock,omitempty"`
	Variants    []ProductVariantResponse   `json:"variants,omitempty"`
	CreatedAt   time.Time                  `json:"created_at"`
	UpdatedAt   time.Time                  `json:"updated_at"`
}

// ProductWithVariantsResponse representa un producto con todas sus variantes
type ProductWithVariantsResponse struct {
	Product        ProductResponse          `json:"product"`
	DefaultVariant *ProductVariantResponse  `json:"default_variant,omitempty"`
	Variants       []ProductVariantResponse `json:"variants"`
	TotalVariants  int                      `json:"total_variants"`
}

// CategoryReferenceResponse representa una referencia a categoría
type CategoryReferenceResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// BrandReferenceResponse representa una referencia a marca
type BrandReferenceResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ProductListResponse representa una lista paginada de productos
type ProductListResponse struct {
	Products   []*ProductResponse `json:"products"`
	Pagination PaginationResponse `json:"pagination"`
}
