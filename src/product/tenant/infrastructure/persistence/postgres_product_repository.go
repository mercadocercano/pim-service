package persistence

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"pim/src/product/tenant/domain/entity"
	"pim/src/product/tenant/domain/port"
	"pim/src/product/tenant/domain/value_object"
	"pim/src/shared/domain/criteria"
	sharedCriteria "pim/src/shared/infrastructure/criteria"

	"github.com/google/uuid"
)

// PostgresProductRepository implementa ProductCriteriaRepository usando PostgreSQL
type PostgresProductRepository struct {
	db *sql.DB
}

// NewPostgresProductRepository crea una nueva instancia del repositorio
func NewPostgresProductRepository(db *sql.DB) port.ProductCriteriaRepository {
	return &PostgresProductRepository{
		db: db,
	}
}

// Save guarda un nuevo producto
func (r *PostgresProductRepository) Save(ctx context.Context, product *entity.Product) error {
	query := `
		INSERT INTO products (
			id, tenant_id, name, description, sku, 
			category_id, category_name, brand_id, brand_name,
			status, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`

	var sku *string
	if product.HasSKU() {
		skuValue := product.SKU().Value()
		sku = &skuValue
	}

	var categoryID, categoryName *string
	if product.HasCategory() {
		catID := product.CategoryReference().ID()
		catName := product.CategoryReference().Name()
		categoryID = &catID
		categoryName = &catName
	}

	var brandID, brandName *string
	if product.HasBrand() {
		brID := product.BrandReference().ID()
		brName := product.BrandReference().Name()
		brandID = &brID
		brandName = &brName
	}

	_, err := r.db.ExecContext(ctx, query,
		product.ID(),
		product.TenantID(),
		product.Name(),
		product.Description(),
		sku,
		categoryID,
		categoryName,
		brandID,
		brandName,
		product.Status().Value(),
		product.CreatedAt(),
		product.UpdatedAt(),
	)

	return err
}

// FindByID busca un producto por ID
func (r *PostgresProductRepository) FindByID(ctx context.Context, id uuid.UUID, tenantID string) (*entity.Product, error) {
	query := `
		SELECT id, tenant_id, name, description, sku,
			   category_id, category_name, brand_id, brand_name,
			   status, created_at, updated_at
		FROM products 
		WHERE id = $1 AND tenant_id = $2 AND status != 'deleted'
	`

	row := r.db.QueryRowContext(ctx, query, id, tenantID)
	return r.scanProduct(row)
}

// FindByIDWithVariants busca un producto por ID incluyendo sus variantes
func (r *PostgresProductRepository) FindByIDWithVariants(ctx context.Context, id uuid.UUID, tenantID string) (*entity.Product, error) {
	// Primero obtener el producto
	product, err := r.FindByID(ctx, id, tenantID)
	if err != nil {
		return nil, err
	}

	// Luego cargar sus variantes
	variants, err := r.LoadVariantsForProduct(ctx, id)
	if err != nil {
		return nil, err
	}

	// Cargar las variantes en el producto
	product.LoadVariants(variants)

	return product, nil
}

// FindBySKU busca un producto por SKU
func (r *PostgresProductRepository) FindBySKU(ctx context.Context, sku, tenantID string) (*entity.Product, error) {
	query := `
		SELECT id, tenant_id, name, description, sku,
			   category_id, category_name, brand_id, brand_name,
			   status, created_at, updated_at
		FROM products 
		WHERE sku = $1 AND tenant_id = $2 AND status != 'deleted'
	`

	row := r.db.QueryRowContext(ctx, query, sku, tenantID)
	return r.scanProduct(row)
}

// Update actualiza un producto existente
func (r *PostgresProductRepository) Update(ctx context.Context, product *entity.Product) error {
	query := `
		UPDATE products SET 
			name = $3, description = $4, sku = $5,
			category_id = $6, category_name = $7, 
			brand_id = $8, brand_name = $9,
			status = $10, updated_at = $11
		WHERE id = $1 AND tenant_id = $2
	`

	var sku *string
	if product.HasSKU() {
		skuValue := product.SKU().Value()
		sku = &skuValue
	}

	var categoryID, categoryName *string
	if product.HasCategory() {
		catID := product.CategoryReference().ID()
		catName := product.CategoryReference().Name()
		categoryID = &catID
		categoryName = &catName
	}

	var brandID, brandName *string
	if product.HasBrand() {
		brID := product.BrandReference().ID()
		brName := product.BrandReference().Name()
		brandID = &brID
		brandName = &brName
	}

	_, err := r.db.ExecContext(ctx, query,
		product.ID(),
		product.TenantID(),
		product.Name(),
		product.Description(),
		sku,
		categoryID,
		categoryName,
		brandID,
		brandName,
		product.Status().Value(),
		product.UpdatedAt(),
	)

	return err
}

