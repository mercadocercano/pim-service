package response

import (
	"time"

	"github.com/google/uuid"
)

// ProductVariantResponse representa la respuesta de una variante de producto
type ProductVariantResponse struct {
	ID         uuid.UUID                  `json:"id"`
	ProductID  uuid.UUID                  `json:"product_id"`
	Name       string                     `json:"name"`
	SKU        *string                    `json:"sku,omitempty"`
	Status     string                     `json:"status"`
	IsDefault  bool                       `json:"is_default"`
	SortOrder  int                        `json:"sort_order"`
	Price      float64                    `json:"price"`
	Stock      int                        `json:"stock"`
	Attributes []VariantAttributeResponse `json:"attributes"`
	CreatedAt  time.Time                  `json:"created_at"`
	UpdatedAt  time.Time                  `json:"updated_at"`
}

// VariantAttributeResponse representa un atributo de variante en la respuesta
type VariantAttributeResponse struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// ProductVariantListResponse representa una lista paginada de variantes de productos
type ProductVariantListResponse struct {
	Variants   []ProductVariantResponse `json:"variants"`
	Pagination PaginationResponse       `json:"pagination"`
}

// PaginationResponse representa la información de paginación
type PaginationResponse struct {
	Page       int `json:"page"`
	PageSize   int `json:"page_size"`
	TotalItems int `json:"total_items"`
	TotalPages int `json:"total_pages"`
}

// ProductVariantsByProductResponse representa las variantes agrupadas por producto
type ProductVariantsByProductResponse struct {
	ProductID      uuid.UUID                `json:"product_id"`
	ProductName    string                   `json:"product_name"`
	DefaultVariant *ProductVariantResponse  `json:"default_variant,omitempty"`
	Variants       []ProductVariantResponse `json:"variants"`
	TotalVariants  int                      `json:"total_variants"`
}

// VariantStatusChangeResponse representa la respuesta de un cambio de estado
type VariantStatusChangeResponse struct {
	ID        uuid.UUID `json:"id"`
	Status    string    `json:"status"`
	UpdatedAt time.Time `json:"updated_at"`
	Message   string    `json:"message"`
}
