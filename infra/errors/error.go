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
