package service

import (
	"strings"
)

// CategoryDeductionService deduces product categories from product names
// using a keyword dictionary and tenant's existing categories.
type CategoryDeductionService struct {
	keywordMap map[string]string // keyword → category name
}

func NewCategoryDeductionService() *CategoryDeductionService {
	return &CategoryDeductionService{
		keywordMap: buildKeywordDictionary(),
	}
}

// DeduceCategory returns the best category match for a product name.
// It first checks tenant categories (if provided), then the keyword dictionary.
func (s *CategoryDeductionService) DeduceCategory(productName string, tenantCategories []string) string {
	lower := strings.ToLower(productName)

	// Priority 1: match against tenant's existing categories
	for _, cat := range tenantCategories {
		if strings.Contains(lower, strings.ToLower(cat)) {
			return cat
		}
	}

	// Priority 2: keyword dictionary
	bestCategory := ""
	bestLen := 0
	for keyword, category := range s.keywordMap {
		if strings.Contains(lower, keyword) && len(keyword) > bestLen {
			bestCategory = category
			bestLen = len(keyword)
		}
	}

	return bestCategory
}

func buildKeywordDictionary() map[string]string {
	entries := map[string][]string{
		"Bebidas": {
			"coca cola", "pepsi", "fanta", "sprite", "agua mineral", "gaseosa",
			"jugo", "cerveza", "vino", "soda", "energizante", "red bull",
			"powerade", "gatorade", "terma", "mate cocido", "té", "cafe",
		},
		"Golosinas": {
			"alfajor", "chocolate", "caramelo", "chicle", "gomita", "chupetin",
			"rocklet", "beldent", "bon o bon", "shot", "mantecol", "turron",
		},
		"Galletitas": {
			"galletita", "galleta", "oreo", "toddy", "pepito", "traviata",
			"criollita", "express", "cerealita",
		},
		"Snacks": {
			"papa", "lays", "doritos", "cheetos", "palitos", "chizito",
			"mani", "nachos", "pochoclo",
		},
		"Lácteos": {
			"leche", "yogur", "yogurt", "queso cremoso", "crema",
			"manteca", "dulce de leche", "serenisima", "sancor",
		},
		"Fiambrería": {
			"jamon", "queso", "salame", "mortadela", "paladini",
			"fiambre", "salchicha", "chorizo",
		},
		"Panadería": {
			"pan lactal", "pan de", "bimbo", "medialunas", "facturas",
			"tostada", "bizcocho",
		},
		"Limpieza": {
			"lavandina", "detergente", "desengrasante", "limpiador",
			"trapo", "esponja", "jabon liquido", "suavizante",
			"magistral", "skip", "ala", "papel higienico", "servilleta",
		},
		"Higiene": {
			"shampoo", "jabon", "desodorante", "cepillo dental",
			"pasta dental", "head & shoulders", "dove", "rexona",
			"crema de enjuague", "acondicionador", "pañal",
		},
		"Alimentos": {
			"arroz", "fideo", "aceite", "azucar", "sal", "harina",
			"atun", "sardina", "puré", "tomate", "mermelada", "miel",
			"caldo", "polenta", "lenteja",
		},
		"Ferretería": {
			"tornillo", "clavo", "tuerca", "arandela", "tarugo",
			"brida", "grampa", "cinta aisladora",
		},
		"Herramientas Manuales": {
			"martillo", "destornillador", "pinza", "llave", "alicate",
			"sierra", "serrucho", "nivel", "cinta metrica",
		},
		"Herramientas Eléctricas": {
			"taladro", "amoladora", "atornillador", "caladora",
			"sierra circular", "lijadora", "soldadora", "compresor",
		},
		"Electricidad": {
			"cable", "llave termica", "disyuntor", "enchufe", "toma",
			"portalampara", "cano luz", "caño luz", "unipolar",
		},
		"Pinturas": {
			"pintura", "latex", "esmalte sintetico", "barniz", "sellador",
			"enduido", "rodillo", "pincel", "brocha",
		},
		"Librería": {
			"cuaderno", "lapiz", "lapicera", "birome", "carpeta",
			"cartulina", "resma", "marcador", "fibra", "goma de borrar",
		},
		"Bazar": {
			"plato", "vaso", "taza", "olla", "sarten", "jarra",
			"cubierto", "cuchillo", "tabla", "tupperware",
		},
		"Indumentaria": {
			"remera", "pantalon", "campera", "buzo", "camisa",
			"jean", "medias", "ropa interior", "boxer",
		},
		"Calzado": {
			"zapatilla", "zapato", "sandalia", "bota", "ojotas",
			"alpargata", "mocasin",
		},
		"Juguetería": {
			"juguete", "muñeca", "auto de juguete", "peluche",
			"rompecabezas", "puzzle", "lego", "pelota",
		},
		"Electrónica": {
			"auricular", "cargador", "cable usb", "pendrive",
			"parlante", "funda celular", "bateria",
		},
		"Mascotas": {
			"alimento perro", "alimento gato", "dog chow", "whiskas",
			"arena gato", "correa", "comedero",
		},
		"Verdulería": {
			"papa", "tomate", "cebolla", "zanahoria", "lechuga",
			"banana", "manzana", "naranja", "limon",
		},
	}

	result := make(map[string]string)
	for category, keywords := range entries {
		for _, kw := range keywords {
			result[kw] = category
		}
	}
	return result
}
