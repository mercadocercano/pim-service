# Catálogo Global de Productos - Diseño y Arquitectura

## 🎯 Visión

Crear un catálogo global de productos argentinos que permita a los sellers:
- **Escanear EAN** → Producto reconocido instantáneamente
- **Solo agregar precio/stock** → Producto online en 30 segundos
- **Datos técnicos exactos** → Comparación directa entre sellers
- **Información consistente** → Sin duplicados o variaciones

## 🏗️ Arquitectura de Solución

### Servicios Involucrados

```
┌─────────────────────────────────────────────────────────────────┐
│                    SAAS-MT-GLOBAL-CATALOG-SERVICE              │
│                    (Nuevo microservicio)                       │
├─────────────────────────────────────────────────────────────────┤
│ • Catálogo global de productos                                 │
│ • APIs de búsqueda por EAN/nombre                              │
│ • Sistema de scraping ético                                    │
│ • Gestión de imágenes y specs                                  │
└─────────────────────────────────────────────────────────────────┘
                              ↕
┌─────────────────────────────────────────────────────────────────┐
│                    SAAS-MT-PIM-SERVICE                         │
│                    (Servicio existente)                        │
├─────────────────────────────────────────────────────────────────┤
│ • Templates de quickstart (existente)                          │
│ • Vinculación con catálogo global                              │
│ • Catálogos específicos por tenant                             │
└─────────────────────────────────────────────────────────────────┘
                              ↕
┌─────────────────────────────────────────────────────────────────┐
│                    SAAS-MT-SCRAPER-SERVICE                     │
│                    (Worker independiente)                      │
├─────────────────────────────────────────────────────────────────┤
│ • Scraping ético de sitios argentinos                          │
│ • Rate limiting inteligente                                    │
│ • Procesamiento de imágenes                                    │
│ • Enriquecimiento de datos                                     │
└─────────────────────────────────────────────────────────────────┘
```

## 🗄️ Estructura de Datos

### Tabla: `global_products`

```sql
CREATE TABLE global_products (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    
    -- Identificadores únicos
    ean VARCHAR(13) UNIQUE,                     -- Código EAN-13 (obligatorio)
    gtin VARCHAR(14),                           -- GTIN alternativo
    sku_global VARCHAR(100) UNIQUE,             -- SKU interno global
    
    -- Información básica
    name VARCHAR(500) NOT NULL,                 -- "Coca Cola 600ml"
    brand VARCHAR(200) NOT NULL,                -- "Coca Cola"
    manufacturer VARCHAR(200),                  -- "Coca Cola FEMSA"
    
    -- Categorización
    marketplace_category_id UUID REFERENCES marketplace_categories(id),
    product_type ENUM('industrialized', 'fmcg', 'local_brand'),
    
    -- Especificaciones técnicas (JSONB para flexibilidad)
    specifications JSONB DEFAULT '{}',          -- Peso, dimensiones, ingredientes, etc.
    technical_specs JSONB DEFAULT '{}',         -- Para electrónicos: CPU, RAM, etc.
    nutritional_info JSONB DEFAULT '{}',        -- Para alimentos: calorías, etc.
    
    -- Imágenes y multimedia
    images JSONB DEFAULT '[]',                  -- URLs de imágenes
    primary_image_url TEXT,                     -- Imagen principal
    video_urls JSONB DEFAULT '[]',              -- Videos del producto
    
    -- Metadatos de origen
    source_type ENUM('scraping', 'api', 'manual', 'partnership'),
    source_url TEXT,                            -- URL original
    source_confidence DECIMAL(3,2),             -- 0.00-1.00 confianza en los datos
    
    -- Regulatorio argentino
    anmat_code VARCHAR(50),                     -- Código ANMAT para alimentos/farmacos
    senasa_code VARCHAR(50),                    -- Código SENASA para productos agropecuarios
    
    -- Estado y calidad
    verification_status ENUM('pending', 'verified', 'rejected'),
    quality_score DECIMAL(3,2),                 -- 0.00-1.00 calidad de datos
    is_active BOOLEAN DEFAULT true,
    
    -- Timestamps
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    last_scraped_at TIMESTAMP WITH TIME ZONE
);

-- Índices para performance
CREATE INDEX idx_global_products_ean ON global_products(ean);
CREATE INDEX idx_global_products_brand ON global_products(brand);
CREATE INDEX idx_global_products_category ON global_products(marketplace_category_id);
CREATE INDEX idx_global_products_type ON global_products(product_type);
CREATE INDEX idx_global_products_verification ON global_products(verification_status);
CREATE INDEX idx_global_products_quality ON global_products(quality_score);

-- Índice de texto completo para búsqueda
CREATE INDEX idx_global_products_search ON global_products 
USING gin(to_tsvector('spanish', name || ' ' || brand));
```

### Tabla: `global_product_variations`

