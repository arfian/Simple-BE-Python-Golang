package controllers

import (
	"net/http"
	"encoding/json"
	"time"
	"efishery-task/auth-app/interfaces"
	"efishery-task/auth-app/transformers"
	jwt "github.com/dgrijalva/jwt-go"
	"strings"
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
		claims := &jwt.MapClaims{
			"exp": time.Now().Add(time.Hour * 72).Unix(),
			"data": map[string]string{
				"phone": user.Phone,
				"name": user.Name,
				"role": user.Role,
				"username": user.Username,
				"timestamp": user.CreatedAt.Format(time.RFC850),
			},
		}
		token := jwt.NewWithClaims(
			jwt.SigningMethodHS256,
			claims)
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

func (controller *AuthController) CheckJwt(w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Authorization")
	jwtString := strings.Split(header, "Bearer ")[1]
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(jwtString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(mySigningKey), nil
	})

	if !token.Valid || err != nil {
		json.NewEncoder(w).Encode(transformers.ResData{
			ErrorMessage: "Token tidak valid",
			IsError: true,
		})
	} else {
		data := claims["data"].(map[string]interface{})
		json.NewEncoder(w).Encode(transformers.ResData{
			ErrorMessage: "",
			IsError: false,
			Data: data,
		})
	}
}