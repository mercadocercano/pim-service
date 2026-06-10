package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"saas-mt-pim-service/src/quickstart/domain/entity"
	"saas-mt-pim-service/src/quickstart/domain/port"
)

// QuickstartHistoryPostgresRepository implementa el repositorio usando PostgreSQL
type QuickstartHistoryPostgresRepository struct {
	db *sql.DB
}

// NewQuickstartHistoryPostgresRepository crea una nueva instancia del repositorio
func NewQuickstartHistoryPostgresRepository(db *sql.DB) port.QuickstartHistoryRepository {
	return &QuickstartHistoryPostgresRepository{
		db: db,
	}
}

// FindActiveByTenantID busca un wizard activo (no completado) para un tenant
func (r *QuickstartHistoryPostgresRepository) FindActiveByTenantID(ctx context.Context, tenantID string) (*entity.TenantQuickstartHistory, error) {
	query := `
		SELECT id, tenant_id, business_type_id, template_id, setup_completed, 
		       setup_data, created_at, updated_at
		FROM tenant_quickstart_history
		WHERE tenant_id = $1 AND setup_completed = false
		ORDER BY created_at DESC
		LIMIT 1
	`

	var history entity.TenantQuickstartHistory
	var templateID sql.NullString
	var setupDataBytes []byte

	err := r.db.QueryRowContext(ctx, query, tenantID).Scan(
		&history.ID,
		&history.TenantID,
		&history.BusinessTypeID,
		&templateID,
		&history.SetupCompleted,
		&setupDataBytes,
		&history.CreatedAt,
		&history.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("error al buscar wizard activo: %w", err)
	}

	if templateID.Valid {
		history.TemplateID = &templateID.String
	}

	history.SetupData = json.RawMessage(setupDataBytes)

	return &history, nil
}

// FindByID busca un registro por su ID
func (r *QuickstartHistoryPostgresRepository) FindByID(ctx context.Context, id string) (*entity.TenantQuickstartHistory, error) {
	query := `
		SELECT id, tenant_id, business_type_id, template_id, setup_completed, 
		       setup_data, created_at, updated_at
		FROM tenant_quickstart_history
		WHERE id = $1
	`

	var history entity.TenantQuickstartHistory
	var templateID sql.NullString
	var setupDataBytes []byte

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&history.ID,
		&history.TenantID,
		&history.BusinessTypeID,
		&templateID,
		&history.SetupCompleted,
		&setupDataBytes,
		&history.CreatedAt,
		&history.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("wizard no encontrado")
	}

	if err != nil {
		return nil, fmt.Errorf("error al buscar wizard: %w", err)
	}

	if templateID.Valid {
		history.TemplateID = &templateID.String
	}

	history.SetupData = json.RawMessage(setupDataBytes)

	return &history, nil
}

// Create crea un nuevo registro de historial
func (r *QuickstartHistoryPostgresRepository) Create(ctx context.Context, history *entity.TenantQuickstartHistory) error {
	query := `
		INSERT INTO tenant_quickstart_history (
			id, tenant_id, business_type_id, template_id, 
			setup_completed, setup_data, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	var templateID sql.NullString
	if history.TemplateID != nil {
		templateID = sql.NullString{String: *history.TemplateID, Valid: true}
	}

	_, err := r.db.ExecContext(ctx, query,
		history.ID,
		history.TenantID,
		history.BusinessTypeID,
		templateID,
		history.SetupCompleted,
		history.SetupData,
		history.CreatedAt,
		history.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("error al crear wizard: %w", err)
	}

	return nil
}

// Update actualiza un registro existente
func (r *QuickstartHistoryPostgresRepository) Update(ctx context.Context, history *entity.TenantQuickstartHistory) error {
	query := `
		UPDATE tenant_quickstart_history
		SET template_id = $2, setup_completed = $3, setup_data = $4, updated_at = $5
		WHERE id = $1
	`

	var templateID sql.NullString
	if history.TemplateID != nil {
		templateID = sql.NullString{String: *history.TemplateID, Valid: true}
	}

	result, err := r.db.ExecContext(ctx, query,
		history.ID,
		templateID,
		history.SetupCompleted,
		history.SetupData,
		time.Now(),
	)

	if err != nil {
		return fmt.Errorf("error al actualizar wizard: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar filas afectadas: %w", err)
	}

	if rowsAffected == 0 {
		return errors.New("wizard no encontrado")
	}

	return nil
}

// MarkAsCompleted marca un wizard como completado
func (r *QuickstartHistoryPostgresRepository) MarkAsCompleted(ctx context.Context, id string) error {
	query := `
		UPDATE tenant_quickstart_history
		SET setup_completed = true, updated_at = $2
		WHERE id = $1
	`

	fmt.Printf("🔍 DEBUG Repository: Ejecutando UPDATE para wizard ID: %s\n", id)
	now := time.Now()
	result, err := r.db.ExecContext(ctx, query, id, now)
	if err != nil {
		fmt.Printf("❌ ERROR Repository: Error en UPDATE: %v\n", err)
		return fmt.Errorf("error al marcar wizard como completado: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("❌ ERROR Repository: Error al verificar filas: %v\n", err)
		return fmt.Errorf("error al verificar filas afectadas: %w", err)
	}

	fmt.Printf("📊 DEBUG Repository: Filas afectadas: %d\n", rowsAffected)

	if rowsAffected == 0 {
		fmt.Printf("⚠️  WARNING Repository: No se encontró el wizard con ID: %s\n", id)
		return errors.New("wizard no encontrado")
	}

	fmt.Printf("✅ DEBUG Repository: Wizard %s marcado como completado exitosamente\n", id)
	return nil
}

// FindAllByTenantID obtiene todo el historial de un tenant
func (r *QuickstartHistoryPostgresRepository) FindAllByTenantID(ctx context.Context, tenantID string) ([]*entity.TenantQuickstartHistory, error) {
	query := `
		SELECT id, tenant_id, business_type_id, template_id, setup_completed, 
		       setup_data, created_at, updated_at
		FROM tenant_quickstart_history
		WHERE tenant_id = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, tenantID)
	if err != nil {
		return nil, fmt.Errorf("error al buscar historial: %w", err)
	}
	defer rows.Close()

	var histories []*entity.TenantQuickstartHistory

	for rows.Next() {
		var history entity.TenantQuickstartHistory
		var templateID sql.NullString
		var setupDataBytes []byte

		err := rows.Scan(
			&history.ID,
			&history.TenantID,
			&history.BusinessTypeID,
			&templateID,
			&history.SetupCompleted,
			&setupDataBytes,
			&history.CreatedAt,
			&history.UpdatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("error al escanear registro: %w", err)
		}

		if templateID.Valid {
			history.TemplateID = &templateID.String
		}

		history.SetupData = json.RawMessage(setupDataBytes)
		histories = append(histories, &history)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error al iterar registros: %w", err)
	}

	return histories, nil
}