```sql
CREATE TABLE global_product_variations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    global_product_id UUID NOT NULL REFERENCES global_products(id),
    
    -- Variaciones (color, tamaño, sabor, etc.)
    variation_type VARCHAR(50) NOT NULL,        -- "color", "size", "flavor"
    variation_value VARCHAR(100) NOT NULL,      -- "rojo", "500ml", "frutilla"
    
    -- Identificadores específicos de la variación
    ean VARCHAR(13),                            -- EAN específico de esta variación
    sku_variation VARCHAR(100),                 -- SKU específico
    
    -- Diferencias de specs
    specifications_diff JSONB DEFAULT '{}',     -- Solo campos que difieren
    price_difference DECIMAL(10,2),             -- Diferencia de precio base
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
```

### Tabla: `business_type_product_templates`

```sql
-- Extensión de los templates existentes para incluir productos
CREATE TABLE business_type_product_templates (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    business_type_template_id UUID NOT NULL REFERENCES business_type_templates(id),
    
    -- Productos sugeridos por tipo de negocio
    suggested_products JSONB DEFAULT '[]',      -- Array de global_product IDs
    
    -- Configuración de sugerencias
    max_products_per_category INTEGER DEFAULT 50,
    priority_brands JSONB DEFAULT '[]',         -- Marcas prioritarias
    exclude_brands JSONB DEFAULT '[]',          -- Marcas a excluir
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
```

## 🔄 Flujo de Datos y Procesos

### 1. Proceso de Scraping

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Target Sites  │ -> │  Scraper Jobs   │ -> │  Data Pipeline  │
│                 │    │                 │    │                 │
│ • Disco         │    │ • Rate Limited  │    │ • Validation    │
│ • Carrefour     │    │ • Respectful    │    │ • Deduplication │
│ • Jumbo         │    │ • Rotating IPs  │    │ • Quality Score │
│ • Fravega       │    │ • Headers       │    │ • Categories    │
└─────────────────┘    └─────────────────┘    └─────────────────┘
                                                        │
                                                        v
                                              ┌─────────────────┐
                                              │ Global Catalog  │
                                              │    Database     │
                                              └─────────────────┘
```

### 2. Proceso de Quickstart con Productos

```
1. Seller selecciona tipo negocio: "Almacén"
        │
        v
2. Sistema obtiene template con categorías Y productos sugeridos
        │
        v
3. Muestra productos populares por categoría:
   • Bebidas: Coca Cola 600ml, Pepsi 500ml, Sprite 500ml
   • Lácteos: La Serenísima Leche 1L, Sancor Yogur, etc.
        │
        v
4. Seller selecciona productos base
        │
        v
5. Sistema crea vinculaciones: tenant_products -> global_products
        │
        v
6. Seller solo agrega precio/stock para cada producto
```

## 🛠️ APIs del Catálogo Global

### Búsqueda por EAN (Proceso de escaneado)

```http
GET /api/v1/global-catalog/ean/{ean}

Response:
{
  "found": true,
  "product": {
    "id": "uuid",
    "ean": "7790895000126",
    "name": "Coca Cola 600ml",
    "brand": "Coca Cola",
    "category": {
      "id": "uuid",
      "name": "Bebidas",
      "slug": "bebidas"
    },
    "specifications": {
      "volume": "600ml",
      "container": "botella plástica",
      "calories_per_100ml": 42
    },
    "images": [
      "https://cdn.global-catalog.com/coca-cola-600ml/primary.jpg"
    ],
    "quality_score": 0.95,
    "verification_status": "verified"
  }
}
```

### Productos Sugeridos por Tipo de Negocio

```http
GET /api/v1/global-catalog/suggestions/{business_type_code}
?category={category_slug}
&limit=50

Response:
{
  "business_type": "almacen",
  "category": "bebidas",
  "suggested_products": [
    {
      "id": "uuid",
      "ean": "7790895000126",
      "name": "Coca Cola 600ml",
      "brand": "Coca Cola",
      "priority_score": 0.98,
      "popularity_rank": 1,
      "estimated_margin": 0.25
    }
  ],
  "total": 25
}
```

## 🕷️ Sistema de Scraping Ético

### Targets Prioritarios Argentina

```yaml
# scraping-config.yml
targets:
  disco:
    base_url: "https://www.disco.com.ar"
    categories:
      - "/bebidas"
      - "/lacteos"
      - "/limpieza"
    rate_limit: 2  # requests per second
    respect_robots: true
    
  carrefour:
    base_url: "https://www.carrefour.com.ar"
    categories:
      - "/electronica"
      - "/hogar"
    rate_limit: 1.5
    
  fravega:
    base_url: "https://www.fravega.com"
    categories:
      - "/celulares"
      - "/computacion"
    rate_limit: 1
