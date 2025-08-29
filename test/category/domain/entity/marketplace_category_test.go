package entity

import (
	"testing"
	"time"

	marketplaceEntity "saas-mt-pim-service/src/category/domain/entity"
)

func TestNewMarketplaceCategory(t *testing.T) {
	tests := []struct {
		name         string
		categoryName string
		slug         string
		description  string
		parentID     *string
		wantErr      bool
	}{
		{
			name:         "Valid root category",
			categoryName: "Electronics",
			slug:         "electronics",
			description:  "Electronic devices and gadgets",
			parentID:     nil,
			wantErr:      false,
		},
		{
			name:         "Valid child category",
			categoryName: "Smartphones",
			slug:         "smartphones",
			description:  "Mobile phones and accessories",
			parentID:     stringPtr("cat-001"),
			wantErr:      false,
		},
		{
			name:         "Empty name should fail",
			categoryName: "",
			slug:         "electronics",
			description:  "Description",
			parentID:     nil,
			wantErr:      true,
		},
		{
			name:         "Empty slug should fail",
			categoryName: "Electronics",
			slug:         "",
			description:  "Description",
			parentID:     nil,
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			category, err := marketplaceEntity.NewMarketplaceCategory(tt.categoryName, tt.slug, tt.description, tt.parentID)

			if tt.wantErr {
				if err == nil {
					t.Errorf("NewMarketplaceCategory() expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("NewMarketplaceCategory() unexpected error: %v", err)
				return
			}

			if category.Name != tt.categoryName {
				t.Errorf("NewMarketplaceCategory() Name = %v, want %v", category.Name, tt.categoryName)
			}

			if category.Slug != tt.slug {
				t.Errorf("NewMarketplaceCategory() Slug = %v, want %v", category.Slug, tt.slug)
			}

			if category.Description != tt.description {
				t.Errorf("NewMarketplaceCategory() Description = %v, want %v", category.Description, tt.description)
			}

			if !category.IsActive {
				t.Errorf("NewMarketplaceCategory() IsActive should be true by default")
			}

			if category.SortOrder != 0 {
				t.Errorf("NewMarketplaceCategory() SortOrder should be 0 by default")
			}
		})
	}
}

func TestMarketplaceCategory_Update(t *testing.T) {
	category, _ := marketplaceEntity.NewMarketplaceCategory("Electronics", "electronics", "Description", nil)

	tests := []struct {
		name    string
		newName string
		newSlug string
		newDesc string
		wantErr bool
	}{
		{
			name:    "Valid update",
			newName: "Electronics & Gadgets",
			newSlug: "electronics-gadgets",
			newDesc: "Updated description",
			wantErr: false,
		},
		{
			name:    "Empty name should fail",
			newName: "",
			newSlug: "electronics",
			newDesc: "Description",
			wantErr: true,
		},
		{
			name:    "Empty slug should fail",
			newName: "Electronics",
			newSlug: "",
			newDesc: "Description",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := category.Update(tt.newName, tt.newSlug, tt.newDesc)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Update() expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Update() unexpected error: %v", err)
				return
			}

			if category.Name != tt.newName {
				t.Errorf("Update() Name = %v, want %v", category.Name, tt.newName)
			}
		})
	}
}

func TestMarketplaceCategory_Activate(t *testing.T) {
	category, _ := marketplaceEntity.NewMarketplaceCategory("Electronics", "electronics", "Description", nil)
	category.IsActive = false

	oldUpdatedAt := category.UpdatedAt
	time.Sleep(10 * time.Millisecond)

	category.Activate()

	if !category.IsActive {
		t.Errorf("Activate() IsActive should be true after activation")
	}

	if !category.UpdatedAt.After(oldUpdatedAt) {
		t.Errorf("Activate() should update UpdatedAt timestamp")
	}
}

func TestMarketplaceCategory_Deactivate(t *testing.T) {
	category, _ := marketplaceEntity.NewMarketplaceCategory("Electronics", "electronics", "Description", nil)

	oldUpdatedAt := category.UpdatedAt
	time.Sleep(10 * time.Millisecond)

	category.Deactivate()

	if category.IsActive {
		t.Errorf("Deactivate() IsActive should be false after deactivation")
	}

	if !category.UpdatedAt.After(oldUpdatedAt) {
		t.Errorf("Deactivate() should update UpdatedAt timestamp")
	}
}

func TestMarketplaceCategory_IsRoot(t *testing.T) {
	rootCategory, _ := marketplaceEntity.NewMarketplaceCategory("Electronics", "electronics", "Description", nil)
	childCategory, _ := marketplaceEntity.NewMarketplaceCategory("Smartphones", "smartphones", "Description", stringPtr("cat-001"))

	if !rootCategory.IsRoot() {
		t.Errorf("IsRoot() should return true for root category")
	}

	if childCategory.IsRoot() {
		t.Errorf("IsRoot() should return false for child category")
	}
}

func TestMarketplaceCategory_ValidateLevel(t *testing.T) {
	category := &marketplaceEntity.MarketplaceCategory{Level: 4}

	err := category.ValidateLevel(3)
	if err == nil {
		t.Errorf("ValidateLevel() should fail when level exceeds maximum")
	}

	category.Level = 2
	err = category.ValidateLevel(3)
	if err != nil {
		t.Errorf("ValidateLevel() should not fail when level is within limit: %v", err)
	}
}

// Helper function for creating string pointers
func stringPtr(s string) *string {
	return &s
}
