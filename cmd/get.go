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

var depth int
var isOrigin bool

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a stuffed Item",
	Run: func(cmd *cobra.Command, args []string) {
		d := data.Load()

		for _, name := range args {
			item, err := d.GetItem(name)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			facility := a(item)
			if isOrigin {
				filteredFacility := &model.Facility{
					Recipe:   facility.Recipe,
					Amount:   facility.Amount,
					Children: &[]model.FacilityChild{},
				}
				cli.Draw(b(filteredFacility, facility), depth)
			} else {
				cli.Draw(facility, depth)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	getCmd.PersistentFlags().IntVarP(&depth, "depth", "d", 10, "depth")
	getCmd.PersistentFlags().BoolVar(&isOrigin, "origin", false, "get only origin")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func a(item *model.Item) *model.Facility {
	recipe := (*item.Recipes)[0]
	product := (*recipe.Products)[0]
	ingredients := *recipe.Ingredients

	if len(ingredients) < 1 {
		return &model.Facility{Recipe: &recipe, Amount: product.Amount, Children: &[]model.FacilityChild{}}
	}

	children := []model.FacilityChild{}

	for _, ingredient := range ingredients {
		ingredientItem := ingredient.Item
		ingredientRecipe := (*ingredientItem.Recipes)[0]
		ingredientProduct := (*ingredientRecipe.Products)[0]

		if ingredient.Amount <= ingredientProduct.Amount {
			children = append(
				children,
				model.FacilityChild{
					Remain:   ingredientProduct.Amount - ingredient.Amount,
					Facility: a(ingredientItem),
				},
			)
		} else {
			magnification := math.Ceil(ingredient.Amount / ingredientProduct.Amount)
			ingredientRecipe.Name = fmt.Sprintf("%s x %d", ingredientRecipe.Name, int(magnification))
			stuffedFacility := &model.Facility{
				Recipe:   &ingredientRecipe,
				Amount:   ingredientProduct.Amount * magnification,
				Children: &[]model.FacilityChild{},
			}
			for i := 0; i < int(magnification); i++ {
				*stuffedFacility.Children = append(
					*stuffedFacility.Children,
					model.FacilityChild{
						Remain:   ingredientProduct.Amount - ingredient.Amount,
						Facility: a(ingredientItem),
					},
				)
			}
			children = append(
				children,
				model.FacilityChild{
					Remain:   stuffedFacility.Amount - ingredient.Amount,
					Facility: stuffedFacility,
				},
			)
		}
	}

	return &model.Facility{Recipe: &recipe, Amount: product.Amount, Children: &children}
}

func b(filteredFacility *model.Facility, facility *model.Facility) *model.Facility {
	pp.Println(facility.Recipe.Name)
	if len(*facility.Children) < 1 {
		*filteredFacility.Children = append(
			*filteredFacility.Children,
			model.FacilityChild{
				Remain:   0,
				Facility: facility,
			},
		)
	}

	for _, v := range *facility.Children {
		b(filteredFacility, v.Facility)
	}

	return filteredFacility
}
