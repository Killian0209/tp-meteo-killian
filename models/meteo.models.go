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
	Daily struct {
		Time        []string  `json:"time"`
		WeatherCode []int     `json:"weather_code"`
		TempMax     []float64 `json:"temperature_2m_max"`
		TempMin     []float64 `json:"temperature_2m_min"`
		Rain        []float64 `json:"precipitation_sum"`
	} `json:"daily"`
}

func (m Meteo) Icon(code int) string {
	switch code {
	case 0:
		return "☀️"
	case 1, 2:
		return "🌤️"
	case 3:
		return "☁️"
	case 45, 48:
		return "🌫️"
	case 51, 53, 55:
		return "🌦️"
	case 61, 63, 65, 80, 81, 82:
		return "🌧️"
	case 71, 73, 75, 85, 86:
		return "❄️"
	case 95, 96, 97, 99:
		return "⛈️"
	default:
		return "🌡️"
	}
}

func (m Meteo) TempClass() string {
	switch {
	case m.Current.Temp < 5:
		return "cold"
	case m.Current.Temp < 20:
		return "mild"
	default:
		return "hot"
	}
}