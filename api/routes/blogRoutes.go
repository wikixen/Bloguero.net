package routes

import (
	"net/http"

	"github.com/wikixen/blogapp/api/handlers"
	"github.com/wikixen/blogapp/api/middleware"
)

func BlogRoutes(app *http.ServeMux) (*http.ServeMux){
	app.HandleFunc("GET /", handlers.GetBlogs)
	app.HandleFunc("GET /{id}", handlers.GetABlog)

	// Requires a token
	createBlog := http.HandlerFunc(handlers.CreateBlog)
	editBlog := http.HandlerFunc(handlers.EditBlog)
	delBlog := http.HandlerFunc(handlers.DeleteBlog)
	
	app.Handle("POST /", middleware.AccessHandler(createBlog))
	app.Handle("PATCH /{id}", middleware.AccessHandler(editBlog))
	app.Handle("DELETE /{id}", middleware.AccessHandler(delBlog))

	return app
}