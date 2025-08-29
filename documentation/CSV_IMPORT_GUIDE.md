# 📊 Guía de Importación CSV para Productos

## 🎯 Descripción General

El sistema de importación CSV permite cargar múltiples productos de forma masiva desde archivos CSV. El sistema procesa cada fila individualmente, permitiendo que algunos productos se importen exitosamente mientras otros pueden fallar sin afectar toda la importación.

## 🚀 Endpoint

```
POST /api/v1/products/import-csv
```

### Headers Requeridos
- `Authorization: Bearer <token>`
- `X-Tenant-ID: <tenant-uuid>`
- `Content-Type: multipart/form-data`

### Parámetros
- `file`: Archivo CSV con los productos (requerido)
- `create_variants`: Boolean para crear variantes por defecto (opcional)

## 📋 Formato del CSV

### Columnas Requeridas
- `name`: Nombre del producto
- `sku`: Código único del producto
- `price`: Precio del producto (número positivo)

### Columnas Opcionales
- `description`: Descripción del producto
- `stock`: Cantidad en inventario
- `category_id`: UUID de la categoría
- `category_name`: Nombre de la categoría
- `brand_id`: UUID de la marca
- `brand_name`: Nombre de la marca

### Columnas Adicionales
Cualquier columna adicional será tratada como atributo personalizado del producto.

## 📄 Ejemplo de CSV

```csv
name,sku,price,description,stock,category_id,category_name,brand_id,brand_name,color,size
"Camiseta Básica","TSH-001",29.99,"Camiseta de algodón",50,"uuid-cat-1","Ropa","uuid-brand-1","Nike","Negro","M"
"Pantalón Jeans","JNS-001",89.99,"Jeans clásico",30,"uuid-cat-1","Ropa","uuid-brand-2","Levi's","Azul","32"
```

## 🔄 Respuesta del Endpoint

### Respuesta Exitosa (200 OK)
```json
{
  "success": true,
  "summary": {
    "total_rows": 10,
    "successful_imports": 8,
    "failed_imports": 2,
    "saved_products": 8,
    "processing_errors": 0
  },
  "imported_products": [
    {
      "id": "uuid",
      "name": "Camiseta Básica",
      "sku": "TSH-001",
      "status": "draft"
    }
  ],
  "errors": [
    {
      "row": 5,
      "data": {
        "name": "Producto Sin SKU",
        "price": "29.99"
      },
      "errors": ["SKU es requerido"]
    }
  ]
}
```

## ⚠️ Validaciones

### Validaciones de Campos
1. **Nombre**: No puede estar vacío
2. **SKU**: No puede estar vacío
3. **Precio**: Debe ser un número positivo
4. **Stock**: Si se proporciona, debe ser un número entero no negativo
5. **IDs**: Si se proporcionan, deben ser UUIDs válidos

### Estados de Productos
- Los productos se crean con estado `draft` por defecto
- Pueden ser actualizados posteriormente a otros estados

## 🛠️ Uso con cURL

```bash
curl -X POST http://localhost:8090/api/v1/products/import-csv \
  -H "Authorization: Bearer <token>" \
  -H "X-Tenant-ID: <tenant-uuid>" \
  -F "file=@productos.csv"
```

## 🔧 Arquitectura del Sistema

### Componentes Principales

1. **FileImporter Interface** (`shared/domain/port/file_importer.go`)
   - Define la interfaz genérica para importar archivos
   - Permite diferentes implementaciones (CSV, Excel, etc.)

2. **BaseCSVFileImporter** (`shared/infrastructure/adapters/csv_file_importer.go`)
   - Implementación base para procesar archivos CSV
   - Maneja la lógica común de lectura y validación

3. **ProductCSVFileImporter** (`product/tenant/infrastructure/adapters/product_csv_file_importer.go`)
   - Implementación específica para productos
   - Define el mapeo de columnas a entidades Product

4. **ImportProductsFromCSVUseCase** (`product/tenant/application/usecase/import_products_from_csv_usecase.go`)
   - Caso de uso que orquesta la importación
   - Intenta guardar cada producto válido individualmente

### Flujo de Procesamiento

```
1. Usuario sube archivo CSV
2. Controller valida el archivo
3. CSVFileImporter parsea cada fila
4. Cada fila se valida y convierte en Product
5. UseCase intenta guardar cada producto
6. Se recopilan éxitos y errores
7. Se devuelve resumen completo
```

## 🔄 Extensibilidad

El sistema está diseñado para ser extensible:

### Agregar Nuevos Formatos
```go
type ExcelFileImporter struct {
    *BaseExcelFileImporter[entity.Product]
}

func (e *ExcelFileImporter) ParseRow(...) (*entity.Product, []string) {
    // Lógica específica para Excel
}
```

### Importar Otras Entidades
```go
type CategoryCSVFileImporter struct {
    *BaseCSVFileImporter[entity.Category]
}

func (c *CategoryCSVFileImporter) ParseRow(...) (*entity.Category, []string) {
    // Lógica específica para categorías
}
```

## 🚨 Consideraciones

1. **Sin Transacciones Globales**: Cada producto se guarda individualmente
2. **Reintento de Errores**: Los productos fallidos pueden ser corregidos y reintentados
3. **Validación de Referencias**: category_id y brand_id deben existir en el sistema
4. **Límite de Archivo**: Considerar límites de tamaño para archivos grandes
5. **Procesamiento Asíncrono**: Para archivos muy grandes, considerar procesamiento en background

## 📊 Mejoras Futuras

1. **Validación de Referencias**: Implementar validación real de categorías y marcas
2. **Importación de Variantes**: Soporte para múltiples variantes por producto
3. **Modo Update**: Actualizar productos existentes por SKU
4. **Plantillas Descargables**: Generar CSV de ejemplo con categorías/marcas válidas
5. **Procesamiento en Lotes**: Para mejorar performance con archivos grandes