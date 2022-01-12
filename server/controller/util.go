package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/rcdmk/shortest-flight-path/domain"
	"github.com/rcdmk/shortest-flight-path/infra/errors"
	"github.com/rcdmk/shortest-flight-path/server/viewmodel"
)

// errorResult returns a standard API error result
func errorResult(c echo.Context, err error) error {
	statusCode := http.StatusServiceUnavailable
	errorMessage := http.StatusText(statusCode)

	switch err {
	case domain.ErrSameRouteSourceAndDestination,
		domain.ErrInvalidRouteDestination,
		domain.ErrInvalidRouteOrigin:
		err = errors.NewValidationError(err.Error())
	}

	if err == domain.ErrNotFound {
		statusCode = http.StatusNotFound
		errorMessage = err.Error()
	} else {
		switch e := err.(type) {
		case *errors.NotFoundError:
			statusCode = http.StatusNotFound
			errorMessage = e.Message

		case *errors.ValidationError:
			statusCode = http.StatusUnprocessableEntity
			errorMessage = e.Message

		case *echo.HTTPError:
			statusCode = e.Code
			errorMessage = e.Error()
		}
	}

	result := viewmodel.ErrorResponse{
		Error:   true,
		Message: errorMessage,
	}

	return c.JSON(statusCode, result)
}
