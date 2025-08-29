package entity

import (
	"time"
	"github.com/gofrs/uuid/v5"
)

// GlobalProduct represents a product from the global catalog
type GlobalProduct struct {
	ID           uuid.UUID
	Name         string
	CategoryID   string
	CategoryName string
	BrandID      string
	BrandName    string
	Price        float64
	IsBulk       bool
	IsPerishable bool
	Tags         []string
	Attributes   map[string]interface{}
	CreatedAt    time.Time
	UpdatedAt    time.Time
}