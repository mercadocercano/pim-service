#!/usr/bin/env python3
"""
Script de migración de productos desde MongoDB (scraper_products) a PostgreSQL (global_products)
"""

import os
import sys
import json
import hashlib
from datetime import datetime
from pymongo import MongoClient
import psycopg2
from psycopg2.extras import RealDictCursor
import uuid

# Configuración de conexiones
MONGO_URI = os.getenv('MONGO_URI', 'mongodb://admin:admin123@localhost:27017/')
MONGO_DB = os.getenv('MONGO_DB', 'pim_marketplace')
POSTGRES_HOST = os.getenv('POSTGRES_HOST', 'localhost')
POSTGRES_PORT = os.getenv('POSTGRES_PORT', '5432')
POSTGRES_DB = os.getenv('POSTGRES_DB', 'pim_db')
POSTGRES_USER = os.getenv('POSTGRES_USER', 'postgres')
POSTGRES_PASSWORD = os.getenv('POSTGRES_PASSWORD', 'postgres')

# Lista expandida de marcas conocidas
KNOWN_BRANDS = [
    # Electronics
    'Samsung', 'LG', 'Sony', 'Philips', 'Panasonic', 'Sharp', 'Toshiba',
    'Apple', 'Motorola', 'Nokia', 'Xiaomi', 'Huawei', 'TCL', 'RCA',
    'BGH', 'Noblex', 'JVC', 'Pioneer', 'Hitachi',
    # Appliances
    'Whirlpool', 'Electrolux', 'Ariston', 'Drean', 'Patrick', 'Longvie',
    'Bosch', 'Liliana', 'Atma', 'Ultracomb', 'Oster', 'Black & Decker',
    # Food & Beverages
    'Coca Cola', 'Pepsi', 'Quilmes', 'Brahma', 'Stella Artois', 'Corona',
    'La Serenisima', 'Sancor', 'Arcor', 'Bagley', 'Terrabusi', 'Molinos',
    'Marolio', 'Molto', 'Knorr', 'Hellmanns', 'Natura',
    # Fashion & Sports
    'Nike', 'Adidas', 'Puma', 'Reebok', 'Topper', 'Fila', 'New Balance',
    'Under Armour', 'Converse', 'Vans', 'Levis', 'Wrangler', 'Lee',
    # Tools & Hardware
    'Stanley', 'Makita', 'DeWalt', 'Skil', 'Tramontina', 'Carol',
    'Durax', 'Colombraro', 'Bremen', 'Gamma', 'Lusqtoff',
    # Personal Care
    'Dove', 'Nivea', 'Colgate', 'Oral-B', 'Head & Shoulders', 'Pantene',
    'Sedal', 'Elvive', 'Rexona', 'Axe', 'Gillette', 'Ponds'
]

# Mapeo de categorías
CATEGORY_MAPPINGS = {
    'electronica': 'Electronics',
    'electronicos': 'Electronics',
    'tecnologia': 'Electronics',
    'celulares': 'Electronics > Mobile Phones',
    'smartphones': 'Electronics > Mobile Phones',
    'tv': 'Electronics > TVs & Video',
    'televisores': 'Electronics > TVs & Video',
    'audio': 'Electronics > Audio',
    'computacion': 'Electronics > Computers',
    'notebooks': 'Electronics > Computers > Laptops',
    'electrodomesticos': 'Home & Garden > Appliances',
    'cocina': 'Home & Garden > Kitchen',
    'heladeras': 'Home & Garden > Appliances > Refrigerators',
    'lavarropas': 'Home & Garden > Appliances > Washing Machines',
    'alimentos': 'Food & Beverages',
    'bebidas': 'Food & Beverages > Beverages',
    'almacen': 'Food & Beverages > Pantry',
    'cuidado personal': 'Health & Beauty > Personal Care',
    'perfumeria': 'Health & Beauty > Fragrances',
    'farmacia': 'Health & Beauty > Pharmacy',
    'moda': 'Fashion',
    'ropa': 'Fashion > Clothing',
    'calzado': 'Fashion > Shoes',
    'deportes': 'Sports & Outdoors',
    'construccion': 'Home & Garden > Construction',
    'herramientas': 'Home & Garden > Tools',
    'limpieza': 'Home & Garden > Cleaning',
    'mascotas': 'Pets',
    'juguetes': 'Toys & Games',
    'libreria': 'Books & Stationery',
    'automotor': 'Automotive',
    'bebe': 'Baby & Kids'
}

