package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"efishery-task/auth-app/controllers"
	"efishery-task/auth-app/services"
	"efishery-task/auth-app/repositories"
	"efishery-task/auth-app/helpers"
)

func InitRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	authRepository := &repositories.AuthRepository{&helpers.SQLiteHandler{}}
	authService := &services.AuthService{
		Repo: authRepository,
		Helper: &helpers.HelperFunc{},
	}
	auth := &controllers.AuthController{authService}

	r.Route("/auth", func(r chi.Router) {
		r.Get("/checkjwt", auth.CheckJwt)
		r.Post("/register", auth.Register)
		r.Post("/login", auth.Login)
	})

	return r
}