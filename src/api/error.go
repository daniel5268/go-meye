package api

import (
	"net/http"

	"github.com/daniel5268/go-meye/src/domain"
	"github.com/labstack/echo/v4"
)

const internalError = "Internal server error, please contanct administrador"

const (
	CodeInternalError   = "internal_error"
	CodeBindError       = "bind_error"
	CodeBadRequest      = "bad_request"
	CodeUnauthorized    = "unauthorized"
	CodeSignError       = "sign_token_error"
	CodeNotFound        = "not_found"
	CodeRepositoryError = "repository_error"
	CodeForbiddenError  = "forbidden"
	CodeHashError       = "hash_error"
)

// APIError structure that represents an API error
type APIError struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// ErrorHandler manages errors on application level
func ErrorHandler(err error, c echo.Context) {
	c.Logger().Error(err)
	aErr := mapToAPIError(err)
	_ = c.JSON(aErr.Status, aErr)
}

func NewAPIError(s int, c, m string) APIError {
	return APIError{
		Status:  s,
		Code:    c,
		Message: m,
	}
}

func mapToAPIError(err error) APIError {
	dErr, isDomainErr := err.(*domain.DomainError)
	if !isDomainErr {
		httpErr, isHttpErr := err.(*echo.HTTPError)
		if isHttpErr && httpErr.Code == http.StatusNotFound {
			return NewAPIError(http.StatusNotFound, CodeNotFound, "Resource not found")
		}
		return NewAPIError(http.StatusInternalServerError, CodeInternalError, internalError)
	}

	switch dErr.Code {
	case domain.CodeBindError:
		return NewAPIError(http.StatusNotAcceptable, CodeBindError, "Not acceptable request, please check your inputs")
	case domain.CodeValidationError:
		return NewAPIError(http.StatusBadRequest, CodeBadRequest, dErr.Error())
	case domain.CodeInvalidCredentialsError:
		return NewAPIError(http.StatusUnauthorized, CodeUnauthorized, "Invalid credentials")
	case domain.CodeSignTokenError:
		return NewAPIError(http.StatusInternalServerError, CodeSignError, internalError)
	case domain.CodeUserNotFoundError:
		return NewAPIError(http.StatusNotFound, CodeNotFound, "User not found")
	case domain.CodeRepositoryError:
		return NewAPIError(http.StatusInternalServerError, CodeRepositoryError, internalError)
	case domain.CodeForbiddenError:
		return NewAPIError(http.StatusForbidden, CodeForbiddenError, "Forbidden")
	case domain.CodeUserAlreadyCreatedError:
		return NewAPIError(http.StatusBadRequest, CodeBadRequest, "User already created")
	case domain.CodeHashError:
		return NewAPIError(http.StatusInternalServerError, CodeHashError, internalError)
	default:
		return NewAPIError(http.StatusInternalServerError, CodeInternalError, internalError)
	}
}
