package repository

import (
	"errors"

	"github.com/daniel5268/go-meye/src/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (r UserRepository) FindByUsername(username string) (domain.User, error) {
	section := "UserRepository.FindByUsername"
	u := domain.User{}
	result := r.db.First(&u, "username = ?", username)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return u, domain.NewDomainError(section, domain.CodeUserNotFoundError, result.Error)
	}

	if result.Error != nil {
		return u, domain.NewDomainError(section, domain.CodeRepositoryError, result.Error)
	}

	return u, nil
}

func (r UserRepository) Create(u ...*domain.User) error {
	section := "UserRepository.Create"
	err := r.db.Create(u).Error
	if err != nil {
		return domain.NewDomainError(section, domain.CodeRepositoryError, err)
	}

	return nil
}
