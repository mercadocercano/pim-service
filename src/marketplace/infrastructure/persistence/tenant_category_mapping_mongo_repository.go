package persistence

import (
	"context"
	"fmt"
	"time"

	"pim/src/marketplace/domain/entity"
	"pim/src/marketplace/domain/port"
	domainCriteria "pim/src/shared/domain/criteria"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TenantCategoryMappingMongoRepository implementa el repositorio usando MongoDB
type TenantCategoryMappingMongoRepository struct {
	collection *mongo.Collection
}

// NewTenantCategoryMappingMongoRepository crea una nueva instancia del repositorio MongoDB
func NewTenantCategoryMappingMongoRepository(db *mongo.Database) port.TenantCategoryMappingRepository {
	return &TenantCategoryMappingMongoRepository{
		collection: db.Collection("tenant_category_mappings"),
	}
}

// mongoMapping representa la estructura en MongoDB
type mongoMapping struct {
	ID                    primitive.ObjectID `bson:"_id,omitempty"`
	TenantID              string             `bson:"tenant_id"`
	CategoryID            string             `bson:"category_id"`
	MarketplaceCategoryID string             `bson:"marketplace_category_id"`
	CustomName            *string            `bson:"custom_name,omitempty"`
	CreatedAt             time.Time          `bson:"created_at"`
	UpdatedAt             time.Time          `bson:"updated_at"`
	DeletedAt             *time.Time         `bson:"deleted_at,omitempty"`
}

// Save guarda un mapeo de categoría tenant
func (r *TenantCategoryMappingMongoRepository) Save(ctx context.Context, mapping *entity.TenantCategoryMapping) error {
	// Convertir entidad a documento MongoDB
	doc := r.entityToMongo(mapping)

	// Si no tiene ID, es una inserción
	if mapping.ID == "" {
		result, err := r.collection.InsertOne(ctx, doc)
		if err != nil {
			return fmt.Errorf("failed to insert mapping: %w", err)
		}

		// Actualizar el ID en la entidad
		if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
			mapping.ID = oid.Hex()
		}
		return nil
	}

	// Si tiene ID, es una actualización (upsert)
	objectID, err := primitive.ObjectIDFromHex(mapping.ID)
	if err != nil {
		return fmt.Errorf("invalid object ID: %w", err)
	}

	filter := bson.M{"_id": objectID, "deleted_at": bson.M{"$exists": false}}
	update := bson.M{"$set": doc}

	opts := options.Update().SetUpsert(true)
	_, err = r.collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return fmt.Errorf("failed to upsert mapping: %w", err)
	}

	return nil
}

// GetByID obtiene un mapeo por su ID
func (r *TenantCategoryMappingMongoRepository) GetByID(ctx context.Context, id string) (*entity.TenantCategoryMapping, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid object ID: %w", err)
	}

	filter := bson.M{
		"_id":        objectID,
		"deleted_at": bson.M{"$exists": false},
	}

	var doc mongoMapping
	err = r.collection.FindOne(ctx, filter).Decode(&doc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find mapping: %w", err)
	}

	return r.mongoToEntity(&doc), nil
}

// GetByTenantAndMarketplaceCategory obtiene un mapeo específico
func (r *TenantCategoryMappingMongoRepository) GetByTenantAndMarketplaceCategory(ctx context.Context, tenantID, marketplaceCategoryID string) (*entity.TenantCategoryMapping, error) {
	filter := bson.M{
		"tenant_id":               tenantID,
		"marketplace_category_id": marketplaceCategoryID,
		"deleted_at":              bson.M{"$exists": false},
	}

	var doc mongoMapping
	err := r.collection.FindOne(ctx, filter).Decode(&doc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find mapping: %w", err)
	}

	return r.mongoToEntity(&doc), nil
}

