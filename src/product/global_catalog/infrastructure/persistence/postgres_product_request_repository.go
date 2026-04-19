package persistence

import (
	"context"
	"database/sql"
	"fmt"

	"saas-mt-pim-service/src/product/global_catalog/domain/entity"

	"github.com/google/uuid"
)

type PostgresProductRequestRepository struct {
	db *sql.DB
}

func NewPostgresProductRequestRepository(db *sql.DB) *PostgresProductRequestRepository {
	return &PostgresProductRequestRepository{db: db}
}

func (r *PostgresProductRequestRepository) Save(ctx context.Context, req *entity.ProductRequest) error {
	query := `
		INSERT INTO global_product_requests
			(id, tenant_id, name, brand, category, description, business_type, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`
	_, err := r.db.ExecContext(ctx, query,
		req.ID(), req.TenantID(), req.Name(), req.Brand(), req.Category(),
		req.Description(), req.BusinessType(), req.Status(),
		req.CreatedAt(), req.UpdatedAt(),
	)
	if err != nil {
		return fmt.Errorf("error saving product request: %w", err)
	}
	return nil
}

func (r *PostgresProductRequestRepository) FindByID(ctx context.Context, id string) (*entity.ProductRequest, error) {
	query := `
		SELECT id, tenant_id, name, brand, category, description, business_type,
		       status, admin_notes, global_product_id, created_at, updated_at
		FROM global_product_requests WHERE id = $1
	`
	return r.scanOne(ctx, query, id)
}

func (r *PostgresProductRequestRepository) Update(ctx context.Context, req *entity.ProductRequest) error {
	query := `
		UPDATE global_product_requests
		SET status = $1, admin_notes = $2, global_product_id = $3, updated_at = $4
		WHERE id = $5
	`
	_, err := r.db.ExecContext(ctx, query,
		req.Status(), req.AdminNotes(), req.GlobalProductID(), req.UpdatedAt(), req.ID(),
	)
	if err != nil {
		return fmt.Errorf("error updating product request: %w", err)
	}
	return nil
}

func (r *PostgresProductRequestRepository) FindPending(ctx context.Context, limit, offset int) ([]*entity.ProductRequest, error) {
	query := `
		SELECT id, tenant_id, name, brand, category, description, business_type,
		       status, admin_notes, global_product_id, created_at, updated_at
		FROM global_product_requests
		WHERE status = 'pending'
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`
	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("error querying pending requests: %w", err)
	}
	defer rows.Close()

	var results []*entity.ProductRequest
	for rows.Next() {
		req, err := r.scanRow(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, req)
	}
	return results, nil
}

func (r *PostgresProductRequestRepository) CountPending(ctx context.Context) (int, error) {
	var count int
	err := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM global_product_requests WHERE status = 'pending'").Scan(&count)
	return count, err
}

func (r *PostgresProductRequestRepository) scanOne(ctx context.Context, query string, args ...interface{}) (*entity.ProductRequest, error) {
	row := r.db.QueryRowContext(ctx, query, args...)

	var (
		id              uuid.UUID
		tenantID, name  string
		brand, category sql.NullString
		description     sql.NullString
		businessType    sql.NullString
		status          string
		adminNotes      sql.NullString
		globalProductID sql.NullString
		createdAt       sql.NullTime
		updatedAt       sql.NullTime
	)

	err := row.Scan(&id, &tenantID, &name, &brand, &category, &description,
		&businessType, &status, &adminNotes, &globalProductID, &createdAt, &updatedAt)
	if err != nil {
		return nil, fmt.Errorf("error scanning product request: %w", err)
	}

	return r.buildEntity(id, tenantID, name, brand, category, description,
		businessType, status, adminNotes, globalProductID, createdAt, updatedAt), nil
}

func (r *PostgresProductRequestRepository) scanRow(rows *sql.Rows) (*entity.ProductRequest, error) {
	var (
		id              uuid.UUID
		tenantID, name  string
		brand, category sql.NullString
		description     sql.NullString
		businessType    sql.NullString
		status          string
		adminNotes      sql.NullString
		globalProductID sql.NullString
		createdAt       sql.NullTime
		updatedAt       sql.NullTime
	)

	err := rows.Scan(&id, &tenantID, &name, &brand, &category, &description,
		&businessType, &status, &adminNotes, &globalProductID, &createdAt, &updatedAt)
	if err != nil {
		return nil, fmt.Errorf("error scanning product request row: %w", err)
	}

	return r.buildEntity(id, tenantID, name, brand, category, description,
		businessType, status, adminNotes, globalProductID, createdAt, updatedAt), nil
}

func (r *PostgresProductRequestRepository) buildEntity(
	id uuid.UUID, tenantID, name string,
	brand, category, description, businessType sql.NullString,
	status string, adminNotes, globalProductID sql.NullString,
	createdAt, updatedAt sql.NullTime,
) *entity.ProductRequest {
	var brandPtr, categoryPtr, descPtr, btPtr, notesPtr *string
	var gpIDPtr *uuid.UUID

	if brand.Valid {
		brandPtr = &brand.String
	}
	if category.Valid {
		categoryPtr = &category.String
	}
	if description.Valid {
		descPtr = &description.String
	}
	if businessType.Valid {
		btPtr = &businessType.String
	}
	if adminNotes.Valid {
		notesPtr = &adminNotes.String
	}
	if globalProductID.Valid {
		parsed, err := uuid.Parse(globalProductID.String)
		if err == nil {
			gpIDPtr = &parsed
		}
	}

	return entity.NewProductRequestFromRepository(
		id, tenantID, name, brandPtr, categoryPtr, descPtr, btPtr,
		entity.RequestStatus(status), notesPtr, gpIDPtr,
		createdAt.Time, updatedAt.Time,
	)
}
