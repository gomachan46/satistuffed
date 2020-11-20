/*
Copyright © 2020 gomachan46 <shiro.gomachan46@gmail.com>

*/
package cmd

import (
	"fmt"
	"github.com/gomachan46/satistuffed/cli"
	"github.com/gomachan46/satistuffed/model"
	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
	"math"
)

// getCmd represents the get command
var maxCmd = &cobra.Command{
	Use:   "max",
	Short: "Get a max Facility",
	Run: func(cmd *cobra.Command, args []string) {
		hoge()
	},
}

func init() {
	rootCmd.AddCommand(maxCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func fuga() {
	ironOreRecipe := &model.Recipe{Name: "鉄鉱石のレシピ1"}
	ironOre := &model.Item{Name: "鉄鉱石"}
	ironOreRecipe.Ingredients = &[]model.Ingredient{}
	ironOreRecipe.Products = &[]model.Product{{Item: ironOre, Amount: 60}}
	ironOre.Recipes = &[]model.Recipe{*ironOreRecipe}

	ironIngotRecipe := &model.Recipe{Name: "鉄インゴットのレシピ1"}
	ironIngot := &model.Item{Name: "鉄インゴット"}
	ironIngotRecipe.Ingredients = &[]model.Ingredient{{
		Item:   ironOre,
		Amount: 30,
	}}
	ironIngotRecipe.Products = &[]model.Product{{
		Item:   ironIngot,
		Amount: 30,
	}}
	ironIngot.Recipes = &[]model.Recipe{*ironIngotRecipe}

	ironPlateRecipe := &model.Recipe{Name: "鉄板のレシピ1"}
	ironPlate := &model.Item{Name: "鉄板"}
	ironPlateRecipe.Ingredients = &[]model.Ingredient{{
		Item:   ironIngot,
		Amount: 30,
	}}
	ironPlateRecipe.Products = &[]model.Product{{
		Item:   ironPlate,
		Amount: 20,
	}}
	ironPlate.Recipes = &[]model.Recipe{*ironPlateRecipe}

	ironRodRecipe := &model.Recipe{Name: "鉄のロッドのレシピ1"}
	ironRod := &model.Item{Name: "鉄のロッド"}
	ironRodRecipe.Ingredients = &[]model.Ingredient{{
		Item:   ironIngot,
		Amount: 15,
	}}
	ironRodRecipe.Products = &[]model.Product{{
		Item:   ironRod,
		Amount: 15,
	}}
	ironRod.Recipes = &[]model.Recipe{*ironRodRecipe}

	screwRecipe := &model.Recipe{Name: "ネジのレシピ1"}
	screw := &model.Item{Name: "ネジ"}
	screwRecipe.Ingredients = &[]model.Ingredient{{
		Item:   ironRod,
		Amount: 10,
	}}
	screwRecipe.Products = &[]model.Product{{
		Item:   screw,
		Amount: 40,
	}}
	screw.Recipes = &[]model.Recipe{*screwRecipe}

	reinforcedIronPlateRecipe := &model.Recipe{Name: "強化鉄板のレシピ1"}
	reinforcedIronPlate := &model.Item{Name: "強化鉄板"}
	reinforcedIronPlateRecipe.Ingredients = &[]model.Ingredient{
		{
			Item:   ironPlate,
			Amount: 30,
		},
		{
			Item:   screw,
			Amount: 60,
		},
	}
	reinforcedIronPlateRecipe.Products = &[]model.Product{{
		Item:   reinforcedIronPlate,
		Amount: 5,
	}}
	reinforcedIronPlate.Recipes = &[]model.Recipe{*reinforcedIronPlateRecipe}

	pp.Println("テスト1: 鉄鉱石")
	cli.Draw(b(ironOre))
	pp.Println("テスト2: 鉄インゴット")
	cli.Draw(b(ironIngot))
	pp.Println("テスト3: 鉄板")
	cli.Draw(b(ironPlate))
	pp.Println("テスト4: 強化鉄板")
	cli.Draw(b(reinforcedIronPlate))
}

func b(item *model.Item) *model.Facility {
	recipe := (*item.Recipes)[0]
	ingredients := *recipe.Ingredients

	if len(ingredients) < 1 {
		return &model.Facility{Recipe: &recipe, Children: &[]model.Facility{}}
	}

	children := []model.Facility{}

	for _, ingredient := range *recipe.Ingredients {
		ingredientItem := ingredient.Item
		ingredientRecipe := (*ingredientItem.Recipes)[0]
		ingredientProduct := (*ingredientRecipe.Products)[0]

		if ingredient.Amount <= ingredientProduct.Amount {
			return &model.Facility{Recipe: &recipe, Children: &[]model.Facility{*a(ingredientItem)}}
		}

		magnification := int(math.Ceil(float64(ingredient.Amount) / float64(ingredientProduct.Amount)))
		ingredientRecipe.Name = fmt.Sprintf("%s x %d", ingredientRecipe.Name, magnification)
		stuffedFacility := &model.Facility{Recipe: &ingredientRecipe, Children: &[]model.Facility{}}
		for i := 0; i < magnification; i++ {
			*stuffedFacility.Children = append(*stuffedFacility.Children, *a(ingredientItem))
		}
		children = append(children, *stuffedFacility)
	}

	return &model.Facility{Recipe: &recipe, Children: &children}
}
