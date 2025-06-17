-- Seed: 004_business_type_quickstart_templates.sql
-- Purpose: Poblar templates de quickstart vinculando business_types con marketplace_categories relevantes
-- Date: 2025-01-27
-- Beneficio: Facilitar onboarding de sellers con categorías pre-sugeridas

-- Limpiar templates existentes
DELETE FROM business_type_templates;

-- Función auxiliar para obtener IDs de categorías por slug
CREATE OR REPLACE FUNCTION get_category_ids(category_slugs TEXT[])
RETURNS JSONB AS $$
DECLARE
    result JSONB := '[]';
    category_id UUID;
    current_slug TEXT;
BEGIN
    FOREACH current_slug IN ARRAY category_slugs
    LOOP
        SELECT id INTO category_id 
        FROM marketplace_categories 
        WHERE slug = current_slug;
        
        IF category_id IS NOT NULL THEN
            result := result || jsonb_build_object('id', category_id, 'slug', current_slug);
        END IF;
    END LOOP;
    
    RETURN result;
END;
$$ LANGUAGE plpgsql;

-- Poblar templates de quickstart para cada tipo de negocio
DO $$
DECLARE
    bt_record RECORD;
BEGIN
    FOR bt_record IN SELECT id, code, name FROM business_types ORDER BY sort_order LOOP
        
        -- COMERCIOS ALIMENTARIOS
        IF bt_record.code = 'almacen' THEN
            INSERT INTO business_type_templates (business_type_id, name, description, region, categories, is_default)
            VALUES (bt_record.id, 'Almacén Argentino', 'Configuración típica para almacén de barrio', 'AR', 
                   get_category_ids(ARRAY['alimentos-bebidas', 'alimentos-frescos', 'bebidas', 'productos-secos', 'lacteos', 'limpieza']), true);
                   
        ELSIF bt_record.code = 'supermercado' THEN
            INSERT INTO business_type_templates (business_type_id, name, description, region, categories, is_default)
            VALUES (bt_record.id, 'Supermercado Completo', 'Amplia variedad de productos para supermercado', 'AR',
                   get_category_ids(ARRAY['alimentos-bebidas', 'alimentos-frescos', 'bebidas', 'productos-secos', 'lacteos', 'limpieza', 'belleza-cuidado-personal', 'hogar-decoracion']), true);
                   
        ELSIF bt_record.code = 'carniceria' THEN
            INSERT INTO business_type_templates (business_type_id, name, description, region, categories, is_default)
            VALUES (bt_record.id, 'Carnicería Tradicional', 'Especializada en carnes y productos cárnicos', 'AR',
                   get_category_ids(ARRAY['alimentos-bebidas', 'alimentos-frescos', 'lacteos']), true);
                   
        ELSIF bt_record.code = 'panaderia' THEN
            INSERT INTO business_type_templates (business_type_id, name, description, region, categories, is_default)
            VALUES (bt_record.id, 'Panadería y Confitería', 'Productos de panadería y confitería', 'AR',
                   get_category_ids(ARRAY['alimentos-bebidas', 'panaderia-confiteria']), true);
                   
        ELSIF bt_record.code = 'verduleria' THEN
            INSERT INTO business_type_templates (business_type_id, name, description, region, categories, is_default)
            VALUES (bt_record.id, 'Verdulería Fresca', 'Frutas y verduras de estación', 'AR',
                   get_category_ids(ARRAY['alimentos-bebidas', 'alimentos-frescos']), true);
                   
        ELSIF bt_record.code = 'fiambreria' THEN
            INSERT INTO business_type_templates (business_type_id, name, description, region, categories, is_default)
            VALUES (bt_record.id, 'Fiambrería Gourmet', 'Fiambres, quesos y productos delicatessen', 'AR',
                   get_category_ids(ARRAY['alimentos-bebidas', 'alimentos-frescos', 'lacteos']), true);
                   
        ELSIF bt_record.code = 'heladeria' THEN
            INSERT INTO business_type_templates (business_type_id, name, description, region, categories, is_default)
            VALUES (bt_record.id, 'Heladería Artesanal', 'Helados y postres helados', 'AR',
                   get_category_ids(ARRAY['alimentos-bebidas', 'panaderia-confiteria']), true);
                   
        -- SALUD Y FARMACIA
        ELSIF bt_record.code = 'farmacia' THEN
            INSERT INTO business_type_templates (business_type_id, name, description, region, categories, is_default)
            VALUES (bt_record.id, 'Farmacia Completa', 'Medicamentos y productos de cuidado personal', 'AR',
                   get_category_ids(ARRAY['salud-farmacia', 'medicamentos', 'equipos-medicos', 'primeros-auxilios', 'belleza-cuidado-personal']), true);
                   
        ELSIF bt_record.code = 'perfumeria' THEN
            INSERT INTO business_type_templates (business_type_id, name, description, region, categories, is_default)
            VALUES (bt_record.id, 'Perfumería y Cosmética', 'Perfumes, cosméticos y cuidado personal', 'AR',
                   get_category_ids(ARRAY['belleza-cuidado-personal', 'cosmeticos', 'perfumes', 'cuidado-facial', 'cuidado-corporal']), true);
                   
        -- MODA E INDUMENTARIA
        ELSIF bt_record.code = 'ropa' THEN
            INSERT INTO business_type_templates (business_type_id, name, description, region, categories, is_default)
            VALUES (bt_record.id, 'Tienda de Ropa', 'Indumentaria para toda la familia', 'AR',
                   get_category_ids(ARRAY['moda-indumentaria', 'ropa-mujer', 'ropa-hombre', 'ropa-infantil', 'accesorios-moda']), true);
                   
        ELSIF bt_record.code = 'zapateria' THEN
            INSERT INTO business_type_templates (business_type_id, name, description, region, categories, is_default)
            VALUES (bt_record.id, 'Zapatería Familiar', 'Calzado para toda la familia', 'AR',
                   get_category_ids(ARRAY['moda-indumentaria', 'calzado']), true);
                   
        ELSIF bt_record.code = 'deportes' THEN
            INSERT INTO business_type_templates (business_type_id, name, description, region, categories, is_default)
            VALUES (bt_record.id, 'Artículos Deportivos', 'Equipamiento y ropa deportiva', 'AR',
                   get_category_ids(ARRAY['deportes-fitness', 'ropa-deportiva', 'futbol', 'fitness']), true);
                   
        -- HOGAR Y CONSTRUCCIÓN
        ELSIF bt_record.code = 'ferreteria' THEN
            INSERT INTO business_type_templates (business_type_id, name, description, region, categories, is_default)
            VALUES (bt_record.id, 'Ferretería Industrial', 'Herramientas y materiales de construcción', 'AR',
                   get_category_ids(ARRAY['herramientas-construccion', 'herramientas-manuales', 'herramientas-electricas', 'materiales-construccion', 'pintura', 'electricidad', 'plomeria']), true);
                   
        ELSIF bt_record.code = 'muebleria' THEN
            INSERT INTO business_type_templates (business_type_id, name, description, region, categories, is_default)
            VALUES (bt_record.id, 'Mueblería del Hogar', 'Muebles y decoración para el hogar', 'AR',
                   get_category_ids(ARRAY['hogar-decoracion', 'muebles', 'decoracion', 'textiles-hogar']), true);
                   
        ELSIF bt_record.code = 'bazar' THEN
            INSERT INTO business_type_templates (business_type_id, name, description, region, categories, is_default)
            VALUES (bt_record.id, 'Bazar del Hogar', 'Artículos y utensilios para el hogar', 'AR',
                   get_category_ids(ARRAY['hogar-decoracion', 'cocina-comedor', 'decoracion', 'organizacion']), true);
                   
        -- TECNOLOGÍA
        ELSIF bt_record.code = 'electronica' THEN
            INSERT INTO business_type_templates (business_type_id, name, description, region, categories, is_default)
            VALUES (bt_record.id, 'Casa de Electrodomésticos', 'Electrodomésticos grandes y pequeños', 'AR',
                   get_category_ids(ARRAY['electrodomesticos', 'tecnologia-electronicos', 'audio-imagen']), true);
                   
        ELSIF bt_record.code = 'celulares' THEN
            INSERT INTO business_type_templates (business_type_id, name, description, region, categories, is_default)
            VALUES (bt_record.id, 'Casa de Celulares', 'Smartphones y accesorios móviles', 'AR',
                   get_category_ids(ARRAY['tecnologia-electronicos', 'celulares-telefonia', 'accesorios-tecnologicos']), true);
                   
        ELSIF bt_record.code = 'computacion' THEN
            INSERT INTO business_type_templates (business_type_id, name, description, region, categories, is_default)
            VALUES (bt_record.id, 'Computación y Sistemas', 'Computadoras, notebooks y periféricos', 'AR',
                   get_category_ids(ARRAY['tecnologia-electronicos', 'computacion', 'accesorios-tecnologicos']), true);
                   
        -- AUTOMOTRIZ
        ELSIF bt_record.code = 'repuestos' THEN
            INSERT INTO business_type_templates (business_type_id, name, description, region, categories, is_default)
            VALUES (bt_record.id, 'Casa de Repuestos', 'Repuestos y accesorios automotrices', 'AR',
                   get_category_ids(ARRAY['automotriz', 'repuestos', 'accesorios-exterior', 'accesorios-interior']), true);
                   
        ELSIF bt_record.code = 'lubricentro' THEN
            INSERT INTO business_type_templates (business_type_id, name, description, region, categories, is_default)
            VALUES (bt_record.id, 'Lubricentro Completo', 'Aceites, lubricantes y mantenimiento', 'AR',
                   get_category_ids(ARRAY['automotriz', 'aceites-lubricantes', 'repuestos']), true);
                   
        -- SERVICIOS ESPECIALIZADOS
        ELSIF bt_record.code = 'optica' THEN
            INSERT INTO business_type_templates (business_type_id, name, description, region, categories, is_default)
            VALUES (bt_record.id, 'Óptica Profesional', 'Anteojos, lentes y servicios ópticos', 'AR',
                   get_category_ids(ARRAY['salud-farmacia', 'equipos-medicos']), true);
                   
        ELSIF bt_record.code = 'relojeria' THEN
            INSERT INTO business_type_templates (business_type_id, name, description, region, categories, is_default)
            VALUES (bt_record.id, 'Relojería y Joyería', 'Relojes, joyas y accesorios', 'AR',
                   get_category_ids(ARRAY['moda-indumentaria', 'joyeria-relojes']), true);
                   
        ELSIF bt_record.code = 'libreria' THEN
            INSERT INTO business_type_templates (business_type_id, name, description, region, categories, is_default)
            VALUES (bt_record.id, 'Librería y Papelería', 'Libros, útiles escolares y papelería', 'AR',
                   get_category_ids(ARRAY['libros-entretenimiento', 'libros', 'papeleria']), true);
                   
        ELSIF bt_record.code = 'jugueteria' THEN
            INSERT INTO business_type_templates (business_type_id, name, description, region, categories, is_default)
            VALUES (bt_record.id, 'Juguetería Infantil', 'Juguetes y juegos para todas las edades', 'AR',
                   get_category_ids(ARRAY['libros-entretenimiento', 'juguetes', 'juegos-mesa']), true);
                   
        -- MASCOTAS
        ELSIF bt_record.code = 'veterinaria' THEN
            INSERT INTO business_type_templates (business_type_id, name, description, region, categories, is_default)
            VALUES (bt_record.id, 'Veterinaria y Pet Shop', 'Productos y servicios para mascotas', 'AR',
                   get_category_ids(ARRAY['mascotas', 'alimento-mascotas', 'accesorios-perros', 'accesorios-gatos', 'salud-veterinaria']), true);
                   
        -- SERVICIOS DIVERSOS
        ELSIF bt_record.code = 'kiosco' THEN
            INSERT INTO business_type_templates (business_type_id, name, description, region, categories, is_default)
            VALUES (bt_record.id, 'Kiosco Completo', 'Golosinas, bebidas y productos de impulso', 'AR',
                   get_category_ids(ARRAY['alimentos-bebidas', 'bebidas', 'productos-secos', 'libros-entretenimiento']), true);
                   
        ELSIF bt_record.code = 'floreria' THEN
            INSERT INTO business_type_templates (business_type_id, name, description, region, categories, is_default)
            VALUES (bt_record.id, 'Florería y Plantas', 'Flores, plantas y accesorios de jardín', 'AR',
                   get_category_ids(ARRAY['hogar-decoracion', 'jardin-exterior', 'decoracion']), true);
                   
        ELSIF bt_record.code = 'polirubro' THEN
            INSERT INTO business_type_templates (business_type_id, name, description, region, categories, is_default)
            VALUES (bt_record.id, 'Polirubro Completo', 'Productos diversos para múltiples necesidades', 'AR',
                   get_category_ids(ARRAY['alimentos-bebidas', 'hogar-decoracion', 'belleza-cuidado-personal', 'libros-entretenimiento', 'limpieza']), true);
                   
        -- SERVICIOS DIGITALES Y MODERNOS
        ELSIF bt_record.code = 'agencia_viajes' THEN
            INSERT INTO business_type_templates (business_type_id, name, description, region, categories, is_default)
            VALUES (bt_record.id, 'Agencia de Viajes', 'Paquetes turísticos y servicios de viaje', 'AR',
                   get_category_ids(ARRAY['turismo-viajes']), true);
                   
        ELSIF bt_record.code = 'delivery' THEN
            INSERT INTO business_type_templates (business_type_id, name, description, region, categories, is_default)
            VALUES (bt_record.id, 'Delivery de Comidas', 'Servicio de entrega de comidas', 'AR',
                   get_category_ids(ARRAY['alimentos-bebidas', 'alimentos-frescos', 'bebidas']), true);
                   
        ELSIF bt_record.code = 'servicios_digitales' THEN
            INSERT INTO business_type_templates (business_type_id, name, description, region, categories, is_default)
            VALUES (bt_record.id, 'Servicios Digitales', 'Software, apps y contenido digital', 'AR',
                   get_category_ids(ARRAY['servicios-digitales']), true);
                   
        ELSIF bt_record.code = 'centro_educativo' THEN
            INSERT INTO business_type_templates (business_type_id, name, description, region, categories, is_default)
            VALUES (bt_record.id, 'Centro Educativo', 'Cursos, capacitaciones y servicios educativos', 'AR',
                   get_category_ids(ARRAY['servicios-digitales', 'libros-entretenimiento', 'libros']), true);
                   
        ELSIF bt_record.code = 'entretenimiento' THEN
            INSERT INTO business_type_templates (business_type_id, name, description, region, categories, is_default)
            VALUES (bt_record.id, 'Entretenimiento', 'Espectáculos, eventos y entretenimiento', 'AR',
                   get_category_ids(ARRAY['libros-entretenimiento', 'servicios-digitales']), true);
                   
        ELSIF bt_record.code = 'lavadero' THEN
            INSERT INTO business_type_templates (business_type_id, name, description, region, categories, is_default)
            VALUES (bt_record.id, 'Lavadero de Autos', 'Servicios de lavado y detailing automotriz', 'AR',
                   get_category_ids(ARRAY['automotriz', 'accesorios-exterior', 'accesorios-interior']), true);
                   
        ELSIF bt_record.code = 'servicios_profesionales' THEN
            INSERT INTO business_type_templates (business_type_id, name, description, region, categories, is_default)
            VALUES (bt_record.id, 'Servicios Profesionales', 'Consultoría y servicios profesionales', 'AR',
                   get_category_ids(ARRAY['servicios-digitales']), true);
                   
        END IF;
    END LOOP;
END $$;

-- Limpiar función auxiliar
DROP FUNCTION get_category_ids(TEXT[]);

-- Comentarios
COMMENT ON TABLE business_type_templates IS 'Templates de quickstart que vinculan tipos de negocio con categorías sugeridas para facilitar el onboarding';

-- Verificar resultados
SELECT 
    bt.name as tipo_negocio,
    btt.name as template_name,
    jsonb_array_length(btt.categories) as num_categorias
FROM business_types bt
JOIN business_type_templates btt ON bt.id = btt.business_type_id
ORDER BY bt.sort_order; 