package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"saas-mt-pim-service/src/quickstart/domain/port"
)

// AttributeServiceImpl implementa AttributeService para crear atributos desde el quickstart
type AttributeServiceImpl struct {
	db *sql.DB
}

// NewAttributeService crea una nueva instancia del servicio de atributos
func NewAttributeService(db *sql.DB) port.AttributeService {
	return &AttributeServiceImpl{
		db: db,
	}
}

// CreateFromTemplate crea atributos desde los datos de template del quickstart
func (s *AttributeServiceImpl) CreateFromTemplate(ctx context.Context, tenantID string, templateData interface{}) error {
	// Los atributos predefinidos para retail/polirubro
	attributes := []struct {
		Name        string
		Description string
		Type        string
	}{
		{"color", "Color", "select"},
		{"material", "Material", "select"},
		{"marca", "Marca", "text"},
		{"peso", "Peso", "number"},
		{"dimensiones", "Dimensiones", "text"},
		{"garantia", "Garantía", "text"},
		{"pais-origen", "País de Origen", "select"},
		{"eco-friendly", "Eco-Amigable", "boolean"},
		{"precio", "Precio", "number"},
		{"disponible", "Disponible", "boolean"},
	}

	// Verificar si ya existen atributos para este tenant
	var count int
	checkQuery := "SELECT COUNT(*) FROM attributes WHERE tenant_id = $1"
	err := s.db.QueryRowContext(ctx, checkQuery, tenantID).Scan(&count)
	if err != nil {
		return fmt.Errorf("error verificando atributos existentes: %w", err)
	}

	// Si ya existen atributos para este tenant, no crear duplicados
	if count > 0 {
		return nil
	}

	// Crear los atributos para el tenant específico
	insertQuery := `
		INSERT INTO attributes (tenant_id, name, description, type, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, 'active', $5, $5)
	`

	now := time.Now()
	for _, attribute := range attributes {
		_, err := s.db.ExecContext(ctx, insertQuery, tenantID, attribute.Name, attribute.Description, attribute.Type, now)
		if err != nil {
			return fmt.Errorf("error creando atributo %s: %w", attribute.Name, err)
		}
	}

	return nil
}
