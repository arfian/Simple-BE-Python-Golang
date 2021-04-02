package interfaces

import (
	"efishery-task/auth-app/models"
)

type IAuthRepository interface {
	SaveUser(models.UserModel) (models.UserModel, error)
	GetUsername(username string) (*models.UserModel, error)
}