// GetByTenantID obtiene todos los mapeos de un tenant
func (r *TenantCategoryMappingMongoRepository) GetByTenantID(ctx context.Context, tenantID string) ([]*entity.TenantCategoryMapping, error) {
	filter := bson.M{
		"tenant_id":  tenantID,
		"deleted_at": bson.M{"$exists": false},
	}

	opts := options.Find().SetSort(bson.D{bson.E{Key: "created_at", Value: -1}})
	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to find mappings: %w", err)
	}
	defer cursor.Close(ctx)

	return r.cursorToEntities(ctx, cursor)
}

// GetActiveByTenant obtiene mapeos activos de un tenant
func (r *TenantCategoryMappingMongoRepository) GetActiveByTenant(ctx context.Context, tenantID string) ([]*entity.TenantCategoryMapping, error) {
	filter := bson.M{
		"tenant_id":  tenantID,
		"is_active":  true,
		"deleted_at": bson.M{"$exists": false},
	}

	opts := options.Find().SetSort(bson.D{bson.E{Key: "created_at", Value: -1}})
	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to find active mappings: %w", err)
	}
	defer cursor.Close(ctx)

	return r.cursorToEntities(ctx, cursor)
}

// GetByMarketplaceCategoryID obtiene mapeos por categoría marketplace
func (r *TenantCategoryMappingMongoRepository) GetByMarketplaceCategoryID(ctx context.Context, marketplaceCategoryID string) ([]*entity.TenantCategoryMapping, error) {
	filter := bson.M{
		"marketplace_category_id": marketplaceCategoryID,
		"deleted_at":              bson.M{"$exists": false},
	}

	opts := options.Find().SetSort(bson.D{bson.E{Key: "created_at", Value: -1}})
	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to find mappings by marketplace category: %w", err)
	}
	defer cursor.Close(ctx)

	return r.cursorToEntities(ctx, cursor)
}

// GetTenantCategoriesForMarketplace obtiene categorías tenant mapeadas a una categoría marketplace
func (r *TenantCategoryMappingMongoRepository) GetTenantCategoriesForMarketplace(ctx context.Context, tenantID, marketplaceCategoryID string) ([]*entity.TenantCategoryMapping, error) {
	filter := bson.M{
		"tenant_id":               tenantID,
		"marketplace_category_id": marketplaceCategoryID,
		"is_active":               true,
		"deleted_at":              bson.M{"$exists": false},
	}

	opts := options.Find().SetSort(bson.D{bson.E{Key: "created_at", Value: -1}})
	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to find tenant categories for marketplace: %w", err)
	}
	defer cursor.Close(ctx)

	return r.cursorToEntities(ctx, cursor)
}

// FindByCriteria busca mapeos según criterios (implementación básica)
func (r *TenantCategoryMappingMongoRepository) FindByCriteria(ctx context.Context, crit domainCriteria.Criteria) ([]*entity.TenantCategoryMapping, error) {
	// Implementación básica - se puede mejorar con un convertidor de criterios a MongoDB
	filter := bson.M{"deleted_at": bson.M{"$exists": false}}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to find mappings by criteria: %w", err)
	}
	defer cursor.Close(ctx)

	return r.cursorToEntities(ctx, cursor)
}

// CountByCriteria cuenta mapeos según criterios
func (r *TenantCategoryMappingMongoRepository) CountByCriteria(ctx context.Context, crit domainCriteria.Criteria) (int, error) {
	filter := bson.M{"deleted_at": bson.M{"$exists": false}}

	count, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, fmt.Errorf("failed to count mappings: %w", err)
	}

	return int(count), nil
}

// Update actualiza un mapeo
func (r *TenantCategoryMappingMongoRepository) Update(ctx context.Context, mapping *entity.TenantCategoryMapping) error {
	objectID, err := primitive.ObjectIDFromHex(mapping.ID)
	if err != nil {
		return fmt.Errorf("invalid object ID: %w", err)
	}

	filter := bson.M{
		"_id":        objectID,
		"deleted_at": bson.M{"$exists": false},
	}

	doc := r.entityToMongo(mapping)
	update := bson.M{"$set": doc}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update mapping: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("mapping with id %s not found", mapping.ID)
	}

	return nil
}

