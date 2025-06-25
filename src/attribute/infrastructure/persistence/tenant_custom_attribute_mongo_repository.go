package persistence

import (
	"context"
	"fmt"
	"time"

	"pim/src/attribute/domain/entity"
	"pim/src/attribute/domain/port"
	domainCriteria "pim/src/shared/domain/criteria"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TenantCustomAttributeMongoRepository implementa el repositorio usando MongoDB
type TenantCustomAttributeMongoRepository struct {
	collection *mongo.Collection
}

// NewTenantCustomAttributeMongoRepository crea una nueva instancia del repositorio MongoDB
func NewTenantCustomAttributeMongoRepository(db *mongo.Database) port.TenantCustomAttributeRepository {
	return &TenantCustomAttributeMongoRepository{
		collection: db.Collection("tenant_custom_attributes"),
	}
}

// mongoAttribute representa la estructura en MongoDB
type mongoAttribute struct {
	ID                    primitive.ObjectID     `bson:"_id,omitempty"`
	TenantID              string                 `bson:"tenant_id"`
	MarketplaceCategoryID *string                `bson:"marketplace_category_id,omitempty"`
	Name                  string                 `bson:"name"`
	Slug                  string                 `bson:"slug"`
	Type                  string                 `bson:"type"`
	IsFilterable          bool                   `bson:"is_filterable"`
	IsSearchable          bool                   `bson:"is_searchable"`
	ValidationRules       map[string]interface{} `bson:"validation_rules"` // ✅ MongoDB maneja esto nativamente
	SortOrder             int                    `bson:"sort_order"`
	CreatedAt             time.Time              `bson:"created_at"`
	UpdatedAt             time.Time              `bson:"updated_at"`
	DeletedAt             *time.Time             `bson:"deleted_at,omitempty"`
}

// Save guarda un atributo personalizado tenant
func (r *TenantCustomAttributeMongoRepository) Save(ctx context.Context, attribute *entity.TenantCustomAttribute) error {
	// Convertir entidad a documento MongoDB
	doc := r.entityToMongo(attribute)

	// Si no tiene ID, es una inserción
	if attribute.ID == "" {
		result, err := r.collection.InsertOne(ctx, doc)
		if err != nil {
			return fmt.Errorf("failed to insert attribute: %w", err)
		}

		// Actualizar el ID en la entidad
		if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
			attribute.ID = oid.Hex()
		}
		return nil
	}

	// Si tiene ID, es una actualización (upsert)
	objectID, err := primitive.ObjectIDFromHex(attribute.ID)
	if err != nil {
		return fmt.Errorf("invalid object ID: %w", err)
	}

	filter := bson.M{"_id": objectID, "deleted_at": bson.M{"$exists": false}}
	update := bson.M{"$set": doc}

	opts := options.Update().SetUpsert(true)
	_, err = r.collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return fmt.Errorf("failed to upsert attribute: %w", err)
	}

	return nil
}

// GetByID obtiene un atributo por su ID
func (r *TenantCustomAttributeMongoRepository) GetByID(ctx context.Context, id string) (*entity.TenantCustomAttribute, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid object ID: %w", err)
	}

	filter := bson.M{
		"_id":        objectID,
		"deleted_at": bson.M{"$exists": false},
	}

	var doc mongoAttribute
	err = r.collection.FindOne(ctx, filter).Decode(&doc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find attribute: %w", err)
	}

	return r.mongoToEntity(&doc), nil
}

// GetByTenantAndSlug obtiene un atributo por tenant y slug
func (r *TenantCustomAttributeMongoRepository) GetByTenantAndSlug(ctx context.Context, tenantID, slug string) (*entity.TenantCustomAttribute, error) {
	filter := bson.M{
		"tenant_id":  tenantID,
		"slug":       slug,
		"deleted_at": bson.M{"$exists": false},
	}

	var doc mongoAttribute
	err := r.collection.FindOne(ctx, filter).Decode(&doc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find attribute: %w", err)
	}

	return r.mongoToEntity(&doc), nil
}

// GetByTenantID obtiene todos los atributos de un tenant
func (r *TenantCustomAttributeMongoRepository) GetByTenantID(ctx context.Context, tenantID string) ([]*entity.TenantCustomAttribute, error) {
	filter := bson.M{
		"tenant_id":  tenantID,
		"deleted_at": bson.M{"$exists": false},
	}

	opts := options.Find().SetSort(bson.D{bson.E{Key: "sort_order", Value: 1}, bson.E{Key: "name", Value: 1}})
	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to find attributes: %w", err)
	}
	defer cursor.Close(ctx)

	return r.cursorToEntities(ctx, cursor)
}