def extract_brand(name, description=''):
    """Extrae la marca del nombre del producto"""
    text = (name + ' ' + description).upper()
    
    # Buscar marcas conocidas
    for brand in KNOWN_BRANDS:
        if brand.upper() in text:
            return brand
    
    # Buscar patrones comunes
    import re
    brand_pattern = r'^([A-Z][A-Z0-9\s&]+)\s*[-:]\s*'
    match = re.match(brand_pattern, name)
    if match and len(match.group(1)) <= 30:
        return match.group(1).strip()
    
    # Primera palabra si es mayúscula
    first_word = name.split()[0] if name else ''
    if first_word and first_word == first_word.upper() and len(first_word) > 2:
        return first_word
    
    return 'Generic'

def normalize_category(category):
    """Normaliza la categoría del producto"""
    if not category:
        return 'Other'
    
    category_lower = category.lower()
    for key, value in CATEGORY_MAPPINGS.items():
        if key in category_lower:
            return value
    
    return 'Other'

def generate_ean13(seed):
    """Genera un EAN-13 válido basado en una semilla"""
    # Prefijos por categoría (Argentina usa 779)
    category_prefixes = {
        'Electronics': '779100',
        'Food & Beverages': '779200',
        'Health & Beauty': '779300',
        'Fashion': '779400',
        'Home & Garden': '779500',
        'Sports & Outdoors': '779600',
        'Other': '779900'
    }
    
    # Generar hash único
    hash_obj = hashlib.md5(seed.encode())
    hash_hex = hash_obj.hexdigest()
    
    # Convertir letras a números
    unique_part = ''
    for char in hash_hex[:6]:
        if char.isdigit():
            unique_part += char
        else:
            unique_part += str(ord(char) % 10)
    
    # Obtener prefijo según categoría
    prefix = '779900'  # Default
    
    # Combinar prefijo + parte única (12 dígitos)
    ean12 = prefix + unique_part
    
    # Calcular dígito de verificación
    total = 0
    for i in range(12):
        digit = int(ean12[i])
        if i % 2 == 0:
            total += digit
        else:
            total += digit * 3
    
    check_digit = (10 - (total % 10)) % 10
    
    return ean12 + str(check_digit)

def extract_price(price_str):
    """Extrae el valor numérico del precio"""
    if not price_str:
        return 0.0, 'ARS'
    
    import re
    price_str = str(price_str)
    
    # Detectar moneda
    currency = 'ARS'
    if 'U$S' in price_str or 'USD' in price_str:
        currency = 'USD'
    elif '€' in price_str or 'EUR' in price_str:
        currency = 'EUR'
    
    # Extraer números
    numbers = re.findall(r'[\d.,]+', price_str)
    if not numbers:
        return 0.0, currency
    
    # Convertir a float
    prices = []
    for num in numbers:
        try:
            # Remover puntos de miles y convertir coma decimal
            clean_num = num.replace('.', '').replace(',', '.')
            price = float(clean_num)
            if price > 0:
                prices.append(price)
        except:
            continue
    
    if prices:
        return max(prices), currency
    
    return 0.0, currency

def calculate_quality_score(product):
    """Calcula el puntaje de calidad del producto"""
    score = 0.5  # Base
    
    if product.get('brand') and product['brand'] != 'Generic':
        score += 0.1
    if product.get('ean') and not product.get('ean_generated'):
        score += 0.1
    if product.get('image_url') or product.get('image'):
        score += 0.1
    if product.get('description') and len(str(product.get('description', ''))) > 50:
        score += 0.1
    if product.get('price_value', 0) > 0:
        score += 0.1
    
    return min(score, 1.0)

