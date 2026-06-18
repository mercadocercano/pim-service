package value_object_test

import (
	"strings"
	"testing"

	"saas-mt-pim-service/src/product/global_catalog/domain/value_object"
)

// T-013: MaxRows=0 → error de validación.
func TestNewReclassifyScope_ZeroMaxRows_ReturnsError(t *testing.T) {
	_, err := value_object.NewReclassifyScope("scraper", 0)
	if err == nil {
		t.Fatal("T-013: expected error for MaxRows=0, got nil")
	}
}

// T-014: MaxRows=50001 → error con código SCOPE_EXCEEDS_CAP.
func TestNewReclassifyScope_ExceedsCap_ReturnsScopeExceedsCapError(t *testing.T) {
	_, err := value_object.NewReclassifyScope("scraper", value_object.ReclassifyMaxRowsCap+1)
	if err == nil {
		t.Fatal("T-014: expected error for MaxRows > cap, got nil")
	}
	if !strings.Contains(err.Error(), value_object.ErrCodeScopeExceedsCap) {
		t.Fatalf("T-014: expected error code %q in %q", value_object.ErrCodeScopeExceedsCap, err.Error())
	}
}

// T-015: MaxRows=50000 → scope válido.
func TestNewReclassifyScope_AtCap_ReturnsValidScope(t *testing.T) {
	scope, err := value_object.NewReclassifyScope("scraper", value_object.ReclassifyMaxRowsCap)
	if err != nil {
		t.Fatalf("T-015: expected no error, got %v", err)
	}
	if scope.MaxRows != value_object.ReclassifyMaxRowsCap {
		t.Fatalf("T-015: expected MaxRows=%d, got %d", value_object.ReclassifyMaxRowsCap, scope.MaxRows)
	}
	if scope.SourcePrefix != "scraper" {
		t.Fatalf("T-015: expected SourcePrefix=scraper, got %q", scope.SourcePrefix)
	}
}
