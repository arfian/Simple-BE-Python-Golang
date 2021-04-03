package interfaces

import (
	"efishery-task/auth-app/models"
)

type IAuthRepository interface {
	SaveUser(models.UserModel) (*models.UserModel, error)
	CheckUsername(username string) (bool, error)
	GetLogin(phone string, password string) (models.UserModel, error)
}
