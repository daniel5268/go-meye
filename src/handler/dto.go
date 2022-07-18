package handler

import "github.com/daniel5268/go-meye/src/domain"

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

func (ur CreateUserRequest) toUser() (domain.User, error) {
	return domain.NewUser(ur.Username, ur.Secret, ur.IsAdmin, ur.IsMaster, ur.IsPlayer)
}
