# Ejemplos cURL - Catálogo Global de Productos Argentinos

Este archivo contiene ejemplos de cURL para probar las APIs del catálogo global. Estos comandos son útiles para:
- Testing de APIs
- Scraping desde scripts
- Integración con el frontend

## URLs Base
```bash
# Base URL del servicio PIM
BASE_URL="http://localhost:8084/api/v1"

# URLs públicas (sin autenticación)
PUBLIC_URL="$BASE_URL/public/global-catalog"

# URLs privadas (para administración/scraping)
PRIVATE_URL="$BASE_URL/global-catalog"
```

## 1. Health Check

```bash
# Verificar estado del servicio
curl -X GET "$PUBLIC_URL/health" \
  -H "Content-Type: application/json"
```

## 2. Crear Producto (Para Scrapers)

```bash
# Crear un producto desde scraper - Disco Argentina
curl -X POST "$PRIVATE_URL/products" \
  -H "Content-Type: application/json" \
  -d '{
    "ean": "7791234567890",
    "name": "Coca Cola 500ml Botella",
    "description": "Bebida cola sabor original en botella de 500ml. Refrescante y con el sabor clásico de Coca Cola.",
    "brand": "Coca Cola",
    "category": "Bebidas y Gaseosas",
    "price": 280.50,
    "image_url": "https://disco.com.ar/images/coca-cola-500ml.jpg",
    "source": "disco",
    "source_url": "https://www.disco.com.ar/coca-cola-500ml",
    "reliability": 0.8,
    "business_type": "kiosco",
    "tags": ["bebida", "gaseosa", "cola", "500ml", "botella"],
    "metadata": {
      "scraped_from": "disco.com.ar",
      "scraper_version": "1.0",
      "category_path": "Bebidas > Gaseosas > Cola"
    }
  }'
```

```bash
# Crear producto desde scraper - Carrefour Argentina
curl -X POST "$PRIVATE_URL/products" \
  -H "Content-Type: application/json" \
  -d '{
    "ean": "7791111222333",
    "name": "Pan Lactal Bimbo Blanco 450g",
    "description": "Pan de molde lactal blanco, ideal para tostadas y sandwiches. Envase de 450g.",
    "brand": "Bimbo",
    "category": "Panadería y Repostería",
    "price": 320.00,
    "image_url": "https://carrefour.com.ar/images/pan-lactal-bimbo-450g.jpg",
    "source": "carrefour",
    "source_url": "https://www.carrefour.com.ar/pan-lactal-bimbo-450g",
    "reliability": 0.9,
    "business_type": "almacen",
    "tags": ["pan", "lactal", "bimbo", "450g", "panaderia"],
    "metadata": {
      "scraped_from": "carrefour.com.ar",
      "scraper_version": "1.0",
      "promotion": "2x1"
    }
  }'
```

## 3. Búsqueda por EAN (Pública)

```bash
# Buscar producto por EAN - API pública
curl -X GET "$PUBLIC_URL/search?ean=7791234567890" \
  -H "Content-Type: application/json"

# Buscar por EAN con URL alternativa
curl -X GET "$PUBLIC_URL/products/ean/7791234567890" \
  -H "Content-Type: application/json"
```

## 4. Sugerencias por Tipo de Negocio (Pública)

```bash
# Sugerencias para kiosco
curl -X GET "$PUBLIC_URL/suggestions?business_type=kiosco&limit=10" \
  -H "Content-Type: application/json"

# Sugerencias para almacén
curl -X GET "$PUBLIC_URL/suggestions?business_type=almacen&limit=15" \
  -H "Content-Type: application/json"

# Sugerencias para supermercado
curl -X GET "$PUBLIC_URL/suggestions?business_type=supermercado&limit=20" \
  -H "Content-Type: application/json"
```

## 5. Listar Productos con Filtros (Privado)

