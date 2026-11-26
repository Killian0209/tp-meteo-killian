package routes

import (
	"meteo/controllers"
	"net/http"
)

func communeRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /{$}", controllers.HomeHandler)
	mux.HandleFunc("GET /commune", controllers.ListCommune)
	mux.HandleFunc("Get /commune/search", controllers.SearchCommune)
	mux.HandleFunc("GET /commune/{id}", controllers.ShowCommune)
}


