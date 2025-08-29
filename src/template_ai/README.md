# 🤖 AI Template Module

Módulo de templates inteligentes con AI para el servicio PIM. Permite generar catálogos de productos optimizados usando inteligencia artificial basándose en el tipo de negocio, región y preferencias específicas.

## 📋 Características

- **Generación Inteligente**: Crea templates optimizados usando AI basándose en:
  - Tipo de negocio (almacén, supermercado, farmacia, etc.)
  - Tamaño objetivo (pequeño, mediano, grande)
  - Preferencias regionales y locales
  - Restricciones de presupuesto

- **Aplicación Dinámica**: Aplica templates al catálogo del tenant con:
  - Customización de precios y cantidades
  - Exclusión de productos específicos
  - Creación automática de categorías y marcas

- **Análisis de Performance**: Monitorea y analiza:
  - Métricas de uso y adopción
  - Modificaciones comunes
  - Satisfacción del usuario
  - ROI y crecimiento

- **Aprendizaje Continuo**: Mejora templates basándose en:
  - Feedback de usuarios
  - Datos de ventas
  - Tendencias regionales

## 🏗️ Arquitectura

```
template_ai/
├── domain/                    # Lógica de negocio
│   ├── entity/               # Entidades del dominio
│   ├── value_object/         # Objetos de valor
│   ├── port/                 # Interfaces (puertos)
│   ├── service/              # Servicios de dominio
│   └── exception/            # Excepciones personalizadas
│
├── application/              # Casos de uso
│   ├── usecase/             # Implementación de casos de uso
│   ├── request/             # DTOs de entrada
│   ├── response/            # DTOs de salida
│   └── mapper/              # Mappers entre capas
│
└── infrastructure/           # Implementaciones técnicas
    ├── controller/          # Controladores HTTP
    ├── persistence/         # Repositorios PostgreSQL
    ├── client/              # Cliente AI Gateway
    ├── service/             # Template Engine
    └── config/              # Configuración y DI
```

## 🚀 Setup Rápido

### 1. Ejecutar Migración

```bash
cd scripts
./run_migration_034.sh
```

### 2. Configurar Variables de Entorno

```bash
# .env
AI_GATEWAY_URL=http://ai-gateway:8000
AI_GATEWAY_API_KEY=your-api-key
```

### 3. Ejecutar Setup End-to-End

```bash
cd scripts
./setup-ai-templates-e2e.sh
```

## 📡 API Endpoints

### Generar Template Inteligente

```bash
POST /api/v1/templates/generate
```

**Request:**
```json
{
  "business_type_id": "4f4e9b9e-7b8a-4c6a-9c5a-3e5f7a8b9c1d",
  "name": "Mi Almacén Premium",
  "target_size": "medium",
  "preferences": {
    "price_range": "standard",
    "include_generics": true,
    "generic_percentage": 25,
    "categories_focus": ["bebidas", "snacks", "limpieza"],
    "exclude_brands": ["marca_x"],
    "regional_preferences": "buenos_aires"
  }
}
```

**Response:**
```json
{
  "template_id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "Mi Almacén Premium",
  "products": [...],
  "summary": {
    "total_products": 75,
    "estimated_investment": 450000
  }
}
```

### Aplicar Template

```bash
POST /api/v1/templates/{id}/apply
```

**Request:**
```json
{
  "customizations": {
    "price_multiplier": 1.1,
    "exclude_products": ["prod_123"]
  },
  "apply_options": {
    "create_categories": true,
    "create_products": true
  }
}
```

### Obtener Performance

```bash
GET /api/v1/templates/{id}/performance
```

**Response:**
```json
{
  "template_id": "550e8400-e29b-41d4-a716-446655440000",
  "performance_metrics": {
    "usage_count": 45,
    "average_satisfaction": 4.2,
    "modification_rate": 0.15
  }
}
```

## 🧪 Testing

### Tests de Integración

```bash
cd test-integration
./test-ai-templates.sh
```

### Tests Unitarios

```bash
go test ./src/template_ai/...
```

## 📊 Métricas y Monitoreo

Las métricas están disponibles en Grafana:

- **Template Generation Rate**: Cantidad de templates generados
- **Application Success Rate**: Tasa de éxito en aplicación
- **User Satisfaction Score**: Puntuación de satisfacción
- **Performance Metrics**: Métricas de rendimiento por template

## 🔧 Configuración Avanzada

### Reglas de Negocio

Editar `domain/service/template_engine.go` para personalizar:

- Mix de productos por defecto (esencial/recomendado/opcional)
- Porcentaje máximo de genéricos
- Reglas de selección por categoría
- Adaptaciones regionales

### Integración AI Gateway

El cliente AI Gateway puede ser configurado para diferentes modelos:

```go
// infrastructure/config/wire.go
aiGatewayURL: getEnvOrDefault("AI_GATEWAY_URL", "http://ai-gateway:8000"),
aiAPIKey:     getEnvOrDefault("AI_GATEWAY_API_KEY", ""),
```

## 🛠️ Troubleshooting

### Error: "AI Gateway not available"

El sistema funciona en modo mock si AI Gateway no está disponible. Para habilitar AI real:

1. Verificar que AI Gateway esté corriendo
2. Configurar `AI_GATEWAY_URL` correctamente
3. Proveer `AI_GATEWAY_API_KEY` válida

### Error: "No products in global catalog"

Los templates requieren productos en el catálogo global:

1. Cargar productos usando importación CSV
2. Ejecutar scraper para obtener productos
3. Usar seeds de prueba

### Performance Issues

Si la generación es lenta:

1. Verificar índices en `template_global_products`
2. Optimizar queries en repositorios
3. Implementar cache para templates populares

## 📚 Documentación Adicional

- [Arquitectura Hexagonal](../../documentation/ARCHITECTURE_FINAL.md)
- [Catálogo Global](../../documentation/GLOBAL_CATALOG_DESIGN.md)
- [Sistema de Templates](../../documentation/QUICKSTART_INTEGRATION_GUIDE.md)

## 🤝 Contribuir

1. Seguir arquitectura hexagonal estricta
2. Agregar tests para nuevas funcionalidades
3. Documentar cambios en OpenAPI
4. Actualizar README si es necesario