package persistence

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"saas-mt-pim-service/src/product/global_catalog/domain/entity"
	"saas-mt-pim-service/src/product/global_catalog/domain/port"
	"saas-mt-pim-service/src/product/global_catalog/domain/value_object"
	cr "github.com/mercadocercano/criteria"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

// PostgresGlobalProductRepository implementa GlobalProductRepository usando PostgreSQL
type PostgresGlobalProductRepository struct {
	db        *sql.DB
	converter *cr.SQLCriteriaConverter
}

// NewPostgresGlobalProductRepository crea una nueva instancia del repositorio
func NewPostgresGlobalProductRepository(db *sql.DB) port.GlobalProductRepository {
	return &PostgresGlobalProductRepository{
		db:        db,
		converter: cr.NewSQLCriteriaConverter(),
	}
}

// Save guarda un nuevo producto global
func (r *PostgresGlobalProductRepository) Save(globalProduct *entity.GlobalProduct) (*entity.GlobalProduct, error) {
	query := `
		INSERT INTO global_products (
			id, ean, name, description, brand, category, price, image_url, image_urls,
			source, source_url, source_reliability, quality_score, is_verified, is_active,
			business_type, tags, metadata, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20)
		RETURNING id, ean, name, description, brand, category, price, image_url, image_urls,
				  source, source_url, source_reliability, quality_score, is_verified, is_active,
				  business_type, tags, metadata, created_at, updated_at, scraped_at, last_scraped_at
	`

	// Preparar datos para inserción
	metadataJSON, err := json.Marshal(globalProduct.Metadata())
	if err != nil {
		return nil, fmt.Errorf("error al serializar metadata: %w", err)
	}

	var imageURLsArray pq.StringArray
	if imageURLs := globalProduct.ImageURLs(); len(imageURLs) > 0 {
		imageURLsArray = pq.StringArray(imageURLs)
	}

	var tagsArray pq.StringArray
	if tags := globalProduct.Tags(); len(tags) > 0 {
		tagsArray = pq.StringArray(tags)
	}

	var eanValue *string
	if globalProduct.EAN() != nil {
		ean := globalProduct.EAN().Value()
		eanValue = &ean
	}
	
	row := r.db.QueryRow(query,
		globalProduct.ID(),
		eanValue,
		globalProduct.Name(),
		globalProduct.Description(),
		globalProduct.Brand(),
		globalProduct.Category(),
		globalProduct.Price(),
		globalProduct.ImageURL(),
		imageURLsArray,
		globalProduct.Source().Source(),
		globalProduct.Source().SourceURL(),
		globalProduct.Source().Reliability(),
		globalProduct.QualityScore().Value(),
		globalProduct.IsVerified(),
		globalProduct.IsActive(),
		globalProduct.BusinessType(),
		tagsArray,
		string(metadataJSON),
		globalProduct.CreatedAt(),
		globalProduct.UpdatedAt(),
	)

	return r.scanGlobalProduct(row)
}

// Update actualiza un producto global existente
func (r *PostgresGlobalProductRepository) Update(globalProduct *entity.GlobalProduct) (*entity.GlobalProduct, error) {
	query := `
		UPDATE global_products SET 
			name = $2, description = $3, brand = $4, category = $5, price = $6,
			image_url = $7, image_urls = $8, source = $9, source_url = $10,
			source_reliability = $11, quality_score = $12, is_verified = $13,
			is_active = $14, business_type = $15, tags = $16, metadata = $17,
			updated_at = $18
		WHERE id = $1
		RETURNING id, ean, name, description, brand, category, price, image_url, image_urls,
				  source, source_url, source_reliability, quality_score, is_verified, is_active,
				  business_type, tags, metadata, created_at, updated_at, scraped_at, last_scraped_at
	`

	metadataJSON, err := json.Marshal(globalProduct.Metadata())
	if err != nil {
		return nil, fmt.Errorf("error al serializar metadata: %w", err)
	}

	var imageURLsArray pq.StringArray
	if imageURLs := globalProduct.ImageURLs(); len(imageURLs) > 0 {
		imageURLsArray = pq.StringArray(imageURLs)
	}

	var tagsArray pq.StringArray
	if tags := globalProduct.Tags(); len(tags) > 0 {
		tagsArray = pq.StringArray(tags)
	}

	row := r.db.QueryRow(query,
		globalProduct.ID(),
		globalProduct.Name(),
		globalProduct.Description(),
		globalProduct.Brand(),
		globalProduct.Category(),
		globalProduct.Price(),
		globalProduct.ImageURL(),
		imageURLsArray,
		globalProduct.Source().Source(),
		globalProduct.Source().SourceURL(),
		globalProduct.Source().Reliability(),
		globalProduct.QualityScore().Value(),
		globalProduct.IsVerified(),
		globalProduct.IsActive(),
		globalProduct.BusinessType(),
		tagsArray,
		string(metadataJSON),
		time.Now(),
	)

	return r.scanGlobalProduct(row)
}

