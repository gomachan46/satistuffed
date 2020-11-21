package model

type Facility struct {
	Recipe   *Recipe
	Amount   uint8
	Children *[]FacilityChild
}

type FacilityChild struct {
	Remain   uint8
	Facility *Facility
}
