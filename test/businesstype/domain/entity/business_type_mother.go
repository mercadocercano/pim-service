package entity

import (
	"time"

	"saas-mt-pim-service/src/businesstype/domain/entity"

	"github.com/google/uuid"
)

// BusinessTypeMother implementa el patrón Object Mother para crear entities BusinessType de prueba
type BusinessTypeMother struct{}

// WithDefaults crea un tipo de negocio con valores por defecto
func (BusinessTypeMother) WithDefaults() *entity.BusinessType {
	now := time.Now()
	return &entity.BusinessType{
		ID:          uuid.New().String(),
		Code:        "DEFAULT",
		Name:        "Tipo de Negocio de Prueba",
		Description: "Descripción del tipo de negocio de prueba",
		Icon:        "store",
		Color:       "#007bff",
		IsActive:    true,
		SortOrder:   1,
		Metadata:    make(map[string]interface{}),
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

// WithID crea un tipo de negocio con un ID específico
func (bt BusinessTypeMother) WithID(id string) *entity.BusinessType {
	businessType := bt.WithDefaults()
	businessType.ID = id
	return businessType
}

// WithCode crea un tipo de negocio con un código específico
func (bt BusinessTypeMother) WithCode(code string) *entity.BusinessType {
	businessType := bt.WithDefaults()
	businessType.Code = code
	return businessType
}

// WithName crea un tipo de negocio con un nombre específico
func (bt BusinessTypeMother) WithName(name string) *entity.BusinessType {
	businessType := bt.WithDefaults()
	businessType.Name = name
	return businessType
}

// WithDescription crea un tipo de negocio con una descripción específica
func (bt BusinessTypeMother) WithDescription(description string) *entity.BusinessType {
	businessType := bt.WithDefaults()
	businessType.Description = description
	return businessType
}

// WithIcon crea un tipo de negocio con un icono específico
func (bt BusinessTypeMother) WithIcon(icon string) *entity.BusinessType {
	businessType := bt.WithDefaults()
	businessType.Icon = icon
	return businessType
}

// WithColor crea un tipo de negocio con un color específico
func (bt BusinessTypeMother) WithColor(color string) *entity.BusinessType {
	businessType := bt.WithDefaults()
	businessType.Color = color
	return businessType
}

// WithSortOrder crea un tipo de negocio con un orden específico
func (bt BusinessTypeMother) WithSortOrder(sortOrder int) *entity.BusinessType {
	businessType := bt.WithDefaults()
	businessType.SortOrder = sortOrder
	return businessType
}

// Inactive crea un tipo de negocio inactivo
func (bt BusinessTypeMother) Inactive() *entity.BusinessType {
	businessType := bt.WithDefaults()
	businessType.IsActive = false
	return businessType
}

// RestaurantBusinessType crea un tipo de negocio para restaurante
func (bt BusinessTypeMother) RestaurantBusinessType() *entity.BusinessType {
	businessType := bt.WithDefaults()
	businessType.Code = "RESTAURANT"
	businessType.Name = "Restaurante"
	businessType.Description = "Negocio de comida y bebidas"
	businessType.Icon = "restaurant"
	businessType.Color = "#ff5722"
	return businessType
}

// RetailBusinessType crea un tipo de negocio para retail
func (bt BusinessTypeMother) RetailBusinessType() *entity.BusinessType {
	businessType := bt.WithDefaults()
	businessType.Code = "RETAIL"
	businessType.Name = "Retail"
	businessType.Description = "Venta al por menor"
	businessType.Icon = "shopping_cart"
	businessType.Color = "#2196f3"
	return businessType
}

// ElectronicsBusinessType crea un tipo de negocio para electrónicos
func (bt BusinessTypeMother) ElectronicsBusinessType() *entity.BusinessType {
	businessType := bt.WithDefaults()
	businessType.Code = "ELECTRONICS"
	businessType.Name = "Electrónicos"
	businessType.Description = "Venta de productos electrónicos"
	businessType.Icon = "devices"
	businessType.Color = "#9c27b0"
	return businessType
}

// ClothingBusinessType crea un tipo de negocio para ropa
func (bt BusinessTypeMother) ClothingBusinessType() *entity.BusinessType {
	businessType := bt.WithDefaults()
	businessType.Code = "CLOTHING"
	businessType.Name = "Ropa y Moda"
	businessType.Description = "Venta de ropa y accesorios"
	businessType.Icon = "checkroom"
	businessType.Color = "#e91e63"
	return businessType
}

// BookstoreBusinessType crea un tipo de negocio para librería
func (bt BusinessTypeMother) BookstoreBusinessType() *entity.BusinessType {
	businessType := bt.WithDefaults()
	businessType.Code = "BOOKSTORE"
	businessType.Name = "Librería"
	businessType.Description = "Venta de libros y material educativo"
	businessType.Icon = "menu_book"
	businessType.Color = "#795548"
	return businessType
}

// PharmacyBusinessType crea un tipo de negocio para farmacia
func (bt BusinessTypeMother) PharmacyBusinessType() *entity.BusinessType {
	businessType := bt.WithDefaults()
	businessType.Code = "PHARMACY"
	businessType.Name = "Farmacia"
	businessType.Description = "Venta de medicamentos y productos de salud"
	businessType.Icon = "local_pharmacy"
	businessType.Color = "#4caf50"
	return businessType
}

// SupermarketBusinessType crea un tipo de negocio para supermercado
func (bt BusinessTypeMother) SupermarketBusinessType() *entity.BusinessType {
	businessType := bt.WithDefaults()
	businessType.Code = "SUPERMARKET"
	businessType.Name = "Supermercado"
	businessType.Description = "Venta de productos alimentarios y de consumo"
	businessType.Icon = "local_grocery_store"
	businessType.Color = "#ff9800"
	return businessType
}

// FurnitureBusinessType crea un tipo de negocio para muebles
func (bt BusinessTypeMother) FurnitureBusinessType() *entity.BusinessType {
	businessType := bt.WithDefaults()
	businessType.Code = "FURNITURE"
	businessType.Name = "Muebles"
	businessType.Description = "Venta de muebles y decoración"
	businessType.Icon = "chair"
	businessType.Color = "#607d8b"
	return businessType
}

// Complete crea un tipo de negocio con todos los parámetros especificados
func (BusinessTypeMother) Complete(
	id, code, name, description, icon, color string,
	isActive bool,
	sortOrder int,
) *entity.BusinessType {
	now := time.Now()
	return &entity.BusinessType{
		ID:          id,
		Code:        code,
		Name:        name,
		Description: description,
		Icon:        icon,
		Color:       color,
		IsActive:    isActive,
		SortOrder:   sortOrder,
		Metadata:    make(map[string]interface{}),
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

// Create retorna una nueva instancia de BusinessTypeMother
func Create() BusinessTypeMother {
	return BusinessTypeMother{}
}
