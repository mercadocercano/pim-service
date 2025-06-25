package entity

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Marketplacebrand representa la entidad marketplace_brand (tabla global)
type Marketplacebrand struct {
	ID                 string    `json:"id"`
	Name               string    `json:"name"`
	Description        string    `json:"description"`
	LogoURL            string    `json:"logo_url"`
	Website            string    `json:"website"`
	Aliases            []string  `json:"aliases"`
	CategoryTags       []string  `json:"category_tags"`
	QualityScore       float64   `json:"quality_score"`
	ProductCount       int       `json:"product_count"`
	Sources            []string  `json:"sources"`
	IsVerified         bool      `json:"is_verified"`
	VerificationStatus string    `json:"verification_status"`
	IsActive           bool      `json:"is_active"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

// NewMarketplacebrand crea una nueva instancia de Marketplacebrand
func NewMarketplacebrand(name string) (*Marketplacebrand, error) {
	if name == "" {
		return nil, fmt.Errorf("name es requerido")
	}

	now := time.Now()
	return &Marketplacebrand{
		ID:                 uuid.New().String(),
		Name:               name,
		Description:        "",
		LogoURL:            "",
		Website:            "",
		Aliases:            []string{},
		CategoryTags:       []string{},
		QualityScore:       0.5, // Default del schema
		ProductCount:       0,
		Sources:            []string{},
		IsVerified:         false,
		VerificationStatus: "unverified",
		IsActive:           true, // Default activo
		CreatedAt:          now,
		UpdatedAt:          now,
	}, nil
}

// UpdateFields actualiza los campos de la entidad
func (e *Marketplacebrand) UpdateFields(name, description, logoURL, website string, aliases, categoryTags, sources []string, qualityScore float64, isActive bool) {
	e.Name = name
	e.Description = description
	e.LogoURL = logoURL
	e.Website = website
	e.Aliases = aliases
	e.CategoryTags = categoryTags
	e.Sources = sources
	e.QualityScore = qualityScore
	e.IsActive = isActive
	e.UpdatedAt = time.Now()

	// Actualizar verification status basado en is_verified
	if e.IsVerified {
		e.VerificationStatus = "verified"
	} else {
		e.VerificationStatus = "unverified"
	}
}

// Update actualiza los campos de la entidad
func (e *Marketplacebrand) Update() {
	e.UpdatedAt = time.Now()
}
