package models

import "time"

type UserModel struct {
	Id			int
	Name		string
	Phone		string
	Role		string
	Password	string
	Username 	string
	CreatedAt	time.Time
	UpdatedAt	time.Time
}
