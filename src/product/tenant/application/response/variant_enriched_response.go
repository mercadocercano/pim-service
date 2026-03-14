package response

import "github.com/google/uuid"

// VariantEnrichedResponse is a slim variant response enriched with product and category data,
// designed for cross-service consumption (e.g., catalog-bff inventory report).
type VariantEnrichedResponse struct {
	VariantID    uuid.UUID `json:"variant_id"`
	ProductID    uuid.UUID `json:"product_id"`
	SKU          string    `json:"sku"`
	VariantName  string    `json:"variant_name"`
	ProductName  string    `json:"product_name"`
	CategoryID   *string   `json:"category_id,omitempty"`
	CategoryName string    `json:"category_name"`
	Price        float64   `json:"price"`
}
