package api

import (
	"errors"
	"strings"

	"github.com/daniel5268/go-meye/src/config"
	"github.com/daniel5268/go-meye/src/domain"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

const tokenHeader = "Authorization"

var errForbidden = errors.New("Forbidden")

type UserRepository interface {
	FindByID(userID int) (*domain.User, error)
}

func AuthAdmin(ur UserRepository) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			section := "middleware.AuthAdmin"
			errForbidden := domain.NewDomainError(section, domain.CodeForbiddenError, errForbidden)
			authorization := c.Request().Header.Get(tokenHeader)
			if authorization == "" {
				return errForbidden
			}
			authSplit := strings.Split(authorization, "Bearer ")
			if len(authSplit) != 2 {
				return errForbidden
			}
			tokenString := authSplit[1]

			claims := domain.TokenClaims{}
			token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(config.JwtSecret), nil
			})

			if err != nil || !token.Valid {
				return errForbidden
			}

			user, err := ur.FindByID(claims.ID)
			if err != nil || !user.IsAdmin {
				return errForbidden
			}

			return next(c)
		}
	}
}
