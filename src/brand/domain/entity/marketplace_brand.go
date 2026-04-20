package entity

import (
	"fmt"
	"regexp"
	"time"

	"github.com/google/uuid"
)

// hexColorRegexp valida formato #RRGGBB
var hexColorRegexp = regexp.MustCompile(`^#[0-9A-Fa-f]{6}$`)

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
	BackgroundColor    string    `json:"background_color"`
	TextColor          string    `json:"text_color"`
	Typography         string    `json:"typography"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

// VisualIdentityParams agrupa los campos de identidad visual para evitar listas largas de parámetros.
type VisualIdentityParams struct {
	BackgroundColor string
	TextColor       string
	Typography      string
}

// validateHexColor devuelve true si s es "" (sin color) o un hex válido #RRGGBB.
func validateHexColor(s string) bool {
	if s == "" {
		return true
	}
	return hexColorRegexp.MatchString(s)
}

// validateTypography devuelve true si s es "" o tiene entre 1 y 100 caracteres no vacíos.
func validateTypography(s string) bool {
	if s == "" {
		return true
	}
	return len([]rune(s)) <= 100
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
		BackgroundColor:    "",
		TextColor:          "",
		Typography:         "",
		CreatedAt:          now,
		UpdatedAt:          now,
	}, nil
}

// SetVisualIdentity valida y asigna los campos de identidad visual.
// String vacío ("") es válido — significa usar el fallback del design system.
func (e *Marketplacebrand) SetVisualIdentity(params VisualIdentityParams) error {
	if !validateHexColor(params.BackgroundColor) {
		return fmt.Errorf("background_color inválido: debe ser #RRGGBB o vacío, recibido: %q", params.BackgroundColor)
	}
	if !validateHexColor(params.TextColor) {
		return fmt.Errorf("text_color inválido: debe ser #RRGGBB o vacío, recibido: %q", params.TextColor)
	}
	if !validateTypography(params.Typography) {
		return fmt.Errorf("typography inválido: máximo 100 caracteres, recibido: %d", len([]rune(params.Typography)))
	}

	e.BackgroundColor = params.BackgroundColor
	e.TextColor = params.TextColor
	e.Typography = params.Typography
	e.UpdatedAt = time.Now()
	return nil
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
