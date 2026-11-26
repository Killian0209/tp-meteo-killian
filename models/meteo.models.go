package models

type Meteo struct {
	Current struct {
		Time        string  `json:"time"`
		Temp        float64 `json:"temperature_2m"`
		Feels       float64 `json:"apparent_temperature"`
		Humidity    int     `json:"relative_humidity_2m"`
		Wind        float64 `json:"wind_speed_10m"`
		WeatherCode int     `json:"weather_code"`
	} `json:"current"`
}

func (m Meteo) Description() string {
	switch m.Current.WeatherCode {
	case 0:
		return "Ciel dégagé"
	case 1:
		return "Peu nuageux"
	case 2:
		return "Partiellement nuageux"
	case 3:
		return "Couvert"
	case 45, 48:
		return "Brouillard"
	case 51, 52, 53, 54, 55, 56, 57:
		return "Bruine"
	case 61, 62, 63, 64, 65, 66, 67:
		return "Pluie"
	case 71, 72, 73, 74, 75, 76, 77:
		return "Neige"
	case 80, 81, 82:
		return "Averses"
	case 85, 86:
		return "Averses de neige"
	case 95, 96, 97, 99:
		return "Orage"
	default:
		return "Conditions inconnues"
	}
}