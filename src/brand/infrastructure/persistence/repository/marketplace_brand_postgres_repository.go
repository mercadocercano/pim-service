package repository

import (
	"context"
	"database/sql"
	"fmt"
	cr "github.com/hornosg/go-shared/criteria"
	"log"
	"saas-mt-pim-service/src/brand/domain/entity"
	"saas-mt-pim-service/src/brand/domain/exception"

	"github.com/lib/pq"
)

// MarketplacebrandPostgresRepository implementa el repositorio usando PostgreSQL
type MarketplacebrandPostgresRepository struct {
	db        *sql.DB
	converter *cr.SQLCriteriaConverter
}

// NewMarketplacebrandPostgresRepository crea una nueva instancia del repositorio
func NewMarketplacebrandPostgresRepository(db *sql.DB) *MarketplacebrandPostgresRepository {
	return &MarketplacebrandPostgresRepository{
		db:        db,
		converter: cr.NewSQLCriteriaConverter(),
	}
}

// Create crea un nuevo marketplace_brand
func (r *MarketplacebrandPostgresRepository) Create(ctx context.Context, marketplace_brand *entity.Marketplacebrand) error {
	query := `
		INSERT INTO marketplace_brands (
			name, description, logo_url, website, aliases, category_tags,
			quality_score, product_count, verification_status, is_active,
			background_color, text_color, typography,
			created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15
		)
		RETURNING id
	`

	// Convertir strings vacíos a NULL para campos opcionales (requerido por restricciones de BD)
	var logoURL interface{} = nil
	if marketplace_brand.LogoURL != "" {
		logoURL = marketplace_brand.LogoURL
	}

	var website interface{} = nil
	if marketplace_brand.Website != "" {
		website = marketplace_brand.Website
	}

	var backgroundColor interface{} = nil
	if marketplace_brand.BackgroundColor != "" {
		backgroundColor = marketplace_brand.BackgroundColor
	}

	var textColor interface{} = nil
	if marketplace_brand.TextColor != "" {
		textColor = marketplace_brand.TextColor
	}

	var typography interface{} = nil
	if marketplace_brand.Typography != "" {
		typography = marketplace_brand.Typography
	}

	var id string
	err := r.db.QueryRowContext(ctx, query,
		marketplace_brand.Name,
		marketplace_brand.Description,
		logoURL,
		website,
		pq.Array(marketplace_brand.Aliases),
		pq.Array(marketplace_brand.CategoryTags),
		marketplace_brand.QualityScore,
		marketplace_brand.ProductCount,
		marketplace_brand.VerificationStatus,
		marketplace_brand.IsActive,
		backgroundColor,
		textColor,
		typography,
		marketplace_brand.CreatedAt,
		marketplace_brand.UpdatedAt,
	).Scan(&id)

	if err != nil {
		log.Printf("Error creando marketplace_brand: %v", err)
		return fmt.Errorf("%w: %v", exception.ErrMarketplacebrandCreateFailed, err)
	}

	// Actualizar el ID en la entidad
	marketplace_brand.ID = id
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
			verification_status = $10,
			is_active = $11,
			background_color = $12,
			text_color = $13,
			typography = $14,
			updated_at = $15
		WHERE id = $1
	`

	// Convertir strings vacíos a NULL para campos opcionales (requerido por restricciones de BD)
	var logoURL interface{} = nil
	if marketplace_brand.LogoURL != "" {
		logoURL = marketplace_brand.LogoURL
	}

	var website interface{} = nil
	if marketplace_brand.Website != "" {
		website = marketplace_brand.Website
	}

	var backgroundColor interface{} = nil
	if marketplace_brand.BackgroundColor != "" {
		backgroundColor = marketplace_brand.BackgroundColor
	}

	var textColor interface{} = nil
	if marketplace_brand.TextColor != "" {
		textColor = marketplace_brand.TextColor
	}

	var typography interface{} = nil
	if marketplace_brand.Typography != "" {
		typography = marketplace_brand.Typography
	}

	result, err := r.db.ExecContext(ctx, query,
		marketplace_brand.ID,
		marketplace_brand.Name,
		marketplace_brand.Description,
		logoURL,
		website,
		pq.Array(marketplace_brand.Aliases),
		pq.Array(marketplace_brand.CategoryTags),
		marketplace_brand.QualityScore,
		marketplace_brand.ProductCount,
		marketplace_brand.VerificationStatus,
		marketplace_brand.IsActive,
		backgroundColor,
		textColor,
		typography,
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
		       quality_score, product_count, verification_status, is_active,
		       background_color, text_color, typography,
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
		       quality_score, product_count, verification_status, is_active,
		       background_color, text_color, typography,
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
func (r *MarketplacebrandPostgresRepository) SearchByCriteria(ctx context.Context, crit cr.Criteria) ([]*entity.Marketplacebrand, error) {
	baseQuery := `
		SELECT id, name, description, logo_url, website, aliases, category_tags,
		       quality_score, product_count, verification_status, is_active,
		       background_color, text_color, typography,
		       created_at, updated_at
		FROM marketplace_brands
	`

	query, params, err := r.converter.ToSelectSQL(baseQuery, crit)
	if err != nil {
		return nil, fmt.Errorf("invalid criteria: %w", err)
	}

	rows, err := r.db.QueryContext(ctx, query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanMarketplacebrands(rows)
}

// CountByCriteria cuenta marketplace_brands usando criterios
func (r *MarketplacebrandPostgresRepository) CountByCriteria(ctx context.Context, crit cr.Criteria) (int, error) {
	baseCountQuery := "SELECT COUNT(*) FROM marketplace_brands"

	query, params, err := r.converter.ToCountSQL(baseCountQuery, crit)
	if err != nil {
		return 0, fmt.Errorf("invalid criteria: %w", err)
	}

	var count int
	err = r.db.QueryRowContext(ctx, query, params...).Scan(&count)
	return count, err
}

// scanMarketplacebrand escanea una fila y devuelve un marketplace_brand
func (r *MarketplacebrandPostgresRepository) scanMarketplacebrand(row *sql.Row) (*entity.Marketplacebrand, error) {
	var marketplace_brand entity.Marketplacebrand
	var aliases pq.StringArray
	var categoryTags pq.StringArray
	var description, logoURL, website sql.NullString
	var backgroundColor, textColor, typography sql.NullString

	err := row.Scan(
		&marketplace_brand.ID,
		&marketplace_brand.Name,
		&description,
		&logoURL,
		&website,
		&aliases,
		&categoryTags,
		&marketplace_brand.QualityScore,
		&marketplace_brand.ProductCount,
		&marketplace_brand.VerificationStatus,
		&marketplace_brand.IsActive,
		&backgroundColor,
		&textColor,
		&typography,
		&marketplace_brand.CreatedAt,
		&marketplace_brand.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	marketplace_brand.Description = description.String
	marketplace_brand.LogoURL = logoURL.String
	marketplace_brand.Website = website.String
	marketplace_brand.Aliases = []string(aliases)
	marketplace_brand.CategoryTags = []string(categoryTags)
	marketplace_brand.Sources = []string{} // Campo no existe en DB, inicializar vacío
	marketplace_brand.IsVerified = marketplace_brand.VerificationStatus == "verified"
	marketplace_brand.BackgroundColor = backgroundColor.String
	marketplace_brand.TextColor = textColor.String
	marketplace_brand.Typography = typography.String

	return &marketplace_brand, nil
}

// scanMarketplacebrands escanea múltiples filas y devuelve una lista de marketplace_brands
func (r *MarketplacebrandPostgresRepository) scanMarketplacebrands(rows *sql.Rows) ([]*entity.Marketplacebrand, error) {
	var marketplace_brands []*entity.Marketplacebrand

	for rows.Next() {
		var marketplace_brand entity.Marketplacebrand
		var aliases pq.StringArray
		var categoryTags pq.StringArray
		var description, logoURL, website sql.NullString
		var backgroundColor, textColor, typography sql.NullString

		err := rows.Scan(
			&marketplace_brand.ID,
			&marketplace_brand.Name,
			&description,
			&logoURL,
			&website,
			&aliases,
			&categoryTags,
			&marketplace_brand.QualityScore,
			&marketplace_brand.ProductCount,
			&marketplace_brand.VerificationStatus,
			&marketplace_brand.IsActive,
			&backgroundColor,
			&textColor,
			&typography,
			&marketplace_brand.CreatedAt,
			&marketplace_brand.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		marketplace_brand.Description = description.String
		marketplace_brand.LogoURL = logoURL.String
		marketplace_brand.Website = website.String
		marketplace_brand.Aliases = []string(aliases)
		marketplace_brand.CategoryTags = []string(categoryTags)
		marketplace_brand.Sources = []string{} // Campo no existe en DB, inicializar vacío
		marketplace_brand.IsVerified = marketplace_brand.VerificationStatus == "verified"
		marketplace_brand.BackgroundColor = backgroundColor.String
		marketplace_brand.TextColor = textColor.String
		marketplace_brand.Typography = typography.String

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
