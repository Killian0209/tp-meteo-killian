package routes

import (
	"meteo/controllers"
	"net/http"
)

func MainRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// Fichiers statiques : /static/css/style.css → assets/css/style.css
	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fs))

	communeRoutes(mux)

	// Toute autre URL : page 404 du site
	mux.HandleFunc("/", controllers.NotFound)
	return mux
}