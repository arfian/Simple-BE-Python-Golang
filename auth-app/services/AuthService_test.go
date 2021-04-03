package services

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"efishery-task/auth-app/interfaces/mocks"
	"efishery-task/auth-app/models"
)

func TestRegisterUser(t *testing.T) {
	helperFunc := new(mocks.IHelperFunc)
	authRepository := new(mocks.IAuthRepository)
	helperFunc.On("GenerateUsername", "arfian bagus").Return("arfianbagus")
	helperFunc.On("GeneratePass").Return("abcd")
	authRepository.On("CheckUsername", "arfianbagus", "12345678").Return(false, nil)
	user := models.UserModel{
		Name: "arfian bagus",
		Phone: "12345678",
		Role: "admin",
		Username: "arfianbagus",
		Password: "abcd",
	}
	authRepository.On("SaveUser", user).Return(&user, nil)

	authService := AuthService{Repo: authRepository, Helper: helperFunc}

	actualResult, _ := authService.RegisterUser("12345678", "arfian bagus", "admin")

	assert.Equal(t, "arfianbagus", actualResult.Username)
}

func TestRegisterUserFail(t *testing.T) {
	helperFunc := new(mocks.IHelperFunc)
	authRepository := new(mocks.IAuthRepository)
	helperFunc.On("GenerateUsername", "arfian bagus").Return("arfianbagus")
	authRepository.On("CheckUsername", "arfianbagus", "12345678").Return(true, nil)

	authService := AuthService{Repo: authRepository, Helper: helperFunc}

	_, err := authService.RegisterUser("12345678", "arfian bagus", "admin")

	assert.Equal(t, "Maaf nama arfian bagus sudah teregister dengan username arfianbagus dan phone 12345678, silakan daftar kembali dengan nama yang berbeda", err.Error())
}

func TestCheckLogin(t *testing.T) {
	helperFunc := new(mocks.IHelperFunc)
	authRepository := new(mocks.IAuthRepository)
	user := models.UserModel{
		Name: "arfian bagus",
		Phone: "12345678",
		Role: "admin",
		Username: "arfianbagus",
		Password: "abcd",
	}
	authRepository.On("GetLogin", "12345678", "abcd").Return(user, nil)

	authService := AuthService{Repo: authRepository, Helper: helperFunc}

	actualResult, _ := authService.CheckLogin("12345678", "abcd")

	assert.Equal(t, "arfianbagus", actualResult.Username)
}

func TestCheckLoginFail(t *testing.T) {
	helperFunc := new(mocks.IHelperFunc)
	authRepository := new(mocks.IAuthRepository)
	authRepository.On("GetLogin", "12345678", "abcd").Return(models.UserModel{}, nil)

	authService := AuthService{Repo: authRepository, Helper: helperFunc}

	_, err := authService.CheckLogin("12345678", "abcd")

	assert.Equal(t, "Maaf phone number dan password tidak cocok", err.Error())
}