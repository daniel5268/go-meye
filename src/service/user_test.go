package service_test

import (
	"errors"
	"os"
	"testing"

	"github.com/daniel5268/go-meye/src/domain"
	"github.com/daniel5268/go-meye/src/service"
	"github.com/daniel5268/go-meye/src/service/mocks"
	"github.com/stretchr/testify/assert"
)

func setEnv(t *testing.T) {
	err := os.Setenv("JWT_ISSUER", "issuer")
	if err != nil {
		assert.Fail(t, "Error setting JWT_ISSUER")
	}
	err = os.Setenv("JWT_SECRET", "secret")
	if err != nil {
		assert.Fail(t, "Error setting JWT_SECRET")
	}
}

func unSetEnv() {
	os.Unsetenv("JWT_ISSUER")
	os.Unsetenv("JWT_SECRET")
}

func TestUserServiceGetToken(t *testing.T) {
	setEnv(t)
	defer unSetEnv()

	username := "rocket"
	secret := "league"
	errTest := errors.New("test_error")
	tests := []struct {
		name       string
		repository service.UserRepository
		wantToken  bool
		wantErr    error
	}{
		{
			name: "should return the token",
			repository: func() service.UserRepository {
				repositoryMock := &mocks.UserRepository{}
				user, err := domain.NewUser(username, secret, false, false, false)
				if err != nil {
					assert.Fail(t, "Error creating user")
				}
				repositoryMock.On("FindByUsername", username).Return(
					user,
					nil,
				)
				return repositoryMock
			}(),
			wantToken: true,
			wantErr:   nil,
		},
		{
			name: "should return an error when the repository fails",
			repository: func() service.UserRepository {
				repositoryMock := &mocks.UserRepository{}
				repositoryMock.On("FindByUsername", username).Return(
					domain.User{},
					errTest,
				)
				return repositoryMock
			}(),
			wantErr: errTest,
		},
		{
			name: "should return an error when the secret is incorrect",
			repository: func() service.UserRepository {
				repositoryMock := &mocks.UserRepository{}
				user, err := domain.NewUser(username, "other", false, false, false)
				if err != nil {
					assert.Fail(t, "Error creating user")
				}
				repositoryMock.On("FindByUsername", username).Return(
					user,
					nil,
				)
				return repositoryMock
			}(),
			wantToken: false,
			wantErr: domain.NewDomainError(
				"UserService.GetToken",
				domain.CodeInvalidCredentialsError,
				service.ErrInvalidCredentials,
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service.NewUserService(tt.repository)
			gotToken, gotErr := s.GetToken(username, secret)
			assert.Equal(t, tt.wantToken, gotToken != "")
			assert.Equal(t, tt.wantErr, gotErr)
		})
	}
}
