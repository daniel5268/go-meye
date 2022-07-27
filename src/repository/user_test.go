package repository_test

import (
	"testing"

	"github.com/daniel5268/go-meye/src/config"
	"github.com/daniel5268/go-meye/src/domain"
	"github.com/daniel5268/go-meye/src/infrastructure"
	"github.com/daniel5268/go-meye/src/repository"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
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

func compareUser(expectedUser, actualUser *domain.User) bool {
	if expectedUser == nil && actualUser != nil {
		return false
	}
	if expectedUser != nil && actualUser == nil {
		return false
	}
	if expectedUser == nil && actualUser == nil {
		return true
	}
	c1 := expectedUser.IsAdmin == actualUser.IsAdmin
	c2 := expectedUser.IsMaster == actualUser.IsMaster
	c3 := expectedUser.IsPlayer == actualUser.IsPlayer
	c4 := expectedUser.Username == actualUser.Username
	c5 := expectedUser.HashedSecret == actualUser.HashedSecret
	return c1 && c2 && c3 && c4 && c5
}

func TestUserRepositoryFindByUserName(t *testing.T) {
	section := "UserRepository.FindByUsername"
	username := "admin"
	tests := []struct {
		name     string
		username string
		wantUser *domain.User
		wantErr  error
	}{
		{
			name:     "Returns a not found error when the User doesn't exist",
			username: "ups",
			wantErr:  domain.NewDomainError(section, domain.CodeUserNotFoundError, gorm.ErrRecordNotFound),
		},
		{
			name:     "Returns the user",
			username: username,
			wantUser: &domain.User{
				Username:     username,
				HashedSecret: "$2a$08$ad8/41F7A5qAXJc17A31UeWBKoWiZDJ/323YHUI28pMaCw9BSUSNm",
				IsAdmin:      true,
			},
		},
	}
	config.LoadConfig(config.Test)
	db := infrastructure.NewGormPostgresClient()
	r := repository.NewUserRepository(db)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUser, gotErr := r.FindByUsername(tt.username)
			assert.True(t, compareUser(tt.wantUser, gotUser))
			assert.True(t, compareErr(tt.wantErr, gotErr))
		})
	}
}

func TestUserRepositoryFindByID(t *testing.T) {
	section := "UserRepository.FindByID"
	userID := 1
	tests := []struct {
		name     string
		ID       int
		wantUser *domain.User
		wantErr  error
	}{
		{
			name:    "Returns a not found error when the User doesn't exist",
			ID:      321,
			wantErr: domain.NewDomainError(section, domain.CodeUserNotFoundError, gorm.ErrRecordNotFound),
		},
		{
			name: "Returns the user",
			ID:   userID,
			wantUser: &domain.User{
				Username:     "admin",
				HashedSecret: "$2a$08$ad8/41F7A5qAXJc17A31UeWBKoWiZDJ/323YHUI28pMaCw9BSUSNm",
				IsAdmin:      true,
			},
		},
	}
	config.LoadConfig(config.Test)
	db := infrastructure.NewGormPostgresClient()
	r := repository.NewUserRepository(db)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUser, gotErr := r.FindByID(tt.ID)
			assert.True(t, compareUser(tt.wantUser, gotUser))
			assert.True(t, compareErr(tt.wantErr, gotErr))
		})
	}
}

func TestUserRepositoryUpdate(t *testing.T) {
	tests := []struct {
		name    string
		updates map[string]interface{}
		user    *domain.User
		wantErr error
	}{
		{
			name: "Updates correctly",
			updates: map[string]interface{}{
				"is_admin": true,
			},
			user: &domain.User{
				ID: 1,
			},
			wantErr: nil,
		},
	}
	config.LoadConfig(config.Test)
	db := infrastructure.NewGormPostgresClient()
	r := repository.NewUserRepository(db)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := r.Update(tt.updates, tt.user)
			assert.Equal(t, tt.wantErr, gotErr)
		})
	}
}

func TestUserRepositoryDelete(t *testing.T) {
	tests := []struct {
		name    string
		userID  int
		wantErr error
	}{
		{
			name:    "Deletes the user",
			userID:  123,
			wantErr: nil,
		},
	}
	config.LoadConfig(config.Test)
	db := infrastructure.NewGormPostgresClient()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repository.NewUserRepository(db)
			gotErr := r.Delete(tt.userID)
			assert.Equal(t, tt.wantErr, gotErr)
		})
	}
}
