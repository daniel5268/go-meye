package handler

import (
	"github.com/daniel5268/go-meye/src/domain"
	"github.com/daniel5268/go-meye/src/util"
)

type SignInRequest struct {
	Username string `json:"username" validate:"required,min=1,alphanum"`
	Secret   string `json:"secret" validate:"required,min=4"`
}

type SignInResponse struct {
	Token string `json:"token"`
}

type CreateUserRequest struct {
	Username string `json:"username" validate:"required,min=1,alphanum"`
	Secret   string `json:"secret" validate:"required,min=4"`
	IsAdmin  bool   `json:"is_admin"`
	IsMaster bool   `json:"is_master"`
	IsPlayer bool   `json:"is_player"`
}

func (cr CreateUserRequest) toUser() (*domain.User, error) {
	return domain.NewUser(cr.Username, cr.Secret, cr.IsAdmin, cr.IsMaster, cr.IsPlayer)
}

type UpdateUserRequest struct {
	Secret   string `json:"secret" validate:"min=4"`
	IsAdmin  *bool  `json:"is_admin"`
	IsMaster *bool  `json:"is_master"`
	IsPlayer *bool  `json:"is_player"`
}

func (ur UpdateUserRequest) toUpdates() (map[string]interface{}, error) {
	updates := map[string]interface{}{}

	if ur.Secret != "" {
		hashedSecret, err := util.HashSecret(ur.Secret)
		if err != nil {
			return updates, err
		}
		updates["hashed_secret"] = hashedSecret
	}

	if ur.IsAdmin != nil {
		updates["is_admin"] = ur.IsAdmin
	}

	if ur.IsMaster != nil {
		updates["is_master"] = ur.IsMaster
	}

	if ur.IsPlayer != nil {
		updates["is_player"] = ur.IsPlayer
	}

	return updates, nil
}
