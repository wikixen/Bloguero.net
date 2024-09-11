package routes

import (
	"net/http"

	"github.com/wikixen/blogapp/api/handlers"
)

func UserRoutes(app *http.ServeMux) (*http.ServeMux){
	app.HandleFunc("POST /", handlers.CreateUser)
	app.HandleFunc("GET /", handlers.GetAllUsers)
	app.HandleFunc("GET /{id}", handlers.GetAUser)
	app.HandleFunc("PATCH /{id}", handlers.EditUser)
	app.HandleFunc("DELETE /{id}", handlers.DeleteUser)

	return app
}