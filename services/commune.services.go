package services

import (
	"strconv"
	"strings"

	"meteo/api"
	"meteo/models"
)

// GetListCommune renvoie la liste des Communes.
func GetListCommune() ([]models.Commune, error) {
	return api.GetListCommune()
}

func GetCommune(idParam string) (models.Commune, error) {
	id, idErr := strconv.Atoi(idParam)
	if len(idParam) != 5 || idErr != nil {
		return models.Commune{}, models.ErrBadRequest
	}
	return api.GetCommune(id)
}

func SearchCommune(query string) ([]models.Commune, error) {
	query = strings.TrimSpace(query)
	if query == "" {
		return nil, models.ErrBadRequest
	}
	return api.SearchCommuneByName(query)
}





