package model

type Facility struct {
	Recipe   *Recipe
	Children *[]Facility
}
