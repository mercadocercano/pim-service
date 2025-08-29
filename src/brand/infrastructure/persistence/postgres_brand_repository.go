package persistence

import (
	"context"
	"database/sql"
	"fmt"

	"saas-mt-pim-service/src/brand/domain/entity"
	"saas-mt-pim-service/src/brand/domain/exception"
	"saas-mt-pim-service/src/brand/domain/value_object"
	"saas-mt-pim-service/src/shared/domain/criteria"
	sharedCriteria "saas-mt-pim-service/src/shared/infrastructure/criteria"
)

// PostgresBrandRepository implementa BrandCriteriaRepository usando PostgreSQL
type PostgresBrandRepository struct {
	db *sql.DB
}

// NewPostgresBrandRepository crea una nueva instancia del repositorio
func NewPostgresBrandRepository(db *sql.DB) *PostgresBrandRepository {
	return &PostgresBrandRepository{
		db: db,
	}
}

// Create guarda una nueva marca en la base de datos
func (r *PostgresBrandRepository) Create(ctx context.Context, brand *entity.Brand) error {
	query := `
		INSERT INTO brands (id, tenant_id, name, description, logo_url, website, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	_, err := r.db.ExecContext(ctx, query,
		brand.ID,
		brand.TenantID,
		brand.Name,
		brand.Description,
		brand.LogoURL,
		brand.Website,
		brand.Status.String(),
		brand.CreatedAt,
		brand.UpdatedAt,
	)

	if err != nil {
		// Verificar si es un error de duplicado
		if isUniqueViolation(err) {
			return exception.ErrBrandAlreadyExists
		}
		return fmt.Errorf("error al crear marca: %w", err)
	}

	return nil
}

// FindByID busca una marca por su ID y tenantID
func (r *PostgresBrandRepository) FindByID(ctx context.Context, id string, tenantID string) (*entity.Brand, error) {
	query := `
		SELECT id, tenant_id, name, description, logo_url, website, status, created_at, updated_at
		FROM brands
		WHERE id = $1 AND tenant_id = $2 AND status != 'deleted'
	`

	row := r.db.QueryRowContext(ctx, query, id, tenantID)
	return r.scanBrand(row)
}

// FindByName busca una marca por su nombre y tenantID
func (r *PostgresBrandRepository) FindByName(ctx context.Context, name string, tenantID string) (*entity.Brand, error) {
	query := `
		SELECT id, tenant_id, name, description, logo_url, website, status, created_at, updated_at
		FROM brands
		WHERE name = $1 AND tenant_id = $2 AND status != 'deleted'
	`

	row := r.db.QueryRowContext(ctx, query, name, tenantID)
	return r.scanBrand(row)
}

// FindAll recupera todas las marcas de un tenant
func (r *PostgresBrandRepository) FindAll(ctx context.Context, tenantID string) ([]*entity.Brand, error) {
	query := `
		SELECT id, tenant_id, name, description, logo_url, website, status, created_at, updated_at
		FROM brands
		WHERE tenant_id = $1 AND status != 'deleted'
		ORDER BY name ASC
	`

	rows, err := r.db.QueryContext(ctx, query, tenantID)
	if err != nil {
		return nil, fmt.Errorf("error al buscar marcas: %w", err)
	}
	defer rows.Close()

	return r.scanBrands(rows)
}

// Update actualiza una marca existente
func (r *PostgresBrandRepository) Update(ctx context.Context, brand *entity.Brand) error {
	query := `
		UPDATE brands
		SET name = $3, description = $4, logo_url = $5, website = $6, status = $7, updated_at = $8
		WHERE id = $1 AND tenant_id = $2
	`

	result, err := r.db.ExecContext(ctx, query,
		brand.ID,
		brand.TenantID,
		brand.Name,
		brand.Description,
		brand.LogoURL,
		brand.Website,
		brand.Status.String(),
		brand.UpdatedAt,
	)

	if err != nil {
		if isUniqueViolation(err) {
			return exception.ErrBrandAlreadyExists
		}
		return fmt.Errorf("error al actualizar marca: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar filas afectadas: %w", err)
	}

	if rowsAffected == 0 {
		return exception.ErrBrandNotFound
	}

	return nil
}

// Delete elimina una marca por su ID y tenantID (hard delete)
func (r *PostgresBrandRepository) Delete(ctx context.Context, id string, tenantID string) error {
	query := `DELETE FROM brands WHERE id = $1 AND tenant_id = $2`

	result, err := r.db.ExecContext(ctx, query, id, tenantID)
	if err != nil {
		return fmt.Errorf("error al eliminar marca: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar filas afectadas: %w", err)
	}

	if rowsAffected == 0 {
		return exception.ErrBrandNotFound
	}

	return nil
}

// ExistsByName verifica si existe una marca con el nombre dado
func (r *PostgresBrandRepository) ExistsByName(ctx context.Context, name string, tenantID string, excludeID *string) (bool, error) {
	query := `
		SELECT COUNT(*)
		FROM brands
		WHERE name = $1 AND tenant_id = $2 AND status != 'deleted'
	`
	args := []interface{}{name, tenantID}

	if excludeID != nil {
		query += ` AND id != $3`
		args = append(args, *excludeID)
	}

	var count int
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("error al verificar existencia de marca: %w", err)
	}

	return count > 0, nil
}

// SearchByCriteria implementa la búsqueda con criteria pattern
func (r *PostgresBrandRepository) SearchByCriteria(ctx context.Context, crit criteria.Criteria) ([]*entity.Brand, error) {
	baseQuery := `
		SELECT id, tenant_id, name, description, logo_url, website, status, created_at, updated_at
		FROM brands
	`

	converter := sharedCriteria.NewSQLCriteriaConverter()
	query, params := converter.ToSelectSQL(baseQuery, crit)

	rows, err := r.db.QueryContext(ctx, query, params...)
	if err != nil {
		return nil, fmt.Errorf("error al buscar marcas con criterios: %w", err)
	}
	defer rows.Close()

	return r.scanBrands(rows)
}

// CountByCriteria cuenta las marcas que coinciden con los criterios
func (r *PostgresBrandRepository) CountByCriteria(ctx context.Context, crit criteria.Criteria) (int, error) {
	baseQuery := "SELECT COUNT(*) FROM brands"

	converter := sharedCriteria.NewSQLCriteriaConverter()
	query, params := converter.ToCountSQL(baseQuery, crit)

	var count int
	err := r.db.QueryRowContext(ctx, query, params...).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("error al contar marcas con criterios: %w", err)
	}

	return count, nil
}

// ListByCriteria combina búsqueda y conteo para generar respuesta de listado
func (r *PostgresBrandRepository) ListByCriteria(ctx context.Context, crit criteria.Criteria) (*criteria.ListResponse[entity.Brand], error) {
	// Obtener elementos
	items, err := r.SearchByCriteria(ctx, crit)
	if err != nil {
		return nil, err
	}

	// Obtener conteo total
	total, err := r.CountByCriteria(ctx, crit)
	if err != nil {
		return nil, err
	}

	// Crear respuesta
	return criteria.NewListResponse(items, total, crit), nil
}

// scanBrand escanea una fila de la base de datos a una entidad Brand
func (r *PostgresBrandRepository) scanBrand(row *sql.Row) (*entity.Brand, error) {
	var brand entity.Brand
	var statusStr string

	err := row.Scan(
		&brand.ID,
		&brand.TenantID,
		&brand.Name,
		&brand.Description,
		&brand.LogoURL,
		&brand.Website,
		&statusStr,
		&brand.CreatedAt,
		&brand.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error al escanear marca: %w", err)
	}

	status, err := value_object.NewBrandStatus(statusStr)
	if err != nil {
		return nil, fmt.Errorf("error al convertir estado de marca: %w", err)
	}

	brand.Status = status
	return &brand, nil
}

// scanBrands escanea múltiples filas de la base de datos a entidades Brand
func (r *PostgresBrandRepository) scanBrands(rows *sql.Rows) ([]*entity.Brand, error) {
	var brands []*entity.Brand

	for rows.Next() {
		var brand entity.Brand
		var statusStr string

		err := rows.Scan(
			&brand.ID,
			&brand.TenantID,
			&brand.Name,
			&brand.Description,
			&brand.LogoURL,
			&brand.Website,
			&statusStr,
			&brand.CreatedAt,
			&brand.UpdatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("error al escanear marca: %w", err)
		}

		status, err := value_object.NewBrandStatus(statusStr)
		if err != nil {
			return nil, fmt.Errorf("error al convertir estado de marca: %w", err)
		}

		brand.Status = status
		brands = append(brands, &brand)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error al iterar filas: %w", err)
	}

	return brands, nil
}

// isUniqueViolation verifica si el error es una violación de constraint único
func isUniqueViolation(err error) bool {
	// Implementación específica para PostgreSQL
	// Código de error 23505 es unique_violation en PostgreSQL
	return err != nil && (err.Error() == "pq: duplicate key value violates unique constraint \"brands_tenant_id_name_key\"" ||
		err.Error() == "ERROR: duplicate key value violates unique constraint \"brands_tenant_id_name_key\" (SQLSTATE 23505)")
}
