package value_object

import "errors"

// GenerationStatus represents the status of an AI generation
type GenerationStatus string

const (
	GenerationStatusPending    GenerationStatus = "pending"
	GenerationStatusProcessing GenerationStatus = "processing"
	GenerationStatusCompleted  GenerationStatus = "completed"
	GenerationStatusFailed     GenerationStatus = "failed"
)

// Valid checks if the generation status is valid
func (s GenerationStatus) Valid() error {
	switch s {
	case GenerationStatusPending, GenerationStatusProcessing, GenerationStatusCompleted, GenerationStatusFailed:
		return nil
	default:
		return errors.New("invalid generation status")
	}
}

// String returns the string representation
func (s GenerationStatus) String() string {
	return string(s)
}

// GenerationType represents how a template was generated
type GenerationType string

const (
	GenerationTypeAI     GenerationType = "ai"
	GenerationTypeManual GenerationType = "manual"
	GenerationTypeHybrid GenerationType = "hybrid"
)

// Valid checks if the generation type is valid
func (t GenerationType) Valid() error {
	switch t {
	case GenerationTypeAI, GenerationTypeManual, GenerationTypeHybrid:
		return nil
	default:
		return errors.New("invalid generation type")
	}
}

// String returns the string representation
func (t GenerationType) String() string {
	return string(t)
}

// ProductPriority represents the priority level of a product in a template
type ProductPriority int

const (
	ProductPriorityEssential   ProductPriority = 1
	ProductPriorityRecommended ProductPriority = 2
	ProductPriorityOptional    ProductPriority = 3
)

// Valid checks if the priority is valid
func (p ProductPriority) Valid() error {
	if p < ProductPriorityEssential || p > ProductPriorityOptional {
		return errors.New("invalid product priority: must be between 1 and 3")
	}
	return nil
}

// String returns a string representation of the priority
func (p ProductPriority) String() string {
	switch p {
	case ProductPriorityEssential:
		return "essential"
	case ProductPriorityRecommended:
		return "recommended"
	case ProductPriorityOptional:
		return "optional"
	default:
		return "unknown"
	}
}

// FeedbackAction represents the type of feedback action
type FeedbackAction string

const (
	FeedbackActionKept            FeedbackAction = "kept"
	FeedbackActionRemoved         FeedbackAction = "removed"
	FeedbackActionQuantityChanged FeedbackAction = "quantity_changed"
	FeedbackActionReplaced        FeedbackAction = "replaced"
)

// Valid checks if the feedback action is valid
func (a FeedbackAction) Valid() error {
	switch a {
	case FeedbackActionKept, FeedbackActionRemoved, FeedbackActionQuantityChanged, FeedbackActionReplaced:
		return nil
	default:
		return errors.New("invalid feedback action")
	}
}

// String returns the string representation
func (a FeedbackAction) String() string {
	return string(a)
}

// MetricType represents the type of performance metric
type MetricType string

const (
	MetricTypeUsageCount        MetricType = "usage_count"
	MetricTypeSatisfactionScore MetricType = "satisfaction_score"
	MetricTypeModificationRate  MetricType = "modification_rate"
	MetricTypeAdoptionRate      MetricType = "adoption_rate"
	MetricTypeProductRetention  MetricType = "product_retention"
)

// Valid checks if the metric type is valid
func (m MetricType) Valid() error {
	switch m {
	case MetricTypeUsageCount, MetricTypeSatisfactionScore, MetricTypeModificationRate, 
		MetricTypeAdoptionRate, MetricTypeProductRetention:
		return nil
	default:
		return errors.New("invalid metric type")
	}
}

// String returns the string representation
func (m MetricType) String() string {
	return string(m)
}