package routes

import (
	"net/http"

	"github.com/wikixen/blogapp/api/handlers"
)

func BlogRoutes(app *http.ServeMux) (*http.ServeMux){
	app.HandleFunc("POST /", handlers.CreateBlog)
	app.HandleFunc("GET /", handlers.GetBlogs)
	app.HandleFunc("GET /{id}", handlers.GetABlog)
	app.HandleFunc("PATCH /{id}", handlers.EditABlog)
	app.HandleFunc("DELETE /{id}", handlers.DeleteABlog)

	return app
}