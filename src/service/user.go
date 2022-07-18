package service

import (
	"errors"

	"github.com/daniel5268/go-meye/src/domain"
)

var (
	ErrInvalidCredentials = errors.New(domain.CodeInvalidCredentialsError)
)

type UserService struct {
	userRepository UserRepository
}

type UserRepository interface {
	FindByUsername(username string) (domain.User, error)
	Create(u ...*domain.User) error
}

func NewUserService(ur UserRepository) UserService {
	return UserService{
		userRepository: ur,
	}
}

func (us UserService) GetToken(username string, secret string) (string, error) {
	section := "UserService.GetToken"
	u, err := us.userRepository.FindByUsername(username)
	if err != nil {
		return "", err
	}

	if !u.ValidateSecret(secret) {
		return "", domain.NewDomainError(section, domain.CodeInvalidCredentialsError, ErrInvalidCredentials)
	}

	token, err := u.GetToken()
	if err != nil {
		return "", err
	}

	return token, nil
}

func (us UserService) Create(user *domain.User) error {
	return us.userRepository.Create(user)
}
