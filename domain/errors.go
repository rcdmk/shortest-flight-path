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

	// ErrInvalidRouteOrigin indicates that de origin for a route is not known or invalid
	ErrInvalidRouteOrigin errorString = "invalid origin"

	// ErrInvalidRouteDestination indicates that de destination for a route is not known or invalid
	ErrInvalidRouteDestination errorString = "invalid destination"

	// ErrSameRouteSourceAndDestination indicates that the route's source and destination are the same
	ErrSameRouteSourceAndDestination errorString = "source and destination can't be the same"

	// ErrCacheMiss indicates that a cache entry is missing or is already expired
	ErrCacheMiss errorString = "cache miss"
)
