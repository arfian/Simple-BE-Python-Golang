package repositories

import (
	"efishery-task/auth-app/interfaces"
	"efishery-task/auth-app/models"
	"fmt"
)

type AuthRepository struct {
	interfaces.IDbHandler
}

func (repository *AuthRepository) CheckUsername(username string, phone string) (bool, error) {
	row, err :=repository.Query(fmt.Sprintf("SELECT * FROM users WHERE username = '%s' AND phone = '%s'", username, phone))
	if err != nil {
		fmt.Println(err)
		return true, err
	}

	return row.Next(), nil
}

func (repository *AuthRepository) SaveUser(user models.UserModel) (*models.UserModel, error) {
	_, err := repository.Prepare("INSERT INTO users (name, role, username, phone, password) VALUES (?, ?, ?, ?, ?)", user.Name, user.Role, user.Username, user.Phone, user.Password)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &user, nil
}

func (repository *AuthRepository) GetLogin(phone string, password string) (models.UserModel, error){
	row, err :=repository.Query(fmt.Sprintf("SELECT name, phone, role, username, created_at FROM users WHERE phone = '%s' AND password = '%s'", phone, password))
	if err != nil {
		fmt.Println(err)
		return models.UserModel{}, err
	}
	var user models.UserModel
	for row.Next() {
		row.Scan(&user.Name, &user.Phone, &user.Role, &user.Username, &user.CreatedAt)
	}

	return user, nil
}