```

### Scraper Architecture

```go
// Scraper Service Structure
type ScraperService struct {
    targets    []ScrapingTarget
    rateLimiter *RateLimiter
    imageProcessor *ImageProcessor
    dataValidator *DataValidator
    globalCatalogAPI *GlobalCatalogAPI
}

type ScrapingTarget struct {
    Name        string
    BaseURL     string
    Categories  []string
    RateLimit   float64
    Headers     map[string]string
    Selectors   ProductSelectors
}

type ProductSelectors struct {
    Title       string
    Price       string
    EAN         string
    Brand       string
    Image       string
    Description string
    Specs       map[string]string
}
```

## 📊 Fuentes de Datos Recomendadas

### 1. APIs Oficiales (60% del catálogo)

```yaml
apis:
  open_food_facts:
    url: "https://world.openfoodfacts.org/api/v0/product/{ean}.json"
    coverage: "Alimentos argentinos"
    cost: "Gratuita"
    quality: "Alta"
    
  gs1_argentina:
    url: "https://www.gs1.org.ar/api/gtin/{gtin}"
    coverage: "Todos los productos con EAN"
    cost: "Partnership requerido"
    quality: "Oficial"
    
  google_shopping:
    url: "https://www.googleapis.com/customsearch/v1"
    coverage: "Universal"
    cost: "$5/1000 queries"
    quality: "Variable"
```

### 2. Scraping Ético (30% del catálogo)

```yaml
scraping_sources:
  tier_1:
    - "disco.com.ar"
    - "carrefour.com.ar"
    - "jumbo.com.ar"
  
  tier_2:
    - "fravega.com"
    - "garbarino.com"
    - "compumundo.com.ar"
  
  tier_3:
    - "coto.com.ar"
    - "lacaseidad.com.ar"
```

### 3. Partnerships (10% del catálogo)

```yaml
partnerships:
  distributors:
    - "Arcor"
    - "Unilever Argentina"
    - "Coca Cola FEMSA"
  
  data_providers:
    - "Nielsen Argentina"
    - "Euromonitor"
    - "Kantar"
```

## 🔧 Implementación Técnica

### Global Catalog Service (Go)

```go
// Service Structure
type GlobalCatalogService struct {
    db          *sql.DB
    redis       *redis.Client
    imageStore  ImageStore
    scraper     ScraperClient
}

// Product Search
func (s *GlobalCatalogService) FindByEAN(ctx context.Context, ean string) (*GlobalProduct, error) {
    // 1. Check cache first
    if product := s.getFromCache(ean); product != nil {
        return product, nil
    }
    
    // 2. Check database
    product, err := s.findInDatabase(ctx, ean)
    if err != nil {
        return nil, err
    }
    
    if product != nil {
        s.cacheProduct(product)
        return product, nil
    }
    
    // 3. Trigger async scraping if not found
    s.triggerScrapingJob(ean)
    return nil, ErrProductNotFound
}

// Product Suggestions for Business Type
func (s *GlobalCatalogService) GetSuggestedProducts(ctx context.Context, businessTypeCode, categorySlug string, limit int) ([]GlobalProduct, error) {
    query := `
        SELECT gp.* FROM global_products gp
        JOIN business_type_product_templates btpt ON jsonb_exists(btpt.suggested_products, gp.id::text)
        JOIN business_type_templates btt ON btpt.business_type_template_id = btt.id
        JOIN business_types bt ON btt.business_type_id = bt.id
        JOIN marketplace_categories mc ON gp.marketplace_category_id = mc.id
        WHERE bt.code = $1 
        AND mc.slug = $2
        AND gp.verification_status = 'verified'
        ORDER BY gp.quality_score DESC, gp.popularity_rank ASC
        LIMIT $3
    `
    
    // Implementation...
}
```

### Scraper Worker (Python)

```python
# scraper/main.py
import asyncio
import aiohttp
from scrapy import Spider
from ratelimit import limits, sleep_and_retry

class EthicalScraper:
    def __init__(self, config):
        self.config = config
        self.session = None
        
    @sleep_and_retry
    @limits(calls=2, period=1)  # 2 calls per second
    async def scrape_product(self, url):
        headers = {
            'User-Agent': 'GlobalCatalog-Bot/1.0 (+https://marketplace.com/robots)',
            'Accept': 'text/html,application/xhtml+xml',
            'Accept-Language': 'es-AR,es;q=0.9',
        }
        
        async with self.session.get(url, headers=headers) as response:
            if response.status == 200:
                return await self.parse_product(await response.text())
            
    async def parse_product(self, html):
        # Extract product data using selectors
        # Validate EAN format
        # Process images
        # Return structured data
        pass
