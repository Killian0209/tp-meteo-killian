package routes

import (
	"meteo/controllers"
	"net/http"
)

func communeRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /{$}", controllers.ListCommune)
	mux.HandleFunc("GET /pokemon/{id}", controllers.ShowCommune)
}


