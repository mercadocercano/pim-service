package port

import (
	"context"
	cr "github.com/hornosg/go-shared/criteria"
	"saas-mt-pim-service/src/product/global_catalog/domain/entity"
)

// GlobalProductRepository define los métodos para persistir GlobalProduct
type GlobalProductRepository interface {
	// Operaciones básicas CRUD
	Save(globalProduct *entity.GlobalProduct) (*entity.GlobalProduct, error)
	Update(globalProduct *entity.GlobalProduct) (*entity.GlobalProduct, error)
	FindByID(id string) (*entity.GlobalProduct, error)
	Delete(id string) error

	// Búsquedas específicas del catálogo global
	FindByEAN(ean string) (*entity.GlobalProduct, error)
	FindActiveByEAN(ean string) (*entity.GlobalProduct, error)
	FindByBusinessType(businessType string, limit int) ([]*entity.GlobalProduct, error)
	FindBySource(source string, limit int) ([]*entity.GlobalProduct, error)
	FindByQualityScoreRange(minScore, maxScore int, limit int) ([]*entity.GlobalProduct, error)

	// Búsquedas avanzadas
	SearchByName(name string, limit int) ([]*entity.GlobalProduct, error)
	SearchByBrand(brand string, limit int) ([]*entity.GlobalProduct, error)
	SearchByCategory(category string, limit int) ([]*entity.GlobalProduct, error)
	SearchByTags(tags []string, limit int) ([]*entity.GlobalProduct, error)

	// Listados con paginación
	FindAll(offset, limit int) ([]*entity.GlobalProduct, error)
	FindActive(offset, limit int) ([]*entity.GlobalProduct, error)
	FindVerified(offset, limit int) ([]*entity.GlobalProduct, error)

	// Productos argentinos específicamente
	FindArgentineProducts(offset, limit int) ([]*entity.GlobalProduct, error)
	FindHighQualityProducts(offset, limit int) ([]*entity.GlobalProduct, error)

	// Productos que necesitan actualización
	FindNeedingUpdate(maxAgeHours int, limit int) ([]*entity.GlobalProduct, error)

	// Productos que necesitan enrichment (sin imagen, precio o marca, o quality_score < 70)
	FindNeedingEnrichment(businessType *string, limit, offset int) ([]*entity.GlobalProduct, error)
	CountNeedingEnrichment(businessType *string) (int, error)

	// Agregaciones por business_type
	FindDistinctBrandsByBusinessType(businessType string) ([]string, error)
	FindDistinctCategoriesByBusinessType(businessType string) ([]string, error)
	FindDistinctBusinessTypes() ([]string, error)

	// Estadísticas
	CountTotal() (int, error)
	CountBySource(source string) (int, error)
	CountByQualityScore(minScore int) (int, error)
	CountArgentineProducts() (int, error)

	// Búsqueda con criterios (para compatibilidad con el sistema existente)
	SearchByCriteria(ctx context.Context, crit cr.Criteria) ([]*entity.GlobalProduct, error)
	CountByCriteria(ctx context.Context, crit cr.Criteria) (int, error)

	// Backfill de imágenes: match por nombre+marca, retorna solo si tiene imagen
	FindByNameAndBrand(ctx context.Context, name, brand string) (*entity.GlobalProduct, error)

	// Búsqueda por múltiples IDs (on-demand enrichment)
	FindByIDs(ctx context.Context, ids []string) ([]*entity.GlobalProduct, error)

	// CountTenantLinks retorna cuántos tenants tienen una referencia al producto global
	CountTenantLinks(ctx context.Context, productID string) (int, error)
}
