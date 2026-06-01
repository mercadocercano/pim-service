package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"saas-mt-pim-service/src/attribute/domain/entity"
)

// AttributeValuePostgresRepository implementa AttributeValueRepository usando PostgreSQL
type AttributeValuePostgresRepository struct {
	db *sql.DB
}

// NewAttributeValuePostgresRepository crea una nueva instancia del repositorio
func NewAttributeValuePostgresRepository(db *sql.DB) *AttributeValuePostgresRepository {
	return &AttributeValuePostgresRepository{db: db}
}

// Create inserta un nuevo valor en la tabla marketplace_attribute_values
func (r *AttributeValuePostgresRepository) Create(ctx context.Context, v *entity.AttributeValue) error {
	query := `
		INSERT INTO marketplace_attribute_values
			(id, attribute_id, value, slug, sort_order, is_active, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.db.ExecContext(ctx, query,
		v.ID, v.AttributeID, v.Value, v.Slug, v.SortOrder, v.IsActive, v.CreatedAt,
	)
	if err != nil {
		return fmt.Errorf("error creating attribute value: %w", err)
	}
	return nil
}

// Update actualiza value y sort_order de un registro existente
func (r *AttributeValuePostgresRepository) Update(ctx context.Context, id, newValue string, sortOrder int) (*entity.AttributeValue, error) {
	slug := generateValueSlug(newValue)
	query := `
		UPDATE marketplace_attribute_values
		SET value = $2, slug = $3, sort_order = $4
		WHERE id = $1
		RETURNING id, attribute_id, value, slug, sort_order, is_active, created_at
	`
	row := r.db.QueryRowContext(ctx, query, id, newValue, slug, sortOrder)
	return scanAttributeValue(row)
}

// FindByID busca un valor por su ID
func (r *AttributeValuePostgresRepository) FindByID(ctx context.Context, id string) (*entity.AttributeValue, error) {
	query := `
		SELECT id, attribute_id, value, slug, sort_order, is_active, created_at
		FROM marketplace_attribute_values
		WHERE id = $1
	`
	row := r.db.QueryRowContext(ctx, query, id)
	v, err := scanAttributeValue(row)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// FindByAttributeID retorna todos los valores activos de un atributo ordenados por sort_order
func (r *AttributeValuePostgresRepository) FindByAttributeID(ctx context.Context, attributeID string) ([]*entity.AttributeValue, error) {
	query := `
		SELECT id, attribute_id, value, slug, sort_order, is_active, created_at
		FROM marketplace_attribute_values
		WHERE attribute_id = $1
		ORDER BY sort_order ASC, value ASC
	`
	rows, err := r.db.QueryContext(ctx, query, attributeID)
	if err != nil {
		return nil, fmt.Errorf("error listing attribute values: %w", err)
	}
	defer rows.Close()

	var values []*entity.AttributeValue
	for rows.Next() {
		var v entity.AttributeValue
		if err := rows.Scan(&v.ID, &v.AttributeID, &v.Value, &v.Slug, &v.SortOrder, &v.IsActive, &v.CreatedAt); err != nil {
			return nil, fmt.Errorf("error scanning attribute value: %w", err)
		}
		values = append(values, &v)
	}
	return values, nil
}

// Delete elimina un valor por ID
func (r *AttributeValuePostgresRepository) Delete(ctx context.Context, id string) error {
	result, err := r.db.ExecContext(ctx, `DELETE FROM marketplace_attribute_values WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("error deleting attribute value: %w", err)
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("attribute value not found")
	}
	return nil
}

// IsAttributeInUse verifica si alguna variante usa el atributo dado
func (r *AttributeValuePostgresRepository) IsAttributeInUse(ctx context.Context, attributeID string) (bool, error) {
	var count int
	query := `SELECT COUNT(*) FROM variant_marketplace_attributes WHERE marketplace_attribute_id = $1`
	err := r.db.QueryRowContext(ctx, query, attributeID).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("error checking attribute usage: %w", err)
	}
	return count > 0, nil
}

// scanAttributeValue escanea una fila en AttributeValue; retorna nil sin error si no existe
func scanAttributeValue(row *sql.Row) (*entity.AttributeValue, error) {
	var v entity.AttributeValue
	var createdAt time.Time
	err := row.Scan(&v.ID, &v.AttributeID, &v.Value, &v.Slug, &v.SortOrder, &v.IsActive, &createdAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error scanning attribute value: %w", err)
	}
	v.CreatedAt = createdAt
	return &v, nil
}

// generateValueSlug genera el slug de un value (misma lógica que la entidad)
func generateValueSlug(value string) string {
	slug := strings.ToLower(strings.TrimSpace(value))
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.ReplaceAll(slug, "á", "a")
	slug = strings.ReplaceAll(slug, "é", "e")
	slug = strings.ReplaceAll(slug, "í", "i")
	slug = strings.ReplaceAll(slug, "ó", "o")
	slug = strings.ReplaceAll(slug, "ú", "u")
	slug = strings.ReplaceAll(slug, "ñ", "n")
	return slug
}
