-- Fix: alinear categorías de business_type_templates con slugs reales de global_products
-- Ejecutar después de seed 112. Idempotente.

-- BAZAR
UPDATE business_type_templates SET categories = '[
  {"name": "Vajilla", "slug": "vajilla", "level": 0},
  {"name": "Vajilla y Vidriería", "slug": "vajilla-vidrieria", "level": 0},
  {"name": "Utensilios de Cocina", "slug": "utensilios-cocina", "level": 0},
  {"name": "Almacenamiento Hogar", "slug": "almacenamiento-hogar", "level": 0},
  {"name": "Organizacion Hogar", "slug": "organizacion-hogar", "level": 0},
  {"name": "Textiles Hogar", "slug": "textiles-hogar", "level": 0},
  {"name": "Mesa", "slug": "mesa", "level": 0},
  {"name": "Cubiertos", "slug": "cubiertos", "level": 0},
  {"name": "Decoracion", "slug": "decoracion", "level": 0},
  {"name": "Iluminacion", "slug": "iluminacion", "level": 0},
  {"name": "Limpieza", "slug": "limpieza", "level": 0},
  {"name": "Cocina", "slug": "cocina", "level": 0}
]'::jsonb
WHERE id = 'c278befb-7131-4bc4-a7f4-416b14999c5d';

-- ELECTRODOMÉSTICOS
UPDATE business_type_templates SET categories = '[
  {"name": "Línea Blanca Pequeña", "slug": "linea-blanca-pequena", "level": 0},
  {"name": "Pequeños Cocina", "slug": "pequenos-cocina", "level": 0},
  {"name": "Pequeños Limpieza", "slug": "pequenos-limpieza", "level": 0},
  {"name": "Audio y Video", "slug": "audio-video", "level": 0},
  {"name": "TV y Audio", "slug": "tv-audio", "level": 0},
  {"name": "Climatizacion", "slug": "climatizacion", "level": 0},
  {"name": "Aires Acondicionados", "slug": "aires-acondicionados", "level": 0},
  {"name": "Calefaccion", "slug": "calefaccion", "level": 0},
  {"name": "Ventilacion", "slug": "ventilacion", "level": 0},
  {"name": "Iluminacion LED", "slug": "iluminacion-led", "level": 0},
  {"name": "Heladeras", "slug": "heladeras", "level": 0},
  {"name": "Lavarropas", "slug": "lavarropas", "level": 0},
  {"name": "Cocinas", "slug": "cocinas", "level": 0},
  {"name": "Tecnologia", "slug": "tecnologia", "level": 0}
]'::jsonb
WHERE id = '4f1816ac-8e3f-40f3-8d96-e2dc54e23d33';

-- FERRETERÍA (fix cerrajeria→cerrajeria-seguridad, pintura→pinturas-latex)
UPDATE business_type_templates SET categories = '[
  {"name": "Herramientas Manuales", "slug": "herramientas-manuales", "level": 0},
  {"name": "Herramientas Eléctricas", "slug": "herramientas-electricas", "level": 0},
  {"name": "Taladros y Percutores", "slug": "taladros-percutores", "level": 0},
  {"name": "Sierras Eléctricas", "slug": "sierras-electricas", "level": 0},
  {"name": "Plomería", "slug": "sanitarios-plomeria", "level": 0},
  {"name": "Caños PVC", "slug": "canos-pvc", "level": 0},
  {"name": "Cables Electricidad", "slug": "cables-electricidad", "level": 0},
  {"name": "Fijaciones", "slug": "fijaciones", "level": 0},
  {"name": "Cerrajería y Seguridad", "slug": "cerrajeria-seguridad", "level": 0},
  {"name": "Pinturas Látex", "slug": "pinturas-latex", "level": 0},
  {"name": "Esmaltes y Barnices", "slug": "esmaltes-barnices", "level": 0},
  {"name": "Adhesivos y Selladores", "slug": "adhesivos-selladores", "level": 0},
  {"name": "Accesorios Pintura", "slug": "accesorios-pintura", "level": 0},
  {"name": "Medición", "slug": "medicion", "level": 0},
  {"name": "Iluminacion", "slug": "iluminacion", "level": 0}
]'::jsonb
WHERE id = 'b2000001-0000-4000-8000-000000000003';

