package models

//structure commune

type Commune struct {
	Name string `json:"nom"`
	Code string `json:"code"`
	Center Center `json:"centre"`
	Score float64 `json:"_score"`
}

type Center struct {
	Type string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}


