package domain

// errorString is the base type that allows error constants
type errorString string

// Error returns the error message
func (es errorString) Error() string {
	return string(es)
}

const (
	// ErrNotFound represents mising data error
	ErrNotFound errorString = "not found"
)
