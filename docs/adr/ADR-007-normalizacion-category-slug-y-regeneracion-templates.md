---
adr: ADR-007
status: accepted
skills:
  implement:
    - dev/hexagonal-go
    - dev/postgres-data-modeling
    - dev/canonical-logs-go
  verify:
    - dev/go-hex-audit
    - dev/inter-service-contracts
---

# ADR-007: Normalización de `category_slug` en global_products y regeneración del surtido de templates

**Estado**: Aceptado
**Fecha**: 2026-06-18 (sign-off Fase 1a: 2026-06-19)
**Deciders**: @architect (diseño + sign-off), owner (decisión: regenerar desde global_products + backfill como operación S2S)
**Épica**: E31 — Saneamiento de surtido de templates del Quickstart

> **Sign-off @dev-architect (2026-06-19) — acotado a Fase 1a:** GO al resolver `domain/category` en go-shared
> (gemelo de `domain/businesstype`, ADR-005). Código completo, tests verdes, contrato público estable.
> Releaseado en **go-shared v0.12.0** (MINOR, aditivo) y consumido por pim-service. Las fases siguientes
> (backfill S2S, switch del read-path del quickstart, regeneración de templates) quedan dentro del scope
> aprobado de E31 y se ejecutan bajo esa épica.
**Ceremony level**: L3 (catálogo global, no money/auth/PII; sin sub-gate L4)
**Relación**: gemelo estructural de [ADR-005](ADR-005-reclasificacion-business-type-s2s.md) sobre el eje
`category_slug` en lugar de `business_type`; reusa su maquinaria (S2S + snapshot + audit).

## Contexto

El owner, validando templates del Quickstart en mc_admin (`/quickstart/templates`), detectó que el surtido
mostrado no refleja la realidad: la tira "Marcas (12)" no representa el surtido y el template "Almacén"
muestra 305 productos de baja calidad. La auditoría (`scripts/audit/template_assortment_audit.sql`,
read-only) confirmó dos problemas distintos.

Hechos verificados contra la data y el código (no supuestos):

- **El `products` JSONB de los templates es un snapshot SINTÉTICO mal apareado.** En
  `business_type_templates.products` el `brand` está barajado fila por fila respecto al nombre
  (ej. `Fideos coditos Marolio 500g`→marca `"Pastasole"`; `Cerveza Quilmes 1L`→marca `"Upa Lala"`,
  categoría `CAKE_TOPPERS`). No se corrige con un mapa de categorías: el dato de marca es ruido.
- **`global_products` (catálogo real) está bien apareado.** 63.829 filas (4.914 con `is_verified=true`);
  brand/nombre correctos (`Cerveza Quilmes`→`QUILMES`). Su único problema es `category`, en 3 formatos
  sucios mezclados: **2.283 valores distintos = 1.334 paths VTEX** (`/Bebidas/Vinos/Vinos tintos/`),
  **727 en MAYÚSCULAS / constantes EN de MercadoLibre** (`BEBIDAS`, `BEERS`, `SOFT_DRINKS`) y solo
  **165 en kebab-case**.
- **El pipeline computado ya existe y ya corrió.** `RefreshProductTemplates`
  (`src/s2s/infrastructure/persistence/postgres_template_repository.go`, `const refreshQuery`) llena
  `business_type_product_templates.suggested_products` (UUIDs de `global_products`) + `priority_brands`
  (top-10 marcas reales del surtido). Último refresh 2026-06-17: almacen computed=61, congelados=223,
  perfumeria=175.
- **El join del refresh es exact-match** `gp.category = tc.category_slug` (slug kebab declarado), por lo
  que ignora ~2.000 categorías VTEX/upper → el surtido computado queda chico. Además exige
  `gp.is_verified = true`.
- **El read-path de la UI lee el editorial, no el computado.**
  `GET /pim/api/v1/quickstart/templates` (`list_templates_postgres_repository.go`) devuelve marcas desde
  el JSONB `brands` curado y cuenta `total_products` desde el JSONB `products` editorial. El apply-flow
  (`apply_template_postgres_repository.go`) también consume el editorial. Por eso la UI muestra la data mala.
- **Origen del ruido de categorías.** Las constantes EN vienen del adapter de MercadoLibre
  (`webdata-service/src/enrichment/infrastructure/sources/mercadolibre/adapter.go`:
  `MLA-SOFT_DRINKS`→`SOFT_DRINKS`), propagadas sin normalizar en `run_enrichment_batch.go`. Los paths
  VTEX vienen de scrapers VTEX. Ya existe el patrón de resolver de dominio compartido:
  `ResolveBusinessTypeFromProductCategory` (promovido a go-shared en ADR-005 §5).

La tensión estructural: **el surtido real existe y está bien apareado en `global_products`; lo único que
falta es un eje de categoría limpio para que el join del refresh lo alcance, y que la UI lea el computado.**

## Decisión

