package main

import (
	"net/http"

	"github.com/wikixen/blogapp/api/routes"
	"github.com/wikixen/blogapp/config"
)

func main() {
	env := config.GetConfig()
	app := http.NewServeMux()

	app = routes.BlogRoutes(app)
	app = routes.UserRoutes(app)

	http.ListenAndServe(env.Port, app)
}