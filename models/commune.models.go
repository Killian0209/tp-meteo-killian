package models

//structure commune

type Commune struct {
	Name   string `json:"nom"`
	Code   string `json:"code"`
	Center Center `json:"centre"`
	Pop    int    `json:"population"`
}

type Center struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

func (c Commune) Longitude() float64 {
	if len(c.Center.Coordinates) < 2 {
		return 0
	}
	return c.Center.Coordinates[0]
}

func (c Commune) Latitude() float64 {
	if len(c.Center.Coordinates) < 2 {
		return 0
	}
	return c.Center.Coordinates[1]
}