1. **Columna nueva `global_products.category_slug` (normalizada), preservando `category` raw.** No se
   sobreescribe `category` (auditoría + otros consumidores). El join del refresh pasa a usar
   `gp.category_slug`.

2. **Resolver de categoría determinístico, promovido a go-shared** como gemelo del resolver de
   business_type, + **tabla DB `category_slug_overrides`** para el long tail curable. Precedencia:
   `override (DB) > resolver determinístico > 'sin-clasificar'`. **Sin fuzzy matching.**

3. **La normalización vive en go-shared y se invoca en la ingestión de webdata.** El adapter ML sigue
   mandando `category` raw; `category_slug` es un derivado mecánico que **se recalcula siempre** (no
   aplica la safe-transition §8 de business_type, porque no hay curación manual que proteger).

4. **El backfill sobre las 63k filas es una SEGUNDA OPERACIÓN del endpoint S2S de ADR-005** (decisión
   del owner), reusando snapshot (`global_products_bkp_<ts>`) + audit
   (`global_product_reclassification_audit`) + `dry_run` + cap. No se crea maquinaria nueva.

5. **El read-path del Quickstart migra de editorial → computed con feature-flag + fallback al editorial.**
   Preview (list endpoint) primero; apply-flow (onboarding) después, una vez validado.

6. **Regeneración**: tras el backfill, correr `RefreshProductTemplates` (join ya apuntando a
   `category_slug`) para repoblar `suggested_products` + `priority_brands` con el surtido completo.

## Detalle de las decisiones

### 1. Columna `category_slug` (no sobreescribir `category`)

`ALTER TABLE global_products ADD COLUMN category_slug VARCHAR(200)` + índice parcial
`WHERE category_slug IS NOT NULL`. Se preserva `category` raw como rastro de la fuente (VTEX path / ML
constant), permitiendo re-derivar si el mapeo cambia, auditar y debuggear. El refresh
(`refreshQuery`) cambia `ON gp.category = tc.category_slug` → `ON gp.category_slug = tc.category_slug`.
Otros consumidores de `category` (búsquedas, facets) no se rompen porque la columna raw sigue existiendo;
se migran a `category_slug` incrementalmente si conviene (fuera de scope inmediato).

Alternativa descartada: sobreescribir `category`. Rechazada — pierde el dato de origen (irreversible),
y mezcla "valor crudo de la fuente" con "valor canónico del ecosistema" en una sola columna.

### 2. Resolver determinístico + tabla de overrides (sin fuzzy)

- **Resolver en go-shared** (`libs/go-shared/domain/category/` — gemelo de
  `domain/businesstype/resolver.go`): reglas determinísticas que toman el **leaf-segment** del path VTEX
  (`/Bebidas/Vinos/Vinos tintos/` → `Vinos tintos` → `vinos-tintos`) y un diccionario de constantes
  EN→slug (`BEERS`→`cervezas-vinos`, `SOFT_DRINKS`→`gaseosas-aguas`, `MILK`→`lacteos`). Orden de reglas
  load-bearing como en el resolver de business_type (guards de colisión). Con su test.
- **Tabla `category_slug_overrides`** en `pim_db` (`raw_category TEXT PRIMARY KEY, category_slug
  VARCHAR(200), note TEXT`): para el long tail de paths que el resolver no cubre limpio (cientos con
  1-2 productos). Seedeable por migración y editable sin redeploy.
- **Precedencia**: override (DB) > resolver (código) > `'sin-clasificar'` (catch-all explícito, nunca
  NULL silencioso). `sin-clasificar` no entra a ningún template (no hay categoría declarada con ese
  slug) → queda visible y medible como deuda de curación.
- **Sin fuzzy**: el matching aproximado introduce no-determinismo y falsos positivos sobre 63k filas;
  el long tail se cubre con overrides explícitos.

### 3. Anti-recaída: normalización en ingestión

La función de normalización vive en go-shared y se invoca en webdata
(`run_enrichment_batch.go` y/o el upsert de `global_products`) para poblar `category_slug` en cada
ingreso/actualización. El adapter ML no cambia (sigue emitiendo `category` raw). Como `category_slug`
es derivado mecánico, **se recalcula siempre** — no hay equivalente a la guarda §8 de business_type
(que protege curación manual); acá no hay nada manual que proteger en la columna derivada (los
overrides viven en su propia tabla, no en la fila).

### 4. Backfill como segunda operación S2S (reusa ADR-005)

```
POST /api/v1/s2s/global-products/normalize-category-slugs
Auth: API-Key (Kong), sin JWT, sin X-Tenant-ID
Header: X-Operator-Id: <user_id del admin>

Request:  { "dry_run": true, "scope": { "max_rows": 70000 }, "confirm": false }
Response: { "mode", "snapshot_ref", "summary": { total_evaluados, resueltos_por_resolver,
            resueltos_por_override, sin_clasificar, sin_cambio }, "antes_despues": [...] }
```

