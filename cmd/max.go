/*
Copyright © 2020 gomachan46 <shiro.gomachan46@gmail.com>

*/
package cmd

import (
	"fmt"
	"github.com/gomachan46/satistuffed/cli"
	"github.com/gomachan46/satistuffed/data"
	"github.com/gomachan46/satistuffed/model"
	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
	"math"
	"os"
)

// getCmd represents the get command
var maxCmd = &cobra.Command{
	Use:   "max",
	Short: "Get a max Facility",
	Run: func(cmd *cobra.Command, args []string) {
		fuga()
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
	d := data.Load()
	screw, err := d.GetItem("ネジ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pp.Println("テスト: ネジ")
	cli.Draw(b(screw))
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
			return &model.Facility{Recipe: &recipe, Children: &[]model.Facility{*b(ingredientItem)}}
		}

		magnification := int(math.Ceil(float64(ingredient.Amount) / float64(ingredientProduct.Amount)))
		ingredientRecipe.Name = fmt.Sprintf("%s x %d", ingredientRecipe.Name, magnification)
		stuffedFacility := &model.Facility{Recipe: &ingredientRecipe, Children: &[]model.Facility{}}
		for i := 0; i < magnification; i++ {
			*stuffedFacility.Children = append(*stuffedFacility.Children, *b(ingredientItem))
		}
		children = append(children, *stuffedFacility)
	}

	return &model.Facility{Recipe: &recipe, Children: &children}
}
