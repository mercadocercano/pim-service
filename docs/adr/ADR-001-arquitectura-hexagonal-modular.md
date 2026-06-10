# ADR-001: Arquitectura Hexagonal Modular

**Estado**: Aceptado  
**Fecha**: 2026-02-01  
**Contexto**: El PIM Service gestiona múltiples dominios (productos, categorías, marcas, atributos, quickstart) dentro de un SaaS multi-tenant. Se necesitaba una estructura que aislara la lógica de negocio de los detalles de infraestructura (PostgreSQL, Gin, CSV importers) y que permitiera testear cada dominio de forma independiente.

## Decisión

Adoptamos arquitectura hexagonal con capas `domain / application / infrastructure` replicadas por módulo. Cada módulo vive en `src/<módulo>/` y es autónomo: define sus propios puertos (interfaces), casos de uso y adaptadores. El módulo `main.go` hace el wiring de todos los módulos.

## Alternativas consideradas

| Opción | Por qué no |
|--------|-----------|
| MVC en capas planas | Acoplamiento entre lógica de negocio e infraestructura — dificulta tests unitarios y cambios de BD |
| Microservicios por dominio | Overhead operativo desproporcionado para la escala actual del proyecto |

## Consecuencias

**Positivas**: Lógica de negocio testeable sin infraestructura; los casos de uso son agnósticos al framework HTTP y a la BD; incorporar un nuevo módulo sigue un patrón predecible.  
**Negativas / trade-offs**: Mayor boilerplate por módulo (domain/port, mapper, config); la curva de incorporación es más pronunciada.  
**Neutral**: La regla de dependencias (domain ← application ← infrastructure) se aplica en code review, no en tiempo de compilación.