```

## 🎯 Roadmap de Implementación

### Fase 1: MVP (1-2 meses)
- ✅ Crear microservicio global-catalog
- ✅ Implementar scraping de 3 sitios principales
- ✅ 500 productos top de bebidas/lácteos
- ✅ API básica de búsqueda por EAN
- ✅ Integración con templates existentes

### Fase 2: Expansión (2-3 meses)
- ✅ Scraping de 10 sitios
- ✅ 5,000 productos en 5 categorías
- ✅ Sistema de calidad y verificación
- ✅ APIs de Open Food Facts
- ✅ Panel de administración

### Fase 3: Escala (3-6 meses)
- ✅ Partnerships con distribuidores
- ✅ 50,000+ productos
- ✅ Machine learning para calidad
- ✅ APIs premium (Google Shopping)
- ✅ Sistema de contribuciones de usuarios

## ⚖️ Consideraciones Legales

### Scraping Ético
```yaml
guidelines:
  respect_robots_txt: true
  rate_limiting: "Máximo 2 req/seg por sitio"
  no_personal_data: "Solo datos públicos de productos"
  attribution: "Credited sources when possible"
  no_redistribution: "Uso interno del marketplace"
  
legal_compliance:
  argentina_law: "Ley de Protección de Datos Personales"
  copyright: "Solo datos factuales, no contenido creativo"
  terms_of_service: "Verificar ToS de cada sitio"
```

## 🎉 Impacto Esperado

### Para Sellers
- ⚡ **Onboarding 50x más rápido** - Escanear EAN vs crear desde cero
- 📊 **Datos consistentes** - Sin duplicados tipo "Coca cola" vs "Coca-Cola"
- 🎯 **Productos populares** - Sugerencias basadas en datos reales
- 💰 **Mejor posicionamiento** - Comparación directa por precio

### Para Compradores
- 🔍 **Búsqueda más efectiva** - Productos idénticos agrupados
- 💲 **Comparación de precios** - Entre diferentes sellers
- ✅ **Información confiable** - Specs técnicas verificadas
- 🇦🇷 **Productos locales** - Catálogo especializado en Argentina

### Para el Marketplace
- 📈 **Mayor conversión** - Sellers suben productos más rápido
- 🎯 **Mejor experiencia** - Catálogo consistente y completo
- 💼 **Ventaja competitiva** - Funcionalidad única en Argentina
- 📊 **Data intelligence** - Insights de productos y precios

## 🤖 Integración con AI Templates

### Flujo de Generación Inteligente

El catálogo global es la fuente principal para el sistema de Templates Inteligentes con AI:

```mermaid
graph LR
    A[Global Catalog] --> B[AI Template Engine]
    B --> C[Business Analysis]
    C --> D[Product Selection]
    D --> E[Optimized Template]
    E --> F[Tenant Catalog]
```

### Características de la Integración

#### 1. **Selección Inteligente de Productos**
- AI analiza el catálogo global completo
- Selecciona productos óptimos según:
  - Tipo de negocio
  - Ubicación geográfica
  - Preferencias del usuario
  - Datos históricos de éxito

#### 2. **Scoring y Priorización**
```sql
-- Ejemplo de query para scoring
SELECT 
    gp.*,
    CASE 
        WHEN gp.popularity_rank < 100 THEN 3  -- Essential
        WHEN gp.popularity_rank < 500 THEN 2  -- Recommended
        ELSE 1                                 -- Optional
    END as priority,
    (gp.quality_score * 0.4 + 
     (1000 - gp.popularity_rank)/1000 * 0.6) as ai_score
FROM global_products gp
WHERE gp.business_type = 'almacen'
AND gp.is_active = true
ORDER BY ai_score DESC;
```

#### 3. **Enriquecimiento con AI**
- Categorización automática mejorada
- Sugerencias de productos complementarios
- Detección de tendencias regionales
- Predicción de demanda

### Beneficios de la Integración

1. **Para Nuevos Tenants**
   - Setup completo en minutos
   - Catálogo optimizado desde el día 1
   - Reducción de curva de aprendizaje

2. **Para el Sistema**
   - Mayor adopción del catálogo global
   - Datos de feedback para mejora continua
   - Reducción de duplicados

3. **Para el Negocio**
   - Diferenciador competitivo
   - Mayor valor agregado
   - Insights de mercado

### Métricas de Éxito

- **Tiempo de onboarding**: < 5 minutos
- **Productos por template**: 50-150 optimizados
- **Tasa de retención**: > 85% de productos sugeridos
- **Satisfacción**: > 4.5/5 en templates generados

### Próximas Mejoras

1. **Machine Learning Avanzado**
   - Predicción de ventas por producto
   - Recomendaciones personalizadas
   - Ajuste dinámico de inventario

2. **Integración con Proveedores**
   - Sugerencias basadas en disponibilidad
   - Precios actualizados en tiempo real
   - Alertas de nuevos productos

¡El catálogo global sería un game-changer para el marketplace argentino! 🚀🇦🇷 