package api

import (
	"fmt"

	"meteo/models"
)

const meteoURL = "https://api.open-meteo.com/v1/forecast"

func GetMeteo(lat, lon float64) (models.Meteo, error) {
	var meteo models.Meteo
	endpoint := fmt.Sprintf("%s?latitude=%f&longitude=%f&current=temperature_2m,apparent_temperature,relative_humidity_2m,wind_speed_10m,weather_code&timezone=auto", meteoURL, lat, lon)
	err := get(endpoint, &meteo)
	return meteo, err
}

