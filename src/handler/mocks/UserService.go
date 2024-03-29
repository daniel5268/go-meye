// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	domain "github.com/daniel5268/go-meye/src/domain"

	mock "github.com/stretchr/testify/mock"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// Create provides a mock function with given fields: user
func (_m *UserService) Create(user *domain.User) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: userID
func (_m *UserService) Delete(userID int) error {
	ret := _m.Called(userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetToken provides a mock function with given fields: username, secret
func (_m *UserService) GetToken(username string, secret string) (string, error) {
	ret := _m.Called(username, secret)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(username, secret)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(username, secret)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: userID, updates
func (_m *UserService) Update(userID int, updates map[string]interface{}) (*domain.User, error) {
	ret := _m.Called(userID, updates)

	var r0 *domain.User
	if rf, ok := ret.Get(0).(func(int, map[string]interface{}) *domain.User); ok {
		r0 = rf(userID, updates)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, map[string]interface{}) error); ok {
		r1 = rf(userID, updates)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserService interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserService creates a new instance of UserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserService(t mockConstructorTestingTNewUserService) *UserService {
	mock := &UserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
