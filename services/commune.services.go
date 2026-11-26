package services

import (
	"strconv"

	"meteo/api"
	"meteo/models"
)

// GetListCommune renvoie la liste des Communes.
func GetListCommune() ([]models.Commune, error) {
	return api.ListCommune()
}

func GetCommune(idParam string) (models.Commune, error) {
	id, idErr := strconv.Atoi(idParam)
	if len(idParam) != 5 && idErr != nil {
		return models.Commune{}, models.ErrBadRequest
	}
	return api.GetCommune(id)
}

