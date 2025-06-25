package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"pim/src/brand/domain/entity"
	"pim/src/brand/domain/exception"
	"pim/src/shared/domain/criteria"
	sharedCriteria "pim/src/shared/infrastructure/criteria"
	"strconv"

	"github.com/lib/pq"
)

// MarketplacebrandPostgresRepository implementa el repositorio usando PostgreSQL
type MarketplacebrandPostgresRepository struct {
	db        *sql.DB
	converter *sharedCriteria.SQLCriteriaConverter
}

// NewMarketplacebrandPostgresRepository crea una nueva instancia del repositorio
func NewMarketplacebrandPostgresRepository(db *sql.DB) *MarketplacebrandPostgresRepository {
	return &MarketplacebrandPostgresRepository{
		db:        db,
		converter: sharedCriteria.NewSQLCriteriaConverter(),
	}
}

// Create crea un nuevo marketplace_brand
func (r *MarketplacebrandPostgresRepository) Create(ctx context.Context, marketplace_brand *entity.Marketplacebrand) error {
	query := `
		INSERT INTO marketplace_brands (
			name, description, logo_url, website, aliases, category_tags, 
			quality_score, product_count, sources, is_verified, is_active,
			created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
		)
		RETURNING id
	`

	var id int
	err := r.db.QueryRowContext(ctx, query,
		marketplace_brand.Name,
		marketplace_brand.Description,
		marketplace_brand.LogoURL,
		marketplace_brand.Website,
		pq.Array(marketplace_brand.Aliases),
		pq.Array(marketplace_brand.CategoryTags),
		marketplace_brand.QualityScore,
		marketplace_brand.ProductCount,
		pq.Array(marketplace_brand.Sources),
		marketplace_brand.IsVerified,
		marketplace_brand.IsActive,
		marketplace_brand.CreatedAt,
		marketplace_brand.UpdatedAt,
	).Scan(&id)

	if err != nil {
		log.Printf("Error creando marketplace_brand: %v", err)
		return fmt.Errorf("%w: %v", exception.ErrMarketplacebrandCreateFailed, err)
	}

	// Actualizar el ID en la entidad
	marketplace_brand.ID = strconv.Itoa(id)
	return nil
}

// Update actualiza un marketplace_brand existente
func (r *MarketplacebrandPostgresRepository) Update(ctx context.Context, marketplace_brand *entity.Marketplacebrand) error {
	query := `
		UPDATE marketplace_brands SET
			name = $2,
			description = $3,
			logo_url = $4,
			website = $5,
			aliases = $6,
			category_tags = $7,
			quality_score = $8,
			product_count = $9,
			sources = $10,
			is_verified = $11,
			is_active = $12,
			updated_at = $13
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query,
		marketplace_brand.ID,
		marketplace_brand.Name,
		marketplace_brand.Description,
		marketplace_brand.LogoURL,
		marketplace_brand.Website,
		pq.Array(marketplace_brand.Aliases),
		pq.Array(marketplace_brand.CategoryTags),
		marketplace_brand.QualityScore,
		marketplace_brand.ProductCount,
		pq.Array(marketplace_brand.Sources),
		marketplace_brand.IsVerified,
		marketplace_brand.IsActive,
		marketplace_brand.UpdatedAt,
	)

	if err != nil {
		log.Printf("Error actualizando marketplace_brand: %v", err)
		return fmt.Errorf("%w: %v", exception.ErrMarketplacebrandUpdateFailed, err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return exception.ErrMarketplacebrandNotFound
	}

	return nil
}

// FindByID busca un marketplace_brand por su ID
func (r *MarketplacebrandPostgresRepository) FindByID(ctx context.Context, id string) (*entity.Marketplacebrand, error) {
	query := `
		SELECT id, name, description, logo_url, website, aliases, category_tags,
		       quality_score, product_count, sources, is_verified, is_active,
		       created_at, updated_at
		FROM marketplace_brands 
		WHERE id = $1
	`

	row := r.db.QueryRowContext(ctx, query, id)
	return r.scanMarketplacebrand(row)
}

// FindAll obtiene todos los marketplace_brands
func (r *MarketplacebrandPostgresRepository) FindAll(ctx context.Context) ([]*entity.Marketplacebrand, error) {
	query := `
		SELECT id, name, description, logo_url, website, aliases, category_tags,
		       quality_score, product_count, sources, is_verified, is_active,
		       created_at, updated_at
		FROM marketplace_brands 
		ORDER BY quality_score DESC, name ASC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanMarketplacebrands(rows)
}

