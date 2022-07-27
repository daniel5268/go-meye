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
	FindByUsername(username string) (*domain.User, error)
	FindByID(ID int) (*domain.User, error)
	Create(u ...*domain.User) error
	Update(updates map[string]interface{}, u ...*domain.User) error
	Delete(userID int) error
}

func NewUserService(ur UserRepository) *UserService {
	return &UserService{
		userRepository: ur,
	}
}

func (us *UserService) GetToken(username string, secret string) (string, error) {
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

func (us *UserService) Create(user *domain.User) error {
	section := "UserService.Create"
	_, err := us.userRepository.FindByUsername(user.Username)
	if err == nil {
		return domain.NewDomainError(section, domain.CodeUserAlreadyCreatedError, errors.New(domain.CodeUserAlreadyCreatedError))
	}

	return us.userRepository.Create(user)
}

func (us *UserService) Update(userID int, updates map[string]interface{}) (*domain.User, error) {
	u, err := us.userRepository.FindByID(userID)
	if err != nil {
		return u, err
	}

	err = us.userRepository.Update(updates, u)

	return u, err
}

func (us *UserService) Delete(userID int) error {
	return us.userRepository.Delete(userID)
}
