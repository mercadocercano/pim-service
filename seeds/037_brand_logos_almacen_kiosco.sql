-- =============================================================================
-- SEED 037: Logos de marcas para almacén y kiosco
-- =============================================================================

-- ALMACÉN
UPDATE business_type_templates
SET brands = '[
    {"name": "La Serenísima", "logo_url": "https://logos-api.apistemic.com/domain:laserenisima.com.ar", "suggested_for_categories": ["leches","yogures","quesos-manteca","lacteos-frescos"]},
    {"name": "Arcor", "logo_url": "https://logos-api.apistemic.com/domain:arcor.com", "suggested_for_categories": ["golosinas-snacks","galletitas-almacen","conservas-enlatados","alfajores-chocolates"]},
    {"name": "Marolio", "logo_url": "https://logos-api.apistemic.com/domain:marolio.com.ar", "suggested_for_categories": ["aceites-vinagres","conservas-enlatados","arroz-legumbres"]},
    {"name": "Coca-Cola", "logo_url": "https://logos-api.apistemic.com/domain:coca-cola.com", "suggested_for_categories": ["gaseosas-aguas","bebidas-almacen"]},
    {"name": "Molinos Río de la Plata", "logo_url": "https://logos-api.apistemic.com/domain:molinos.com.ar", "suggested_for_categories": ["pastas-secas","harinas-premezclas","aceites-vinagres"]},
    {"name": "Bagley", "logo_url": "https://logos-api.apistemic.com/domain:bagley.com.ar", "suggested_for_categories": ["galletitas-almacen"]},
    {"name": "Sancor", "logo_url": "https://logos-api.apistemic.com/domain:sancor.com", "suggested_for_categories": ["leches","yogures","quesos-manteca"]},
    {"name": "La Campagnola", "logo_url": "https://logos-api.apistemic.com/domain:lacampagnola.com.ar", "suggested_for_categories": ["conservas-enlatados"]},
    {"name": "Hellmann''s", "logo_url": "https://logos-api.apistemic.com/domain:hellmanns.com", "suggested_for_categories": ["almacen-seco"]},
    {"name": "Skip", "logo_url": "https://logos-api.apistemic.com/domain:skip.com.ar", "suggested_for_categories": ["detergentes-jabones"]},
    {"name": "Dove", "logo_url": "https://logos-api.apistemic.com/domain:dove.com", "suggested_for_categories": ["jabones-desodorantes","shampoo-acondicionador"]},
    {"name": "Higienol", "logo_url": "https://logos-api.apistemic.com/domain:higienol.com.ar", "suggested_for_categories": ["papel-higienico"]},
    {"name": "Paladini", "logo_url": "https://logos-api.apistemic.com/domain:paladini.com.ar", "suggested_for_categories": ["fiambres-embutidos","fiambreria"]},
    {"name": "Fargo", "logo_url": "https://logos-api.apistemic.com/domain:fargo.com.ar", "suggested_for_categories": ["pan-envasado","panaderia-reposteria"]},
    {"name": "Ledesma", "logo_url": "https://logos-api.apistemic.com/domain:ledesma.com.ar", "suggested_for_categories": ["almacen-seco"]},
    {"name": "Natura", "logo_url": "https://logos-api.apistemic.com/domain:natura.com.ar", "suggested_for_categories": ["perfumeria"]},
    {"name": "Quilmes", "logo_url": "https://logos-api.apistemic.com/domain:quilmes.com.ar", "suggested_for_categories": ["cervezas-vinos"]},
    {"name": "Unilever", "logo_url": "https://logos-api.apistemic.com/domain:unilever.com", "suggested_for_categories": ["detergentes-jabones","jabones-desodorantes","shampoo-acondicionador"]},
    {"name": "Pampers", "logo_url": "https://logos-api.apistemic.com/domain:pampers.com", "suggested_for_categories": ["panales"]},
    {"name": "Knorr", "logo_url": "https://logos-api.apistemic.com/domain:knorr.com", "suggested_for_categories": ["almacen-seco","conservas-enlatados"]}
  ]'::jsonb,
