---
adr: ADR-006
status: accepted
skills:
  implement:
    - dev/hexagonal-go
    - dev/postgres-data-modeling
    - dev/inter-service-contracts
    - dev/canonical-logs-go
  verify:
    - dev/go-hex-audit
    - dev/inter-service-contracts
    - dev/concurrency-transactions
---

# ADR-006: Propagación de business_type en el re-sync webdata→PIM (updateExisting)

**Estado**: Propuesto
**Fecha**: 2026-06-18
**Deciders**: @technical-leader, owner (política de propagación)
**Épica**: E25 — Re-sync de catálogo propaga el business_type resuelto
**Ceremony level**: L2 (feature con spec clara; extiende una política ya formalizada, no la redefine)
**ADRs relacionados**: ADR-005 (hermano — define la política de corrección segura y el ownership del dato en pim). Origen del gap: E17 (re-clasificación one-off que dejó el sync sin propagar `business_type`).

## Contexto

El sync de catálogo webdata→PIM tiene dos caminos: **crear producto nuevo** (`createNew`) y
**actualizar producto existente** (`updateExisting`). En el camino de creación el `business_type`
resuelto SÍ se propaga; en el camino de actualización **no**. Consecuencia: cuando un producto ya
existe en `global_products` y vuelve a entrar por scraping, el re-sync refresca precio e imagen pero
deja la clasificación congelada. La clasificación solo se corrige corriendo a mano la re-clasificación
masiva (el endpoint S2S de ADR-005 / E24). E25 cierra ese gap: el re-sync debe propagar el
`business_type` resuelto, aplicando **la misma política de corrección segura ya formalizada en ADR-005 §8**.

E25 **no redefine** la política de propagación. La política está decidida por el owner y es idéntica
a la invariante de ADR-005 §8:

> "solo se rellena vacío o se mueve DESDE `almacen` a rubro específico; nunca se toca un producto ya
> en rubro específico (skip), nunca se degrada un rubro curado a mano."

Este ADR **extiende** esa política desde el endpoint batch (ADR-005, E24) al camino del sync producto-a-producto.

### Hechos verificados contra el código (no supuestos)

1. **`updateExisting` hoy solo propaga `Price` + `ImageURL`** — confirmado.
   `SyncProductToPIMUseCase.updateExisting`
   (`services/webdata-service/src/product/application/usecase/sync_product_to_pim_usecase.go:65-74`)
   arma `port.UpdatePIMProductRequest{ Price, ImageURL }` y nada más. El puerto
   `UpdatePIMProductRequest`
   (`services/webdata-service/src/product/domain/port/pim_catalog_syncer.go:50-53`) tiene
   exactamente dos campos: `Price *float64` y `ImageURL string`. **NO existe campo `BusinessType`**
   en el request de update (sí existe en `CreatePIMProductRequest`, línea 47). La épica acierta.

2. **El DTO de update del PIM YA acepta y persiste `business_type` — pero a ciegas.**
   Esto **contradice parcialmente el supuesto de la épica** ("el DTO lo ignora"). El use case
   `UpdateGlobalProductByID`
   (`services/pim-service/src/product/global_catalog/application/usecase/update_global_product_by_id.go`)
   ya tiene `BusinessType *string` en su `UpdateGlobalProductByIDRequest` (línea 25), y en las
   líneas 207-210 hace:
   ```go
   if request.BusinessType != nil {
       existingProduct.SetBusinessType(*request.BusinessType)
   }
   ```
   `GlobalProduct.SetBusinessType(string)` (`domain/entity/global_product.go:241`) sobreescribe el
   valor sin condición alguna. **No es que el PIM ignore el campo: es que lo aplica SIN la invariante
   de corrección segura — pisaría un rubro curado a mano.** El gap real no es de deserialización; es
   de *política*: el update normal del PIM no protege la invariante §8.

3. **El endpoint de update es `PUT /api/v1/global-catalog/products/:id`, grupo `private`** (no S2S).
   Registrado en `GlobalCatalogController.RegisterRoutes`
   (`infrastructure/controller/http_handler.go:119,128`) → `UpdateProductByID`. Es exactamente la
   ruta que webdata ya consume: `PIMCatalogSyncerAdapter.Update`
   (`webdata-service/.../adapter/pim_catalog_syncer_adapter.go:161-185`) hace `PUT
   {baseURL}/api/v1/global-catalog/products/{id}`. webdata se identifica con `X-User-Role:
   marketplace_admin` y `X-Tenant-ID` (no usa el grupo `/s2s`).

