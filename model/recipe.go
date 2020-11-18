package model

type Recipe struct {
	Name        string
	Ingredients *[]Ingredient
	Products    *[]Product
}

type Ingredient struct {
	Item   *Item
	Amount uint8
}

type Product struct {
	Item   *Item
	Amount uint8
}
