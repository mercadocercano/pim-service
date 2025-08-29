package entity

import (
	"time"
	"github.com/gofrs/uuid/v5"
)

// ProductFeedback represents feedback on a product in a template
type ProductFeedback struct {
	ID                   uuid.UUID
	TenantID             uuid.UUID
	TemplateID           uuid.UUID
	GlobalProductID      uuid.UUID
	ProductName          string
	Action               string // add, remove, increase, decrease
	Reason               string
	OriginalQuantity     int
	PreviousQuantity     int
	NewQuantity          int
	ReplacementProductID *uuid.UUID
	FeedbackReason       string
	CreatedBy            string
	CreatedAt            time.Time
}