4. **El resolver de taxonomía YA vive en go-shared — E24 está cerrada.**
   `ResolveBusinessTypeFromProductCategory` está en
   `libs/go-shared/domain/businesstype/resolver.go` (single source of truth, tabla de reglas con
   orden load-bearing). Ambos servicios están en `go-shared v0.9.0`. webdata lo consume vía un
   thin wrapper (`webdata-service/.../value_object/product_category_resolver.go` re-exporta desde
   go-shared) y el use case de reclasificación del PIM lo importa directo
   (`reclassify_business_types_usecase.go:8,172`). **No hay que tocar go-shared**: el resolver ya
   está donde tiene que estar y E25 lo consume tal cual.

5. **La invariante §8 ya está implementada en el PIM — pero en otro use case.**
   `ReclassifyBusinessTypesUseCase.classifyCandidates`
   (`reclassify_business_types_usecase.go:160-224`) implementa la corrección segura: `no resuelve →
   skip`; `ya en rubro específico (no nil, no "almacen") → skip ya_especifico`; `mismo rubro → skip
   ya_correcto`; `nil/almacen → rubro específico → update (kind relleno|correccion)`. Esta lógica
   **no se debe duplicar**: E25 debe REUSAR exactamente este criterio en el camino del update normal.

### La tensión central de este ADR

ADR-005 dejó la invariante de corrección segura viviendo en el **use case batch** del PIM. El update
producto-a-producto (que webdata invoca en cada re-sync) **no pasa por esa invariante** — la
sobreescritura es directa. Si E25 hiciera que webdata mande el `business_type` resuelto al update
actual, el PIM lo aplicaría a ciegas y **degradaría rubros curados**. Por eso E25 no es "agregar un
campo al request de webdata": es **mover la decisión de propagación al dueño del dato (pim) y aplicar
ahí la misma política §8 que ya existe en batch.**

## Decisión

1. **La invariante de corrección segura se aplica en pim-service**, dueño del dato `global_products`,
   consistente con ADR-005 §1. webdata propaga el `business_type` resuelto como **candidate**; PIM
   decide si rellena, corrige (`almacen`→específico) o skipea (ya específico / ya correcto / no resuelve).

2. **webdata agrega `BusinessType` (candidate) al `UpdatePIMProductRequest`** y lo completa en
   `updateExisting` con el `business_type` ya resuelto del `ScrapedProduct` (el mismo valor que ya usa
   `createNew`). webdata **no** decide la política — solo aporta el candidato.

3. **El use case de update del PIM (`UpdateGlobalProductByID`) aplica la invariante §8 al campo
   `business_type`**, reusando la lógica de `classifyCandidates` extraída a una función de dominio
   compartida — NO se duplica el criterio.

4. **El contrato del request de update se mantiene retrocompatible.** `business_type` sigue siendo
   opcional (`omitempty`). El cambio de comportamiento es server-side (PIM ahora protege la
   invariante); ningún consumidor existente se rompe.

5. **El resolver de go-shared es la única fuente de la taxonomía.** webdata resuelve el candidate con
   el resolver compartido (como ya hace en el upsert y en `createNew`); el PIM no reimplementa
   taxonomía — solo aplica la política de aceptación/skip sobre el candidate recibido.

6. **Idempotencia y observabilidad.** Un re-sync de un producto ya correcto produce 0 cambios de
   `business_type` (skip `ya_correcto`/`ya_especifico`). El skip/aplicación se refleja en el canonical
   log ADR-001 ya cableado en el sync de webdata.

## Detalle de las decisiones que cierra este ADR

### 1. Dónde vive la política de corrección segura — pim-service

**Decisión: en pim.** Fundamento contra el código y contra el bounded context:

- El dato `global_products` es de `pim_db`, con su entity, repo y transacción en
  `src/product/global_catalog/`. Igual criterio que ADR-005 §1: el dueño del dato es el dueño de la
  política sobre el dato.
- La invariante §8 ya está implementada en pim (`classifyCandidates`). Aplicarla también en el update
  es reusar lógica existente en su dueño natural, no exportar reglas a webdata.
