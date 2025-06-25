package entity

// AttributeType define los tipos de atributos disponibles
type AttributeType string

const (
	AttributeTypeText        AttributeType = "text"
	AttributeTypeString      AttributeType = "string"
	AttributeTypeNumber      AttributeType = "number"
	AttributeTypeBoolean     AttributeType = "boolean"
	AttributeTypeDate        AttributeType = "date"
	AttributeTypeEnum        AttributeType = "enum"
	AttributeTypeSelect      AttributeType = "select"
	AttributeTypeMultiSelect AttributeType = "multi_select"
)

// IsValidAttributeType verifica si un tipo de atributo es válido
func IsValidAttributeType(attrType AttributeType) bool {
	validTypes := []AttributeType{
		AttributeTypeText,
		AttributeTypeString,
		AttributeTypeNumber,
		AttributeTypeBoolean,
		AttributeTypeDate,
		AttributeTypeEnum,
		AttributeTypeSelect,
		AttributeTypeMultiSelect,
	}

	for _, validType := range validTypes {
		if attrType == validType {
			return true
		}
	}
	return false
}

// ToString convierte AttributeType a string
func (at AttributeType) ToString() string {
	return string(at)
}

// FromString convierte string a AttributeType
func FromString(s string) AttributeType {
	return AttributeType(s)
}
