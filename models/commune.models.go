package models

import "strings"

//structure commune

type Commune struct {
	Name        string   `json:"nom"`
	Code        string   `json:"code"`
	CodesPostaux []string `json:"codesPostaux"`
	Center      Center   `json:"centre"`
	Pop         int      `json:"population"`
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

func (c Commune) CodePostal() string {
	return strings.Join(c.CodesPostaux, ", ")
}

func (c Commune) LonMin() float64 {
	return c.Longitude() - 0.08 
}

func (c Commune) LatMin() float64 { 
	return c.Latitude() - 0.05 
}

func (c Commune) LonMax() float64 { 
	return c.Longitude() + 0.08 
}

func (c Commune) LatMax() float64 { 
	return c.Latitude() + 0.05 
}