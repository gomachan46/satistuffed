package data

import (
	"fmt"
	"github.com/gomachan46/satistuffed/model"
)

type Data struct {
	Items *[]model.Item
}

func (d Data) GetItem(name string) (*model.Item, error) {
	for _, v := range *d.Items {
		if v.Name == name {
			return &v, nil
		}
	}

	return nil, fmt.Errorf("Error: item not found %s", name)
}

func Load() *Data {
	ironOre := &model.Item{Name: "鉄鉱石"}
	ironIngot := &model.Item{Name: "鉄インゴット"}
	ironPlate := &model.Item{Name: "鉄板"}
	ironRod := &model.Item{Name: "鉄のロッド"}
	screw := &model.Item{Name: "ネジ"}
	reinforcedIronPlate := &model.Item{Name: "強化鉄板"}

	ironOreRecipe := &model.Recipe{Name: "鉄鉱石のレシピ1"}
	ironIngotRecipe := &model.Recipe{Name: "鉄インゴットのレシピ1"}
	ironPlateRecipe := &model.Recipe{Name: "鉄板のレシピ1"}
	ironRodRecipe := &model.Recipe{Name: "鉄のロッドのレシピ1"}
	screwRecipe := &model.Recipe{Name: "ネジのレシピ1"}
	reinforcedIronPlateRecipe := &model.Recipe{Name: "強化鉄板のレシピ1"}

	ironOreRecipe.Ingredients = &[]model.Ingredient{}
	ironOreRecipe.Products = &[]model.Product{{Item: ironOre, Amount: 60}}
	ironOre.Recipes = &[]model.Recipe{*ironOreRecipe}

	ironIngotRecipe.Ingredients = &[]model.Ingredient{{Item: ironOre, Amount: 30}}
	ironIngotRecipe.Products = &[]model.Product{{Item: ironIngot, Amount: 30}}
	ironIngot.Recipes = &[]model.Recipe{*ironIngotRecipe}

	ironPlateRecipe.Ingredients = &[]model.Ingredient{{Item: ironIngot, Amount: 30}}
	ironPlateRecipe.Products = &[]model.Product{{Item: ironPlate, Amount: 20}}
	ironPlate.Recipes = &[]model.Recipe{*ironPlateRecipe}

	ironRodRecipe.Ingredients = &[]model.Ingredient{{Item: ironIngot, Amount: 15}}
	ironRodRecipe.Products = &[]model.Product{{Item: ironRod, Amount: 15}}
	ironRod.Recipes = &[]model.Recipe{*ironRodRecipe}

	screwRecipe.Ingredients = &[]model.Ingredient{{Item: ironRod, Amount: 10}}
	screwRecipe.Products = &[]model.Product{{Item: screw, Amount: 40}}
	screw.Recipes = &[]model.Recipe{*screwRecipe}

	reinforcedIronPlateRecipe.Ingredients = &[]model.Ingredient{{Item: ironPlate, Amount: 30}, {Item: screw, Amount: 60}}
	reinforcedIronPlateRecipe.Products = &[]model.Product{{Item: reinforcedIronPlate, Amount: 5}}
	reinforcedIronPlate.Recipes = &[]model.Recipe{*reinforcedIronPlateRecipe}

	var items []model.Item
	items = append(items, *ironRod)
	items = append(items, *screw)
	items = append(items, *ironPlate)
	items = append(items, *ironOre)
	items = append(items, *ironIngot)
	items = append(items, *reinforcedIronPlate)

	return &Data{Items: &items}
}
