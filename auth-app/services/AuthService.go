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
	checkUser, err:= service.Repo.CheckUsername(username)
	if err != nil {
		return nil, err
	}
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
	user, err := service.Repo.SaveUser(u)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (service *AuthService) CheckLogin(phone string, password string) (*models.UserModel, error) {
	user, err:= service.Repo.GetLogin(phone, password)
	if err != nil {
		return nil, err
	}
	
	if user.Phone == "" {
		return nil, errors.New("Maaf phone number dan password tidak cocok")
	}
	return &user, nil
}