// Delete elimina un mapeo (soft delete)
func (r *TenantCategoryMappingMongoRepository) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid object ID: %w", err)
	}

	filter := bson.M{
		"_id":        objectID,
		"deleted_at": bson.M{"$exists": false},
	}

	update := bson.M{"$set": bson.M{"deleted_at": time.Now()}}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to delete mapping: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("mapping with id %s not found", id)
	}

	return nil
}

// ExistsByTenantAndCategory verifica si ya existe un mapeo para la categoría en el tenant
func (r *TenantCategoryMappingMongoRepository) ExistsByTenantAndCategory(ctx context.Context, tenantID, categoryID string) (bool, error) {
	filter := bson.M{
		"tenant_id":   tenantID,
		"category_id": categoryID,
		"deleted_at":  bson.M{"$exists": false},
	}

	count, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return false, fmt.Errorf("failed to check mapping existence: %w", err)
	}

	return count > 0, nil
}

// GetTenantTaxonomy obtiene la taxonomía completa de un tenant (categorías + mapeos)
func (r *TenantCategoryMappingMongoRepository) GetTenantTaxonomy(ctx context.Context, tenantID string) ([]*entity.TenantCategoryMapping, error) {
	filter := bson.M{
		"tenant_id":  tenantID,
		"deleted_at": bson.M{"$exists": false},
	}

	opts := options.Find().SetSort(bson.D{bson.E{Key: "marketplace_category_id", Value: 1}, bson.E{Key: "created_at", Value: -1}})
	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to find tenant taxonomy: %w", err)
	}
	defer cursor.Close(ctx)

	return r.cursorToEntities(ctx, cursor)
}

// entityToMongo convierte una entidad a documento MongoDB
func (r *TenantCategoryMappingMongoRepository) entityToMongo(mapping *entity.TenantCategoryMapping) *mongoMapping {
	doc := &mongoMapping{
		TenantID:              mapping.TenantID,
		CategoryID:            mapping.CategoryID,
		MarketplaceCategoryID: mapping.MarketplaceCategoryID,
		CustomName:            mapping.CustomName,
		CreatedAt:             mapping.CreatedAt,
		UpdatedAt:             mapping.UpdatedAt,
	}

	// Convertir ID si existe
	if mapping.ID != "" {
		if objectID, err := primitive.ObjectIDFromHex(mapping.ID); err == nil {
			doc.ID = objectID
		}
	}

	return doc
}

// mongoToEntity convierte un documento MongoDB a entidad
func (r *TenantCategoryMappingMongoRepository) mongoToEntity(doc *mongoMapping) *entity.TenantCategoryMapping {
	mapping := &entity.TenantCategoryMapping{
		ID:                    doc.ID.Hex(),
		TenantID:              doc.TenantID,
		CategoryID:            doc.CategoryID,
		MarketplaceCategoryID: doc.MarketplaceCategoryID,
		CustomName:            doc.CustomName,
		CreatedAt:             doc.CreatedAt,
		UpdatedAt:             doc.UpdatedAt,
	}

	return mapping
}

// cursorToEntities convierte un cursor MongoDB a slice de entidades
func (r *TenantCategoryMappingMongoRepository) cursorToEntities(ctx context.Context, cursor *mongo.Cursor) ([]*entity.TenantCategoryMapping, error) {
	var mappings []*entity.TenantCategoryMapping

	for cursor.Next(ctx) {
		var doc mongoMapping
		if err := cursor.Decode(&doc); err != nil {
			return nil, fmt.Errorf("failed to decode mapping: %w", err)
		}

		mappings = append(mappings, r.mongoToEntity(&doc))
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return mappings, nil
}
