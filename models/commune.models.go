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
