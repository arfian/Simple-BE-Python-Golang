package controllers

import (
	"net/http"
	"encoding/json"
	"time"
	"efishery-task/auth-app/interfaces"
	"efishery-task/auth-app/transformers"
	jwt "github.com/dgrijalva/jwt-go"
)

type AuthController struct {
	interfaces.IAuthService
}

const (
    mySigningKey = "efishery123"
)

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
		token := jwt.New(jwt.GetSigningMethod("HS256"))
		claims := token.Claims.(jwt.MapClaims)
		claims["phone"] = user.Phone
		claims["name"] = user.Name
		claims["role"] = user.Role
		claims["username"] = user.Username
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
		tokenString, err := token.SignedString([]byte(mySigningKey))
		if err != nil {
			json.NewEncoder(w).Encode(transformers.ResData{
				ErrorMessage: err.Error(),
				IsError: true,
			})
		}
		json.NewEncoder(w).Encode(transformers.ResData{
			ErrorMessage: "",
			IsError: false,
			Data: transformers.Login{
				Phone: user.Phone,
				Name: user.Name,
				Role: user.Role,
				Username: user.Username,
				Timestamp: user.CreatedAt.Format(time.RFC850),
				BarerToken: tokenString,
			},
		})
	}
}