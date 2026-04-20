package integration_test

// TODO(cycle-002-global-brands-colors/T007): Implementar tests de integración reales cuando
// haya un entorno de test con PostgreSQL disponible (ej. testcontainers o docker-compose de CI).
//
// Escenarios a cubrir:
//   1. marca con background_color, text_color, typography en DB → GET devuelve los strings hex correctos.
//   2. marca con NULL en esas columnas → GET devuelve "" en los 3 campos (backward compatible).
//   3. PUT con color inválido (#ZZZ) → 400 Bad Request.
//   4. consumidores actuales siguen funcionando: campos existentes no se rompen.
//
// Criterios de aceptación del endpoint GET /pim/api/v1/marketplace-brands:
//   - JSON response incluye siempre background_color, text_color, typography (nunca omitidos).
//   - NULL en DB → "" en JSON (no null, no campo ausente).
//   - Hex válido en DB → mismo hex en JSON.
//
// Verificación mínima de compilación — el build de los paquetes depende de este archivo
// como prueba de que las importaciones son correctas.

import (
	"saas-mt-pim-service/src/brand/application/response"
	"saas-mt-pim-service/src/brand/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestMarketplaceBrand_VisualIdentityDTO_NullDbToEmptyString verifica que el mapper
// convierte correctamente los campos vacíos de la entidad al DTO de respuesta.
// Este test no necesita DB — valida la capa de mapeo en aislamiento.
func TestMarketplaceBrand_VisualIdentityDTO_NullDbToEmptyString(t *testing.T) {
	// Arrange — entidad sin colores (como vendría de DB con columnas NULL → "" en scan)
	brand, err := entity.NewMarketplacebrand("MarcaSinColor")
	require.NoError(t, err)
	// BackgroundColor, TextColor, Typography ya son "" por defecto

	// Act
	dto := response.NewMarketplaceBrandResponse(brand)

	// Assert — los campos deben existir en el DTO y ser ""
	assert.Equal(t, "", dto.BackgroundColor, "NULL en DB debe mapearse a string vacío, no omitirse")
	assert.Equal(t, "", dto.TextColor, "NULL en DB debe mapearse a string vacío, no omitirse")
	assert.Equal(t, "", dto.Typography, "NULL en DB debe mapearse a string vacío, no omitirse")
}

// TestMarketplaceBrand_VisualIdentityDTO_WithColors verifica que los colores válidos
// se propagan correctamente del dominio al DTO.
func TestMarketplaceBrand_VisualIdentityDTO_WithColors(t *testing.T) {
	// Arrange
	brand, err := entity.NewMarketplacebrand("CocaCola")
	require.NoError(t, err)

	err = brand.SetVisualIdentity(entity.VisualIdentityParams{
		BackgroundColor: "#FF0000",
		TextColor:       "#FFFFFF",
		Typography:      "Oswald",
	})
	require.NoError(t, err)

	// Act
	dto := response.NewMarketplaceBrandResponse(brand)

	// Assert
	assert.Equal(t, "#FF0000", dto.BackgroundColor)
	assert.Equal(t, "#FFFFFF", dto.TextColor)
	assert.Equal(t, "Oswald", dto.Typography)
}

// TestMarketplaceBrand_VisualIdentityDTO_ExistingFieldsNotBroken verifica backward compatibility:
// los campos previos del DTO no se ven afectados por los campos nuevos.
func TestMarketplaceBrand_VisualIdentityDTO_ExistingFieldsNotBroken(t *testing.T) {
	// Arrange
	brand, err := entity.NewMarketplacebrand("Nestlé")
	require.NoError(t, err)
	brand.Description = "Multinacional alimentaria"
	brand.IsActive = true

	// Act
	dto := response.NewMarketplaceBrandResponse(brand)

	// Assert — campos existentes deben seguir funcionando
	assert.Equal(t, "Nestlé", dto.Name)
	assert.Equal(t, "Multinacional alimentaria", dto.Description)
	assert.True(t, dto.IsActive)
	assert.NotEmpty(t, dto.CreatedAt)
	assert.NotEmpty(t, dto.UpdatedAt)
}
