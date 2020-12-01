package model

type Recipe struct {
	Name        string
	Ingredients []*Ingredient
	Products    []*Product
}

type Ingredient struct {
	Item   *Item
	Amount float64
}

type Product struct {
	Item   *Item
	Amount float64
}
