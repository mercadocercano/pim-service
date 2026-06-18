---
adr: ADR-005
status: proposed
skills:
  implement:
    - dev/hexagonal-go
    - dev/postgres-data-modeling
    - dev/canonical-logs-go
  verify:
    - dev/go-hex-audit
    - dev/owasp-top10
---

# ADR-005: Re-clasificación de business_type como endpoint S2S en pim-service

**Estado**: Propuesto
**Fecha**: 2026-06-17
**Deciders**: @architect, @security (sub-gate L4 sobre guard + audit), owner (D-E24-1, §5)
**Épica**: E24 — Re-clasificación de business_type disparable desde mc_admin
**Ceremony level**: L3 con sub-gate L4 (protección del endpoint + audit)

## Contexto

En E17 la re-clasificación masiva de `business_type` del catálogo global se resolvió con un
script Go one-off (`services/webdata-service/scripts/reclassify_business_types.go`,
`//go:build ignore`) corrido vía `docker exec`. El owner decidió (memoria #149) que estas
operaciones de mantenimiento masivo deben volverse una acción de primera clase, disparable
desde mc_admin, sin entrar a la consola.

Hechos verificados contra el código (no supuestos):

- **El dato vive en pim-service.** `global_products` es una tabla de `pim_db` con módulo
  hexagonal completo en `services/pim-service/src/product/global_catalog/` (entity, repo,
  use cases). El entity `GlobalProduct` ya expone `SetBusinessType(string)`
  (`domain/entity/global_product.go:241`) y `businessType` es `*string` (puede ser vacío/null).
- **Las reglas de clasificación viven en webdata-service.** El value object
  `ResolveBusinessTypeFromProductCategory`
  (`services/webdata-service/src/product/domain/value_object/product_category_resolver.go`)
  es la ÚNICA fuente de verdad de la taxonomía categoría→business_type, con su test
  (`product_category_resolver_test.go`). El orden de las reglas es load-bearing (guards
  de colisión: vinagre antes que vino, conserva antes que carne, etc.).
- **El patrón S2S ya existe en pim.** `main.go:88` excluye `/api/v1/s2s*` de `TenantValidation`;
  el `InternalHandler` (`src/s2s/controller/internal_handler.go`) sirve el grupo `/s2s`
  autenticado por API-Key en Kong, sin JWT ni X-Tenant-ID, con logger canónico inyectado.
- **El logger canónico ADR-001 ya está cableado** (`main.go:150`, `pim_logger.go`). El
  `PIMEvent` (`src/pim/domain/port/pim_event_logger.go`) tiene `UserID` y `Reason`, pero es
  un evento de observabilidad flat (sin antes/después, sin retención garantizada).
- **El constraint existe.** `UNIQUE(name, business_type)` en `global_products`
  (`migrations/202504240001_...sql`, constraint `uq_global_products_name_business_type`).

**Decisión de auth ya tomada por el owner (D-E24-1, NO se re-abre)**: el endpoint se autentica
vía S2S + API-Key en Kong (Opción A), no por JWT de tenant, porque `global_products` es el
catálogo GLOBAL del PIM (no tenant-scoped) y la operación es cross-tenant/global. Reusa el
patrón S2S existente; ruta no expuesta públicamente en Kong; audit log obligatorio que capture
el operador humano; sin cambios en el middleware de auth de go-shared; no replicar el
anti-patrón "ruta excluida sin control".

La tensión estructural central de este ADR: **el dato está en pim, las reglas están en webdata.**

## Decisión

1. **Ownership → pim-service.** El endpoint/use case vive en pim-service, dueño de
   `global_products`/`business_types`. Es quien tiene el repo, el constraint, la transacción y el
   logger canónico sobre `pim_db`. webdata-service no es dueño del dato (el script conectaba a
   `pim_db` desde afuera — un cross-DB write que no debe perpetuarse como producto).

2. **Las reglas se promueven a go-shared como única fuente de verdad y NO se reimplementan.**
   `ResolveBusinessTypeFromProductCategory` + su tabla de reglas + test se mueven a
   `libs/go-shared/domain/businesstype/` (decisión §5, ratificada por el owner). Tanto webdata
   como pim lo consumen; el use case de pim invoca el resolver, nunca reimplementa la taxonomía.

3. **El endpoint es S2S**, montado en el grupo `/s2s` existente del `InternalHandler`, detrás
   de API-Key en Kong, fuera de `TenantValidation`. No se crea una "ruta excluida" nueva ad-hoc.

4. **dry_run por defecto true.** El apply solo procede si el snapshot de rollback se creó OK.

5. **Audit log persistente** en tabla propia de `pim_db` (`global_product_reclassification_audit`),
   separado del canonical log ADR-001 (que se usa en paralelo para observabilidad operacional).

6. **Ejecución síncrona con transacción única y cap de tamaño** para el alcance actual; job async
   queda fuera de scope de E24 (ver Consecuencias 🚫).

## Detalle de las decisiones que cierra este ADR

### 1. Ownership / ubicación — pim-service

Fundamento contra el código: `global_products` es tabla de `pim_db`, con módulo
`src/product/global_catalog/` hexagonal y `GlobalProduct.SetBusinessType()`. El script one-off
vivía en webdata pero **conectaba a `pim_db`** (`DB_NAME=pim_db`, ver cabecera del script): un
write cross-servicio que rompe el bounded context. Llevarlo a pim alinea dato + escritura +
transacción + audit en un solo dueño. webdata mantiene la *autoría de las reglas* (vía el resolver
compartido), no la *operación sobre el dato*.

### 2. Contrato del endpoint

```
POST /api/v1/s2s/global-products/reclassify-business-types
Auth: API-Key (Kong), sin JWT, sin X-Tenant-ID
Header: X-Operator-Id: <user_id del admin>   (propagado por mc_admin/BFF; ver §4)

Request:
{
  "dry_run": true,                  // default true si se omite
  "scope": {                        // criterio explícito, no implícito
    "source_prefix": "scraper",     // replica el WHERE source LIKE 'scraper%' del script
    "max_rows": 50000               // cap de seguridad (ver §6)
  },
  "confirm": false                  // apply real requiere dry_run=false AND confirm=true
}

Response 200:
{
  "mode": "dry_run" | "applied",
  "snapshot_ref": "global_products_bkp_20260617_..." | null,  // null sólo en dry_run
  "summary": {
    "total_evaluados": 12345,
    "candidatos": 1200,
    "updates_por_rubro": { "fiambreria": {"relleno": 800, "correccion": 50}, ... },
    "colisiones_skipeadas": 12,
    "skips": { "no_resuelve": 300, "ya_especifico": 900, "ya_correcto": 200 }
  },
  "antes_despues": [ {"id","name","from","to","kind"} ... ]   // muestra acotada
}
```

**Ruta por Kong**: mc_admin/backend → Kong (route S2S con plugin key-auth, **no** expuesta en
el listener público) → `mc-pim-service:8090`. mc_admin nunca habla directo a pim; el secreto
de API-Key vive server-side en el BFF/backend de mc_admin, nunca en el browser.

### 3. Snapshot de rollback

**Decisión: tabla auxiliar `CREATE TABLE global_products_bkp_<ts> AS SELECT ...`** del subconjunto
afectado (o de la tabla completa filtrada por `scope`), creada DENTRO de la misma transacción del
apply, ANTES de cualquier UPDATE. Si el snapshot falla, la transacción aborta y no se aplica nada
(invariante dura). El `snapshot_ref` se devuelve en la response y se registra en el audit log.
Equivale al `global_products_bkp_*` que el owner ya usaba manualmente, pero ahora es automático y
parte de la operación. Retención/limpieza de snapshots viejos: fuera de scope de E24 (anotar como
deuda operativa).

Alternativa descartada: confiar solo en el audit row con before/after como mecanismo de rollback —
rechazada porque reconstruir miles de UPDATEs inversos desde el audit es frágil; el snapshot
tabular permite un rollback masivo trivial.

### 4. Audit log — captura del operador con auth por API-Key

Problema: la API-Key identifica a mc_admin como sistema, no al admin humano. **Decisión**:
mc_admin/BFF propaga la identidad del operador en un header `X-Operator-Id` (validado server-side
contra su propia sesión admin antes de llamar a pim). El use case persiste un row en
`global_product_reclassification_audit` (`pim_db`) con: `operator_id`, `timestamp`, `mode`
(dry_run/applied), `scope` (JSONB), `snapshot_ref`, `summary` (JSONB con resumen estructurado),
`affected_count`. El detalle before/after por fila vive en el snapshot tabular (§3), no duplicado
en el audit.

**Encaje con ADR-001**: el canonical logger (`PIMEvent`) se usa EN PARALELO para observabilidad
(eventos `catalog.reclassification_completed` / `_failed`, con `user_id`=operator, `count`,
`reason`), pero NO reemplaza el audit persistente. Distinción explícita:
- **Canonical log (ADR-001)** = observabilidad, va a Loki, retención de logs, baja garantía.
- **Audit table** = rastro forense con retención propia, queryable, sobrevive rotación de logs.

Esto NO cambia el `PIMEvent` ni el middleware de go-shared. Usa `UserID`+`Reason`+`Count` que ya existen.

### 5. Reuso de la lógica — resolver promovido a go-shared (ratificado por el owner)

`ResolveBusinessTypeFromProductCategory` vive HOY en
`webdata-service/src/product/domain/value_object/`. pim no puede importar el `internal`/módulo de
otro servicio sin acoplarse a webdata.

**Decisión (ratificada por el owner, 2026-06-17)**: promover el resolver + su tabla de reglas +
test a **`libs/go-shared/domain/businesstype/resolver.go`**, como value object compartido del
ecosistema, y que TANTO webdata como pim lo consuman. Es la única forma de mantener
"single source of truth" sin duplicar reglas entre servicios.

Nota sobre D-E24-1: la condición "sin cambios en go-shared" de D-E24-1 se refiere al **middleware
de auth** de go-shared (que efectivamente no se toca). Agregar un value object de dominio nuevo es
ortogonal a esa condición y es la decisión limpia — explícitamente avalada por el owner.

Alternativa descartada (fallback): copiar el resolver a pim como dueño único y retirarlo de
webdata — inferior porque deja a webdata sin la clasificación inline en su upsert o fuerza
duplicación de reglas.

El use case (`ReclassifyBusinessTypesUseCase` en `src/product/global_catalog/application/usecase/`)
orquesta: query → resolver compartido (sin reimplementar) → detección de colisión
`UNIQUE(name,business_type)` → snapshot → UPDATE → audit. La taxonomía nunca se reescribe en pim.

### 6. Modo de ejecución — síncrono con transacción y cap

**Decisión: síncrono, transacción única, con `max_rows` cap.** El script actual evalúa
`source LIKE 'scraper%'` (orden de miles, no millones). Para ese volumen, una transacción única con
snapshot + UPDATEs batcheados es atómica y simple. El cap (`scope.max_rows`, default 50k) protege
contra una operación desbocada; si se excede, el endpoint responde 422 pidiendo acotar el scope.
Atomicidad: o se aplica todo o nada (snapshot + updates en la misma tx). Idempotencia: correr dos
veces seguidas el apply produce 0 cambios en la segunda (los ya-correctos caen en skip `ya_correcto`).

🚫 **Job async con estado queda fuera de scope de E24** — se adopta solo si el volumen crece a un
punto donde la tx única bloquee la tabla demasiado tiempo. Anotar como revisión futura.

### 7. Destino del script one-off

**Decisión: deprecar al cerrar E24.** El script `reclassify_business_types.go` fue el paso de
validación (E17); una vez que el endpoint S2S replica su lógica y se valida que el resultado coincide
(criterio manual de la épica), el script se retira. Hasta el merge+validación del endpoint, coexisten.
El value object NO se borra (se promueve a go-shared, §5); solo se retira el `main()` del script.

### 8. Invariantes a preservar (verificables en tests)

- `dry_run=true` NO muta y devuelve el MISMO `summary` que produciría el apply.
- Idempotencia: segundo apply consecutivo → 0 updates.
- Ningún rubro específico pierde productos: solo se rellena vacío o se mueve DESDE `almacen` a
  rubro específico; nunca se toca un producto ya en rubro específico (`ya_especifico` → skip).
- Colisión `UNIQUE(name, business_type)` → skip, nunca error que aborte el lote.
- Si el snapshot falla → no se aplica nada.

## Alternativas consideradas

- **Ownership en webdata-service** (donde está el script y las reglas): rechazada. Perpetúa el
  write cross-DB a `pim_db` desde un servicio que no es dueño del dato; rompe el bounded context y
  deja el audit/transacción del catálogo fuera de su dueño natural.
- **Auth por JWT de tenant (Opción B de D-E24-1)**: ya descartada por el owner — `global_products`
  es global, no tenant-scoped. No se re-abre.
- **Audit solo vía canonical log ADR-001**: rechazada. Loki no es un rastro forense con retención
  garantizada; mezclar auditoría de cambios masivos con logs operacionales degrada ambos.
- **Snapshot vía audit row con before/after como mecanismo de rollback**: rechazada a favor de la
  tabla `_bkp_` (rollback masivo trivial).
- **Job async con estado desde el día 1**: rechazada por YAGNI al volumen actual; reservada como
  evolución si la tx única se vuelve problemática.
- **Copiar el resolver a pim (sin go-shared)**: fallback descartado a favor de promover a go-shared
  (§5, decisión del owner); inferior porque deja a webdata sin la clasificación inline o fuerza duplicación.

## Consecuencias

- ✅ Dato + escritura + transacción + snapshot + audit quedan en un solo dueño (pim), eliminando
  el write cross-DB del script.
- ✅ Reusa el patrón S2S existente y el logger canónico ADR-001 — sin inventar mecanismos nuevos.
- ✅ dry_run + snapshot + audit + cap dan una operación masiva segura y reversible, disparable
  desde mc_admin sin consola.
- ✅ Single source of truth de la taxonomía preservada (resolver compartido en go-shared), sin duplicar reglas.
- ⚠️ Promover el resolver a go-shared crea una dependencia nueva del ecosistema sobre ese value
  object: cambios futuros a la taxonomía impactan webdata + pim (deseable, pero requiere recompilar ambos).
- ⚠️ El audit depende de que mc_admin propague `X-Operator-Id` honestamente; la confianza recae en
  la frontera S2S (Kong + backend de mc_admin). Sub-gate L4 debe validar esto.
- 🚫 Fuera de scope E24: job async con estado; retención/limpieza de snapshots viejos; la
  productización de template-base + curación de `is_verified` (se difiere a E23/follow-up).

## Reparto L3 / sub-gate L4

**L3 (diseño + implementación estándar, @architect + @technical-leader):**
- Ownership en pim (§1), contrato del endpoint (§2), modo síncrono/tx/cap (§6), promoción del
  resolver a go-shared (§5), deprecación del script (§7), invariantes y sus tests (§8), snapshot tabular (§3).

**Sub-gate L4 (sign-off obligatorio de @security antes de merge):**
- **Guard del endpoint**: que `/s2s/global-products/reclassify-business-types` esté efectivamente
  fuera del listener público de Kong, con key-auth obligatorio, y que NO replique el anti-patrón
  "ruta excluida sin control" (verificar que la exclusión de `TenantValidation` esté acompañada de
  la API-Key en Kong, no abierta).
- **Audit + captura del operador**: validar que `X-Operator-Id` se valide server-side en mc_admin,
  que el audit row sea inmutable/no-borrable por el flujo, y que un apply sin operador identificable
  sea rechazado (no se permite cambio masivo anónimo).
- **Threat modeling** del disparo masivo: abuso de la API-Key, replay, falta de rate-limit, y el
  riesgo de un apply destructivo sin snapshot. Verificación con `code-reviewer` D6 (Security) y
  `owasp-top10` en modo Gate.

## Revisión prevista

Revisar si: (a) el volumen de `global_products` crece al punto que la tx única bloquee la tabla
demasiado → migrar a job async (§6); (b) aparece un segundo consumidor del resolver que justifique
o invalide la decisión de packaging (§5); (c) se productiza la edición de template-base, que puede
querer compartir el mismo endpoint/patrón S2S.