// GetGlobalByTenant obtiene atributos globales de un tenant
func (r *TenantCustomAttributeMongoRepository) GetGlobalByTenant(ctx context.Context, tenantID string) ([]*entity.TenantCustomAttribute, error) {
	filter := bson.M{
		"tenant_id":               tenantID,
		"marketplace_category_id": bson.M{"$exists": false},
		"deleted_at":              bson.M{"$exists": false},
	}

	opts := options.Find().SetSort(bson.D{bson.E{Key: "sort_order", Value: 1}, bson.E{Key: "name", Value: 1}})
	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to find global attributes: %w", err)
	}
	defer cursor.Close(ctx)

	return r.cursorToEntities(ctx, cursor)
}

// GetByTenantAndCategory obtiene atributos de un tenant para una categoría específica
func (r *TenantCustomAttributeMongoRepository) GetByTenantAndCategory(ctx context.Context, tenantID, marketplaceCategoryID string) ([]*entity.TenantCustomAttribute, error) {
	filter := bson.M{
		"tenant_id":               tenantID,
		"marketplace_category_id": marketplaceCategoryID,
		"deleted_at":              bson.M{"$exists": false},
	}

	opts := options.Find().SetSort(bson.D{bson.E{Key: "sort_order", Value: 1}, bson.E{Key: "name", Value: 1}})
	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to find category attributes: %w", err)
	}
	defer cursor.Close(ctx)

	return r.cursorToEntities(ctx, cursor)
}

// GetFilterableByTenant obtiene atributos filtrables de un tenant
func (r *TenantCustomAttributeMongoRepository) GetFilterableByTenant(ctx context.Context, tenantID string) ([]*entity.TenantCustomAttribute, error) {
	filter := bson.M{
		"tenant_id":     tenantID,
		"is_filterable": true,
		"deleted_at":    bson.M{"$exists": false},
	}

	opts := options.Find().SetSort(bson.D{bson.E{Key: "sort_order", Value: 1}, bson.E{Key: "name", Value: 1}})
	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to find filterable attributes: %w", err)
	}
	defer cursor.Close(ctx)

	return r.cursorToEntities(ctx, cursor)
}

// GetSearchableByTenant obtiene atributos buscables de un tenant
func (r *TenantCustomAttributeMongoRepository) GetSearchableByTenant(ctx context.Context, tenantID string) ([]*entity.TenantCustomAttribute, error) {
	filter := bson.M{
		"tenant_id":     tenantID,
		"is_searchable": true,
		"deleted_at":    bson.M{"$exists": false},
	}

	opts := options.Find().SetSort(bson.D{bson.E{Key: "sort_order", Value: 1}, bson.E{Key: "name", Value: 1}})
	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to find searchable attributes: %w", err)
	}
	defer cursor.Close(ctx)

	return r.cursorToEntities(ctx, cursor)
}

// GetTenantAttributesForCategory obtiene todos los atributos aplicables a una categoría
func (r *TenantCustomAttributeMongoRepository) GetTenantAttributesForCategory(ctx context.Context, tenantID, marketplaceCategoryID string) ([]*entity.TenantCustomAttribute, error) {
	filter := bson.M{
		"tenant_id": tenantID,
		"$or": []bson.M{
			{"marketplace_category_id": marketplaceCategoryID},
			{"marketplace_category_id": bson.M{"$exists": false}},
		},
		"deleted_at": bson.M{"$exists": false},
	}

	opts := options.Find().SetSort(bson.D{
		bson.E{Key: "marketplace_category_id", Value: 1}, // Globales primero (null values)
		bson.E{Key: "sort_order", Value: 1},
		bson.E{Key: "name", Value: 1},
	})

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to find category attributes: %w", err)
	}
	defer cursor.Close(ctx)

	return r.cursorToEntities(ctx, cursor)
}

// FindByCriteria busca atributos según criterios (implementación básica)
func (r *TenantCustomAttributeMongoRepository) FindByCriteria(ctx context.Context, crit domainCriteria.Criteria) ([]*entity.TenantCustomAttribute, error) {
	// Implementación básica - se puede mejorar con un convertidor de criterios a MongoDB
	filter := bson.M{"deleted_at": bson.M{"$exists": false}}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to find attributes by criteria: %w", err)
	}
	defer cursor.Close(ctx)

	return r.cursorToEntities(ctx, cursor)
}

