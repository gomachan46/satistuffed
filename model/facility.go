package model

type Facility struct {
	Recipe   *Recipe
	Amount   float64
	Children *[]FacilityChild
}

type FacilityChild struct {
	Remain   float64
	Facility *Facility
}
