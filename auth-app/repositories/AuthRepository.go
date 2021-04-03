package repositories

import (
	"efishery-task/auth-app/interfaces"
	"efishery-task/auth-app/models"
	"fmt"
)

type AuthRepository struct {
	interfaces.IDbHandler
}

func (repository *AuthRepository) CheckUsername(username string) (bool, error) {
	row, err :=repository.Query(fmt.Sprintf("SELECT * FROM users WHERE username = '%s'", username))
	if err != nil {
		fmt.Println(err)
	}

	return row.Next(), nil
}

func (repository *AuthRepository) SaveUser(user models.UserModel) (models.UserModel, error) {
	statement, err := repository.Prepare("INSERT INTO users (name, role, username, phone, password) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err)
	}

	statement.Exec(user.Name, user.Role, user.Username, user.Phone, user.Password)
	defer repository.CloseDb()

	return user, nil
}
