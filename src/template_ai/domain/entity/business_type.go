package entity

import (
	"time"
	"github.com/gofrs/uuid/v5"
)

// BusinessType represents a type of business
type BusinessType struct {
	ID          uuid.UUID
	Code        string
	Name        string
	Description string
	IsActive    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}