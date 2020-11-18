package main

type Recipe struct {
	Name        string
	Ingredients *[]Ingredient
	Products    *[]Product
}

type Ingredient struct {
	Item   *Item
	amount uint8
}

type Product struct {
	Item   *Item
	amount uint8
}
