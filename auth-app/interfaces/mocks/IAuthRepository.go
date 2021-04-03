package mocks

import (
	mock "github.com/stretchr/testify/mock"
	"efishery-task/auth-app/models"
)

type IAuthRepository struct {
	mock.Mock
}

func (_m *IAuthRepository) SaveUser(user models.UserModel) (*models.UserModel, error) {
	args := _m.Called(user)
	result := args.Get(0)
	return result.(*models.UserModel), args.Error(1)
}

func (_m *IAuthRepository) CheckUsername(username string, phone string) (bool, error) {
	args := _m.Called(username, phone)
	result := args.Get(0)
	return result.(bool), args.Error(1)
}

func (_m *IAuthRepository) GetLogin(phone string, password string) (models.UserModel, error) {
	args := _m.Called(phone, password)
	result := args.Get(0)
	return result.(models.UserModel), args.Error(1)
}