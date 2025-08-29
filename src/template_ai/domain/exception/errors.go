package exception

import (
	"errors"
	"fmt"
)

var (
	// Template errors
	ErrTemplateNotFound      = errors.New("template not found")
	ErrTemplateAlreadyExists = errors.New("template already exists")
	ErrInvalidTemplateData   = errors.New("invalid template data")
	ErrTemplateGenerationFailed = errors.New("template generation failed")

	// Product errors
	ErrProductNotFound      = errors.New("product not found")
	ErrInvalidProductData   = errors.New("invalid product data")
	ErrDuplicateProduct     = errors.New("duplicate product in template")

	// Generation errors
	ErrGenerationInProgress = errors.New("generation already in progress")
	ErrGenerationTimeout    = errors.New("generation timeout")
	ErrInvalidGenerationParams = errors.New("invalid generation parameters")
	ErrAIServiceUnavailable = errors.New("AI service unavailable")

	// Metric errors
	ErrInvalidMetricType = errors.New("invalid metric type")
	ErrMetricNotFound    = errors.New("metric not found")

	// Feedback errors
	ErrInvalidFeedbackAction = errors.New("invalid feedback action")
	ErrFeedbackNotFound      = errors.New("feedback not found")
)

// TemplateGenerationError represents an error during template generation
type TemplateGenerationError struct {
	Reason string
	Err    error
}

func (e *TemplateGenerationError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("template generation failed: %s - %v", e.Reason, e.Err)
	}
	return fmt.Sprintf("template generation failed: %s", e.Reason)
}

func (e *TemplateGenerationError) Unwrap() error {
	return e.Err
}

// NewTemplateGenerationError creates a new template generation error
func NewTemplateGenerationError(reason string, err error) *TemplateGenerationError {
	return &TemplateGenerationError{
		Reason: reason,
		Err:    err,
	}
}

// ValidationError represents a validation error
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error on field '%s': %s", e.Field, e.Message)
}

// NewValidationError creates a new validation error
func NewValidationError(field, message string) *ValidationError {
	return &ValidationError{
		Field:   field,
		Message: message,
	}
}

// IsNotFoundError checks if an error is a not found error
func IsNotFoundError(err error) bool {
	return errors.Is(err, ErrTemplateNotFound) || 
		errors.Is(err, ErrProductNotFound) || 
		errors.Is(err, ErrMetricNotFound) || 
		errors.Is(err, ErrFeedbackNotFound)
}

// IsValidationError checks if an error is a validation error
func IsValidationError(err error) bool {
	_, ok := err.(*ValidationError)
	return ok
}

// IsGenerationError checks if an error is a generation error
func IsGenerationError(err error) bool {
	_, ok := err.(*TemplateGenerationError)
	return ok || errors.Is(err, ErrGenerationInProgress) || 
		errors.Is(err, ErrGenerationTimeout) || 
		errors.Is(err, ErrTemplateGenerationFailed)
}

// InternalError represents an internal server error
type InternalError struct {
	Message string
}

func (e *InternalError) Error() string {
	return e.Message
}

// NewInternalError creates a new internal error
func NewInternalError(message string) error {
	return &InternalError{
		Message: message,
	}
}

// NotFoundError represents a not found error
type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return e.Message
}

// NewNotFoundError creates a new not found error
func NewNotFoundError(message string) error {
	return &NotFoundError{
		Message: message,
	}
}