// Delete elimina un producto (soft delete)
func (r *PostgresProductRepository) Delete(ctx context.Context, id uuid.UUID, tenantID string) error {
	query := `
		UPDATE products SET 
			status = 'deleted', updated_at = NOW()
		WHERE id = $1 AND tenant_id = $2
	`

	_, err := r.db.ExecContext(ctx, query, id, tenantID)
	return err
}

// ExistsByID verifica si existe un producto por ID
func (r *PostgresProductRepository) ExistsByID(ctx context.Context, id uuid.UUID, tenantID string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM products WHERE id = $1 AND tenant_id = $2 AND status != 'deleted')`

	var exists bool
	err := r.db.QueryRowContext(ctx, query, id, tenantID).Scan(&exists)
	return exists, err
}

// ExistsBySKU verifica si existe un producto por SKU
func (r *PostgresProductRepository) ExistsBySKU(ctx context.Context, sku, tenantID string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM products WHERE sku = $1 AND tenant_id = $2 AND status != 'deleted')`

	var exists bool
	err := r.db.QueryRowContext(ctx, query, sku, tenantID).Scan(&exists)
	return exists, err
}

// ExistsByName verifica si existe un producto por nombre
func (r *PostgresProductRepository) ExistsByName(ctx context.Context, name, tenantID string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM products WHERE name = $1 AND tenant_id = $2 AND status != 'deleted')`

	var exists bool
	err := r.db.QueryRowContext(ctx, query, name, tenantID).Scan(&exists)
	return exists, err
}

// ExistsByNameExcludingID verifica si existe un producto por nombre excluyendo un ID
func (r *PostgresProductRepository) ExistsByNameExcludingID(ctx context.Context, name, tenantID string, excludeID uuid.UUID) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM products WHERE name = $1 AND tenant_id = $2 AND id != $3 AND status != 'deleted')`

	var exists bool
	err := r.db.QueryRowContext(ctx, query, name, tenantID, excludeID).Scan(&exists)
	return exists, err
}