// FindByID busca un producto por ID
func (r *PostgresGlobalProductRepository) FindByID(id string) (*entity.GlobalProduct, error) {
	query := `
		SELECT id, ean, name, description, brand, category, price, 
			   image_url, image_urls, source, source_url, source_reliability, 
			   quality_score, is_verified, is_active, business_type, tags, 
			   metadata, created_at, updated_at, scraped_at, last_scraped_at
		FROM global_products 
		WHERE id = $1
	`

	row := r.db.QueryRow(query, id)
	product, err := r.scanGlobalProduct(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No encontrado, pero no es error
		}
		return nil, err
	}

	return product, nil
}

// Delete elimina un producto (soft delete)
func (r *PostgresGlobalProductRepository) Delete(id string) error {
	query := `
		UPDATE global_products SET 
			is_active = false, updated_at = NOW()
		WHERE id = $1
	`

	_, err := r.db.Exec(query, id)
	return err
}

// FindByEAN busca un producto por EAN
func (r *PostgresGlobalProductRepository) FindByEAN(ean string) (*entity.GlobalProduct, error) {
	query := `
		SELECT id, ean, name, description, brand, category, price, 
			   image_url, image_urls, source, source_url, source_reliability, 
			   quality_score, is_verified, is_active, business_type, tags, 
			   metadata, created_at, updated_at, scraped_at, last_scraped_at
		FROM global_products 
		WHERE ean = $1
		ORDER BY quality_score DESC, created_at DESC
		LIMIT 1
	`

	row := r.db.QueryRow(query, ean)
	product, err := r.scanGlobalProduct(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No encontrado, pero no es error
		}
		return nil, err
	}

	return product, nil
}

