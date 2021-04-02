package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"efishery-task/auth-app/controllers"
	"efishery-task/auth-app/services"
	"efishery-task/auth-app/repositories"
)

func InitRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	authService := &services.AuthService{&repositories.AuthRepository{}}
	auth := &controllers.AuthController{authService}

	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", auth.Register)
	})

	return r
}