// ExistsBySKUExcludingID verifica si existe un producto por SKU excluyendo un ID
func (r *PostgresProductRepository) ExistsBySKUExcludingID(ctx context.Context, sku, tenantID string, excludeID uuid.UUID) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM products WHERE sku = $1 AND tenant_id = $2 AND id != $3 AND status != 'deleted')`

	var exists bool
	err := r.db.QueryRowContext(ctx, query, sku, tenantID, excludeID).Scan(&exists)
	return exists, err
}

// SearchByCriteria busca productos usando criterios
func (r *PostgresProductRepository) SearchByCriteria(ctx context.Context, crit criteria.Criteria) ([]*entity.Product, error) {
	baseQuery := `
		SELECT id, tenant_id, name, description, sku,
			   category_id, category_name, brand_id, brand_name,
			   status, created_at, updated_at
		FROM products
	`

	converter := sharedCriteria.NewSQLCriteriaConverter()
	query, params := converter.ToSelectSQL(baseQuery, crit)

	rows, err := r.db.QueryContext(ctx, query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanProducts(rows)
}

// CountByCriteria cuenta productos usando criterios
func (r *PostgresProductRepository) CountByCriteria(ctx context.Context, crit criteria.Criteria) (int, error) {
	baseQuery := "SELECT COUNT(*) FROM products"

	converter := sharedCriteria.NewSQLCriteriaConverter()
	query, params := converter.ToCountSQL(baseQuery, crit)

	var count int
	err := r.db.QueryRowContext(ctx, query, params...).Scan(&count)
	return count, err
}

// scanProduct escanea una fila a una entidad Product
func (r *PostgresProductRepository) scanProduct(row *sql.Row) (*entity.Product, error) {
	var id, tenantID, name, statusStr string
	var description, sku, categoryID, categoryName, brandID, brandName *string
	var createdAt, updatedAt sql.NullTime

	err := row.Scan(
		&id, &tenantID, &name, &description, &sku,
		&categoryID, &categoryName, &brandID, &brandName,
		&statusStr, &createdAt, &updatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("producto no encontrado")
		}
		return nil, err
	}

	return r.buildProductFromScan(
		id, tenantID, name, description, sku,
		categoryID, categoryName, brandID, brandName,
		statusStr, createdAt, updatedAt,
	)
}

// scanProducts escanea múltiples filas a entidades Product
func (r *PostgresProductRepository) scanProducts(rows *sql.Rows) ([]*entity.Product, error) {
	var products []*entity.Product

	for rows.Next() {
		var id, tenantID, name, statusStr string
		var description, sku, categoryID, categoryName, brandID, brandName *string
		var createdAt, updatedAt sql.NullTime

		err := rows.Scan(
			&id, &tenantID, &name, &description, &sku,
			&categoryID, &categoryName, &brandID, &brandName,
			&statusStr, &createdAt, &updatedAt,
		)

		if err != nil {
			return nil, err
		}

		product, err := r.buildProductFromScan(
			id, tenantID, name, description, sku,
			categoryID, categoryName, brandID, brandName,
			statusStr, createdAt, updatedAt,
		)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, rows.Err()
}

// buildProductFromScan construye una entidad Product desde los datos escaneados
func (r *PostgresProductRepository) buildProductFromScan(
	idStr, tenantID, name string,
	description, sku, categoryID, categoryName, brandID, brandName *string,
	statusStr string,
	createdAt, updatedAt sql.NullTime,
) (*entity.Product, error) {
	// Convertir string ID a UUID
	id, err := uuid.Parse(idStr)
	if err != nil {
		return nil, err
	}

	// Crear SKU value object si existe
	var productSKU *value_object.ProductSKU
	if sku != nil && *sku != "" {
		productSKU, err = value_object.NewProductSKU(*sku)
		if err != nil {
			return nil, err
		}
	}

	// Crear referencia de categoría si existe
	var categoryRef *value_object.CategoryReference
	if categoryID != nil && *categoryID != "" {
		categoryRef, err = value_object.NewCategoryReference(*categoryID, *categoryName)
		if err != nil {
			return nil, err
		}
	}

	// Crear referencia de marca si existe
	var brandRef *value_object.BrandReference
	if brandID != nil && *brandID != "" {
		brandRef, err = value_object.NewBrandReference(*brandID, *brandName)
		if err != nil {
			return nil, err
		}
	}

	// Crear status value object
	status, err := value_object.NewProductStatus(statusStr)
	if err != nil {
		return nil, err
	}

	// Manejar timestamps
	var createdAtTime, updatedAtTime time.Time
	if createdAt.Valid {
		createdAtTime = createdAt.Time
	}
	if updatedAt.Valid {
		updatedAtTime = updatedAt.Time
	}

	return entity.NewProductFromRepository(
		id,
		tenantID,
		name,
		description,
		productSKU,
		categoryRef,
		brandRef,
		status,
		createdAtTime,
		updatedAtTime,
	)
}

// SaveVariant guarda una nueva variante de producto
func (r *PostgresProductRepository) SaveVariant(ctx context.Context, productID uuid.UUID, variant *entity.ProductVariant) error {
	// Primero insertar la variante
	variantQuery := `
		INSERT INTO product_variants (
			id, tenant_id, product_id, name, sku, status, 
			is_default, sort_order, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	var sku *string
	if variant.HasSKU() {
		skuValue := variant.SKU().Value()
		sku = &skuValue
	}

	_, err := r.db.ExecContext(ctx, variantQuery,
		variant.ID(),
		variant.TenantID(),
		variant.ProductID(),
		variant.Name(),
		sku,
		variant.Status().Value(),
		variant.IsDefault(),
		variant.SortOrder(),
		variant.CreatedAt(),
		variant.UpdatedAt(),
	)
	if err != nil {
		return err
	}

	// Luego insertar los atributos
	return r.saveVariantAttributes(ctx, variant)
}

// UpdateVariant actualiza una variante existente
func (r *PostgresProductRepository) UpdateVariant(ctx context.Context, variant *entity.ProductVariant) error {
	// Actualizar la variante
	variantQuery := `
		UPDATE product_variants SET 
			name = $3, sku = $4, status = $5, 
			is_default = $6, sort_order = $7, updated_at = $8
		WHERE id = $1 AND tenant_id = $2
	`

	var sku *string
	if variant.HasSKU() {
		skuValue := variant.SKU().Value()
		sku = &skuValue
	}

	_, err := r.db.ExecContext(ctx, variantQuery,
		variant.ID(),
		variant.TenantID(),
		variant.Name(),
		sku,
		variant.Status().Value(),
		variant.IsDefault(),
		variant.SortOrder(),
		variant.UpdatedAt(),
	)
	if err != nil {
		return err
	}

	// Actualizar los atributos (eliminar y recrear)
	deleteAttrsQuery := `DELETE FROM variant_attributes WHERE variant_id = $1`
	_, err = r.db.ExecContext(ctx, deleteAttrsQuery, variant.ID())
	if err != nil {
		return err
	}

	return r.saveVariantAttributes(ctx, variant)
}