Mismas invariantes que ADR-005 §8: `dry_run` no muta y predice el `summary` del apply; snapshot
(`global_products_bkp_<ts>`) creado en la misma tx ANTES de cualquier UPDATE (si falla, aborta todo);
idempotencia (segundo apply → 0 cambios); audit row inmutable en
`global_product_reclassification_audit` (se reusa la tabla; `scope`/`summary` JSONB distinguen la
operación). Cap de tamaño y ejecución síncrona con tx única, como ADR-005 §6.

### 5. Read-path editorial → computed (feature-flag + fallback)

- **Preview** (`GET /pim/api/v1/quickstart/templates`): resolver marcas desde `priority_brands` y conteo
  de productos desde `suggested_products` (join a `global_products` para nombre/marca/imagen) cuando el
  computed existe; **fallback al editorial** si `business_type_product_templates` está vacío para ese
  template. Detrás de feature-flag para rollback instantáneo.
- **Apply-flow** (onboarding de tenant, `apply_template_postgres_repository.go`): migrar a sembrar el
  catálogo del tenant desde `suggested_products` (productos reales). Se hace **después** de validar el
  preview, por ser el flujo de mayor impacto (lo que el comercio ve al arrancar). Mismo flag/fallback.

Riesgo del apply-flow: sembrar un tenant con surtido distinto al actual. Mitigación: flag + fallback +
validación del preview antes de activar; el editorial sigue disponible como red.

### 6. Invariantes a preservar (verificables en tests)

- `category_slug` nunca queda NULL tras el backfill: o resuelve, o cae en `'sin-clasificar'`.
- `dry_run=true` no muta y devuelve el mismo `summary` que el apply.
- Idempotencia: segundo backfill consecutivo → 0 updates.
- El resolver es determinístico: misma `category` raw → mismo `category_slug` siempre.
- Si el snapshot falla → no se aplica nada.
- El read-path con flag OFF se comporta idéntico al actual (editorial), byte por byte.

## Alternativas consideradas

- **Patchear el JSONB editorial (mapa EN→ES sobre 305 filas + arreglo manual de brands)**: rechazada por
  el owner. El `brand` está barajado fila por fila (no hay mapa que lo arregle) y deja un surtido
  sintético chico en vez del catálogo real de miles de productos.
- **Sobreescribir `global_products.category`**: rechazada (§1) — pierde el dato de origen, irreversible.
- **Mapeo fuzzy de categorías**: rechazada (§2) — no-determinístico sobre 63k filas; long tail vía overrides.
- **Backfill como script `//go:build ignore`** (camino E17): descartada por el owner a favor de la
  operación S2S gateada/auditada (reusa snapshot+audit; repetible desde mc_admin).
- **Read-path directo a computed sin fallback ni flag**: rechazada — el apply-flow es onboarding; un
  cambio sin red de seguridad es riesgoso. Flag + fallback al editorial.

## Consecuencias

- ✅ El surtido de los templates pasa a reflejar el catálogo real (miles de productos bien apareados) en
  vez del JSONB sintético; marcas reales caen del `priority_brands` ya computado.
- ✅ Reusa íntegramente la maquinaria de ADR-005 (S2S + snapshot + audit + dry_run) y el resolver
  compartido de go-shared — sin inventar mecanismos nuevos.
- ✅ `category_slug` limpio destraba el join del refresh y habilita futuras facets/búsquedas canónicas.
- ⚠️ El surtido computado está capado por `is_verified=true` (4.914 de 63.829). Aun normalizado, el
  surtido crece pero no usa todo el catálogo. Curación de `is_verified` = follow-up separado.
- ⚠️ `category_slug` agrega un derivado a mantener: cambios al resolver/overrides requieren re-backfill.
- ⚠️ `'sin-clasificar'` hace visible la deuda de curación del long tail (deseable, pero requiere un
  proceso de revisión de overrides).
- 🚫 Fuera de scope: curación de `is_verified`; migrar todos los consumidores de `category`→`category_slug`;
  retención/limpieza de snapshots viejos (heredado de ADR-005).

## Por qué L3 sin sub-gate L4

A diferencia de ADR-005 (que tiene sub-gate L4 por el guard del endpoint + captura de operador), ADR-007
**reusa el endpoint S2S ya gateado y auditado** — no abre superficie nueva: agrega una operación sobre la
misma ruta key-auth, mismo snapshot, mismo audit. El dato es el catálogo GLOBAL (no tenant, no money, no
auth, no PII). El único flujo sensible es el apply-flow de onboarding, mitigado con feature-flag +
fallback (§5). Por eso: **L3, sin sub-gate L4**. El guard/audit del endpoint ya fue validado en ADR-005.

## Revisión prevista

Revisar si: (a) `'sin-clasificar'` queda con volumen alto → priorizar curación de overlays/overrides o
extender el resolver; (b) se decide curar `is_verified` para ampliar el surtido elegible; (c) aparece un
tercer eje de normalización (marca) que justifique generalizar el patrón resolver+override+backfill.
