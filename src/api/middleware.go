package api

import (
	"errors"

	"github.com/daniel5268/go-meye/src/config"
	"github.com/daniel5268/go-meye/src/domain"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

const tokenHeader = "Authorization"

var errForbidden = errors.New("Forbidden")

type UserRepository interface {
	FindByUsername(username string) (domain.User, error)
}

func AuthAdmin(repository UserRepository) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			section := "middleware.AuthAdmin"
			errForbidden := domain.NewDomainError(section, domain.CodeForbiddenError, errForbidden)
			tokenString := c.Request().Header.Get(tokenHeader)
			if tokenString == "" {
				return errForbidden
			}

			claims := &domain.TokenClaims{}
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(config.JwtSecret), nil
			})

			if err != nil || !token.Valid {
				return errForbidden
			}

			user, err := repository.FindByUsername(claims.Username)
			if err != nil || !user.IsAdmin {
				return errForbidden
			}

			return next(c)
		}
	}
}
