//go:build ignore

// normalize_category_dryrun.go — corre el DRY-RUN (read-only) del use case de normalización de
// category_slug contra una DB Postgres, reusando el repo + use case REALES (ADR-007 §4). NO muta:
// dry_run=true nunca crea snapshot ni UPDATE. Sirve para ver la cobertura del resolver y el
// worklist de overrides ANTES de aplicar el backfill vía el endpoint S2S.
//
// Uso:
//
//	DB_DSN="postgres://postgres:postgres@localhost:5432/pim_db?sslmode=disable" \
//	  go run scripts/normalize_category_dryrun.go
package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"sort"

	_ "github.com/lib/pq"

	"github.com/hornosg/go-shared/domain/category"
	uc "saas-mt-pim-service/src/product/global_catalog/application/usecase"
	persistence "saas-mt-pim-service/src/product/global_catalog/infrastructure/persistence"
	vo "saas-mt-pim-service/src/product/global_catalog/domain/value_object"
)

func main() {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "postgres://postgres:postgres@localhost:5432/pim_db?sslmode=disable"
	}
	db, err := sql.Open("postgres", dsn)
	must(err)
	defer db.Close()
	must(db.Ping())

	repo := persistence.NewPostgresNormalizeCategoryRepository(db)
	usecase := uc.NewNormalizeCategorySlugsUseCase(repo, nil)

	scope, err := vo.NewNormalizeCategoryScope("", vo.NormalizeCategoryMaxRowsCap)
	must(err)

	resp, err := usecase.Execute(context.Background(), uc.NormalizeRequest{DryRun: true, Scope: scope})
	must(err)

	declared := loadDeclaredSlugs(db)

	s := resp.Summary
	fmt.Println("=== DRY-RUN normalize-category-slugs (read-only) ===")
	fmt.Printf("mode: %s\n", resp.Mode)
	fmt.Printf("categorias distintas: %d | productos: %d | productos afectados: %d\n",
		s.TotalCategorias, s.TotalProductos, s.ProductosAfectados)

	// Cobertura REAL: productos cuyo slug producido EXISTE en el vocabulario declarado
	// (el join del refresh funcionaría). El resto necesita override (slug-leaf != declarado).
	var prodDeclared, prodUndeclared int
	undeclared := map[string]int{}
	for slug, n := range s.TopSlugs {
		if declared[slug] {
			prodDeclared += n
		} else {
			prodUndeclared += n
			undeclared[slug] = n
		}
	}
	fmt.Printf("\n=== COBERTURA vs vocabulario declarado (%d slugs declarados) ===\n", len(declared))
	fmt.Printf("productos con slug DECLARADO (join OK):    %d (%.1f%%)\n", prodDeclared, pct(prodDeclared, s.TotalProductos))
	fmt.Printf("productos con slug NO declarado (override): %d (%.1f%%)\n", prodUndeclared, pct(prodUndeclared, s.TotalProductos))

	// Lever automático: ¿y si normalizamos TAMBIÉN los slugs declarados con el mismo resolver?
	declaredNorm := map[string]bool{}
	for d := range declared {
		if ns, ok := category.ResolveCategorySlug(d); ok {
			declaredNorm[ns] = true
		}
		declaredNorm[d] = true // conservar el original también
	}
	var prodDeclNorm int
	for slug, n := range s.TopSlugs {
		if declaredNorm[slug] {
			prodDeclNorm += n
		}
	}
	fmt.Printf("\n-- si normalizamos TAMBIEN los slugs declarados (leaf) --\n")
	fmt.Printf("productos con slug declarado-normalizado: %d (%.1f%%)\n", prodDeclNorm, pct(prodDeclNorm, s.TotalProductos))

	// Métrica que decide la riqueza del template: cuántas categorías DECLARADAS quedan con >=1
	// producto (el refresh toma top-30 por categoría, no le importa el % global de productos).
	declaredWithProducts := 0
	for d := range declared {
		if s.TopSlugs[d] > 0 {
			declaredWithProducts++
		}
	}
	fmt.Printf("\n=== COBERTURA A NIVEL CATEGORIA DECLARADA (lo que llena los templates) ===\n")
	fmt.Printf("categorias declaradas con >=1 producto: %d / %d (%.1f%%)\n",
		declaredWithProducts, len(declared), pct(declaredWithProducts, len(declared)))

	fmt.Printf("\n-- top 40 slugs NO declarados (worklist de overrides, por productos) --\n")
	printTop(undeclared, 40)
	fmt.Printf("\n-- sin-clasificar (%d categorias) --\n", len(s.SinClasificarMuestra))
	for _, c := range s.SinClasificarMuestra {
		fmt.Printf("  %s\n", c)
	}
}

func loadDeclaredSlugs(db *sql.DB) map[string]bool {
	rows, err := db.Query(`SELECT DISTINCT c->>'slug' FROM business_type_templates btt,
		jsonb_array_elements(btt.categories) c WHERE btt.is_active`)
	must(err)
	defer rows.Close()
	set := map[string]bool{}
	for rows.Next() {
		var s sql.NullString
		must(rows.Scan(&s))
		if s.Valid {
			set[s.String] = true
		}
	}
	return set
}

func pct(n, total int) float64 {
	if total == 0 {
		return 0
	}
	return 100.0 * float64(n) / float64(total)
}

func printMap(m map[string]int) {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("  %-16s %d\n", k, m[k])
	}
}

func printTop(m map[string]int, n int) {
	type kv struct {
		k string
		v int
	}
	pairs := make([]kv, 0, len(m))
	for k, v := range m {
		pairs = append(pairs, kv{k, v})
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].v > pairs[j].v })
	for i, p := range pairs {
		if i >= n {
			break
		}
		fmt.Printf("  %-32s %d\n", p.k, p.v)
	}
}

func must(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
		os.Exit(1)
	}
}
