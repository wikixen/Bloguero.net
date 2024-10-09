package routes

import (
	"net/http"

	"github.com/wikixen/blogapp/api/handlers"
	"github.com/wikixen/blogapp/api/middleware"
)

func UserRoutes(app *http.ServeMux) (*http.ServeMux){
	app.HandleFunc("POST /user", handlers.CreateUser)
	app.HandleFunc("POST /user/{id}", handlers.LoginUser)
	app.HandleFunc("GET /user", handlers.GetAllUsers)
	app.HandleFunc("GET /user/{id}", handlers.GetAUser)

	// Require a Token
	editUser := http.HandlerFunc(handlers.EditUser)
	delUser := http.HandlerFunc(handlers.DeleteUser)
	
	app.Handle("PATCH /user/{id}", middleware.AccessHandler(editUser))
	app.Handle("DELETE /user/{id}", middleware.AccessHandler(delUser))

	return app
}