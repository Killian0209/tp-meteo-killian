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

func get(endpoint string, dest any) error {

	request, err := http.NewRequest(http.MethodGet, baseURL+endpoint, nil)
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
	var commune []models.Commune
	err := get("/commune", &commune)
	return commune, err
}

func GetCommune(code int) (models.Commune, error) {
	var commune models.Commune
	err := get(fmt.Sprintf("/commune/%d", code), &commune)
	return commune, err
}

func SearchCommuneByName(query string) ([]models.Commune, error) {
	resp, err := http.Get("https://geo.api.gouv.fr/communes?nom=" + url.QueryEscape(query) + "&fields=nom,code,codesPostaux&limit=20")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API communes: statut %d", resp.StatusCode)
	}

	var communes []models.Commune
	if err := json.NewDecoder(resp.Body).Decode(&communes); err != nil {
		return nil, err
	}
	return communes, nil
}
