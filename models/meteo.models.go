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

func (m Meteo) Description(code int) string {
	switch code {
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
	case 51, 53, 55, 56, 57:
		return "Bruine"
	case 61, 63, 65, 66, 67:
		return "Pluie"
	case 71, 73, 75, 77:
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

func (m Meteo) TempClass() string {
	switch {
	case m.Current.Temp < 5:
		return "froid"
	case m.Current.Temp < 20:
		return "doux"
	default:
		return "chaud"
	}
}

var joursSemaine = []string{"Dim", "Lun", "Mar", "Mer", "Jeu", "Ven", "Sam"}

type Jour struct {
	Date  string
	Icone string
	Max   float64
	Min   float64
	Rain  float64
}

func (m Meteo) Jours() []Jour {
	var jours []Jour
	for i := range m.Daily.Time {
		date := m.Daily.Time[i]
		if t, err := time.Parse("2006-01-02", date); err == nil {
			date = fmt.Sprintf("%s %02d/%02d", joursSemaine[t.Weekday()], t.Day(), int(t.Month()))
		}
		if i == 0 {
			date = "Aujourd'hui"
		}
		jours = append(jours, Jour{
			Date:  date,
			Icone: m.Icon(m.Daily.WeatherCode[i]),
			Max:   m.Daily.TempMax[i],
			Min:   m.Daily.TempMin[i],
			Rain:  m.Daily.Rain[i],
		})
	}
	return jours
}
