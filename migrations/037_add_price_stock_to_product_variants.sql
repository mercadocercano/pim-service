-- Migración 037: Agregar price y stock a product_variants
-- Objetivo: Permitir persistir precio y stock por variante
-- Fecha: 2026-02-04

BEGIN;

ALTER TABLE product_variants
ADD COLUMN price NUMERIC(10,2) NOT NULL DEFAULT 0,
ADD COLUMN stock INTEGER NOT NULL DEFAULT 0;

COMMIT;
