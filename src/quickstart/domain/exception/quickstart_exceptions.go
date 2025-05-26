package exception

import "errors"

// Errores específicos del dominio quickstart
var (
	// ErrBusinessTypeNotFound se produce cuando no se encuentra un tipo de negocio
	ErrBusinessTypeNotFound = errors.New("tipo de negocio no encontrado")

	// ErrInvalidSetupData se produce cuando los datos de configuración son inválidos
	ErrInvalidSetupData = errors.New("datos de configuración inválidos")

	// ErrTemplateNotFound se produce cuando no se encuentra una plantilla
	ErrTemplateNotFound = errors.New("plantilla no encontrada")

	// ErrQuickstartAlreadyCompleted se produce cuando el quickstart ya fue completado para un tenant
	ErrQuickstartAlreadyCompleted = errors.New("el quickstart ya fue completado para este tenant")

	// ErrQuickstartInProgress se produce cuando hay un quickstart en progreso para un tenant
	ErrQuickstartInProgress = errors.New("hay un quickstart en progreso para este tenant")

	// ErrYamlFileNotFound se produce cuando no se encuentra un archivo YAML
	ErrYamlFileNotFound = errors.New("archivo YAML no encontrado")

	// ErrInvalidYamlFormat se produce cuando el formato del archivo YAML es inválido
	ErrInvalidYamlFormat = errors.New("formato de archivo YAML inválido")

	// ErrSetupFailed se produce cuando falla la configuración del tenant
	ErrSetupFailed = errors.New("falló la configuración del tenant")

	// ErrTenantNotFound se produce cuando no se encuentra el tenant
	ErrTenantNotFound = errors.New("tenant no encontrado")

	// ErrInvalidBusinessTypeFormat se produce cuando el formato del tipo de negocio es inválido
	ErrInvalidBusinessTypeFormat = errors.New("formato de tipo de negocio inválido")
)
