---
adr: ADR-004
status: accepted
skills:
  implement:
    - dev/hexagonal-go
    - go-shared-migrate
  verify:
    - dev/go-hex-audit
    - dev/code-reviewer
---
# ADR-004: Extracción de Código Compartido a go-shared

**Estado**: Aceptado  
**Fecha**: 2026-06-09  
**Contexto**: Varios tipos de infraestructura (Executor SQL, FileImporter, MetricsRecorder, NotificationGateway, middleware gzip) estaban duplicados entre `iam-service` y `pim-service`. El workspace Go (`go.work`) facilita compartir código local sin publicar una librería externa.

## Decisión

Extraemos los tipos y adaptadores reutilizables a `github.com/mercadocercano/go-shared` (módulo en `libs/go-shared` del workspace). El código que depende de entidades de negocio propias del servicio (ej. `ImportJob`, `NotificationService`) permanece local en `src/shared/`.

## Alternativas consideradas

| Opción | Por qué no |
|--------|-----------|
| Copiar código entre servicios (copy-paste) | Divergencia inevitable; bugs corregidos en un servicio no se propagan al otro |
| Módulo Go publicado en GitHub (tag+versión) | Overhead de release para código en desarrollo activo; el workspace ya resuelve el problema |

## Consecuencias

**Positivas**: Cambios en infraestructura compartida se reflejan en ambos servicios desde el mismo commit; `NoopRecorder` y `NoopNotificationGateway` disponibles para tests en cualquier servicio.  
**Negativas / trade-offs**: Un cambio breaking en `go-shared` requiere actualizar todos los servicios que lo importan antes de compilar.  
**Neutral**: El bump de versión de `go-shared` se gestiona con `go get` dentro del workspace.
