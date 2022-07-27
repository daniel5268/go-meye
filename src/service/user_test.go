package service_test

import (
	"errors"
	"testing"

	"github.com/daniel5268/go-meye/src/config"
	"github.com/daniel5268/go-meye/src/domain"
	"github.com/daniel5268/go-meye/src/service"
	"github.com/daniel5268/go-meye/src/service/mocks"
	"github.com/stretchr/testify/assert"
)

func TestUserServiceGetToken(t *testing.T) {
	config.LoadConfig(config.Test)

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
			name: "Returns the token",
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
			name: "Returns an error when the repository fails",
			repository: func() service.UserRepository {
				repositoryMock := &mocks.UserRepository{}
				repositoryMock.On("FindByUsername", username).Return(
					&domain.User{},
					errTest,
				)
				return repositoryMock
			}(),
			wantErr: errTest,
		},
		{
			name: "Returns an error when the secret is incorrect",
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

func TestUserServiceCreate(t *testing.T) {
	username := "test_username"
	user := domain.User{
		Username: username,
	}
	tests := []struct {
		name       string
		user       *domain.User
		repository service.UserRepository
		wantErr    error
	}{
		{
			name: "Returns an error when the user already exists",
			user: &user,
			repository: func() service.UserRepository {
				repositoriMock := &mocks.UserRepository{}
				repositoriMock.On("FindByUsername", username).Return(&domain.User{}, nil)
				return repositoriMock
			}(),
			wantErr: domain.NewDomainError("UserService.Create", domain.CodeUserAlreadyCreatedError, errors.New(domain.CodeUserAlreadyCreatedError)),
		},
		{
			name: "Creates an user",
			user: &user,
			repository: func() service.UserRepository {
				repositoriMock := &mocks.UserRepository{}
				repositoriMock.On("FindByUsername", username).Return(&domain.User{}, errors.New("not found"))
				repositoriMock.On("Create", &user).Return(nil)
				return repositoriMock
			}(),
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service.NewUserService(tt.repository)
			gotErr := s.Create(tt.user)
			assert.Equal(t, tt.wantErr, gotErr)
		})
	}
}

func TestUserServiceUpdate(t *testing.T) {
	userID := 1
	updates := map[string]interface{}{
		"is_admin": true,
	}
	user := domain.User{
		ID:      userID,
		IsAdmin: true,
	}
	errTest := errors.New("test_error")
	tests := []struct {
		name       string
		repository service.UserRepository
		wantUser   *domain.User
		wantErr    error
	}{
		{
			name: "Returns an error when the user is not found",
			repository: func() service.UserRepository {
				repositoryMock := &mocks.UserRepository{}
				repositoryMock.On("FindByID", userID).Return(nil, errTest)
				return repositoryMock
			}(),
			wantErr: errTest,
		},
		{
			name: "Returns the user",
			repository: func() service.UserRepository {
				repositoryMock := &mocks.UserRepository{}
				repositoryMock.On("FindByID", userID).Return(&user, nil)
				repositoryMock.On("Update", updates, &user).Return(nil)
				return repositoryMock
			}(),
			wantUser: &user,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service.NewUserService(tt.repository)
			gotUser, gotErr := s.Update(userID, updates)
			assert.Equal(t, tt.wantUser, gotUser)
			assert.Equal(t, tt.wantErr, gotErr)
		})
	}
}
