/*
Copyright © 2020 gomachan46 <shiro.gomachan46@gmail.com>

*/
package cmd

import (
	"fmt"
	"github.com/gomachan46/satistuffed/model"
	"github.com/spf13/cobra"
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
	ironOreRecipe.Products = &[]model.Product{{Item: ironOre, Amount: 30}}
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
	fmt.Println(ironIngot.Recipes)
	fuga(ironIngot)

	fmt.Println("get called")
}

func fuga(item *model.Item) {
	fmt.Println(item)
	recipe := (*item.Recipes)[0]
	piyo(&recipe)
}

func piyo(recipe *model.Recipe) {
	ingredients := *recipe.Ingredients
	fmt.Println(ingredients)
	if len(ingredients) < 1 {
		product := (*recipe.Products)[0]
		fmt.Println("finish!")
		fmt.Println(product)
		return
	}

	ingredient := ingredients[0]
	item := ingredient.Item
	fmt.Println(ingredient)
	fmt.Println(ingredient.Item)
	fmt.Println(item)
	fmt.Println(item.Recipes)
	piyo(&(*item.Recipes)[0])
}
