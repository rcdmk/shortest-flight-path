package errors

// errorString is the base type that allows error constants
type errorString string

// Error returns the error message
func (es errorString) Error() string {
	return string(es)
}

const (
	// NotFound represents mising data error
	NotFound errorString = "not found"
)

// ValidationError represents a input validation error
type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

// NewValidationError returns a new ValidationError
func NewValidationError(message string) error {
	return &ValidationError{
		Message: message,
	}
}

// NotFoundError represents a data not found error
type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return e.Message
}

// NewNotFoundError returns a new NotFoundError
func NewNotFoundError(message string) error {
	return &NotFoundError{
		Message: message,
	}
}
