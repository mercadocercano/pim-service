package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"saas-mt-pim-service/src/quickstart/domain/port"
)


// CategoryServiceImpl implementa CategoryService para crear categorías desde el quickstart
type CategoryServiceImpl struct {
	db *sql.DB
}

// NewCategoryService crea una nueva instancia del servicio de categorías
func NewCategoryService(db *sql.DB) port.CategoryService {
	return &CategoryServiceImpl{
		db: db,
	}
}

// CreateFromTemplate crea categorías desde los datos de template del quickstart
func (s *CategoryServiceImpl) CreateFromTemplate(ctx context.Context, tenantID string, templateData interface{}) error {
	// Las categorías predefinidas para retail/polirubro
	categories := []struct {
		Name        string
		Description string
	}{
		{"Hogar y Jardín", "Productos para el hogar y jardinería"},
		{"Salud y Belleza", "Productos de cuidado personal"},
		{"Electrónicos y Electrodomésticos", "Dispositivos electrónicos y electrodomésticos"},
		{"Ropa y Accesorios", "Vestimenta y accesorios"},
		{"Alimentos y Bebidas", "Productos alimenticios y bebidas"},
		{"Limpieza del Hogar", "Productos para limpieza y cuidado del hogar"},
		{"Muebles de Oficina", "Mobiliario y accesorios para oficinas"},
	}

	// Verificar si ya existen categorías para este tenant
	var count int
	checkQuery := "SELECT COUNT(*) FROM categories WHERE tenant_id = $1"
	err := s.db.QueryRowContext(ctx, checkQuery, tenantID).Scan(&count)
	if err != nil {
		return fmt.Errorf("error verificando categorías existentes: %w", err)
	}

	// Si ya existen categorías para este tenant, no crear duplicados
	if count > 0 {
		return nil
	}

	// Crear las categorías para el tenant específico
	insertQuery := `
		INSERT INTO categories (id, tenant_id, name, description, parent_id, status, created_at, updated_at)
		VALUES (gen_random_uuid(), $1, $2, $3, NULL, 'active', $4, $4)
	`

	now := time.Now()
	for _, category := range categories {
		_, err := s.db.ExecContext(ctx, insertQuery, tenantID, category.Name, category.Description, now)
		if err != nil {
			return fmt.Errorf("error creando categoría %s: %w", category.Name, err)
		}
	}

	return nil
}
