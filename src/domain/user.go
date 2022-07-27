package domain

import (
	"time"

	"github.com/daniel5268/go-meye/src/config"
	"github.com/daniel5268/go-meye/src/util"
	"github.com/golang-jwt/jwt"
)

type User struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	HashedSecret string    `json:"-"`
	IsAdmin      bool      `json:"is_admin"`
	IsMaster     bool      `json:"is_master"`
	IsPlayer     bool      `json:"is_player"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type TokenClaims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func NewUser(username, secret string, isAdmin, isMaster, isPlayer bool) (*User, error) {
	section := "NewUser"
	hashedSecret, err := util.HashSecret(secret)
	if err != nil {
		return nil, NewDomainError(section, CodeHashError, err)
	}
	return &User{
		Username:     username,
		IsAdmin:      isAdmin,
		IsMaster:     isMaster,
		IsPlayer:     isPlayer,
		HashedSecret: hashedSecret,
	}, nil
}

func (u User) GetToken() (string, error) {
	section := "User.GetToken"
	claims := TokenClaims{
		u.ID,
		u.Username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 12).Unix(),
			Issuer:    config.JwtIssuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(config.JwtSecret))
	if err != nil {
		return "", NewDomainError(section, CodeSignTokenError, err)
	}

	return signedToken, nil
}

func (u User) ValidateSecret(secret string) bool {
	return util.CompareHashAndSecret(u.HashedSecret, secret)
}