```bash
# Listar todos los productos (paginado)
curl -X GET "$PRIVATE_URL/products?offset=0&limit=20" \
  -H "Content-Type: application/json"

# Listar solo productos activos
curl -X GET "$PRIVATE_URL/products?only_active=true&limit=10" \
  -H "Content-Type: application/json"

# Listar solo productos argentinos
curl -X GET "$PRIVATE_URL/products?only_argentine=true&limit=15" \
  -H "Content-Type: application/json"

# Listar productos de alta calidad
curl -X GET "$PRIVATE_URL/products?only_high_quality=true&limit=10" \
  -H "Content-Type: application/json"

# Filtrar por fuente específica
curl -X GET "$PRIVATE_URL/products?source=disco&limit=20" \
  -H "Content-Type: application/json"

# Filtrar por rango de calidad
curl -X GET "$PRIVATE_URL/products?min_quality=70&max_quality=100&limit=10" \
  -H "Content-Type: application/json"

# Buscar por nombre de producto
curl -X GET "$PRIVATE_URL/products?search_name=coca%20cola&limit=5" \
  -H "Content-Type: application/json"

# Buscar por marca
curl -X GET "$PRIVATE_URL/products?search_brand=bimbo&limit=10" \
  -H "Content-Type: application/json"

# Filtrar por tipo de negocio
curl -X GET "$PRIVATE_URL/products?business_type=kiosco&limit=15" \
  -H "Content-Type: application/json"

# Búsqueda combinada
curl -X GET "$PRIVATE_URL/products?business_type=almacen&min_quality=80&only_active=true&limit=10" \
  -H "Content-Type: application/json"
```

## 6. Búsqueda Avanzada por EAN (Privado)

```bash
# Búsqueda por EAN incluyendo productos inactivos
curl -X GET "$PRIVATE_URL/products/search?ean=7791234567890&only_active=false" \
  -H "Content-Type: application/json"

# Búsqueda por EAN solo productos activos
curl -X GET "$PRIVATE_URL/products/search?ean=7791234567890&only_active=true" \
  -H "Content-Type: application/json"
```

## 7. Scripts de Testing Automatizado

```bash
#!/bin/bash
# Test básico de APIs

echo "🏥 Testing Health Check..."
curl -s "$PUBLIC_URL/health" | jq .

echo -e "\n🔍 Testing Search by EAN..."
curl -s "$PUBLIC_URL/search?ean=7791234567890" | jq .

echo -e "\n💡 Testing Suggestions..."
curl -s "$PUBLIC_URL/suggestions?business_type=kiosco&limit=5" | jq .

echo -e "\n📋 Testing Product List..."
curl -s "$PRIVATE_URL/products?limit=5" | jq .
```

## 8. Ejemplos para Scraper Scripts

```bash
#!/bin/bash
# Script ejemplo para scraper

# Función para crear producto desde scraping
create_product() {
  local ean="$1"
  local name="$2"
  local price="$3"
  local source="$4"
  local source_url="$5"
  
  curl -X POST "$PRIVATE_URL/products" \
    -H "Content-Type: application/json" \
    -d "{
      \"ean\": \"$ean\",
      \"name\": \"$name\",
      \"price\": $price,
      \"source\": \"$source\",
      \"source_url\": \"$source_url\",
      \"reliability\": 0.8,
      \"metadata\": {
        \"scraped_at\": \"$(date -u +%Y-%m-%dT%H:%M:%SZ)\",
        \"scraper\": \"disco-scraper-v1\"
      }
    }"
}

# Uso del scraper
create_product "7790001234567" "Producto Test" 150.50 "disco" "https://disco.com.ar/producto-test"
```

## 9. Testing de Errores

```bash
# EAN inválido
curl -X GET "$PUBLIC_URL/search?ean=123" \
  -H "Content-Type: application/json"

# EAN no encontrado
curl -X GET "$PUBLIC_URL/search?ean=1234567890123" \
  -H "Content-Type: application/json"

# Producto duplicado (crear el mismo EAN dos veces)
curl -X POST "$PRIVATE_URL/products" \
  -H "Content-Type: application/json" \
  -d '{
    "ean": "7791234567890",
    "name": "Producto duplicado",
    "source": "manual"
  }'
```

## 10. Verificación de Base de Datos

```bash
# Después de crear productos, verifica en la base de datos:
docker exec -i saas-postgres psql -U postgres -d pim_db -c "
  SELECT 
    ean, 
    name, 
    brand, 
    quality_score, 
    source, 
    is_active,
    created_at 
  FROM global_products 
  ORDER BY created_at DESC 
  LIMIT 10;
"
```

---

## Variables de Entorno

Para usar estos ejemplos, puedes configurar las siguientes variables:

```bash
export PIM_SERVICE_URL="http://localhost:8084"
export API_VERSION="v1"
export PUBLIC_BASE="$PIM_SERVICE_URL/api/$API_VERSION/public/global-catalog"
export PRIVATE_BASE="$PIM_SERVICE_URL/api/$API_VERSION/global-catalog"
```

¡Estos ejemplos están listos para ser utilizados en el desarrollo y testing del catálogo global! 