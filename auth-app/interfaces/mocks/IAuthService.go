package mocks

import (
	mock "github.com/stretchr/testify/mock"
	"efishery-task/auth-app/models"
)

type IAuthService struct {
	mock.Mock
}

func (_m *IAuthService) RegisterUser(phone string, name string, role string) (*models.UserModel, error) {
	args := _m.Called(phone, name, role)
	result := args.Get(0)
	return result.(*models.UserModel), args.Error(1)
}

func (_m *IAuthService) CheckLogin(phone string, password string) (*models.UserModel, error) {
	args := _m.Called(phone, password)
	result := args.Get(0)
	return result.(*models.UserModel), args.Error(1)
}