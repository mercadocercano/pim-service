package value_object_test

import (
	"testing"

	"saas-mt-pim-service/src/product/global_catalog/domain/value_object"
)

func strPtr(s string) *string { return &s }

// TestResolveSafeBusinessTypeTransition cubre las invariantes del ADR-006 §"Invariantes a preservar".
func TestResolveSafeBusinessTypeTransition(t *testing.T) {
	tests := []struct {
		name      string
		current   *string
		candidate string
		wantApply bool
		wantType  string
		wantKind  string
	}{
		{
			name:      "relleno: current nil → candidate",
			current:   nil,
			candidate: "ferreteria",
			wantApply: true, wantType: "ferreteria", wantKind: value_object.TransitionRelleno,
		},
		{
			name:      "relleno: current vacío → candidate",
			current:   strPtr(""),
			candidate: "veterinaria",
			wantApply: true, wantType: "veterinaria", wantKind: value_object.TransitionRelleno,
		},
		{
			name:      "correccion: almacen → rubro específico",
			current:   strPtr("almacen"),
			candidate: "bazar",
			wantApply: true, wantType: "bazar", wantKind: value_object.TransitionCorreccion,
		},
		{
			name:      "skip ya_especifico: nunca degrada un rubro curado",
			current:   strPtr("ferreteria"),
			candidate: "bazar",
			wantApply: false, wantType: "", wantKind: value_object.TransitionSkipYaEspecifico,
		},
		{
			name:      "skip ya_especifico: aunque el candidate coincida, no toca el rubro curado",
			current:   strPtr("ferreteria"),
			candidate: "ferreteria",
			wantApply: false, wantType: "", wantKind: value_object.TransitionSkipYaEspecifico,
		},
		{
			name:      "skip ya_correcto: almacen == almacen",
			current:   strPtr("almacen"),
			candidate: "almacen",
			wantApply: false, wantType: "", wantKind: value_object.TransitionSkipYaCorrecto,
		},
		{
			name:      "skip sin_candidate: candidate vacío con current nil",
			current:   nil,
			candidate: "",
			wantApply: false, wantType: "", wantKind: value_object.TransitionSkipSinCandidate,
		},
		{
			name:      "skip sin_candidate: candidate vacío no degrada rubro curado",
			current:   strPtr("ferreteria"),
			candidate: "",
			wantApply: false, wantType: "", wantKind: value_object.TransitionSkipSinCandidate,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apply, newType, kind := value_object.ResolveSafeBusinessTypeTransition(tt.current, tt.candidate)
			if apply != tt.wantApply || newType != tt.wantType || kind != tt.wantKind {
				t.Errorf("ResolveSafeBusinessTypeTransition(%v, %q) = (%v, %q, %q); want (%v, %q, %q)",
					tt.current, tt.candidate, apply, newType, kind, tt.wantApply, tt.wantType, tt.wantKind)
			}
		})
	}
}

// TestResolveSafeBusinessTypeTransition_Idempotente: aplicar el candidate y reevaluar → skip (0 cambios).
func TestResolveSafeBusinessTypeTransition_Idempotente(t *testing.T) {
	current := (*string)(nil)
	candidate := "veterinaria"

	apply, newType, _ := value_object.ResolveSafeBusinessTypeTransition(current, candidate)
	if !apply {
		t.Fatalf("primera evaluación debería aplicar")
	}
	// Simular el estado tras aplicar.
	current = &newType

	apply2, _, kind2 := value_object.ResolveSafeBusinessTypeTransition(current, candidate)
	if apply2 {
		t.Errorf("re-sync de un producto ya clasificado debe ser idempotente (0 cambios), got apply=true kind=%q", kind2)
	}
}
