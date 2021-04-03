package models

import "time"

type UserModel struct {
	Id			int			`json:"id"`
	Name		string		`json:"name"`
	Phone		string		`json:"phone"`
	Role		string		`json:"role"`
	Password	string		`json:"password"`
	Username 	string		`json:"username"`
	CreatedAt	time.Time	`json:"create-at"`
	UpdatedAt	time.Time	`json:"update-at"`
}
