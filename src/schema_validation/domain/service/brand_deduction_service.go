package service

import (
	"math"
	"strings"
	"unicode"
)

// BrandCandidate represents a potential brand detected from product names.
type BrandCandidate struct {
	Name       string  `json:"name"`
	Score      float64 `json:"score"`
	Count      int     `json:"count"`
	TotalProducts int  `json:"total_products"`
	Confidence string  `json:"confidence"`
}

// BrandDeductionService infers brand candidates from product names
// using frequency-weighted token analysis.
type BrandDeductionService struct {
	stopWords      map[string]bool
	unitPatterns   map[string]bool
	minFrequencyPct float64
	minScore       float64
}

func NewBrandDeductionService() *BrandDeductionService {
	return &BrandDeductionService{
		stopWords:      buildStopWords(),
		unitPatterns:   buildUnitPatterns(),
		minFrequencyPct: 0.05,
		minScore:       1.5,
	}
}

// DeduceBrands analyzes product names and returns brand candidates sorted by score.
func (s *BrandDeductionService) DeduceBrands(productNames []string, existingCategories []string, existingBrands []string) []BrandCandidate {
	if len(productNames) == 0 {
		return nil
	}

	categorySet := toSetLower(existingCategories)
	brandSet := toSetLower(existingBrands)

	type tokenStats struct {
		count     int
		positions []int
	}

	stats := make(map[string]*tokenStats)
	total := len(productNames)

	for _, name := range productNames {
		tokens := s.tokenize(name)
		seen := make(map[string]bool)
		for pos, token := range tokens {
			lower := strings.ToLower(token)
			if seen[lower] {
				continue
			}
			seen[lower] = true

			if s.isExcluded(lower, categorySet) {
				continue
			}

			if _, ok := stats[lower]; !ok {
				stats[lower] = &tokenStats{}
			}
			stats[lower].count++
			stats[lower].positions = append(stats[lower].positions, pos)
		}
	}

	// Also try bigrams (2-token sequences) for multi-word brands like "Coca Cola"
	for _, name := range productNames {
		tokens := s.tokenize(name)
		for i := 0; i < len(tokens)-1; i++ {
			bigram := strings.ToLower(tokens[i]) + " " + strings.ToLower(tokens[i+1])
			if s.isExcluded(strings.ToLower(tokens[i]), categorySet) || s.isExcluded(strings.ToLower(tokens[i+1]), categorySet) {
				continue
			}
			if _, ok := stats[bigram]; !ok {
				stats[bigram] = &tokenStats{}
			}
			stats[bigram].count++
			stats[bigram].positions = append(stats[bigram].positions, i)
		}
	}

	var candidates []BrandCandidate
	for token, st := range stats {
		freqPct := float64(st.count) / float64(total)
		if freqPct < s.minFrequencyPct {
			continue
		}

		if brandSet[token] {
			continue
		}

		posWeight := s.positionWeight(st.positions)
		consistency := s.positionConsistency(st.positions)
		score := freqPct * posWeight * consistency * float64(st.count)

		if score < s.minScore {
			continue
		}

		confidence := "low"
		if score >= 5.0 {
			confidence = "high"
		} else if score >= 2.5 {
			confidence = "medium"
		}

		displayName := capitalizeToken(token)
		candidates = append(candidates, BrandCandidate{
			Name:          displayName,
			Score:         math.Round(score*100) / 100,
			Count:         st.count,
			TotalProducts: total,
			Confidence:    confidence,
		})
	}

	sortCandidatesByScore(candidates)
	return candidates
}

func (s *BrandDeductionService) tokenize(name string) []string {
	var tokens []string
	for _, word := range strings.FieldsFunc(name, func(r rune) bool {
		return unicode.IsSpace(r) || r == ',' || r == ';' || r == '/' || r == '-'
	}) {
		word = strings.TrimFunc(word, func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsDigit(r)
		})
		if word != "" {
			tokens = append(tokens, word)
		}
	}
	return tokens
}

func (s *BrandDeductionService) isExcluded(token string, categorySet map[string]bool) bool {
	if s.stopWords[token] {
		return true
	}
	if s.unitPatterns[token] {
		return true
	}
	if categorySet[token] {
		return true
	}
	if isNumericOrUnit(token) {
		return true
	}
	return false
}

// positionWeight favors tokens in position 1 (index 1, after category word)
// which is where brands typically appear in Argentine product names.
func (s *BrandDeductionService) positionWeight(positions []int) float64 {
	if len(positions) == 0 {
		return 0
	}
	avgPos := 0.0
	for _, p := range positions {
		avgPos += float64(p)
	}
	avgPos /= float64(len(positions))

	if avgPos <= 1.0 {
		return 2.0
	} else if avgPos <= 2.0 {
		return 1.5
	}
	return 1.0
}

// positionConsistency measures how consistently a token appears in the same position.
func (s *BrandDeductionService) positionConsistency(positions []int) float64 {
	if len(positions) <= 1 {
		return 1.0
	}
	avg := 0.0
	for _, p := range positions {
		avg += float64(p)
	}
	avg /= float64(len(positions))

	variance := 0.0
	for _, p := range positions {
		diff := float64(p) - avg
		variance += diff * diff
	}
	variance /= float64(len(positions))

	stddev := math.Sqrt(variance)
	if stddev == 0 {
		return 2.0
	}
	return math.Max(0.5, 2.0-stddev)
}

func buildStopWords() map[string]bool {
	words := []string{
		"de", "con", "sin", "en", "y", "el", "la", "los", "las", "del",
		"para", "por", "al", "un", "una", "unos", "unas", "su", "sus",
		"tipo", "sabor", "pack", "caja", "bolsa", "lata", "botella", "sobre",
		"grande", "chico", "mediano", "extra", "super", "mega", "mini",
		"light", "diet", "zero", "free", "premium", "especial", "clasico",
		"original", "nuevo", "nueva", "natural", "integral", "organico",
	}
	m := make(map[string]bool, len(words))
	for _, w := range words {
		m[w] = true
	}
	return m
}

func buildUnitPatterns() map[string]bool {
	units := []string{
		"gr", "g", "kg", "mg", "l", "lt", "ml", "cc", "cm", "mm",
		"un", "u", "oz", "lb",
	}
	m := make(map[string]bool, len(units))
	for _, u := range units {
		m[u] = true
	}
	return m
}

func isNumericOrUnit(token string) bool {
	hasDigit := false
	for _, r := range token {
		if unicode.IsDigit(r) {
			hasDigit = true
		}
	}
	return hasDigit
}

func toSetLower(items []string) map[string]bool {
	m := make(map[string]bool, len(items))
	for _, item := range items {
		m[strings.ToLower(item)] = true
	}
	return m
}

func capitalizeToken(token string) string {
	parts := strings.Fields(token)
	for i, p := range parts {
		if len(p) > 0 {
			parts[i] = strings.ToUpper(p[:1]) + strings.ToLower(p[1:])
		}
	}
	return strings.Join(parts, " ")
}

func sortCandidatesByScore(candidates []BrandCandidate) {
	for i := 1; i < len(candidates); i++ {
		key := candidates[i]
		j := i - 1
		for j >= 0 && candidates[j].Score < key.Score {
			candidates[j+1] = candidates[j]
			j--
		}
		candidates[j+1] = key
	}
}
