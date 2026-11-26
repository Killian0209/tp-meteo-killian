package controllers

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"meteo/helpers"
	"meteo/models"
	"meteo/services"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request : %s %s", r.Method, r.URL.Path)

	helpers.Render(w, http.StatusOK, "index", nil)
}

// ListCommune affiche la liste des Communes
func ListCommune(w http.ResponseWriter, r *http.Request) {
	commune, err := services.GetListCommune()
	if err != nil {
		renderErrorSevice(w, err)
		return
	}
	helpers.Render(w, http.StatusOK, "commune", commune)
}

// SearchCommune affiche les communes correspondant au terme recherché.
func SearchCommune(w http.ResponseWriter, r *http.Request) {
	query := strings.TrimSpace(r.URL.Query().Get("name"))
	if query == "" {
		http.Redirect(w, r, "/commune", http.StatusSeeOther)
		return
	}

	communes, err := services.SearchCommune(query)
	if err != nil {
		renderErrorSevice(w, err)
		return
	}
	helpers.Render(w, http.StatusOK, "commune", communes)
}

// ShowCommune affiche la fiche de la Commune dont l'identifiant est dans l'URL.
func ShowCommune(w http.ResponseWriter, r *http.Request) {
	commune, err := services.GetCommune(r.PathValue("id"))
	if err != nil {
		renderErrorSevice(w, err)
		return
	}
	helpers.Render(w, http.StatusOK, "meteo", commune)
}

// NotFound affiche la page 404 pour toute URL inconnue du site.
func NotFound(w http.ResponseWriter, r *http.Request) {
	helpers.RenderError(w, http.StatusNotFound, "Cette page n'existe pas.")
}

// renderAPIError traduit une erreur connue en page d'erreur, sur la route même.
func renderErrorSevice(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, models.ErrNotFound):
		helpers.RenderError(w, http.StatusNotFound, "Cette commune n'existe pas.")
	case errors.Is(err, models.ErrBadRequest):
		helpers.RenderError(w, http.StatusBadRequest, "Code de Commune invalide.")
	case errors.Is(err, models.ErrTimeout), errors.Is(err, models.ErrUnavailable):
		log.Println("API Commune :", err)
		helpers.RenderError(w, http.StatusServiceUnavailable, "Service momentanément indisponible.")
	default:
		log.Println("Erreur inattendue :", err)
		helpers.RenderError(w, http.StatusInternalServerError, "Une erreur est survenue.")
	}
}
