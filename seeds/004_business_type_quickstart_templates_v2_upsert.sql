-- Seed v2 (idempotente): Templates de quickstart basados en categorías realmente disponibles
-- NO borra business_type_templates. Crea/actualiza el template default por (business_type_id, region) donde is_default=true.
--
-- Fuente de disponibilidad: pim_db.global_products (is_verified=true) agrupado por global_products.category (slug curado).
-- Requiere que esos slugs existan en marketplace_categories (ver seeds/011_marketplace_categories_from_curation_upsert.sql).
--
-- Regla: preferir categorías con >=10 productos verificados; garantizar >=3 categorías donde aplique (fallback: 'general').

CREATE OR REPLACE FUNCTION _qs_get_category_ids_verified(category_slugs TEXT[], min_verified INT)
RETURNS JSONB AS $$
DECLARE
  result JSONB := '[]';
  current_slug TEXT;
  category_id UUID;
  verified_count INT;
BEGIN
  FOREACH current_slug IN ARRAY category_slugs LOOP
    SELECT COUNT(*) INTO verified_count
    FROM global_products gp
    WHERE gp.is_active = true
      AND gp.is_verified = true
      AND gp.category = current_slug;

    IF verified_count >= min_verified THEN
      SELECT id INTO category_id FROM marketplace_categories WHERE slug = current_slug;
      IF category_id IS NOT NULL THEN
        result := result || jsonb_build_object('id', category_id, 'slug', current_slug);
      END IF;
    END IF;
  END LOOP;
  RETURN result;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION _qs_category_count(json_categories JSONB)
RETURNS INT AS $$
BEGIN
  RETURN COALESCE(jsonb_array_length(json_categories), 0);
END;
$$ LANGUAGE plpgsql;

DO $$
DECLARE
  bt RECORD;
  region_code TEXT := 'AR';
  min_verified INT := 10;
  cats JSONB;
  fallback JSONB;
  is_default_template BOOLEAN;
BEGIN
  -- Fallback category (si existe y tiene volumen)
  fallback := _qs_get_category_ids_verified(ARRAY['general'], min_verified);

  FOR bt IN SELECT id, code, name FROM business_types WHERE is_active = true ORDER BY sort_order, code LOOP

    -- Selección por rubro (slugs curados)
    IF bt.code IN ('almacen', 'kiosco', 'supermercado') THEN
      cats := _qs_get_category_ids_verified(
        ARRAY[
          'alimentos-envasados',
          'bebidas-sin-alcohol',
          'bebidas-con-alcohol',
          'galletas',
          'golosinas',
          'conservas',
          'pastas-cereales',
          'snacks-salados'
        ],
        min_verified
      );

    ELSIF bt.code = 'panaderia' THEN
      cats := _qs_get_category_ids_verified(ARRAY['panificados', 'alimentos-envasados', 'bebidas-sin-alcohol', 'galletas'], min_verified);

    ELSIF bt.code = 'verduleria' THEN
      cats := _qs_get_category_ids_verified(ARRAY['frutas-verduras', 'alimentos-frescos', 'bebidas-sin-alcohol', 'conservas'], min_verified);

    ELSIF bt.code = 'carniceria' THEN
      cats := _qs_get_category_ids_verified(ARRAY['carnes-pescados', 'lacteos-fiambres', 'alimentos-envasados', 'conservas'], min_verified);

    ELSIF bt.code IN ('farmacia', 'perfumeria') THEN
      cats := _qs_get_category_ids_verified(
        ARRAY['salud-belleza', 'higiene-personal', 'cuidado-cabello', 'cuidado-facial', 'cuidado-corporal', 'cuidado-bucal'],
        min_verified
      );

    ELSIF bt.code IN ('electronica', 'celulares', 'computacion') THEN
      cats := _qs_get_category_ids_verified(ARRAY['tecnologia', 'general'], min_verified);

    ELSIF bt.code IN ('muebleria', 'bazar', 'floreria') THEN
      cats := _qs_get_category_ids_verified(ARRAY['hogar-jardin', 'general'], min_verified);

    ELSIF bt.code IN ('ropa', 'deportes', 'zapateria', 'relojeria') THEN
      cats := _qs_get_category_ids_verified(ARRAY['moda-accesorios', 'general'], min_verified);

    ELSIF bt.code IN ('jugueteria') THEN
      cats := _qs_get_category_ids_verified(ARRAY['bebes-ninos', 'bebes-0-2', 'general'], min_verified);

    ELSIF bt.code IN ('delivery') THEN
      cats := _qs_get_category_ids_verified(ARRAY['alimentos-envasados', 'bebidas-sin-alcohol', 'frutas-verduras'], min_verified);

    ELSE
      cats := _qs_get_category_ids_verified(ARRAY['general'], min_verified);
    END IF;

    -- Garantizar mínimo 3 categorías si es posible (fallback a 'general')
    IF _qs_category_count(cats) < 3 THEN
      cats := cats || fallback;
    END IF;

    is_default_template := (_qs_category_count(cats) >= 3);

    -- Upsert del template default por (business_type_id, region) para is_default=true
    INSERT INTO business_type_templates (
      business_type_id, name, description, region, categories, is_default, is_active, generated_by, version
    ) VALUES (
      bt.id,
      bt.name || ' (Curación)',
      'Template generado desde categorías curadas (Mongo) + disponibilidad real',
      region_code,
      cats,
      true,
      true,
      'curation-sync',
      '2.0.0'
    )
    ON CONFLICT (business_type_id, region) WHERE is_default = true
    DO UPDATE SET
      name = EXCLUDED.name,
      description = EXCLUDED.description,
      categories = EXCLUDED.categories,
      is_active = EXCLUDED.is_active,
      generated_by = EXCLUDED.generated_by,
      version = EXCLUDED.version,
      updated_at = CURRENT_TIMESTAMP;

    -- Nota: si el template no llega a 3 categorías, igual queda default (para no romper UX),
    -- pero el reporte va a marcarlo como \"parcial\".
    -- Si querés forzar is_default=false cuando <3, lo cambiamos (requiere manejar la partial unique index distinto).

  END LOOP;
END $$;

DROP FUNCTION _qs_get_category_ids_verified(TEXT[], INT);
DROP FUNCTION _qs_category_count(JSONB);


