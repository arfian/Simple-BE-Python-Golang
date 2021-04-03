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
	r.ParseForm()

	user, err := controller.RegisterUser(r.FormValue("phone"), r.FormValue("name"), r.FormValue("role"))
	if err != nil {
		json.NewEncoder(w).Encode(transformers.ResData{
			ErrorMessage: err.Error(),
			IsError: true,
		})
	} else {
		json.NewEncoder(w).Encode(transformers.ResData{
			ErrorMessage: "",
			IsError: false,
			Data: transformers.Register{
				Phone: user.Phone,
				Username: user.Username,
				Password: user.Password,
			},
		})
	}
}