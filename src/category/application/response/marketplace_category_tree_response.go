package response

// MarketplaceCategoryTree representa una categoría en formato de árbol
type MarketplaceCategoryTree struct {
	ID       string                     `json:"id"`
	Name     string                     `json:"name"`
	Slug     string                     `json:"slug"`
	Level    int                        `json:"level"`
	ParentID *string                    `json:"parent_id,omitempty"`
	Children []*MarketplaceCategoryTree `json:"children"`
}