def migrate_products(limit=None):
    """Migra productos de MongoDB a PostgreSQL"""
    # Conectar a MongoDB
    print("Conectando a MongoDB...")
    mongo_client = MongoClient(MONGO_URI)
    mongo_db = mongo_client[MONGO_DB]
    scraper_products = mongo_db['scraper_products']
    
    # Conectar a PostgreSQL
    print("Conectando a PostgreSQL...")
    pg_conn = psycopg2.connect(
        host=POSTGRES_HOST,
        port=POSTGRES_PORT,
        database=POSTGRES_DB,
        user=POSTGRES_USER,
        password=POSTGRES_PASSWORD
    )
    pg_cursor = pg_conn.cursor()
    
    # Query para productos no migrados
    query = {'migrated_to_pim': {'$ne': True}}
    products = scraper_products.find(query).limit(limit) if limit else scraper_products.find(query)
    
    migrated_count = 0
    error_count = 0
    
    print(f"Iniciando migración...")
    
    for product in products:
        try:
            # Extraer y normalizar datos
            name = product.get('title', product.get('name', ''))[:500]
            brand = extract_brand(name, product.get('description', ''))
            category = normalize_category(product.get('category', ''))
            # El precio puede venir como dict o string
            price_data = product.get('price')
            if isinstance(price_data, dict):
                price_value = price_data.get('amount', 0)
                currency = price_data.get('currency', 'ARS')
            else:
                price_value, currency = extract_price(price_data)
            
            # Generar EAN si no existe
            ean = product.get('ean') or product.get('barcode', '')
            ean_generated = False
            if not ean or len(ean) != 13:
                seed = str(product.get('_id', '')) or name
                ean = generate_ean13(seed)
                ean_generated = True
            
            # Generar SKU si no existe
            sku = product.get('sku', '')
            if not sku:
                brand_code = brand[:3].upper() if brand != 'Generic' else 'GEN'
                name_code = ''.join(c for c in name[:4].upper() if c.isalnum())[:4] or 'PROD'
                hash_code = hashlib.md5(str(product.get('_id', name)).encode()).hexdigest()[:6].upper()
                sku = f"{brand_code}-{name_code}-{hash_code}"
            
            # Preparar especificaciones
            specifications = {
                'weight': product.get('weight'),
                'dimensions': product.get('dimensions'),
                'color': product.get('color'),
                'material': product.get('material'),
                'warranty': product.get('warranty')
            }
            specifications = {k: v for k, v in specifications.items() if v}
            
            # Preparar imágenes
            images = []
            if product.get('images'):
                images = product['images']
            elif product.get('image_url'):
                images = [product['image_url']]
            elif product.get('image'):
                images = [product['image']]
            
            primary_image = images[0] if images else None
            
            # Calcular scores
            quality_score = calculate_quality_score({
                'brand': brand,
                'ean': ean,
                'ean_generated': ean_generated,
                'image_url': primary_image,
                'description': product.get('description'),
                'price_value': price_value
            })
            
            source_confidence = 0.7 if brand != 'Generic' else 0.4
            
            # Insert en PostgreSQL (adaptado a la estructura real)
            insert_query = """
                INSERT INTO global_products (
                    id, ean, name, description, brand, category,
                    price, image_url, image_urls, source, source_url,
                    source_reliability, quality_score, is_verified, is_active,
                    created_at, updated_at
                ) VALUES (
                    %s, %s, %s, %s, %s, %s,
                    %s, %s, %s, %s, %s,
                    %s, %s, %s, %s,
                    %s, %s
                )
                ON CONFLICT (ean) DO UPDATE SET
                    name = EXCLUDED.name,
                    brand = EXCLUDED.brand,
                    category = EXCLUDED.category,
                    description = EXCLUDED.description,
                    price = EXCLUDED.price,
                    image_url = EXCLUDED.image_url,
                    quality_score = EXCLUDED.quality_score,
                    updated_at = EXCLUDED.updated_at
            """
            
            # Convertir quality_score a entero (0-100)
            quality_score_int = int(quality_score * 100)
            
            pg_cursor.execute(insert_query, (
                str(uuid.uuid4()),  # id
                ean,  # ean
                name,  # name
                product.get('description'),  # description
                brand,  # brand
                category,  # category
                price_value if price_value > 0 else None,  # price
                primary_image,  # image_url
                images if images else None,  # image_urls
                product.get('source', 'unknown'),  # source
                product.get('url'),  # source_url
                source_confidence,  # source_reliability
                quality_score_int,  # quality_score (como entero)
                False,  # is_verified
                True,  # is_active
                datetime.now(),  # created_at
                datetime.now()  # updated_at
            ))
            
            # Marcar como migrado en MongoDB
            scraper_products.update_one(
                {'_id': product['_id']},
                {
                    '$set': {
                        'migrated_to_pim': True,
                        'migrated_at': datetime.now(),
                        'pim_ean': ean,
                        'pim_sku': sku
                    }
                }
            )
            
            migrated_count += 1
            
            if migrated_count % 100 == 0:
                pg_conn.commit()
                print(f"Migrados: {migrated_count} productos")
                
        except Exception as e:
            error_count += 1
            print(f"Error migrando producto {product.get('_id')}: {str(e)}")
            pg_conn.rollback()
    
    # Commit final
    pg_conn.commit()
    
    # Cerrar conexiones
    pg_cursor.close()
    pg_conn.close()
    mongo_client.close()
    
    print(f"\nMigración completada:")
    print(f"- Productos migrados: {migrated_count}")
    print(f"- Errores: {error_count}")
    
    return migrated_count, error_count

if __name__ == "__main__":
    # Verificar argumentos
    limit = None
    if len(sys.argv) > 1:
        try:
            limit = int(sys.argv[1])
            print(f"Limitando migración a {limit} productos")
        except:
            print("Uso: python migrate_scraper_to_global_products.py [limite]")
            sys.exit(1)
    
    # Ejecutar migración
    migrate_products(limit)