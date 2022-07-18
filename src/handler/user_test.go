package handler_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/daniel5268/go-meye/src/api"
	"github.com/daniel5268/go-meye/src/domain"
	"github.com/daniel5268/go-meye/src/handler"
	"github.com/daniel5268/go-meye/src/handler/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func compareErr(expectedErr, actualErr error) bool {
	expectedDomainErr, expectedOk := expectedErr.(*domain.DomainError)
	actualDomainErr, actualOk := actualErr.(*domain.DomainError)

	if expectedOk != actualOk {
		return false
	}

	if expectedOk && actualOk {
		return expectedDomainErr.Code == actualDomainErr.Code && expectedDomainErr.Section == actualDomainErr.Section
	}

	return expectedErr == actualErr
}

func TestUserHandlerSignIn(t *testing.T) {
	e := echo.New()
	e.Validator = api.NewValidator()
	errTest := errors.New("test_error")
	section := "UserHandler.SignIn"
	tests := []struct {
		name    string
		body    []byte
		service handler.UserService
		wantErr error
	}{
		{
			name:    "returns a domain error with code = CodeBindError when there is an unexpected type",
			body:    []byte(`{"username":50,"secret":"secret"}`),
			wantErr: domain.NewDomainError(section, domain.CodeBindError, errTest),
		},
		{
			name:    "returns a domain error when validations fail",
			wantErr: domain.NewDomainError(section, domain.CodeValidationError, errTest),
		},
		{
			name:    "returns an error when the service fails",
			wantErr: errTest,
			body:    []byte(`{"username":"myUsername","secret":"mySecret"}`),
			service: func() handler.UserService {
				serviceMock := &mocks.UserService{}
				serviceMock.On("GetToken", "myUsername", "mySecret").Return("", errTest)
				return serviceMock
			}(),
		},
		{
			name: "works correctly",
			body: []byte(`{"username":"myUsername","secret":"mySecret"}`),
			service: func() handler.UserService {
				serviceMock := &mocks.UserService{}
				serviceMock.On("GetToken", "myUsername", "mySecret").Return("token", nil)
				return serviceMock
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/sign-in", bytes.NewBuffer(tt.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			c := e.NewContext(req, rec)
			h := handler.NewUserHandler(tt.service)

			gotErr := h.SignIn(c)

			assert.True(t, compareErr(tt.wantErr, gotErr))
		})
	}

}
