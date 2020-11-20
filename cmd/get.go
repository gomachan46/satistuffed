/*
Copyright © 2020 gomachan46 <shiro.gomachan46@gmail.com>

*/
package cmd

import (
	"fmt"
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

	pp.Println("Iron Ore")
	draw(a(ironOre))
	pp.Println("Iron Ingot")
	draw(a(ironIngot))
	pp.Println("Iron Plate")
	draw(a(ironPlate))
	pp.Println("Reinforced Iron Plate")
	draw(a(reinforcedIronPlate))
}

func draw(facility *model.Facility) {
	fmt.Printf(" %s\n", facility.Recipe.Name)
	drawChildren("", facility)
	fmt.Println("")
}

func drawChildren(indent string, facility *model.Facility) {
	children := *facility.Children

	for i, v := range children {
		s := indent + " ├─"
		a := " │ "
		if i == len(children)-1 {
			s = indent + " └─"
			a = "   "
		}

		fmt.Printf("%s %s\n", s, v.Recipe.Name)

		drawChildren(indent+a, &v)
	}
}

func a(item *model.Item) *model.Facility {
	recipe := (*item.Recipes)[0]
	ingredients := *recipe.Ingredients

	if len(ingredients) < 1 {
		return &model.Facility{Recipe: &recipe, Children: &[]model.Facility{}}
	}

	ingredient := (*recipe.Ingredients)[0]
	ingredientItem := ingredient.Item
	ingredientRecipe := (*ingredientItem.Recipes)[0]
	ingredientProduct := (*ingredientRecipe.Products)[0]

	if ingredient.Amount <= ingredientProduct.Amount {
		return &model.Facility{Recipe: &recipe, Children: &[]model.Facility{*a(ingredientItem)}}
	}

	magnification := int(math.Ceil(float64(ingredient.Amount) / float64(ingredientProduct.Amount)))
	children := []model.Facility{}
	for i := 0; i < magnification; i++ {
		children = append(children, *a(ingredientItem))
	}

	return &model.Facility{Recipe: &recipe, Children: &children}
}