// Delete elimina un marketplace_brand
func (r *MarketplacebrandPostgresRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM marketplace_brands WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		log.Printf("Error eliminando marketplace_brand: %v", err)
		return fmt.Errorf("%w: %v", exception.ErrMarketplacebrandDeleteFailed, err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return exception.ErrMarketplacebrandNotFound
	}

	return nil
}

// SearchByCriteria busca marketplace_brands usando criterios
func (r *MarketplacebrandPostgresRepository) SearchByCriteria(ctx context.Context, crit criteria.Criteria) ([]*entity.Marketplacebrand, error) {
	baseQuery := `
		SELECT id, name, description, logo_url, website, aliases, category_tags,
		       quality_score, product_count, sources, is_verified, is_active,
		       created_at, updated_at
		FROM marketplace_brands
	`

	query, params := r.converter.ToSelectSQL(baseQuery, crit)

	rows, err := r.db.QueryContext(ctx, query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanMarketplacebrands(rows)
}

// CountByCriteria cuenta marketplace_brands usando criterios
func (r *MarketplacebrandPostgresRepository) CountByCriteria(ctx context.Context, crit criteria.Criteria) (int, error) {
	baseCountQuery := "SELECT COUNT(*) FROM marketplace_brands"

	query, params := r.converter.ToCountSQL(baseCountQuery, crit)

	var count int
	err := r.db.QueryRowContext(ctx, query, params...).Scan(&count)
	return count, err
}

// scanMarketplacebrand escanea una fila y devuelve un marketplace_brand
func (r *MarketplacebrandPostgresRepository) scanMarketplacebrand(row *sql.Row) (*entity.Marketplacebrand, error) {
	var marketplace_brand entity.Marketplacebrand
	var aliases pq.StringArray
	var categoryTags pq.StringArray
	var sources pq.StringArray
	var id int
	var description, logoURL, website sql.NullString

	err := row.Scan(
		&id,
		&marketplace_brand.Name,
		&description,
		&logoURL,
		&website,
		&aliases,
		&categoryTags,
		&marketplace_brand.QualityScore,
		&marketplace_brand.ProductCount,
		&sources,
		&marketplace_brand.IsVerified,
		&marketplace_brand.IsActive,
		&marketplace_brand.CreatedAt,
		&marketplace_brand.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	// Convertir arrays de PostgreSQL a slices de Go
	marketplace_brand.ID = strconv.Itoa(id)
	marketplace_brand.Description = description.String
	marketplace_brand.LogoURL = logoURL.String
	marketplace_brand.Website = website.String
	marketplace_brand.Aliases = []string(aliases)
	marketplace_brand.CategoryTags = []string(categoryTags)
	marketplace_brand.Sources = []string(sources)

	// Actualizar verification status basado en is_verified
	if marketplace_brand.IsVerified {
		marketplace_brand.VerificationStatus = "verified"
	} else {
		marketplace_brand.VerificationStatus = "unverified"
	}

	return &marketplace_brand, nil
}

// scanMarketplacebrands escanea múltiples filas y devuelve una lista de marketplace_brands
func (r *MarketplacebrandPostgresRepository) scanMarketplacebrands(rows *sql.Rows) ([]*entity.Marketplacebrand, error) {
	var marketplace_brands []*entity.Marketplacebrand

	for rows.Next() {
		var marketplace_brand entity.Marketplacebrand
		var aliases pq.StringArray
		var categoryTags pq.StringArray
		var sources pq.StringArray
		var id int
		var description, logoURL, website sql.NullString

		err := rows.Scan(
			&id,
			&marketplace_brand.Name,
			&description,
			&logoURL,
			&website,
			&aliases,
			&categoryTags,
			&marketplace_brand.QualityScore,
			&marketplace_brand.ProductCount,
			&sources,
			&marketplace_brand.IsVerified,
			&marketplace_brand.IsActive,
			&marketplace_brand.CreatedAt,
			&marketplace_brand.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		// Convertir arrays de PostgreSQL a slices de Go
		marketplace_brand.ID = strconv.Itoa(id)
		marketplace_brand.Description = description.String
		marketplace_brand.LogoURL = logoURL.String
		marketplace_brand.Website = website.String
		marketplace_brand.Aliases = []string(aliases)
		marketplace_brand.CategoryTags = []string(categoryTags)
		marketplace_brand.Sources = []string(sources)

		// Actualizar verification status basado en is_verified
		if marketplace_brand.IsVerified {
			marketplace_brand.VerificationStatus = "verified"
		} else {
			marketplace_brand.VerificationStatus = "unverified"
		}

		marketplace_brands = append(marketplace_brands, &marketplace_brand)
	}

	return marketplace_brands, nil
}
