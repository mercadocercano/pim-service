---
adr: ADR-002
status: accepted
skills:
  implement:
    - dev/hexagonal-go
    - dev/prometheus
  verify:
    - dev/go-hex-audit
    - dev/code-reviewer
---
# ADR-002: MetricsRecorder como Dependencia Inyectada

**Estado**: Aceptado  
**Fecha**: 2026-06-09  
**Contexto**: El servicio expone métricas Prometheus. La primera implementación usaba variables globales de Prometheus, lo que acoplaba los casos de uso a la librería de métricas y hacía imposible testearlos sin efectos secundarios en el registro global.

## Decisión

Definimos el puerto `MetricsRecorder` en `go-shared` e inyectamos la implementación (`PrometheusRecorder` o `NoopRecorder`) desde la capa de configuración. Los casos de uso reciben el recorder como parámetro de constructor y llaman a `metrics.Record(MetricEvent{...})`. Las constantes de nombres de métricas viven en `domain/port/metrics.go` de cada bounded context.

## Alternativas consideradas

| Opción | Por qué no |
|--------|-----------|
| Variables globales Prometheus (`prometheus.MustRegister`) | Tests contaminan el registro global entre ejecuciones; imposible usar `NoopRecorder` en pruebas |
| Middleware que intercepta respuestas HTTP | Solo captura métricas de latencia/status — no métricas de negocio (ej. tamaño de CSV importado) |

## Consecuencias

**Positivas**: Tests unitarios usan `NoopRecorder` sin efectos secundarios; la implementación de métricas es intercambiable; las constantes de nombres evitan typos.  
**Negativas / trade-offs**: Cada constructor de use case recibe un parámetro extra; las constantes de métricas deben declararse explícitamente.  
**Neutral**: `PrometheusRecorder` requiere `PROMETHEUS_ENABLED=true` para registrar en el servidor HTTP.