// CountByCriteria cuenta atributos según criterios
func (r *TenantCustomAttributeMongoRepository) CountByCriteria(ctx context.Context, crit domainCriteria.Criteria) (int, error) {
	filter := bson.M{"deleted_at": bson.M{"$exists": false}}

	count, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, fmt.Errorf("failed to count attributes: %w", err)
	}

	return int(count), nil
}

// Update actualiza un atributo
func (r *TenantCustomAttributeMongoRepository) Update(ctx context.Context, attribute *entity.TenantCustomAttribute) error {
	objectID, err := primitive.ObjectIDFromHex(attribute.ID)
	if err != nil {
		return fmt.Errorf("invalid object ID: %w", err)
	}

	filter := bson.M{
		"_id":        objectID,
		"deleted_at": bson.M{"$exists": false},
	}

	doc := r.entityToMongo(attribute)
	update := bson.M{"$set": doc}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update attribute: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("attribute with id %s not found", attribute.ID)
	}

	return nil
}

// Delete elimina un atributo (soft delete)
func (r *TenantCustomAttributeMongoRepository) Delete(ctx context.Context, id string) error {
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
		return fmt.Errorf("failed to delete attribute: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("attribute with id %s not found", id)
	}

	return nil
}

// ExistsByTenantAndSlug verifica si ya existe un atributo con el slug en el tenant
func (r *TenantCustomAttributeMongoRepository) ExistsByTenantAndSlug(ctx context.Context, tenantID, slug string, marketplaceCategoryID *string) (bool, error) {
	filter := bson.M{
		"tenant_id":  tenantID,
		"slug":       slug,
		"deleted_at": bson.M{"$exists": false},
	}

	if marketplaceCategoryID == nil {
		filter["marketplace_category_id"] = bson.M{"$exists": false}
	} else {
		filter["marketplace_category_id"] = *marketplaceCategoryID
	}

	count, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return false, fmt.Errorf("failed to check attribute existence: %w", err)
	}

	return count > 0, nil
}

// entityToMongo convierte una entidad a documento MongoDB
func (r *TenantCustomAttributeMongoRepository) entityToMongo(attribute *entity.TenantCustomAttribute) *mongoAttribute {
	doc := &mongoAttribute{
		TenantID:              attribute.TenantID,
		MarketplaceCategoryID: attribute.MarketplaceCategoryID,
		Name:                  attribute.Name,
		Slug:                  attribute.Slug,
		Type:                  attribute.Type,
		IsFilterable:          attribute.IsFilterable,
		IsSearchable:          attribute.IsSearchable,
		ValidationRules:       attribute.ValidationRules, // ✅ Sin serialización necesaria
		SortOrder:             attribute.SortOrder,
		CreatedAt:             attribute.CreatedAt,
		UpdatedAt:             attribute.UpdatedAt,
	}

	// Convertir ID si existe
	if attribute.ID != "" {
		if objectID, err := primitive.ObjectIDFromHex(attribute.ID); err == nil {
			doc.ID = objectID
		}
	}

	return doc
}

// mongoToEntity convierte un documento MongoDB a entidad
func (r *TenantCustomAttributeMongoRepository) mongoToEntity(doc *mongoAttribute) *entity.TenantCustomAttribute {
	attribute := &entity.TenantCustomAttribute{
		ID:                    doc.ID.Hex(),
		TenantID:              doc.TenantID,
		MarketplaceCategoryID: doc.MarketplaceCategoryID,
		Name:                  doc.Name,
		Slug:                  doc.Slug,
		Type:                  doc.Type,
		IsFilterable:          doc.IsFilterable,
		IsSearchable:          doc.IsSearchable,
		ValidationRules:       doc.ValidationRules, // ✅ Sin deserialización necesaria
		SortOrder:             doc.SortOrder,
		CreatedAt:             doc.CreatedAt,
		UpdatedAt:             doc.UpdatedAt,
	}

	return attribute
}

// cursorToEntities convierte un cursor MongoDB a slice de entidades
func (r *TenantCustomAttributeMongoRepository) cursorToEntities(ctx context.Context, cursor *mongo.Cursor) ([]*entity.TenantCustomAttribute, error) {
	var attributes []*entity.TenantCustomAttribute

	for cursor.Next(ctx) {
		var doc mongoAttribute
		if err := cursor.Decode(&doc); err != nil {
			return nil, fmt.Errorf("failed to decode attribute: %w", err)
		}

		attributes = append(attributes, r.mongoToEntity(&doc))
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return attributes, nil
}