// DeleteVariant elimina una variante (soft delete)
func (r *PostgresProductRepository) DeleteVariant(ctx context.Context, variantID uuid.UUID) error {
	query := `
		UPDATE product_variants SET 
			status = 'deleted', updated_at = NOW()
		WHERE id = $1
	`

	_, err := r.db.ExecContext(ctx, query, variantID)
	return err
}

// LoadVariantsForProduct carga todas las variantes de un producto
func (r *PostgresProductRepository) LoadVariantsForProduct(ctx context.Context, productID uuid.UUID) ([]*entity.ProductVariant, error) {
	query := `
		SELECT id, tenant_id, product_id, name, sku, status,
			   is_default, sort_order, created_at, updated_at
		FROM product_variants 
		WHERE product_id = $1 AND status != 'deleted'
		ORDER BY sort_order, created_at
	`

	rows, err := r.db.QueryContext(ctx, query, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var variants []*entity.ProductVariant
	for rows.Next() {
		variant, err := r.scanVariant(rows)
		if err != nil {
			return nil, err
		}

		// Cargar atributos para cada variante
		attributes, err := r.loadVariantAttributes(ctx, variant.ID())
		if err != nil {
			return nil, err
		}

		// Crear nueva variante con atributos
		variantWithAttrs := entity.NewProductVariantFromRepository(
			variant.ID(),
			variant.TenantID(),
			variant.ProductID(),
			variant.Name(),
			variant.SKU(),
			variant.Status(),
			variant.IsDefault(),
			variant.SortOrder(),
			attributes,
			variant.CreatedAt(),
			variant.UpdatedAt(),
		)

		variants = append(variants, variantWithAttrs)
	}

	return variants, rows.Err()
}

// Métodos de verificación de variantes
func (r *PostgresProductRepository) VariantExistsByName(ctx context.Context, name string, productID uuid.UUID, tenantID string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM product_variants WHERE name = $1 AND product_id = $2 AND tenant_id = $3 AND status != 'deleted')`

	var exists bool
	err := r.db.QueryRowContext(ctx, query, name, productID, tenantID).Scan(&exists)
	return exists, err
}

func (r *PostgresProductRepository) VariantExistsBySKU(ctx context.Context, sku, tenantID string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM product_variants WHERE sku = $1 AND tenant_id = $2 AND status != 'deleted')`

	var exists bool
	err := r.db.QueryRowContext(ctx, query, sku, tenantID).Scan(&exists)
	return exists, err
}

func (r *PostgresProductRepository) VariantExistsByNameExcludingID(ctx context.Context, name string, productID uuid.UUID, tenantID string, excludeID uuid.UUID) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM product_variants WHERE name = $1 AND product_id = $2 AND tenant_id = $3 AND id != $4 AND status != 'deleted')`

	var exists bool
	err := r.db.QueryRowContext(ctx, query, name, productID, tenantID, excludeID).Scan(&exists)
	return exists, err
}

func (r *PostgresProductRepository) VariantExistsBySKUExcludingID(ctx context.Context, sku, tenantID string, excludeID uuid.UUID) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM product_variants WHERE sku = $1 AND tenant_id = $2 AND id != $3 AND status != 'deleted')`

	var exists bool
	err := r.db.QueryRowContext(ctx, query, sku, tenantID, excludeID).Scan(&exists)
	return exists, err
}

