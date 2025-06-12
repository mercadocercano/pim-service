// MongoDB initialization script for marketplace collections
print('🚀 Initializing PIM Marketplace MongoDB...');

// Switch to marketplace database
db = db.getSiblingDB('pim_marketplace');

// Create collections with validation schemas
print('📋 Creating tenant_custom_attributes collection...');
db.createCollection('tenant_custom_attributes', {
  validator: {
    $jsonSchema: {
      bsonType: 'object',
      required: ['tenant_id', 'name', 'slug', 'type'],
      properties: {
        tenant_id: {
          bsonType: 'string',
          description: 'Tenant ID is required'
        },
        marketplace_category_id: {
          bsonType: ['string', 'null'],
          description: 'Marketplace category ID (optional for global attributes)'
        },
        name: {
          bsonType: 'string',
          description: 'Attribute name is required'
        },
        slug: {
          bsonType: 'string',
          description: 'Attribute slug is required'
        },
        type: {
          bsonType: 'string',
          enum: ['text', 'number', 'boolean', 'select', 'multi_select'],
          description: 'Attribute type must be one of the enum values'
        },
        is_filterable: {
          bsonType: 'bool',
          description: 'Whether the attribute is filterable'
        },
        is_searchable: {
          bsonType: 'bool',
          description: 'Whether the attribute is searchable'
        },
        validation_rules: {
          bsonType: 'object',
          description: 'Validation rules as flexible JSON object'
        },
        sort_order: {
          bsonType: 'int',
          description: 'Sort order for display'
        },
        created_at: {
          bsonType: 'date',
          description: 'Creation timestamp'
        },
        updated_at: {
          bsonType: 'date',
          description: 'Last update timestamp'
        },
        deleted_at: {
          bsonType: ['date', 'null'],
          description: 'Soft delete timestamp'
        }
      }
    }
  }
});

print('📋 Creating tenant_category_mappings collection...');
db.createCollection('tenant_category_mappings', {
  validator: {
    $jsonSchema: {
      bsonType: 'object',
      required: ['tenant_id', 'category_id', 'marketplace_category_id'],
      properties: {
        tenant_id: {
          bsonType: 'string',
          description: 'Tenant ID is required'
        },
        category_id: {
          bsonType: 'string',
          description: 'Tenant category ID is required'
        },
        marketplace_category_id: {
          bsonType: 'string',
          description: 'Marketplace category ID is required'
        },
        custom_name: {
          bsonType: ['string', 'null'],
          description: 'Custom name for the mapping'
        },
        custom_description: {
          bsonType: ['string', 'null'],
          description: 'Custom description for the mapping'
        },
        is_active: {
          bsonType: 'bool',
          description: 'Whether the mapping is active'
        },
        metadata: {
          bsonType: 'object',
          description: 'Additional metadata as flexible JSON'
        },
        created_at: {
          bsonType: 'date',
          description: 'Creation timestamp'
        },
        updated_at: {
          bsonType: 'date',
          description: 'Last update timestamp'
        },
        deleted_at: {
          bsonType: ['date', 'null'],
          description: 'Soft delete timestamp'
        }
      }
    }
  }
});

// Create indexes for performance
print('🔍 Creating indexes...');

// Indexes for tenant_custom_attributes
db.tenant_custom_attributes.createIndex(
  { 'tenant_id': 1, 'slug': 1, 'marketplace_category_id': 1 },
  { 
    unique: true, 
    partialFilterExpression: { deleted_at: { $exists: false } },
    name: 'unique_tenant_slug_category'
  }
);

db.tenant_custom_attributes.createIndex(
  { 'tenant_id': 1, 'marketplace_category_id': 1 },
  { name: 'tenant_category_lookup' }
);

db.tenant_custom_attributes.createIndex(
  { 'tenant_id': 1, 'is_filterable': 1 },
  { name: 'tenant_filterable_attributes' }
);

db.tenant_custom_attributes.createIndex(
  { 'tenant_id': 1, 'is_searchable': 1 },
  { name: 'tenant_searchable_attributes' }
);

db.tenant_custom_attributes.createIndex(
  { 'type': 1 },
  { name: 'attribute_type_lookup' }
);

db.tenant_custom_attributes.createIndex(
  { 'created_at': 1 },
  { name: 'creation_time_sort' }
);

// Indexes for tenant_category_mappings
db.tenant_category_mappings.createIndex(
  { 'tenant_id': 1, 'category_id': 1 },
  { 
    unique: true,
    partialFilterExpression: { deleted_at: { $exists: false } },
    name: 'unique_tenant_category'
  }
);

db.tenant_category_mappings.createIndex(
  { 'tenant_id': 1, 'marketplace_category_id': 1 },
  { name: 'tenant_marketplace_lookup' }
);

db.tenant_category_mappings.createIndex(
  { 'marketplace_category_id': 1 },
  { name: 'marketplace_category_lookup' }
);

db.tenant_category_mappings.createIndex(
  { 'tenant_id': 1, 'is_active': 1 },
  { name: 'tenant_active_mappings' }
);

// Create user for application access
print('👤 Creating application user...');
db.createUser({
  user: 'pim_app',
  pwd: 'pim_app_password',
  roles: [
    {
      role: 'readWrite',
      db: 'pim_marketplace'
    }
  ]
});

print('✅ MongoDB initialization completed successfully!');
print('📊 Collections created:');
print('  - tenant_custom_attributes (with validation schema)');
print('  - tenant_category_mappings (with validation schema)');
print('🔍 Indexes created for optimal performance');
print('👤 Application user created: pim_app'); 