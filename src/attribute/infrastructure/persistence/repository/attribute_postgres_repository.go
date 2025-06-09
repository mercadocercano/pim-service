package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"pim/src/attribute/domain/entity"
	"pim/src/attribute/domain/exception"
	"pim/src/shared/domain/criteria"
	sharedCriteria "pim/src/shared/infrastructure/criteria"
)

// AttributePostgresRepository implementa el repositorio usando PostgreSQL
type AttributePostgresRepository struct {
	db        *sql.DB
	converter *sharedCriteria.SQLCriteriaConverter
}

// NewAttributePostgresRepository crea una nueva instancia del repositorio
func NewAttributePostgresRepository(db *sql.DB) *AttributePostgresRepository {
	return &AttributePostgresRepository{
		db:        db,
		converter: sharedCriteria.NewSQLCriteriaConverter(),
	}
}

// Create crea un nuevo attribute
func (r *AttributePostgresRepository) Create(ctx context.Context, attribute *entity.Attribute) error {
	query := `
		INSERT INTO attributes (
			id, tenant_id, name, active, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6
		)
	`

	_, err := r.db.ExecContext(ctx, query,
		attribute.ID,
		attribute.TenantID,
		attribute.Name,
		attribute.Active,
		attribute.CreatedAt,
		attribute.UpdatedAt,
	)

	if err != nil {
		log.Printf("Error creando attribute: %v", err)
		return fmt.Errorf("%w: %v", exception.ErrAttributeCreateFailed, err)
	}

	return nil
}

// Update actualiza un attribute existente
func (r *AttributePostgresRepository) Update(ctx context.Context, attribute *entity.Attribute) error {
	query := `
		UPDATE attributes SET
			name = $3,
			active = $4,
			updated_at = $5
		WHERE id = $1 AND tenant_id = $2
	`

	result, err := r.db.ExecContext(ctx, query,
		attribute.ID,
		attribute.TenantID,
		attribute.Name,
		attribute.Active,
		attribute.UpdatedAt,
	)

	if err != nil {
		log.Printf("Error actualizando attribute: %v", err)
		return fmt.Errorf("%w: %v", exception.ErrAttributeUpdateFailed, err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return exception.ErrAttributeNotFound
	}

	return nil
}

// FindByID busca un attribute por su ID
func (r *AttributePostgresRepository) FindByID(ctx context.Context, id string, tenantID string) (*entity.Attribute, error) {
	query := `
		SELECT id, tenant_id, name, active, created_at, updated_at
		FROM attributes 
		WHERE id = $1 AND tenant_id = $2
	`

	row := r.db.QueryRowContext(ctx, query, id, tenantID)
	return r.scanAttribute(row)
}

// FindByTenant obtiene todos los attributes de un tenant
func (r *AttributePostgresRepository) FindByTenant(ctx context.Context, tenantID string) ([]*entity.Attribute, error) {
	query := `
		SELECT id, tenant_id, name, active, created_at, updated_at
		FROM attributes 
		WHERE tenant_id = $1 AND active = true
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, tenantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanAttributes(rows)
}

// Delete elimina un attribute
func (r *AttributePostgresRepository) Delete(ctx context.Context, id string, tenantID string) error {
	query := `DELETE FROM attributes WHERE id = $1 AND tenant_id = $2`

	result, err := r.db.ExecContext(ctx, query, id, tenantID)
	if err != nil {
		log.Printf("Error eliminando attribute: %v", err)
		return fmt.Errorf("%w: %v", exception.ErrAttributeDeleteFailed, err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return exception.ErrAttributeNotFound
	}

	return nil
}

// SearchByCriteria busca attributes usando criterios
func (r *AttributePostgresRepository) SearchByCriteria(ctx context.Context, crit criteria.Criteria) ([]*entity.Attribute, error) {
	baseQuery := `
		SELECT id, tenant_id, name, active, created_at, updated_at
		FROM attributes
	`

	query, params := r.converter.ToSelectSQL(baseQuery, crit)

	rows, err := r.db.QueryContext(ctx, query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanAttributes(rows)
}

// CountByCriteria cuenta attributes usando criterios
func (r *AttributePostgresRepository) CountByCriteria(ctx context.Context, crit criteria.Criteria) (int, error) {
	baseCountQuery := "SELECT COUNT(*) FROM attributes"

	query, params := r.converter.ToCountSQL(baseCountQuery, crit)

	var count int
	err := r.db.QueryRowContext(ctx, query, params...).Scan(&count)
	return count, err
}

// scanAttribute escanea una fila y devuelve un attribute
func (r *AttributePostgresRepository) scanAttribute(row *sql.Row) (*entity.Attribute, error) {
	var attribute entity.Attribute

	err := row.Scan(
		&attribute.ID,
		&attribute.TenantID,
		&attribute.Name,
		&attribute.Active,
		&attribute.CreatedAt,
		&attribute.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &attribute, nil
}

// scanAttributes escanea múltiples filas y devuelve una lista de attributes
func (r *AttributePostgresRepository) scanAttributes(rows *sql.Rows) ([]*entity.Attribute, error) {
	var attributes []*entity.Attribute

	for rows.Next() {
		var attribute entity.Attribute

		err := rows.Scan(
			&attribute.ID,
			&attribute.TenantID,
			&attribute.Name,
			&attribute.Active,
			&attribute.CreatedAt,
			&attribute.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		attributes = append(attributes, &attribute)
	}

	return attributes, nil
}
