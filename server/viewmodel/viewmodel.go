package viewmodel

// ErrorResponse is the standard error response model
type ErrorResponse struct {
	Error   bool   `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}
