package request

// StartWizardRequest representa la petición para iniciar el wizard
type StartWizardRequest struct {
	BusinessTypeID string `json:"business_type_id" binding:"required" example:"550e8400-e29b-41d4-a716-446655440000"`
}

// UpdateWizardStepRequest representa la petición para actualizar un step del wizard
type UpdateWizardStepRequest struct {
	CurrentStep string                 `json:"current_step" binding:"required" example:"categories_selection"`
	StepData    map[string]interface{} `json:"step_data" example:"{\"selected_categories\": [\"1\", \"2\", \"3\"]}"`
}

// CompleteWizardRequest representa la petición para completar el wizard
type CompleteWizardRequest struct {
	FinalSelections map[string]interface{} `json:"final_selections" binding:"required"`
}