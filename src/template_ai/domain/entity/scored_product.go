package entity

// ScoredProduct represents a product with AI scoring
type ScoredProduct struct {
	Product   *GlobalProduct
	Score     float64
	Reasoning string
	Tags      []string
}