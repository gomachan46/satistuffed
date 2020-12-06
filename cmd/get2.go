/*
Copyright © 2020 gomachan46 <shiro.gomachan46@gmail.com>

*/
package cmd

import (
	"fmt"
	"github.com/gomachan46/satistuffed/cli"
	"github.com/gomachan46/satistuffed/data"
	"github.com/gomachan46/satistuffed/model"
	"github.com/spf13/cobra"
	"math"
	"os"
)

//var isOrigin bool

// get2Cmd represents the get2 command
var get2Cmd = &cobra.Command{
	Use:   "get2",
	Short: "Get stuffed items",
	Run: func(cmd *cobra.Command, args []string) {
		d := data.Load()

		for _, name := range args {
			item, err := d.GetItem(name)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			facility := a2(item)
			hoge := b2([]*model.Product{}, facility)
			sum := map[string]float64{}
			for _, v := range hoge {
				sum[v.Item.Name] += v.Amount
			}

			for name, amount := range sum {
				fmt.Printf("%s: %g\n", name, amount)
			}

			cli.Draw(facility, depth)
		}
	},
}

func init() {
	rootCmd.AddCommand(get2Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//get2Cmd.PersistentFlags().BoolVar(&isOrigin, "origin", false, "get only origin")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func a2(item *model.Item) *model.Facility {
	recipe := item.Recipes[0]
	product := recipe.Products[0]
	ingredients := recipe.Ingredients

	if len(ingredients) < 1 {
		return &model.Facility{Recipe: recipe, Amount: product.Amount, Children: []*model.FacilityChild{}}
	}

	children := []*model.FacilityChild{}

	for _, ingredient := range ingredients {
		ingredientItem := ingredient.Item
		ingredientRecipe := ingredientItem.Recipes[0]
		ingredientProduct := ingredientRecipe.Products[0]

		if ingredient.Amount <= ingredientProduct.Amount {
			children = append(
				children,
				&model.FacilityChild{
					Remain:   ingredientProduct.Amount - ingredient.Amount,
					Facility: a2(ingredientItem),
				},
			)
		} else {
			magnification := math.Ceil(ingredient.Amount / ingredientProduct.Amount)
			for i := 0; i < int(magnification)-1; i++ {
				children = append(
					children,
					&model.FacilityChild{
						Remain:   0,
						Facility: a2(ingredientItem),
					},
				)
			}
			children = append(
				children,
				&model.FacilityChild{
					Remain:   magnification*ingredientProduct.Amount - ingredient.Amount,
					Facility: a2(ingredientItem),
				},
			)
		}
	}

	return &model.Facility{Recipe: recipe, Amount: product.Amount, Children: children}
}

func b2(origins []*model.Product, facility *model.Facility) []*model.Product {
	o := []*model.Product{}
	for _, v := range facility.Children {
		o = append(o, c2(v))
		o = append(o, b2([]*model.Product{}, v.Facility)...)
	}

	return append(origins, o...)
}

func c2(child *model.FacilityChild) *model.Product {
	product := child.Facility.Recipe.Products[0]
	return &model.Product{
		Item:   product.Item,
		Amount: child.Facility.Amount - child.Remain,
	}
}
