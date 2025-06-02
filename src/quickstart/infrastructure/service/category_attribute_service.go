package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"pim/src/quickstart/domain/port"
)

// CategoryAttributeServiceImpl implementa CategoryAttributeService para crear relaciones categoría-atributo
type CategoryAttributeServiceImpl struct {
	db *sql.DB
}

// NewCategoryAttributeService crea una nueva instancia del servicio de relaciones categoría-atributo
func NewCategoryAttributeService(db *sql.DB) port.CategoryAttributeService {
	return &CategoryAttributeServiceImpl{
		db: db,
	}
}

// CreateFromTemplate crea relaciones categoría-atributo desde los datos del quickstart
func (s *CategoryAttributeServiceImpl) CreateFromTemplate(ctx context.Context, tenantID string, templateData interface{}) error {
	// Verificar si ya existen relaciones para este tenant
	var count int
	checkQuery := "SELECT COUNT(*) FROM category_attributes WHERE tenant_id = $1"
	err := s.db.QueryRowContext(ctx, checkQuery, tenantID).Scan(&count)
	if err != nil {
		return fmt.Errorf("error verificando relaciones categoría-atributo existentes: %w", err)
	}

	// Si ya existen relaciones para este tenant, no crear duplicados
	if count > 0 {
		fmt.Printf("Ya existen %d relaciones categoría-atributo para el tenant %s, omitiendo creación\n", count, tenantID)
		return nil
	}

	// Obtener categorías del tenant
	categoriesQuery := "SELECT id, name FROM categories WHERE tenant_id = $1"
	categoryRows, err := s.db.QueryContext(ctx, categoriesQuery, tenantID)
	if err != nil {
		return fmt.Errorf("error obteniendo categorías: %w", err)
	}
	defer categoryRows.Close()

	categories := make(map[string]string) // id -> name
	for categoryRows.Next() {
		var id, name string
		if err := categoryRows.Scan(&id, &name); err != nil {
			return fmt.Errorf("error escaneando categoría: %w", err)
		}
		categories[id] = name
	}

	// Obtener atributos del tenant
	attributesQuery := "SELECT id, name FROM attributes WHERE tenant_id = $1"
	attributeRows, err := s.db.QueryContext(ctx, attributesQuery, tenantID)
	if err != nil {
		return fmt.Errorf("error obteniendo atributos: %w", err)
	}
	defer attributeRows.Close()

	attributes := make(map[string]string) // name -> id
	for attributeRows.Next() {
		var id, name string
		if err := attributeRows.Scan(&id, &name); err != nil {
			return fmt.Errorf("error escaneando atributo: %w", err)
		}
		attributes[name] = id
	}

	// Definir qué atributos pertenecen a cada categoría
	categoryAttributeMap := map[string][]string{
		"Hogar y Jardín":                   {"color", "material", "dimensiones", "peso", "eco-friendly", "precio", "disponible"},
		"Salud y Belleza":                  {"marca", "color", "eco-friendly", "precio", "disponible", "pais-origen"},
		"Electrónicos y Electrodomésticos": {"marca", "color", "peso", "dimensiones", "garantia", "precio", "disponible"},
		"Ropa y Accesorios":                {"color", "material", "marca", "dimensiones", "pais-origen", "precio", "disponible"},
		"Alimentos y Bebidas":              {"marca", "peso", "eco-friendly", "pais-origen", "precio", "disponible"},
		"Limpieza del Hogar":               {"marca", "eco-friendly", "peso", "precio", "disponible"},
		"Muebles de Oficina":               {"color", "material", "dimensiones", "peso", "garantia", "precio", "disponible"},
	}

	// Crear las relaciones
	insertQuery := `
		INSERT INTO category_attributes (tenant_id, category_id, attribute_id, status, created_at, updated_at)
		VALUES ($1, $2, $3, 'active', $4, $4)
	`

	now := time.Now()
	relationsCreated := 0

	for categoryID, categoryName := range categories {
		if attributeNames, exists := categoryAttributeMap[categoryName]; exists {
			for _, attributeName := range attributeNames {
				if attributeID, found := attributes[attributeName]; found {
					_, err := s.db.ExecContext(ctx, insertQuery, tenantID, categoryID, attributeID, now)
					if err != nil {
						return fmt.Errorf("error creando relación categoría-atributo %s-%s: %w", categoryName, attributeName, err)
					}
					relationsCreated++
					fmt.Printf("Relación creada: %s -> %s para tenant %s\n", categoryName, attributeName, tenantID)
				}
			}
		}
	}

	fmt.Printf("Se crearon %d relaciones categoría-atributo para el tenant %s\n", relationsCreated, tenantID)
	return nil
}
