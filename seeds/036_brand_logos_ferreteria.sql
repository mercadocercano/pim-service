-- =============================================================================
-- SEED 036: Logos de marcas para ferretería
-- Actualiza business_type_templates.brands con logo_url via apistemic.com
-- =============================================================================

UPDATE business_type_templates
SET brands = '[
    {"name": "Loma Negra", "logo_url": "https://logos-api.apistemic.com/domain:lomanegra.com.ar", "suggested_for_categories": ["cementos-cales"]},
    {"name": "Avellaneda", "logo_url": "https://logos-api.apistemic.com/domain:cementosavellaneda.com", "suggested_for_categories": ["cementos-cales"]},
    {"name": "Acindar", "logo_url": "https://logos-api.apistemic.com/domain:acindar.com.ar", "suggested_for_categories": ["hierros-acero"]},
    {"name": "Stanley", "logo_url": "https://logos-api.apistemic.com/domain:stanleytools.com", "suggested_for_categories": ["martillos-mazas","destornilladores","llaves-pinzas","medicion-trazado"]},
    {"name": "Bosch", "logo_url": "https://logos-api.apistemic.com/domain:bosch.com", "suggested_for_categories": ["taladros-percutores","amoladoras","accesorios-mechas"]},
    {"name": "Black+Decker", "logo_url": "https://logos-api.apistemic.com/domain:blackanddecker.com", "suggested_for_categories": ["taladros-percutores","sierras-electricas"]},
    {"name": "Makita", "logo_url": "https://logos-api.apistemic.com/domain:makita.com", "suggested_for_categories": ["atornilladores-electricos","taladros-percutores"]},
    {"name": "DeWalt", "logo_url": "https://logos-api.apistemic.com/domain:dewalt.com", "suggested_for_categories": ["amoladoras","sierras-electricas"]},
    {"name": "Tramontina", "logo_url": "https://logos-api.apistemic.com/domain:tramontina.com", "suggested_for_categories": ["llaves-pinzas","sierras-serruchos"]},
    {"name": "Tigre", "logo_url": "https://logos-api.apistemic.com/domain:tigre.com.ar", "suggested_for_categories": ["canos-pvc","conexiones-accesorios-plom"]},
    {"name": "FV", "logo_url": "https://logos-api.apistemic.com/domain:fv.com.ar", "suggested_for_categories": ["griferias-ferret","tanques-flotantes"]},
    {"name": "Alba", "logo_url": "https://logos-api.apistemic.com/domain:alba.com.ar", "suggested_for_categories": ["latex-ferret","esmaltes-ferret"]},
    {"name": "Sinteplast", "logo_url": "https://logos-api.apistemic.com/domain:sinteplast.com.ar", "suggested_for_categories": ["latex-ferret"]},
    {"name": "Sika", "logo_url": "https://logos-api.apistemic.com/domain:sika.com", "suggested_for_categories": ["membranas-aislantes"]},
    {"name": "Schneider", "logo_url": "https://logos-api.apistemic.com/domain:se.com", "suggested_for_categories": ["llaves-termicas"]},
    {"name": "Lusqtoff", "logo_url": "https://logos-api.apistemic.com/domain:lusqtoff.com.ar", "suggested_for_categories": ["taladros-percutores","amoladoras","sierras-electricas","atornilladores-electricos"]},
    {"name": "Bahco", "logo_url": "https://logos-api.apistemic.com/domain:bahco.com", "suggested_for_categories": ["martillos-mazas","destornilladores","llaves-pinzas","sierras-serruchos"]},
    {"name": "Hamilton", "logo_url": "https://logos-api.apistemic.com/domain:hamiltontools.com.ar", "suggested_for_categories": ["martillos-mazas","llaves-pinzas","amoladoras"]},
    {"name": "Philips", "logo_url": "https://logos-api.apistemic.com/domain:philips.com", "suggested_for_categories": ["iluminacion-ferret"]},
    {"name": "Poxipol", "logo_url": "https://logos-api.apistemic.com/domain:poxipol.com", "suggested_for_categories": ["cintas-adhesivos-ferret"]},
    {"name": "La Gotita", "logo_url": "https://logos-api.apistemic.com/domain:lagotita.com.ar", "suggested_for_categories": ["cintas-adhesivos-ferret"]},
    {"name": "Cetol", "logo_url": "https://logos-api.apistemic.com/domain:cetol.com.ar", "suggested_for_categories": ["esmaltes-ferret"]},
    {"name": "Weber", "logo_url": "https://logos-api.apistemic.com/domain:weber.com.ar", "suggested_for_categories": ["cementos-cales"]},
    {"name": "Megaflex", "logo_url": "https://logos-api.apistemic.com/domain:megaflex.com.ar", "suggested_for_categories": ["membranas-aislantes"]},
    {"name": "Fischer", "logo_url": "https://logos-api.apistemic.com/domain:fischer.de", "suggested_for_categories": ["tarugos-bulones"]},
    {"name": "IMSA", "logo_url": "https://logos-api.apistemic.com/domain:imsa.com.ar", "suggested_for_categories": ["cables-ferret"]},
    {"name": "Jeluz", "logo_url": "https://logos-api.apistemic.com/domain:jeluz.com.ar", "suggested_for_categories": ["tomas-interruptores-ferret"]}
  ]'::jsonb,
updated_at = CURRENT_TIMESTAMP
WHERE business_type_id = (SELECT id FROM business_types WHERE code = 'ferreteria') AND is_default = true;
