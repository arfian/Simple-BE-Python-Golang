package services

import (
	"fmt"
	"efishery-task/auth-app/interfaces"
	"efishery-task/auth-app/models"
)

type AuthService struct {
	interfaces.IAuthRepository
}

func (service *AuthService) RegisterUser(phone string, name string, role string) (*models.UserModel, error) {
	u := models.UserModel{
		Name: name,
		Phone: phone,
		Role: role,
	}
	user, err := service.SaveUser(u)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &user, nil
}