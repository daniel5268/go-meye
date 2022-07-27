package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/daniel5268/go-meye/src/domain"
	"github.com/labstack/echo/v4"
)

var errInvalidParam = errors.New("invalid_param")

const (
	userIDParameter = "userID"
)

type UserService interface {
	GetToken(username string, secret string) (string, error)
	Create(user *domain.User) error
	Update(userID int, updates map[string]interface{}) (*domain.User, error)
}

type UserHandler struct {
	userService UserService
}

func NewUserHandler(s UserService) *UserHandler {
	return &UserHandler{
		userService: s,
	}
}

func (h *UserHandler) SignIn(c echo.Context) error {
	req := SignInRequest{}
	section := "UserHandler.SignIn"

	if err := c.Bind(&req); err != nil {
		return domain.NewDomainError(section, domain.CodeBindError, err)
	}

	if err := c.Validate(req); err != nil {
		return domain.NewDomainError(section, domain.CodeValidationError, err)
	}

	token, err := h.userService.GetToken(req.Username, req.Secret)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, SignInResponse{
		Token: fmt.Sprintf("Bearer %s", token),
	})
}

func (h UserHandler) Create(c echo.Context) error {
	req := CreateUserRequest{}
	section := "UserHandler.Create"

	if err := c.Bind(&req); err != nil {
		return domain.NewDomainError(section, domain.CodeBindError, err)
	}

	if err := c.Validate(req); err != nil {
		return domain.NewDomainError(section, domain.CodeValidationError, err)
	}

	user, err := req.toUser()
	if err != nil {
		return err
	}

	err = h.userService.Create(user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, user)
}

func (h UserHandler) Update(c echo.Context) error {
	req := UpdateUserRequest{}
	section := "UserHandler.Update"

	if err := c.Bind(&req); err != nil {
		return domain.NewDomainError(section, domain.CodeBindError, err)
	}

	if err := c.Validate(req); err != nil {
		return domain.NewDomainError(section, domain.CodeValidationError, err)
	}

	userID, err := intParam(c, userIDParameter)
	if err != nil {
		return domain.NewDomainError(section, domain.CodeValidationError, errors.New("userID should be integer"))
	}
	updates, err := req.toUpdates()
	if err != nil {
		return err
	}

	updated, err := h.userService.Update(userID, updates)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, updated)
}

func intParam(c echo.Context, parameter string) (int, error) {
	stringValue := c.Param(parameter)
	intValue, err := strconv.Atoi(stringValue)
	if err != nil {
		return 0, errInvalidParam
	}
	return intValue, nil
}