-- FIAMBRERÍA
UPDATE business_type_templates SET categories = '[
  {"name": "Quesos Frescos", "slug": "quesos-frescos", "level": 0},
  {"name": "Quesos Blandos", "slug": "quesos-blandos", "level": 0},
  {"name": "Quesos Semiduros", "slug": "quesos-semiduros", "level": 0},
  {"name": "Quesos Duros", "slug": "quesos-duros", "level": 0},
  {"name": "Quesos Untables", "slug": "quesos-untables", "level": 0},
  {"name": "Fiambres Cocidos", "slug": "fiambres-cocidos", "level": 0},
  {"name": "Fiambres Curados", "slug": "fiambres-curados", "level": 0},
  {"name": "Jamones", "slug": "jamones", "level": 0},
  {"name": "Fiambres y Embutidos", "slug": "fiambres-embutidos", "level": 0},
  {"name": "Lacteos Frescos", "slug": "lacteos-frescos", "level": 0},
  {"name": "Picadas", "slug": "picadas", "level": 0},
  {"name": "Conservas Gourmet", "slug": "conservas-gourmet", "level": 0},
  {"name": "Pan Fiambreria", "slug": "pan-fiambreria", "level": 0}
]'::jsonb
WHERE id = 'ad55633a-8686-452a-b233-84b760c15e2d';

-- JUGUETERÍA
UPDATE business_type_templates SET categories = '[
  {"name": "Bebés 0-2", "slug": "bebes-0-2", "level": 0},
  {"name": "Bebés", "slug": "bebes", "level": 0},
  {"name": "Niños 3-5", "slug": "ninos-3-5", "level": 0},
  {"name": "Niños 6-10", "slug": "ninos-6-10", "level": 0},
  {"name": "Adolescentes", "slug": "adolescentes", "level": 0},
  {"name": "Juegos de Mesa", "slug": "juegos-mesa", "level": 0},
  {"name": "Cotillón", "slug": "cotillon", "level": 0},
  {"name": "Vehículos", "slug": "vehiculos", "level": 0},
  {"name": "Construcción", "slug": "construccion", "level": 0},
  {"name": "Creativos", "slug": "creativos", "level": 0},
  {"name": "Didácticos", "slug": "didacticos", "level": 0},
  {"name": "Aire Libre", "slug": "aire-libre", "level": 0},
  {"name": "Muñecos y Figuras", "slug": "munecos-figuras", "level": 0},
  {"name": "Electrónico", "slug": "electronico", "level": 0},
  {"name": "Disfraces", "slug": "disfraces", "level": 0}
]'::jsonb
WHERE id = '86001b97-681c-491d-9a0e-0408d811a3ce';

-- KIOSCO
UPDATE business_type_templates SET categories = '[
  {"name": "Gaseosas", "slug": "gaseosas", "level": 0},
  {"name": "Aguas", "slug": "aguas-kiosco", "level": 0},
  {"name": "Jugos", "slug": "jugos-kiosco", "level": 0},
  {"name": "Energéticas", "slug": "energeticas", "level": 0},
  {"name": "Chocolates", "slug": "chocolates", "level": 0},
  {"name": "Alfajores", "slug": "alfajores", "level": 0},
  {"name": "Galletitas Dulces", "slug": "galletitas-dulces", "level": 0},
  {"name": "Galletitas Saladas", "slug": "galletitas-saladas", "level": 0},
  {"name": "Gomitas y Confites", "slug": "gomitas-confites", "level": 0},
  {"name": "Caramelos y Chicles", "slug": "caramelos-chicles", "level": 0},
  {"name": "Chupetines", "slug": "chupetines", "level": 0},
  {"name": "Maní y Snacks", "slug": "mani-snacks", "level": 0},
  {"name": "Cigarrillos", "slug": "cigarrillos", "level": 0},
  {"name": "Pilas y Baterías", "slug": "pilas-baterias", "level": 0},
  {"name": "Bebidas", "slug": "bebidas", "level": 0}
]'::jsonb
WHERE id = 'b2000001-0000-4000-8000-000000000002';

