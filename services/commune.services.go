package services

import (

	"strings"

	"meteo/api"
	
	"meteo/models"
)

// GetListCommune renvoie la liste des Communes.
func GetListCommune() ([]models.Commune, error) {
	return api.GetListCommune()
}

func GetCommune(code string) (models.Commune, error) {
	code = strings.ToUpper(strings.TrimSpace(code))
	if len(code) != 5 {
		return models.Commune{}, models.ErrBadRequest
	}
	return api.GetCommune(code)
}

func SearchCommune(query string) ([]models.Commune, error) {
	query = strings.TrimSpace(query)
	if query == "" {
		return nil, models.ErrBadRequest
	}
	return api.SearchCommuneByName(query)
}

type Fiche struct {
	Commune models.Commune
	Meteo   models.Meteo
}

func GetFiche(idParam string) (Fiche, error) {
	commune, err := GetCommune(idParam)
	if err != nil {
		return Fiche{}, err
	}

	meteo, err := api.GetMeteo(commune.Latitude(), commune.Longitude())
	if err != nil {
		return Fiche{}, err
	}

	return Fiche{Commune: commune, Meteo: meteo}, nil
}





