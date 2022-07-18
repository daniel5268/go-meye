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

func compareUser(expectedUser, actualUser domain.User) bool {
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
		wantUser domain.User
		wantErr  error
	}{
		{
			name:     "Should return a not found error when the User doesn't exist",
			username: "ups",
			wantUser: domain.User{},
			wantErr:  domain.NewDomainError(section, domain.CodeUserNotFoundError, gorm.ErrRecordNotFound),
		},
		{
			name:     "Should return the user",
			username: username,
			wantUser: domain.User{
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