// FindActiveByEAN busca un producto activo por EAN
func (r *PostgresGlobalProductRepository) FindActiveByEAN(ean string) (*entity.GlobalProduct, error) {
	query := `
		SELECT id, ean, name, description, brand, category, price, 
			   image_url, image_urls, source, source_url, source_reliability, 
			   quality_score, is_verified, is_active, business_type, tags, 
			   metadata, created_at, updated_at, scraped_at, last_scraped_at
		FROM global_products 
		WHERE ean = $1 AND is_active = true
		ORDER BY quality_score DESC, created_at DESC
		LIMIT 1
	`

	row := r.db.QueryRow(query, ean)
	product, err := r.scanGlobalProduct(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return product, nil
}

// FindByBusinessType busca productos por tipo de negocio
func (r *PostgresGlobalProductRepository) FindByBusinessType(businessType string, limit int) ([]*entity.GlobalProduct, error) {
	query := `
		SELECT id, ean, name, description, brand, category, price, 
			   image_url, image_urls, source, source_url, source_reliability, 
			   quality_score, is_verified, is_active, business_type, tags, 
			   metadata, created_at, updated_at, scraped_at, last_scraped_at
		FROM global_products 
		WHERE business_type = $1 AND is_active = true
		ORDER BY quality_score DESC, created_at DESC
		LIMIT $2
	`

	rows, err := r.db.Query(query, businessType, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanGlobalProducts(rows)
}

// FindBySource busca productos por fuente
func (r *PostgresGlobalProductRepository) FindBySource(source string, limit int) ([]*entity.GlobalProduct, error) {
	query := `
		SELECT id, ean, name, description, brand, category, price, 
			   image_url, image_urls, source, source_url, source_reliability, 
			   quality_score, is_verified, is_active, business_type, tags, 
			   metadata, created_at, updated_at, scraped_at, last_scraped_at
		FROM global_products 
		WHERE source = $1 AND is_active = true
		ORDER BY quality_score DESC, created_at DESC
		LIMIT $2
	`

	rows, err := r.db.Query(query, source, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanGlobalProducts(rows)
}

// FindByQualityScoreRange busca productos por rango de calidad
func (r *PostgresGlobalProductRepository) FindByQualityScoreRange(minScore, maxScore int, limit int) ([]*entity.GlobalProduct, error) {
	query := `
		SELECT id, ean, name, description, brand, category, price, 
			   image_url, image_urls, source, source_url, source_reliability, 
			   quality_score, is_verified, is_active, business_type, tags, 
			   metadata, created_at, updated_at, scraped_at, last_scraped_at
		FROM global_products 
		WHERE quality_score >= $1 AND quality_score <= $2 AND is_active = true
		ORDER BY quality_score DESC, created_at DESC
		LIMIT $3
	`

	rows, err := r.db.Query(query, minScore, maxScore, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanGlobalProducts(rows)
}

// SearchByName busca productos por nombre
func (r *PostgresGlobalProductRepository) SearchByName(name string, limit int) ([]*entity.GlobalProduct, error) {
	query := `
		SELECT id, ean, name, description, brand, category, price, 
			   image_url, image_urls, source, source_url, source_reliability, 
			   quality_score, is_verified, is_active, business_type, tags, 
			   metadata, created_at, updated_at, scraped_at, last_scraped_at
		FROM global_products 
		WHERE name ILIKE $1 AND is_active = true
		ORDER BY quality_score DESC, created_at DESC
		LIMIT $2
	`

	searchPattern := "%" + name + "%"
	rows, err := r.db.Query(query, searchPattern, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanGlobalProducts(rows)
}

// SearchByBrand busca productos por marca
func (r *PostgresGlobalProductRepository) SearchByBrand(brand string, limit int) ([]*entity.GlobalProduct, error) {
	query := `
		SELECT id, ean, name, description, brand, category, price, 
			   image_url, image_urls, source, source_url, source_reliability, 
			   quality_score, is_verified, is_active, business_type, tags, 
			   metadata, created_at, updated_at, scraped_at, last_scraped_at
		FROM global_products 
		WHERE brand ILIKE $1 AND is_active = true
		ORDER BY quality_score DESC, created_at DESC
		LIMIT $2
	`

	searchPattern := "%" + brand + "%"
	rows, err := r.db.Query(query, searchPattern, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanGlobalProducts(rows)
}

// SearchByCategory busca productos por categoría
func (r *PostgresGlobalProductRepository) SearchByCategory(category string, limit int) ([]*entity.GlobalProduct, error) {
	query := `
		SELECT id, ean, name, description, brand, category, price, 
			   image_url, image_urls, source, source_url, source_reliability, 
			   quality_score, is_verified, is_active, business_type, tags, 
			   metadata, created_at, updated_at, scraped_at, last_scraped_at
		FROM global_products 
		WHERE category ILIKE $1 AND is_active = true
		ORDER BY quality_score DESC, created_at DESC
		LIMIT $2
	`

	searchPattern := "%" + category + "%"
	rows, err := r.db.Query(query, searchPattern, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanGlobalProducts(rows)
}

// SearchByTags busca productos por tags
func (r *PostgresGlobalProductRepository) SearchByTags(tags []string, limit int) ([]*entity.GlobalProduct, error) {
	if len(tags) == 0 {
		return []*entity.GlobalProduct{}, nil
	}

	// Construir query para buscar productos que contengan cualquiera de los tags
	placeholders := make([]string, len(tags))
	args := make([]interface{}, len(tags)+1)

	for i, tag := range tags {
		placeholders[i] = fmt.Sprintf("$%d = ANY(tags)", i+1)
		args[i] = tag
	}
	args[len(tags)] = limit

	query := fmt.Sprintf(`
		SELECT id, ean, name, description, brand, category, price, 
			   image_url, image_urls, source, source_url, source_reliability, 
			   quality_score, is_verified, is_active, business_type, tags, 
			   metadata, created_at, updated_at, scraped_at, last_scraped_at
		FROM global_products 
		WHERE (%s) AND is_active = true
		ORDER BY quality_score DESC, created_at DESC
		LIMIT $%d
	`, strings.Join(placeholders, " OR "), len(tags)+1)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanGlobalProducts(rows)
}

// FindAll lista todos los productos con paginación
func (r *PostgresGlobalProductRepository) FindAll(offset, limit int) ([]*entity.GlobalProduct, error) {
	query := `
		SELECT id, ean, name, description, brand, category, price, 
			   image_url, image_urls, source, source_url, source_reliability, 
			   quality_score, is_verified, is_active, business_type, tags, 
			   metadata, created_at, updated_at, scraped_at, last_scraped_at
		FROM global_products 
		ORDER BY quality_score DESC, created_at DESC
		OFFSET $1 LIMIT $2
	`

	rows, err := r.db.Query(query, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanGlobalProducts(rows)
}

// FindActive lista productos activos con paginación
func (r *PostgresGlobalProductRepository) FindActive(offset, limit int) ([]*entity.GlobalProduct, error) {
	query := `
		SELECT id, ean, name, description, brand, category, price, 
			   image_url, image_urls, source, source_url, source_reliability, 
			   quality_score, is_verified, is_active, business_type, tags, 
			   metadata, created_at, updated_at, scraped_at, last_scraped_at
		FROM global_products 
		WHERE is_active = true
		ORDER BY quality_score DESC, created_at DESC
		OFFSET $1 LIMIT $2
	`

	rows, err := r.db.Query(query, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanGlobalProducts(rows)
}

// FindVerified lista productos verificados con paginación
func (r *PostgresGlobalProductRepository) FindVerified(offset, limit int) ([]*entity.GlobalProduct, error) {
	query := `
		SELECT id, ean, name, description, brand, category, price, 
			   image_url, image_urls, source, source_url, source_reliability, 
			   quality_score, is_verified, is_active, business_type, tags, 
			   metadata, created_at, updated_at, scraped_at, last_scraped_at
		FROM global_products 
		WHERE is_verified = true AND is_active = true
		ORDER BY quality_score DESC, created_at DESC
		OFFSET $1 LIMIT $2
	`

	rows, err := r.db.Query(query, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanGlobalProducts(rows)
}

// FindArgentineProducts lista productos argentinos
func (r *PostgresGlobalProductRepository) FindArgentineProducts(offset, limit int) ([]*entity.GlobalProduct, error) {
	query := `
		SELECT id, ean, name, description, brand, category, price, 
			   image_url, image_urls, source, source_url, source_reliability, 
			   quality_score, is_verified, is_active, business_type, tags, 
			   metadata, created_at, updated_at, scraped_at, last_scraped_at
		FROM global_products 
		WHERE source IN ('disco', 'carrefour', 'fravega', 'coto', 'jumbo') AND is_active = true
		ORDER BY quality_score DESC, created_at DESC
		OFFSET $1 LIMIT $2
	`

	rows, err := r.db.Query(query, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanGlobalProducts(rows)
}

// FindHighQualityProducts lista productos de alta calidad
func (r *PostgresGlobalProductRepository) FindHighQualityProducts(offset, limit int) ([]*entity.GlobalProduct, error) {
	query := `
		SELECT id, ean, name, description, brand, category, price, 
			   image_url, image_urls, source, source_url, source_reliability, 
			   quality_score, is_verified, is_active, business_type, tags, 
			   metadata, created_at, updated_at, scraped_at, last_scraped_at
		FROM global_products 
		WHERE quality_score >= 70 AND is_active = true
		ORDER BY quality_score DESC, created_at DESC
		OFFSET $1 LIMIT $2
	`

	rows, err := r.db.Query(query, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanGlobalProducts(rows)
}

// FindNeedingUpdate busca productos que necesitan actualización
func (r *PostgresGlobalProductRepository) FindNeedingUpdate(maxAgeHours int, limit int) ([]*entity.GlobalProduct, error) {
	query := `
		SELECT id, ean, name, description, brand, category, price, 
			   image_url, image_urls, source, source_url, source_reliability, 
			   quality_score, is_verified, is_active, business_type, tags, 
			   metadata, created_at, updated_at, scraped_at, last_scraped_at
		FROM global_products 
		WHERE updated_at < NOW() - INTERVAL '%d hours' AND is_active = true
		ORDER BY updated_at ASC
		LIMIT $1
	`

	formattedQuery := fmt.Sprintf(query, maxAgeHours)
	rows, err := r.db.Query(formattedQuery, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanGlobalProducts(rows)
}

// CountTotal cuenta el total de productos
func (r *PostgresGlobalProductRepository) CountTotal() (int, error) {
	query := `SELECT COUNT(*) FROM global_products`

	var count int
	err := r.db.QueryRow(query).Scan(&count)
	return count, err
}

// CountBySource cuenta productos por fuente
func (r *PostgresGlobalProductRepository) CountBySource(source string) (int, error) {
	query := `SELECT COUNT(*) FROM global_products WHERE source = $1 AND is_active = true`

	var count int
	err := r.db.QueryRow(query, source).Scan(&count)
	return count, err
}

// CountByQualityScore cuenta productos con calidad mínima
func (r *PostgresGlobalProductRepository) CountByQualityScore(minScore int) (int, error) {
	query := `SELECT COUNT(*) FROM global_products WHERE quality_score >= $1 AND is_active = true`

	var count int
	err := r.db.QueryRow(query, minScore).Scan(&count)
	return count, err
}

// CountArgentineProducts cuenta productos argentinos
func (r *PostgresGlobalProductRepository) CountArgentineProducts() (int, error) {
	query := `
		SELECT COUNT(*) 
		FROM global_products 
		WHERE source IN ('disco', 'carrefour', 'fravega', 'coto', 'jumbo') AND is_active = true
	`

	var count int
	err := r.db.QueryRow(query).Scan(&count)
	return count, err
}

// SearchByCriteria busca productos usando criterios
func (r *PostgresGlobalProductRepository) SearchByCriteria(ctx context.Context, crit cr.Criteria) ([]*entity.GlobalProduct, error) {
	baseQuery := `
		SELECT id, ean, name, description, brand, category, price, image_url, image_urls,
			   source, source_url, source_reliability, quality_score, is_verified, is_active,
			   business_type, tags, metadata, created_at, updated_at, scraped_at, last_scraped_at
		FROM global_products
	`
	
	query, params, err := r.converter.ToSelectSQL(baseQuery, crit)
	if err != nil {
		return nil, fmt.Errorf("invalid criteria: %w", err)
	}

	rows, err := r.db.QueryContext(ctx, query, params...)
	if err != nil {
		return nil, fmt.Errorf("error querying global products by criteria: %w", err)
	}
	defer rows.Close()
	
	return r.scanGlobalProducts(rows)
}

// CountByCriteria cuenta productos usando criterios
func (r *PostgresGlobalProductRepository) CountByCriteria(ctx context.Context, crit cr.Criteria) (int, error) {
	baseCountQuery := "SELECT COUNT(*) FROM global_products"
	
	query, params, err := r.converter.ToCountSQL(baseCountQuery, crit)
	if err != nil {
		return 0, fmt.Errorf("invalid criteria: %w", err)
	}

	var count int
	err = r.db.QueryRowContext(ctx, query, params...).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("error counting global products by criteria: %w", err)
	}
	
	return count, nil
}

// scanGlobalProduct escanea una fila en una entidad GlobalProduct
func (r *PostgresGlobalProductRepository) scanGlobalProduct(row *sql.Row) (*entity.GlobalProduct, error) {
	var (
		idStr, name, source                                  string
		ean                                                  *string
		description, brand, category, imageURL, businessType *string
		price, sourceReliability                             *float64
		sourceURL                                            *string
		qualityScore                                         int
		isVerified, isActive                                 bool
		imageURLsArray, tagsArray                            pq.StringArray
		metadataJSON                                         string
		createdAt, updatedAt                                 time.Time
		scrapedAt, lastScrapedAt                             *time.Time
	)

	err := row.Scan(
		&idStr, &ean, &name, &description, &brand, &category, &price,
		&imageURL, &imageURLsArray, &source, &sourceURL, &sourceReliability,
		&qualityScore, &isVerified, &isActive, &businessType, &tagsArray,
		&metadataJSON, &createdAt, &updatedAt, &scrapedAt, &lastScrapedAt,
	)

	if err != nil {
		return nil, err // Propagar el error para que se maneje en FindByID
	}

	return r.buildGlobalProductFromScan(
		idStr, ean, name, description, brand, category, price,
		sourceReliability, &source, sourceURL, qualityScore,
		isVerified, isActive, businessType, imageURLsArray, tagsArray,
		imageURL, metadataJSON, createdAt, updatedAt, lastScrapedAt,
	)
}

// scanGlobalProducts escanea múltiples filas en entidades GlobalProduct
func (r *PostgresGlobalProductRepository) scanGlobalProducts(rows *sql.Rows) ([]*entity.GlobalProduct, error) {
	var products []*entity.GlobalProduct

	for rows.Next() {
		var (
			idStr, name, source                                  string
		ean                                                  *string
			description, brand, category, imageURL, businessType *string
			price, sourceReliability                             *float64
			sourceURL                                            *string
			qualityScore                                         int
			isVerified, isActive                                 bool
			imageURLsArray, tagsArray                            pq.StringArray
			metadataJSON                                         string
			createdAt, updatedAt                                 time.Time
			scrapedAt, lastScrapedAt                             *time.Time
		)

		err := rows.Scan(
			&idStr, &ean, &name, &description, &brand, &category, &price,
			&imageURL, &imageURLsArray, &source, &sourceURL, &sourceReliability,
			&qualityScore, &isVerified, &isActive, &businessType, &tagsArray,
			&metadataJSON, &createdAt, &updatedAt, &scrapedAt, &lastScrapedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("error al escanear fila de producto global: %w", err)
		}

		product, err := r.buildGlobalProductFromScan(
			idStr, ean, name, description, brand, category, price,
			sourceReliability, &source, sourceURL, qualityScore,
			isVerified, isActive, businessType, imageURLsArray, tagsArray,
			imageURL, metadataJSON, createdAt, updatedAt, lastScrapedAt,
		)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error al iterar filas: %w", err)
	}

	return products, nil
}

// buildGlobalProductFromScan construye una entidad GlobalProduct desde datos escaneados
func (r *PostgresGlobalProductRepository) buildGlobalProductFromScan(
	idStr string, ean *string, name string,
	description, brand, category *string,
	price, sourceReliability *float64,
	source, sourceURL *string,
	qualityScore int,
	isVerified, isActive bool,
	businessType *string,
	imageURLsArray, tagsArray pq.StringArray,
	imageURL *string,
	metadataJSON string,
	createdAt, updatedAt time.Time,
	lastScrapedAt *time.Time,
) (*entity.GlobalProduct, error) {

	// Parsear UUID
	id, err := uuid.Parse(idStr)
	if err != nil {
		return nil, fmt.Errorf("ID inválido: %w", err)
	}

	// Crear EAN si existe
	var eanVO *value_object.EAN13
	if ean != nil && *ean != "" {
		var err error
		eanVO, err = value_object.NewEAN13(*ean)
		if err != nil {
			return nil, fmt.Errorf("EAN inválido: %w", err)
		}
	}

	// Crear QualityScore
	qualityScoreVO, err := value_object.NewQualityScore(qualityScore)
	if err != nil {
		return nil, fmt.Errorf("Quality score inválido: %w", err)
	}

	// Crear ProductSource
	var productSource *value_object.ProductSource
	if source != nil {
		reliability := 0.5
		if sourceReliability != nil {
			reliability = *sourceReliability
		}
		productSource, err = value_object.NewProductSource(*source, sourceURL, nil, reliability)
		if err != nil {
			return nil, fmt.Errorf("Product source inválido: %w", err)
		}
	}

	// Parsear metadata
	var metadata map[string]interface{}
	if metadataJSON != "" {
		if err := json.Unmarshal([]byte(metadataJSON), &metadata); err != nil {
			return nil, fmt.Errorf("error al parsear metadata: %w", err)
		}
	}

	// Convertir arrays a slices
	var imageURLs []string
	if len(imageURLsArray) > 0 {
		imageURLs = []string(imageURLsArray)
	}

	var tags []string
	if len(tagsArray) > 0 {
		tags = []string(tagsArray)
	}

	// Usar el constructor desde repositorio
	globalProduct, err := entity.NewGlobalProductFromRepository(
		id,
		eanVO,
		name,
		description,
		brand,
		category,
		price,
		imageURL,
		imageURLs,
		productSource,
		qualityScoreVO,
		isVerified,
		isActive,
		businessType,
		tags,
		metadata,
		createdAt,
		updatedAt,
		lastScrapedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("error al crear entidad GlobalProduct: %w", err)
	}

	return globalProduct, nil
}
