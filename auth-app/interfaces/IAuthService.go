package interfaces

import (
	"efishery-task/auth-app/models"
)

type IAuthService interface {
	RegisterUser(phone string, name string, role string) (*models.UserModel, error)
}
