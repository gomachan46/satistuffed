/*
Copyright © 2020 gomachan46 <shiro.gomachan46@gmail.com>

*/
package cmd

import (
	"fmt"
	"github.com/gomachan46/satistuffed/cli"
	"github.com/gomachan46/satistuffed/data"
	"github.com/gomachan46/satistuffed/model"
	"github.com/gomachan46/satistuffed/mymath"
	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
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

	children := []model.Facility{}
	magnifications := []int{}
	for _, ingredient := range *recipe.Ingredients {
		ingredientItem := ingredient.Item
		ingredientRecipe := (*ingredientItem.Recipes)[0]
		ingredientProduct := (*ingredientRecipe.Products)[0]

		lcm := mymath.LCM(int(ingredient.Amount), int(ingredientProduct.Amount))
		magnification := lcm / int(ingredient.Amount)
		magnifications = append(magnifications, magnification)
	}

	var m int
	if len(magnifications) == 0 {
		m = 1
	} else if len(magnifications) == 1 {
		m = magnifications[0]
	} else {
		m = mymath.LCM(magnifications[0], magnifications[1], magnifications[2:]...)
	}
	recipe.Name = fmt.Sprintf("%s x %d", recipe.Name, m)
	stuffedFacility := &model.Facility{Recipe: &recipe, Children: &[]model.Facility{}}
	for i := 0; i < m; i++ {
		*stuffedFacility.Children = append(*stuffedFacility.Children, *a(item))
	}
	children = append(children, *stuffedFacility)
	pp.Println(m)

	return stuffedFacility
}
