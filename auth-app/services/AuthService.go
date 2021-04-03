package services

import (
	"fmt"
	"efishery-task/auth-app/interfaces"
	"efishery-task/auth-app/models"
	"time"
	"errors"
)

type AuthService struct {
	Repo interfaces.IAuthRepository
	Helper interfaces.IHelperFunc
}

func (service *AuthService) RegisterUser(phone string, name string, role string) (*models.UserModel, error) {
	username := service.Helper.GenerateUsername(name)
	checkUser, _:= service.Repo.CheckUsername(username)
	if checkUser {
		response := fmt.Sprintf("Maaf nama %s sudah teregister dengan username %s, silakan daftar kembali dengan nama yang berbeda", name, username)
		return nil, errors.New(response)
	}

	u := models.UserModel{
		Name: name,
		Phone: phone,
		Role: role,
		Username: username,
		Password: service.Helper.GeneratePass(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	fmt.Println("=========== checkUser ========")
	fmt.Println(checkUser)
	user, err := service.Repo.SaveUser(u)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &user, nil
}