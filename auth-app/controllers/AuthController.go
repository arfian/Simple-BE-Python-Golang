package controllers

import (
	"net/http"
	"encoding/json"
	"time"
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

func (controller *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := controller.CheckLogin(r.FormValue("phone"), r.FormValue("password"))
	if err != nil {
		json.NewEncoder(w).Encode(transformers.ResData{
			ErrorMessage: err.Error(),
			IsError: true,
		})
	} else {
		json.NewEncoder(w).Encode(transformers.ResData{
			ErrorMessage: "",
			IsError: false,
			Data: transformers.Login{
				Phone: user.Phone,
				Name: user.Name,
				Role: user.Role,
				Username: user.Username,
				Timestamp: user.CreatedAt.Format(time.RFC850),
				BarerToken: "",
			},
		})
	}
}