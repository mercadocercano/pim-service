package model

import (
	"time"

	"github.com/lib/pq"
)

// CategoryAttribute representa el modelo de base de datos para un atributo de categoría
type CategoryAttribute struct {
	ID            string         `db:"id"`
	TenantID      string         `db:"tenant_id"`
	CategoryID    string         `db:"category_id"`
	AttributeID   string         `db:"attribute_id"`
	AllowedValues pq.StringArray `db:"allowed_values"`
	Status        string         `db:"status"`
	CreatedAt     time.Time      `db:"created_at"`
	UpdatedAt     time.Time      `db:"updated_at"`
}
