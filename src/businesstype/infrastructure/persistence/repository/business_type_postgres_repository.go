package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"saas-mt-pim-service/src/businesstype/domain/entity"
	"saas-mt-pim-service/src/businesstype/domain/port"
	cr "github.com/mercadocercano/criteria"
)

// BusinessTypePostgresRepository implementa el repositorio usando PostgreSQL
type BusinessTypePostgresRepository struct {
	db        *sql.DB
	converter *cr.SQLCriteriaConverter
}

// NewBusinessTypePostgresRepository crea una nueva instancia del repositorio
func NewBusinessTypePostgresRepository(db *sql.DB) port.BusinessTypeRepository {
	return &BusinessTypePostgresRepository{
		db:        db,
		converter: cr.NewSQLCriteriaConverter(),
	}
}

// Create crea un nuevo business type
func (r *BusinessTypePostgresRepository) Create(ctx context.Context, businessType *entity.BusinessType) error {
	metadataJSON, err := json.Marshal(businessType.Metadata)
	if err != nil {
		return fmt.Errorf("error serializando metadata: %w", err)
	}

	query := `
		INSERT INTO business_types (
			id, code, name, description, icon, color, 
			is_active, sort_order, metadata, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
		)
	`

	_, err = r.db.ExecContext(ctx, query,
		businessType.ID,
		businessType.Code,
		businessType.Name,
		businessType.Description,
		businessType.Icon,
		businessType.Color,
		businessType.IsActive,
		businessType.SortOrder,
		metadataJSON,
		businessType.CreatedAt,
		businessType.UpdatedAt,
	)

	if err != nil {
		log.Printf("Error creando business type: %v", err)
		return fmt.Errorf("error creando business type: %w", err)
	}

	return nil
}

// Update actualiza un business type existente
func (r *BusinessTypePostgresRepository) Update(ctx context.Context, businessType *entity.BusinessType) error {
	metadataJSON, err := json.Marshal(businessType.Metadata)
	if err != nil {
		return fmt.Errorf("error serializando metadata: %w", err)
	}

	query := `
		UPDATE business_types SET
			name = $2,
			description = $3,
			icon = $4,
			color = $5,
			is_active = $6,
			sort_order = $7,
			metadata = $8,
			updated_at = $9
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query,
		businessType.ID,
		businessType.Name,
		businessType.Description,
		businessType.Icon,
		businessType.Color,
		businessType.IsActive,
		businessType.SortOrder,
		metadataJSON,
		businessType.UpdatedAt,
	)

	if err != nil {
		log.Printf("Error actualizando business type: %v", err)
		return fmt.Errorf("error actualizando business type: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("business type no encontrado")
	}

	return nil
}

// FindByID busca un business type por su ID
func (r *BusinessTypePostgresRepository) FindByID(ctx context.Context, id string) (*entity.BusinessType, error) {
	query := `
		SELECT id, code, name, description, icon, color, 
			   is_active, sort_order, metadata, created_at, updated_at
		FROM business_types 
		WHERE id = $1
	`

	row := r.db.QueryRowContext(ctx, query, id)
	return r.scanBusinessType(row)
}

// FindByCode busca un business type por su código
func (r *BusinessTypePostgresRepository) FindByCode(ctx context.Context, code string) (*entity.BusinessType, error) {
	query := `
		SELECT id, code, name, description, icon, color, 
			   is_active, sort_order, metadata, created_at, updated_at
		FROM business_types 
		WHERE code = $1
	`

	row := r.db.QueryRowContext(ctx, query, code)
	return r.scanBusinessType(row)
}

// FindAll obtiene todos los business types
func (r *BusinessTypePostgresRepository) FindAll(ctx context.Context) ([]*entity.BusinessType, error) {
	query := `
		SELECT id, code, name, description, icon, color, 
			   is_active, sort_order, metadata, created_at, updated_at
		FROM business_types 
		ORDER BY sort_order ASC, name ASC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanBusinessTypes(rows)
}

// FindActive obtiene solo los business types activos
func (r *BusinessTypePostgresRepository) FindActive(ctx context.Context) ([]*entity.BusinessType, error) {
	query := `
		SELECT id, code, name, description, icon, color, 
			   is_active, sort_order, metadata, created_at, updated_at
		FROM business_types 
		WHERE is_active = true
		ORDER BY sort_order ASC, name ASC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanBusinessTypes(rows)
}

// Delete elimina un business type
func (r *BusinessTypePostgresRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM business_types WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		log.Printf("Error eliminando business type: %v", err)
		return fmt.Errorf("error eliminando business type: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("business type no encontrado")
	}

	return nil
}

// SearchByCriteria busca business types usando criterios
func (r *BusinessTypePostgresRepository) SearchByCriteria(ctx context.Context, crit cr.Criteria) ([]*entity.BusinessType, error) {
	baseQuery := `
		SELECT id, code, name, description, icon, color, 
			   is_active, sort_order, metadata, created_at, updated_at
		FROM business_types
	`

	query, params := r.converter.ToSelectSQL(baseQuery, crit)

	rows, err := r.db.QueryContext(ctx, query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanBusinessTypes(rows)
}

// CountByCriteria cuenta business types usando criterios
func (r *BusinessTypePostgresRepository) CountByCriteria(ctx context.Context, crit cr.Criteria) (int, error) {
	baseCountQuery := "SELECT COUNT(*) FROM business_types"

	query, params := r.converter.ToCountSQL(baseCountQuery, crit)

	var count int
	err := r.db.QueryRowContext(ctx, query, params...).Scan(&count)
	return count, err
}

// scanBusinessType escanea una fila y devuelve un business type
func (r *BusinessTypePostgresRepository) scanBusinessType(row *sql.Row) (*entity.BusinessType, error) {
	var businessType entity.BusinessType
	var metadataJSON []byte

	err := row.Scan(
		&businessType.ID,
		&businessType.Code,
		&businessType.Name,
		&businessType.Description,
		&businessType.Icon,
		&businessType.Color,
		&businessType.IsActive,
		&businessType.SortOrder,
		&metadataJSON,
		&businessType.CreatedAt,
		&businessType.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	// Deserializar metadata
	if len(metadataJSON) > 0 {
		if err := json.Unmarshal(metadataJSON, &businessType.Metadata); err != nil {
			log.Printf("Error deserializando metadata: %v", err)
			businessType.Metadata = make(map[string]interface{})
		}
	} else {
		businessType.Metadata = make(map[string]interface{})
	}

	return &businessType, nil
}

// scanBusinessTypes escanea múltiples filas y devuelve una lista de business types
func (r *BusinessTypePostgresRepository) scanBusinessTypes(rows *sql.Rows) ([]*entity.BusinessType, error) {
	var businessTypes []*entity.BusinessType

	for rows.Next() {
		var businessType entity.BusinessType
		var metadataJSON []byte

		err := rows.Scan(
			&businessType.ID,
			&businessType.Code,
			&businessType.Name,
			&businessType.Description,
			&businessType.Icon,
			&businessType.Color,
			&businessType.IsActive,
			&businessType.SortOrder,
			&metadataJSON,
			&businessType.CreatedAt,
			&businessType.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		// Deserializar metadata
		if len(metadataJSON) > 0 {
			if err := json.Unmarshal(metadataJSON, &businessType.Metadata); err != nil {
				log.Printf("Error deserializando metadata: %v", err)
				businessType.Metadata = make(map[string]interface{})
			}
		} else {
			businessType.Metadata = make(map[string]interface{})
		}

		businessTypes = append(businessTypes, &businessType)
	}

	return businessTypes, nil
}
