package interfaces

import (
	"efishery-task/auth-app/models"
)

type IAuthRepository interface {
	SaveUser(user models.UserModel) (*models.UserModel, error)
	CheckUsername(username string, phone string) (bool, error)
	GetLogin(phone string, password string) (models.UserModel, error)
}
