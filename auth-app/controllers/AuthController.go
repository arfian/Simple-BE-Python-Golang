package controllers

import (
	"net/http"
	"encoding/json"
	"efishery-task/auth-app/interfaces"
	"efishery-task/auth-app/transformers"
)

type AuthController struct {
	interfaces.IAuthService
}

func (controller *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	name := r.Context().Value("name").(string)
	phone := r.Context().Value("phone").(string)
	role := r.Context().Value("role").(string)

	user, err := controller.RegisterUser(phone, name, role)
	if err == nil {
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(transformers.Register{
		Phone: user.Phone,
		Username: user.Username,
		Password: user.Password,
	})
}