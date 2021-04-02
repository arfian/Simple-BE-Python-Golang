package repositories

import (
	"efishery-task/auth-app/interfaces"
	"efishery-task/auth-app/models"
)

type AuthRepository struct {
	interfaces.IDbHandler
}

func (repository *AuthRepository) GetUsername(username string) (*models.UserModel, error) {
	var user models.UserModel
	return &user, nil
}

func (repository *AuthRepository) SaveUser(models.UserModel) (models.UserModel, error) {
	var user models.UserModel
	return user, nil
}