- Si webdata "pisara a ciegas" mandando el `business_type` resuelto al update actual, el PIM lo
  aplicaría sin guardas (hecho verificado #2) y degradaría rubros curados. Mover la decisión a webdata
  obligaría a webdata a consultar el estado actual del producto en pim y a duplicar la política §8 —
  doble fuente de verdad y un cross-context write de política. Rechazado.

**Trade-off documentado honestamente**: webdata YA tiene el `business_type` resuelto en mano y el
estado actual del producto lo trae en `findExisting` (el `PIMGlobalProduct` devuelto incluye
`BusinessType`, ver `pim_catalog_syncer.go:35`). Es decir, *técnicamente* webdata podría aplicar la
política antes de llamar. Aun así se decide PIM porque: (a) consistencia con ADR-005 — una sola
implementación de §8, en pim; (b) la política debe valer para CUALQUIER cliente del update, no solo
webdata (ej. mc_admin editando, un futuro importer); poner la guarda en webdata dejaría el endpoint
del PIM desprotegido para los demás. El PIM debe ser seguro por sí mismo.

### 2. Contrato — request de update

**webdata** (`UpdatePIMProductRequest`, `pim_catalog_syncer.go`):
```go
type UpdatePIMProductRequest struct {
    Price        *float64 `json:"price,omitempty"`
    ImageURL     string   `json:"image_url,omitempty"`
    BusinessType string   `json:"business_type,omitempty"`  // NUEVO: candidate, no autoritativo
}
```
`updateExisting` lo completa con el mismo `primaryBusinessType(product)` que ya usa `createNew`
(`sync_product_to_pim_usecase.go:82,116-121`). Si el producto scrapeado no tiene business_type
resuelto, el campo va vacío y el PIM lo trata como "sin candidate" (no toca nada).

**PIM** (`UpdateGlobalProductByIDRequest`): el campo `BusinessType *string` (línea 25) **ya existe**.
No hay cambio de DTO ni de deserialización. El cambio es el reemplazo de la sobreescritura ciega
(líneas 207-210) por la aplicación de la política §8.

**Deserialización tolerante / prioridad documentada**: el request del PIM ya es tolerante (todos los
campos `omitempty`/punteros, semántica "solo actualizo lo que mando"). `business_type` ausente o
vacío = "no propongo candidate" → el PIM no toca el `business_type` actual. Esto preserva el
comportamiento de cualquier consumidor que hoy hace update sin mandar el campo.

### 3. Aplicación de la invariante §8 en el update del PIM

**Decisión**: extraer el criterio de `classifyCandidates`
(`reclassify_business_types_usecase.go:160-224`) a una **función pura de dominio** reutilizable
(p.ej. `value_object.ResolveSafeBusinessTypeTransition(current *string, candidate string) (apply bool, newType string, kind string)`)
y consumirla tanto desde el reclassify batch como desde `UpdateGlobalProductByID`. El update reemplaza:

```go
// ANTES (update_global_product_by_id.go:207-210) — sobreescritura ciega:
if request.BusinessType != nil {
    existingProduct.SetBusinessType(*request.BusinessType)
}
```
por la evaluación de la transición segura: si el candidate corresponde a `relleno`
(actual nil/vacío) o `correccion` (actual `almacen` → específico) **se aplica**; si el producto ya
está en rubro específico o el candidate coincide con el actual, **se skipea** (no se llama
`SetBusinessType`).

La taxonomía no se reimplementa: webdata ya manda el candidate resuelto por el resolver de go-shared;
el PIM solo decide aceptar/skip. (El reclassify batch resuelve el candidate él mismo desde la
categoría; el update lo recibe ya resuelto. La función de transición segura es la **parte común** —
la decisión de aceptación sobre `(current, candidate)` — y es la que se comparte.)

### 4. Idempotencia y observabilidad

- **Idempotente**: re-sync de un producto ya específico o ya correcto → 0 escrituras de
  `business_type` (skip). Verificable en test (invariante §6 de este ADR).
- **Canonical log ADR-001**: el sync de webdata ya tiene logger canónico cableado
  (`sync_product_to_pim_usecase.go:24-33`, `WithLogger`). No se requiere un evento nuevo obligatorio,
  pero el PIM puede emitir el resultado de la transición (`applied`/`skipped` + `kind`) con el
  `PIMEvent` existente si se quiere trazabilidad fina del re-sync. Sin tabla de audit nueva: este
  camino es el flujo normal de sync, no una operación masiva forense como el batch de ADR-005.

### 5. Qué NO toca este ADR

- **go-shared**: intacto. El resolver ya está en `libs/go-shared/domain/businesstype` (v0.9.0) y se
  consume tal cual. Restricción dura del owner respetada.
- **El endpoint S2S de reclasificación** (ADR-005): sigue existiendo para la corrección masiva
  on-demand. E25 reduce su necesidad de uso rutinario (el re-sync ya mantiene la clasificación al
  día) pero no lo reemplaza.
- **Auth / ruta**: se usa el `PUT /api/v1/global-catalog/products/:id` ya existente; no se crea ruta
  nueva ni se cambia el esquema de auth del sync.

## Alternativas consideradas

- **webdata aplica la política antes de llamar al PIM** (con el estado que ya trae de `findExisting`):
  rechazada. Duplica la invariante §8 fuera de su dueño, deja el endpoint de update del PIM
  desprotegido para otros clientes, y vuelve a poner política de catálogo en un servicio que no es
  dueño del dato (mismo anti-patrón cross-context que ADR-005 §1 evita).
- **webdata manda el business_type al update actual sin tocar la política del PIM** (la "solución
  obvia" de la épica): rechazada. El PIM aplica `SetBusinessType` a ciegas (hecho #2) → degradaría
  rubros curados. Viola la política del owner.
- **Reusar el endpoint S2S de reclasificación batch por cada producto del sync**: rechazada. El batch
  está diseñado para operaciones masivas con snapshot+audit; invocarlo por producto es caro y mezcla
  el flujo de sync rutinario con la maquinaria forense de ADR-005.
- **Agregar la guarda solo en el adapter de webdata (skip si el producto ya es específico)**:
  rechazada por la misma razón que la primera alternativa — la fuente de verdad de la política debe
  ser el PIM, server-side, para todos los clientes.
- **Duplicar la lógica de `classifyCandidates` en el update**: rechazada. Se extrae a función pura
  compartida (§3); duplicar la invariante crearía drift entre el camino batch y el camino sync.

## Consecuencias

- ✅ El re-sync mantiene `business_type` al día automáticamente; desaparece la necesidad de correr a
  mano la re-clasificación tras cada scraping.
- ✅ La invariante de corrección segura (§8 de ADR-005) queda con **una sola implementación**,
  server-side en el PIM, válida para TODO cliente del update (webdata, mc_admin, futuros importers).
- ✅ Retrocompatible: `business_type` opcional, deserialización tolerante; ningún consumidor existente
  se rompe.
- ✅ Reusa el resolver de go-shared y el logger canónico ADR-001 — sin inventar mecanismos ni tocar
  go-shared.
- ⚠️ Cambia el comportamiento server-side del `PUT /global-catalog/products/:id`: antes
  `business_type` se aplicaba siempre; ahora se aplica solo si pasa la política §8. Cualquier
  consumidor que dependiera de pisar `business_type` a ciegas vería un skip. (Verificado: el único
  consumidor del campo en el update hoy es el flujo de edición — el sync ni siquiera lo mandaba.)
  Documentar el cambio de contrato en el changelog del PIM.
- ⚠️ Acopla `UpdateGlobalProductByID` a la función de transición segura compartida con el reclassify
  batch; un cambio futuro a la política impacta ambos caminos (deseable: es el punto, una sola
  política).
- 🚫 Fuera de scope: tocar go-shared / la taxonomía; un evento de audit persistente para el re-sync
  (el canonical log alcanza para el flujo rutinario); reemplazar o deprecar el endpoint S2S de
  reclasificación batch.

## Invariantes a preservar (verificables en tests)

- **Relleno**: producto con `business_type` nil/vacío + candidate resuelto → se aplica el candidate
  (`kind=relleno`).
- **Corrección**: producto con `business_type="almacen"` + candidate específico distinto → se aplica
  (`kind=correccion`).
- **Skip ya_especifico**: producto ya en rubro específico (no nil, no `almacen`) → `business_type`
  **no cambia**, aunque el candidate difiera. Nunca se degrada un rubro curado.
- **Skip ya_correcto**: candidate == business_type actual → 0 escrituras.
- **Skip sin candidate**: request sin `business_type` (o vacío) → el PIM no toca `business_type`
  (retrocompatibilidad del update).
- **Idempotencia**: re-sync consecutivo del mismo producto ya clasificado → 0 cambios de
  `business_type`.
- **Paridad con batch**: el criterio aplicado en el update produce el MISMO veredicto
  (apply/skip + kind) que `classifyCandidates` para el mismo `(current, candidate)` — garantizado por
  compartir la función de dominio extraída (§3), verificable con un test de paridad.
- **Sin tocar otros campos**: agregar `business_type` al update no altera la semántica de `Price`,
  `ImageURL` ni el resto de campos del request.

## Revisión prevista

Revisar si: (a) aparece un tercer camino de escritura de `business_type` (ej. importer CSV de
mc_admin) que también deba pasar por la política — confirmaría el acierto de tenerla server-side en
pim; (b) la política §8 necesita variar entre el camino batch y el camino sync (hoy se asume idéntica;
si divergen, la función compartida debe parametrizarse, no duplicarse); (c) el volumen de re-syncs
hace deseable un evento de audit persistente del cambio de clasificación (hoy fuera de scope).
