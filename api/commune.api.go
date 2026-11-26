package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"meteo/models"
	"net"
	"net/http"
	"net/url"
	"time"
)

const (
	baseURL   = "https://geo.api.gouv.fr/communes"
	fields    = "fields=nom,code,centre,population"
	userAgent = "meteo-killian"
)

var client = &http.Client{
	Timeout: 5 * time.Second,
}

func get(fullURL string, dest any) error {
	request, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return fmt.Errorf("création de la requête : %w", err)
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("User-Agent", userAgent)

	response, err := client.Do(request)
	if err != nil {

		var netErr net.Error
		if errors.As(err, &netErr) && netErr.Timeout() {
			return models.ErrTimeout
		}
		return fmt.Errorf("%w : %v", models.ErrUnavailable, err)
	}
	defer response.Body.Close()

	switch {
	case response.StatusCode == http.StatusOK:

	case response.StatusCode == http.StatusNotFound:
		return models.ErrNotFound
	case response.StatusCode == http.StatusBadRequest:
		return models.ErrBadRequest
	case response.StatusCode >= 500:
		return fmt.Errorf("%w : %s", models.ErrUnavailable, response.Status)
	default:
		return fmt.Errorf("réponse inattendue : %s", response.Status)
	}

	if err := json.NewDecoder(response.Body).Decode(dest); err != nil {
		return fmt.Errorf("décodage de la réponse : %w", err)
	}
	return nil
}

func GetListCommune() ([]models.Commune, error) {
	var communes []models.Commune
	err := get(baseURL+"?"+fields, &communes)
	return communes, err
}

func GetCommune(code string) (models.Commune, error) {
	var commune models.Commune
	err := get(fmt.Sprintf("%s/%s?%s", baseURL, url.PathEscape(code), fields), &commune)
	return commune, err
}

func SearchCommuneByName(query string) ([]models.Commune, error) {
	var communes []models.Commune
	err := get(baseURL+"?nom="+url.QueryEscape(query)+"&limit=20&"+fields, &communes)
	return communes, err
}


 
