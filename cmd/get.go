/*
Copyright © 2020 gomachan46 <shiro.gomachan46@gmail.com>

*/
package cmd

import (
	"github.com/gomachan46/satistuffed/model"
	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
	"math"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a stuffed Item",
	Run: func(cmd *cobra.Command, args []string) {
		hoge()
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func hoge() {
	ironOreRecipe := &model.Recipe{Name: "Iron Ore 1"}
	ironOre := &model.Item{Name: "Iron Ore"}
	ironOreRecipe.Ingredients = &[]model.Ingredient{}
	ironOreRecipe.Products = &[]model.Product{{Item: ironOre, Amount: 60}}
	ironOre.Recipes = &[]model.Recipe{*ironOreRecipe}

	ironIngotRecipe := &model.Recipe{Name: "Iron Ingot 1"}
	ironIngot := &model.Item{Name: "Iron Ingot"}
	ironIngotRecipe.Ingredients = &[]model.Ingredient{{
		Item:   ironOre,
		Amount: 30,
	}}
	ironIngotRecipe.Products = &[]model.Product{{
		Item:   ironIngot,
		Amount: 30,
	}}
	ironIngot.Recipes = &[]model.Recipe{*ironIngotRecipe}

	ironPlateRecipe := &model.Recipe{Name: "Iron Plate 1"}
	ironPlate := &model.Item{Name: "Iron Plate"}
	ironPlateRecipe.Ingredients = &[]model.Ingredient{{
		Item:   ironIngot,
		Amount: 30,
	}}
	ironPlateRecipe.Products = &[]model.Product{{
		Item:   ironPlate,
		Amount: 20,
	}}
	ironPlate.Recipes = &[]model.Recipe{*ironPlateRecipe}

	reinforcedIronPlateRecipe := &model.Recipe{Name: "Reinforced Iron Plate 1"}
	reinforcedIronPlate := &model.Item{Name: "Reinforced Iron Plate"}
	reinforcedIronPlateRecipe.Ingredients = &[]model.Ingredient{{
		Item:   ironPlate,
		Amount: 30,
	}}
	reinforcedIronPlateRecipe.Products = &[]model.Product{{
		Item:   reinforcedIronPlate,
		Amount: 5,
	}}
	reinforcedIronPlate.Recipes = &[]model.Recipe{*reinforcedIronPlateRecipe}

	results := fuga(reinforcedIronPlate)
	for i, v := range results {
		pp.Printf("No. %v\n", i)
		pp.Println(v.Item.Name)
	}
}

func fuga(item *model.Item) []model.Product {
	recipe := (*item.Recipes)[0]
	product := (*recipe.Products)[0]
	ingredients := *recipe.Ingredients
	products := []model.Product{product}

	if len(ingredients) < 1 {
		return products
	}

	ingredient := (*recipe.Ingredients)[0]
	ingredientItem := ingredient.Item
	ingredientRecipe := (*ingredientItem.Recipes)[0]
	ingredientProduct := (*ingredientRecipe.Products)[0]

	if ingredient.Amount <= ingredientProduct.Amount {
		if len(*ingredientRecipe.Ingredients) < 1 {
			return append(products, ingredientProduct)
		}

		return append(products, fuga(ingredientItem)...)
	}

	ingredientProducts := []model.Product{}
	if ingredient.Amount > ingredientProduct.Amount {
		magnification := int(math.Ceil(float64(ingredient.Amount) / float64(ingredientProduct.Amount)))
		for i := 0; i < magnification; i++ {
			ingredientProducts = append(ingredientProducts, fuga(ingredientItem)...)
		}
	}

	return append(products, ingredientProducts...)
}

func a(ingredient *model.Ingredient) {
	item := ingredient.Item
	recipe := (*item.Recipes)[0]

	b(ingredient.Amount, &recipe)
}

func b(needed uint8, recipe *model.Recipe) *model.Recipe {
	if needed <= piyo(recipe) {
		pp.Println("needed!")
		return recipe
	}

	return recipe
}

func piyo(recipe *model.Recipe) uint8 {
	ingredients := *recipe.Ingredients
	if len(ingredients) < 1 {
		product := (*recipe.Products)[0]
		pp.Println("piyo!")
		return product.Amount
	}

	ingredient := ingredients[0]
	item := ingredient.Item
	return piyo(&(*item.Recipes)[0])
}