-- LIBRERÍA
UPDATE business_type_templates SET categories = '[
  {"name": "Útiles Escolares", "slug": "utiles-escolares", "level": 0},
  {"name": "Lápices y Lapiceras", "slug": "lapices-lapiceras", "level": 0},
  {"name": "Escritura", "slug": "escritura", "level": 0},
  {"name": "Cuadernos", "slug": "cuadernos", "level": 0},
  {"name": "Carpetas", "slug": "carpetas", "level": 0},
  {"name": "Arte", "slug": "arte", "level": 0},
  {"name": "Papelería", "slug": "papeleria", "level": 0},
  {"name": "Papelería Insumos", "slug": "papeleria-insumos", "level": 0},
  {"name": "Adhesivos", "slug": "adhesivos", "level": 0},
  {"name": "Tecnología Básica", "slug": "tecnologia-basica", "level": 0},
  {"name": "Organización", "slug": "organizacion", "level": 0},
  {"name": "Mochilas", "slug": "mochilas", "level": 0},
  {"name": "Juegos Didácticos", "slug": "juegos-didacticos", "level": 0}
]'::jsonb
WHERE id = '4d53e0e5-bef2-4af5-9f41-7113ca9c1b52';

-- PERFUMERIA
UPDATE business_type_templates SET categories = '[
  {"name": "Cuidado Capilar", "slug": "cuidado-capilar", "level": 0},
  {"name": "Cuidado Corporal", "slug": "cuidado-corporal", "level": 0},
  {"name": "Cuidado Facial", "slug": "cuidado-facial", "level": 0},
  {"name": "Tinturas Capilares", "slug": "tinturas-capilares", "level": 0},
  {"name": "Shampoos Profesionales", "slug": "shampoos-profesionales", "level": 0},
  {"name": "Tratamientos Capilares", "slug": "tratamientos-capilares", "level": 0},
  {"name": "Maquillaje Rostro", "slug": "maquillaje-rostro", "level": 0},
  {"name": "Maquillaje Ojos", "slug": "maquillaje-ojos", "level": 0},
  {"name": "Maquillaje Labios", "slug": "maquillaje-labios", "level": 0},
  {"name": "Esmaltes Uñas", "slug": "esmaltes-unas", "level": 0},
  {"name": "Esmaltes y Maquillaje", "slug": "esmaltes-maquillaje", "level": 0},
  {"name": "Fragancias Mujer", "slug": "fragancias-mujer", "level": 0},
  {"name": "Fragancias Hombre", "slug": "fragancias-hombre", "level": 0},
  {"name": "Cremas Corporales", "slug": "cremas-corporales", "level": 0},
  {"name": "Accesorios Belleza", "slug": "accesorios-belleza", "level": 0}
]'::jsonb
WHERE id = '5301a74e-4cde-449a-90bf-96b87a8a397a';

-- PILETAS
UPDATE business_type_templates SET categories = '[
  {"name": "Cloro y Químicos", "slug": "cloro-quimicos", "level": 0},
  {"name": "Químicos Pileta", "slug": "quimicos-pileta", "level": 0},
  {"name": "Tabletas Cloro", "slug": "tabletas-cloro", "level": 0},
  {"name": "Equipos Filtración", "slug": "equipos-filtracion", "level": 0},
  {"name": "Bombas y Filtros", "slug": "bombas-filtros-pileta", "level": 0},
  {"name": "Elevadores pH", "slug": "elevadores-ph", "level": 0},
  {"name": "Limpieza Pileta", "slug": "limpieza-pileta", "level": 0},
  {"name": "Accesorios Pileta", "slug": "accesorios-pileta", "level": 0},
  {"name": "Mantenimiento", "slug": "mantenimiento-pileta", "level": 0},
  {"name": "Piletas Estructurales", "slug": "piletas-estructurales", "level": 0}
]'::jsonb
WHERE id = 'b5407229-0e0f-4bd7-81e5-dc89b51e1b78';

