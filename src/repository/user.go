package repository

import (
	"errors"

	"github.com/daniel5268/go-meye/src/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) FindByUsername(username string) (*domain.User, error) {
	section := "UserRepository.FindByUsername"
	u := domain.User{}
	result := r.db.First(&u, "username = ?", username)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, domain.NewDomainError(section, domain.CodeUserNotFoundError, result.Error)
	}

	if result.Error != nil {
		return nil, domain.NewDomainError(section, domain.CodeRepositoryError, result.Error)
	}

	return &u, nil
}

func (r *UserRepository) FindByID(ID int) (*domain.User, error) {
	section := "UserRepository.FindByID"
	u := domain.User{}
	result := r.db.First(&u, "id = ?", ID)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, domain.NewDomainError(section, domain.CodeUserNotFoundError, result.Error)
	}

	if result.Error != nil {
		return nil, domain.NewDomainError(section, domain.CodeRepositoryError, result.Error)
	}

	return &u, nil
}

func (r *UserRepository) Create(u ...*domain.User) error {
	section := "UserRepository.Create"
	err := r.db.Create(u).Error
	if err != nil {
		return domain.NewDomainError(section, domain.CodeRepositoryError, err)
	}

	return nil
}

func (r *UserRepository) Update(updates map[string]interface{}, u ...*domain.User) error {
	section := "UserRepository.Update"
	err := r.db.Model(u).Updates(updates).Error
	if err != nil {
		return domain.NewDomainError(section, domain.CodeRepositoryError, err)
	}

	return nil
}
