package api

import (
	"fmt"

	"meteo/models"
)

const meteoURL = "https://api.open-meteo.com/v1/forecast"

func GetMeteo(lat, lon float64) (models.Meteo, error) {
	var meteo models.Meteo
	endpoint := fmt.Sprintf(
	"%s?latitude=%f&longitude=%f&current=temperature_2m,apparent_temperature,relative_humidity_2m,wind_speed_10m,weather_code&daily=weather_code,temperature_2m_max,temperature_2m_min,precipitation_sum&forecast_days=7&timezone=auto", meteoURL, lat, lon)
	err := get(endpoint, &meteo)
	return meteo, err
}