updated_at = CURRENT_TIMESTAMP
WHERE business_type_id = (SELECT id FROM business_types WHERE code = 'almacen') AND is_default = true;

-- KIOSCO
UPDATE business_type_templates
SET brands = '[
    {"name": "Arcor", "logo_url": "https://logos-api.apistemic.com/domain:arcor.com", "suggested_for_categories": ["golosinas","alfajores","caramelos-chicles","chocolates"]},
    {"name": "Coca-Cola", "logo_url": "https://logos-api.apistemic.com/domain:coca-cola.com", "suggested_for_categories": ["gaseosas","bebidas"]},
    {"name": "Quilmes", "logo_url": "https://logos-api.apistemic.com/domain:quilmes.com.ar", "suggested_for_categories": ["cervezas"]},
    {"name": "Pepsi", "logo_url": "https://logos-api.apistemic.com/domain:pepsi.com", "suggested_for_categories": ["gaseosas","bebidas"]},
    {"name": "Bagley", "logo_url": "https://logos-api.apistemic.com/domain:bagley.com.ar", "suggested_for_categories": ["galletitas","galletitas-dulces","galletitas-saladas"]},
    {"name": "Mondelez", "logo_url": "https://logos-api.apistemic.com/domain:mondelezinternational.com", "suggested_for_categories": ["chocolates","galletitas","golosinas"]},
    {"name": "Felfort", "logo_url": "https://logos-api.apistemic.com/domain:felfort.com.ar", "suggested_for_categories": ["chocolates","golosinas"]},
    {"name": "Guaymallén", "logo_url": "https://logos-api.apistemic.com/domain:guaymallenok.com.ar", "suggested_for_categories": ["alfajores"]},
    {"name": "Lay''s", "logo_url": "https://logos-api.apistemic.com/domain:lays.com", "suggested_for_categories": ["papas-fritas","snacks"]},
    {"name": "Beldent", "logo_url": "https://logos-api.apistemic.com/domain:beldent.com.ar", "suggested_for_categories": ["caramelos-chicles"]},
    {"name": "Red Bull", "logo_url": "https://logos-api.apistemic.com/domain:redbull.com", "suggested_for_categories": ["energizantes"]},
    {"name": "Villavicencio", "logo_url": "https://logos-api.apistemic.com/domain:villavicencio.com.ar", "suggested_for_categories": ["aguas-saborizadas"]},
    {"name": "Bonafide", "logo_url": "https://logos-api.apistemic.com/domain:bonafide.com.ar", "suggested_for_categories": ["chocolates","alfajores"]},
    {"name": "Georgalos", "logo_url": "https://logos-api.apistemic.com/domain:georgalos.com", "suggested_for_categories": ["chocolates","golosinas"]},
    {"name": "Levité", "logo_url": "https://logos-api.apistemic.com/domain:levite.com.ar", "suggested_for_categories": ["aguas-saborizadas","jugos"]},
    {"name": "Marlboro", "logo_url": "https://logos-api.apistemic.com/domain:pmi.com", "suggested_for_categories": ["cigarrillos-tabaco"]},
    {"name": "Terrabusi", "logo_url": "https://logos-api.apistemic.com/domain:terrabusi.com.ar", "suggested_for_categories": ["galletitas","galletitas-dulces"]},
    {"name": "Milka", "logo_url": "https://logos-api.apistemic.com/domain:milka.com", "suggested_for_categories": ["chocolates","alfajores"]},
    {"name": "Bic", "logo_url": "https://logos-api.apistemic.com/domain:bicworld.com", "suggested_for_categories": ["encendedores","varios-kiosco"]},
    {"name": "Duracell", "logo_url": "https://logos-api.apistemic.com/domain:duracell.com", "suggested_for_categories": ["pilas-baterias"]}
  ]'::jsonb,
updated_at = CURRENT_TIMESTAMP
WHERE business_type_id = (SELECT id FROM business_types WHERE code = 'kiosco') AND is_default = true;
