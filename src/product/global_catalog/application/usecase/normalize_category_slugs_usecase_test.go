package usecase

import (
	"testing"

	"saas-mt-pim-service/src/product/global_catalog/domain/value_object"
)

func strptr(s string) *string { return &s }

func TestResolveMappings(t *testing.T) {
	overrides := map[string]string{
		"/Bebidas/Cervezas/": "cervezas-vinos", // override gana sobre el resolver (que daría "cervezas")
	}

	cats := []value_object.RawCategoryCount{
		{RawCategory: strptr("/Bebidas/Cervezas/"), ProductCount: 10, CurrentSlug: nil},                 // override
		{RawCategory: strptr("/Bebidas/Vinos/Vinos tintos/"), ProductCount: 5, CurrentSlug: nil},        // resolver
		{RawCategory: strptr("BEERS"), ProductCount: 3, CurrentSlug: strptr("cervezas-vinos")},          // resolver, sin cambio
		{RawCategory: nil, ProductCount: 2, CurrentSlug: nil},                                            // sin-clasificar
	}

	mappings, summary := resolveMappings(cats, overrides)

	if len(mappings) != 4 {
		t.Fatalf("expected 4 mappings, got %d", len(mappings))
	}

	// 1) override
	if mappings[0].NewSlug != "cervezas-vinos" || mappings[0].Source != value_object.SourceOverride {
		t.Errorf("override mapping wrong: %+v", mappings[0])
	}
	// 2) resolver leaf
	if mappings[1].NewSlug != "vinos-tintos" || mappings[1].Source != value_object.SourceResolver {
		t.Errorf("resolver mapping wrong: %+v", mappings[1])
	}
	// 3) resolver, ya correcto → no cambia
	if mappings[2].NewSlug != "cervezas-vinos" || mappings[2].Changed {
		t.Errorf("expected unchanged BEERS mapping, got %+v", mappings[2])
	}
	// 4) NULL category → sin-clasificar
	if mappings[3].NewSlug != "sin-clasificar" || mappings[3].Source != value_object.SourceUnclassified {
		t.Errorf("unclassified mapping wrong: %+v", mappings[3])
	}

	// Summary
	if summary.TotalCategorias != 4 || summary.TotalProductos != 20 {
		t.Errorf("summary totals wrong: %+v", summary)
	}
	if summary.CategoriasPorFuente[value_object.SourceOverride] != 1 ||
		summary.CategoriasPorFuente[value_object.SourceResolver] != 2 ||
		summary.CategoriasPorFuente[value_object.SourceUnclassified] != 1 {
		t.Errorf("categorias_por_fuente wrong: %+v", summary.CategoriasPorFuente)
	}
	// ProductosAfectados: todos cambian salvo BEERS (3) → 20-3 = 17
	if summary.ProductosAfectados != 17 {
		t.Errorf("expected 17 productos afectados, got %d", summary.ProductosAfectados)
	}
}