-- ROPA
UPDATE business_type_templates SET categories = '[
  {"name": "Remeras", "slug": "remeras", "level": 0},
  {"name": "Remeras Básicas", "slug": "remeras-basicas", "level": 0},
  {"name": "Pantalones", "slug": "pantalones", "level": 0},
  {"name": "Ropa Interior Hombre", "slug": "ropa-interior-hombre", "level": 0},
  {"name": "Ropa Interior Mujer", "slug": "ropa-interior-mujer", "level": 0},
  {"name": "Medias", "slug": "medias", "level": 0},
  {"name": "Calzado", "slug": "calzado", "level": 0},
  {"name": "Calzado Básico", "slug": "calzado-basico", "level": 0},
  {"name": "Camperas", "slug": "camperas", "level": 0},
  {"name": "Buzos", "slug": "buzos", "level": 0},
  {"name": "Deportiva", "slug": "deportiva", "level": 0},
  {"name": "Niños", "slug": "ninos", "level": 0},
  {"name": "Accesorios", "slug": "accesorios", "level": 0}
]'::jsonb
WHERE id = '027b96f7-e3bc-4d3f-a236-50c4e8d66d32';

-- VERDULERÍA
UPDATE business_type_templates SET categories = '[
  {"name": "Frutas Frescas", "slug": "frutas-frescas", "level": 0},
  {"name": "Frutas Secas", "slug": "frutas-secas", "level": 0},
  {"name": "Verduras Hoja", "slug": "verduras-hoja", "level": 0},
  {"name": "Verduras Raíz", "slug": "verduras-raiz", "level": 0},
  {"name": "Verduras Fruto", "slug": "verduras-fruto", "level": 0},
  {"name": "Hierbas Aromáticas", "slug": "hierbas-aromaticas", "level": 0},
  {"name": "Huevos", "slug": "huevos", "level": 0}
]'::jsonb
WHERE id = 'b2000001-0000-4000-8000-000000000005';

-- VETERINARIA
UPDATE business_type_templates SET categories = '[
  {"name": "Alimento Perro", "slug": "alimento-perro", "level": 0},
  {"name": "Alimento Gato", "slug": "alimento-gato", "level": 0},
  {"name": "Alimento Húmedo", "slug": "alimento-humedo", "level": 0},
  {"name": "Snacks Mascotas", "slug": "snacks-mascotas", "level": 0},
  {"name": "Accesorios Perro", "slug": "accesorios-perro", "level": 0},
  {"name": "Accesorios Gato", "slug": "accesorios-gato", "level": 0},
  {"name": "Higiene Mascotas", "slug": "higiene-mascotas", "level": 0},
  {"name": "Ropa Mascotas", "slug": "ropa-mascotas", "level": 0},
  {"name": "Medicamentos OTC", "slug": "medicamentos-otc", "level": 0},
  {"name": "Otros Animales", "slug": "otros-animales", "level": 0}
]'::jsonb
WHERE id = 'b608c693-6c70-428f-a308-2524a5e7f00d';

-- VINOTECA
UPDATE business_type_templates SET categories = '[
  {"name": "Vinos Tintos", "slug": "vinos-tintos", "level": 0},
  {"name": "Vinos Blancos", "slug": "vinos-blancos", "level": 0},
  {"name": "Vinos Rosados", "slug": "vinos-rosados", "level": 0},
  {"name": "Espumantes", "slug": "espumantes", "level": 0},
  {"name": "Whisky", "slug": "whisky", "level": 0},
  {"name": "Gin", "slug": "gin", "level": 0},
  {"name": "Vodka", "slug": "vodka", "level": 0},
  {"name": "Rum y Brandy", "slug": "rum-brandy", "level": 0},
  {"name": "Aperitivos y Licores", "slug": "aperitivos-licores", "level": 0},
  {"name": "Vermouths", "slug": "vermouths-aperitivos", "level": 0},
  {"name": "Cervezas Artesanales", "slug": "cervezas-artesanales", "level": 0},
  {"name": "Cervezas Importadas", "slug": "cervezas-importadas", "level": 0},
  {"name": "Accesorios Vino", "slug": "accesorios-vino", "level": 0}
]'::jsonb
WHERE id = '3c546ca9-0be2-4a30-97bb-f917dd6b251f';

-- También verificar productos de kiosco, piletas, ropa, veterinaria, vinoteca con quality_score >= 40
UPDATE global_products SET is_verified=true
WHERE is_active=true AND is_verified=false AND quality_score >= 40
  AND business_type IN ('kiosco','piletas','ropa','veterinaria','vinoteca','bazar','jugueteria','libreria','electrodomesticos','fiambreria','perfumeria','carniceria','verduleria');
