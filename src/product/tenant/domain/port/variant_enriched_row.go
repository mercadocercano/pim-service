package port

import "github.com/google/uuid"

// VariantEnrichedRow represents a read-model row for variant + product + category data.
// Used by outbound ports that need to return denormalized data for cross-service queries.
type VariantEnrichedRow struct {
	VariantID    uuid.UUID
	ProductID    uuid.UUID
	SKU          string
	VariantName  string
	ProductName  string
	CategoryID   *string
	CategoryName string
	Price        float64
}
