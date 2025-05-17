# PIM (Product Information Management)

Sistema de gestión de información de productos, desarrollado con Go y siguiendo arquitectura hexagonal y Domain-Driven Design (DDD).

## Estructura del Proyecto

El proyecto sigue una arquitectura hexagonal (puertos y adaptadores) con principios de DDD:

```
src/
├── category/           # Módulo de categorías
│   ├── application/    # Capa de aplicación (casos de uso, DTOs)
│   │   ├── mapper/
│   │   ├── request/
│   │   ├── response/
│   │   └── usecase/
│   ├── domain/         # Capa de dominio (entidades, reglas de negocio)
│   │   ├── entity/
│   │   ├── exception/
│   │   ├── port/
│   │   └── value_object/
│   └── infrastructure/ # Capa de infraestructura (adaptadores)
│       ├── config/
│       ├── controller/
│       ├── event/
│       └── persistence/
├── product/            # Módulo de productos
│   ├── application/
│   ├── domain/
│   └── infrastructure/
└── stock_location/     # Módulo de ubicaciones de stock
    ├── application/
    ├── domain/
    └── infrastructure/
```

## Requisitos

- Go 1.21 o superior
- Gin Web Framework

## Instalación

1. Clonar el repositorio
2. Instalar dependencias:

```bash
go mod tidy
```

## Ejecutar

```bash
go run main.go
```

El servidor API estará accesible en http://localhost:8080

## Endpoints API

### Health Check
- GET /health

### API v1 Categorías
- GET /api/v1/categories
- GET /api/v1/categories/:id
- POST /api/v1/categories

## Desarrollo

Este proyecto sigue los principios de arquitectura hexagonal:

- **Dominio**: Contiene las entidades y reglas de negocio.
- **Puertos**: Interfaces que definen cómo se comunica el dominio con el exterior.
- **Adaptadores**: Implementaciones de los puertos para tecnologías específicas.

Cada módulo (categorías, productos, ubicaciones) tiene su propia estructura hexagonal completa, permitiendo evolucionar de forma independiente. 