// FindVariantsByCriteria busca variantes por criterios
func (r *PostgresProductRepository) FindVariantsByCriteria(ctx context.Context, criteria *criteria.Criteria) ([]*entity.ProductVariant, error) {
	// Por ahora implementamos una búsqueda básica sin el CriteriaBuilder
	// TODO: Implementar CriteriaBuilder para variantes
	query := `
		SELECT id, tenant_id, product_id, name, sku, status,
			   is_default, sort_order, created_at, updated_at
		FROM product_variants 
		WHERE status != 'deleted'
		ORDER BY sort_order, created_at
		LIMIT $1 OFFSET $2
	`

	limit := criteria.Pagination.PageSize
	offset := (criteria.Pagination.Page - 1) * criteria.Pagination.PageSize

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var variants []*entity.ProductVariant
	for rows.Next() {
		variant, err := r.scanVariant(rows)
		if err != nil {
			return nil, err
		}

		// Cargar atributos
		attributes, err := r.loadVariantAttributes(ctx, variant.ID())
		if err != nil {
			return nil, err
		}

		// Crear variante con atributos
		variantWithAttrs := entity.NewProductVariantFromRepository(
			variant.ID(),
			variant.TenantID(),
			variant.ProductID(),
			variant.Name(),
			variant.SKU(),
			variant.Status(),
			variant.IsDefault(),
			variant.SortOrder(),
			attributes,
			variant.CreatedAt(),
			variant.UpdatedAt(),
		)

		variants = append(variants, variantWithAttrs)
	}

	return variants, rows.Err()
}

// CountVariantsByCriteria cuenta variantes por criterios
func (r *PostgresProductRepository) CountVariantsByCriteria(ctx context.Context, criteria *criteria.Criteria) (int, error) {
	// Por ahora implementamos un conteo básico
	// TODO: Implementar CriteriaBuilder para variantes
	query := `SELECT COUNT(*) FROM product_variants WHERE status != 'deleted'`

	var count int
	err := r.db.QueryRowContext(ctx, query).Scan(&count)
	return count, err
}

// Métodos auxiliares para variantes

// saveVariantAttributes guarda los atributos de una variante
func (r *PostgresProductRepository) saveVariantAttributes(ctx context.Context, variant *entity.ProductVariant) error {
	if !variant.HasAttributes() {
		return nil
	}

	query := `
		INSERT INTO variant_attributes (tenant_id, variant_id, attribute_name, attribute_value)
		VALUES ($1, $2, $3, $4)
	`

	for _, attr := range variant.Attributes().Attributes() {
		_, err := r.db.ExecContext(ctx, query,
			variant.TenantID(),
			variant.ID(),
			attr.Name(),
			attr.Value(),
		)
		if err != nil {
			return err
		}
	}

	return nil
}

// loadVariantAttributes carga los atributos de una variante
func (r *PostgresProductRepository) loadVariantAttributes(ctx context.Context, variantID uuid.UUID) (*value_object.VariantAttributeCollection, error) {
	query := `
		SELECT attribute_name, attribute_value
		FROM variant_attributes
		WHERE variant_id = $1
		ORDER BY attribute_name
	`

	rows, err := r.db.QueryContext(ctx, query, variantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var attributes []*value_object.VariantAttribute
	for rows.Next() {
		var name, value string
		if err := rows.Scan(&name, &value); err != nil {
			return nil, err
		}

		attr, err := value_object.NewVariantAttribute(name, value)
		if err != nil {
			return nil, err
		}

		attributes = append(attributes, attr)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return value_object.NewVariantAttributeCollection(attributes)
}

// scanVariant escanea una fila de variante desde la base de datos
func (r *PostgresProductRepository) scanVariant(rows interface{}) (*entity.ProductVariant, error) {
	var id, tenantID, productID uuid.UUID
	var name, statusStr string
	var sku *string
	var isDefault bool
	var sortOrder int
	var createdAt, updatedAt time.Time

	var err error
	switch v := rows.(type) {
	case *sql.Rows:
		err = v.Scan(&id, &tenantID, &productID, &name, &sku, &statusStr, &isDefault, &sortOrder, &createdAt, &updatedAt)
	case *sql.Row:
		err = v.Scan(&id, &tenantID, &productID, &name, &sku, &statusStr, &isDefault, &sortOrder, &createdAt, &updatedAt)
	default:
		return nil, errors.New("tipo de scanner no soportado")
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	// Crear SKU value object si existe
	var productSKU *value_object.ProductSKU
	if sku != nil && *sku != "" {
		productSKU, err = value_object.NewProductSKU(*sku)
		if err != nil {
			return nil, err
		}
	}

	// Crear status value object
	status, err := value_object.NewVariantStatus(statusStr)
	if err != nil {
		return nil, err
	}

	// Crear colección de atributos vacía (se cargará por separado)
	attributes, _ := value_object.NewVariantAttributeCollection([]*value_object.VariantAttribute{})

	return entity.NewProductVariantFromRepository(
		id,
		tenantID.String(),
		productID,
		name,
		productSKU,
		status,
		isDefault,
		sortOrder,
		attributes,
		createdAt,
		updatedAt,
	), nil
}
