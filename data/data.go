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
	var items []model.Item
	ironOreRecipe := &model.Recipe{Name: "鉄鉱石のレシピ1"}
	ironOre := &model.Item{Name: "鉄鉱石"}
	ironOreRecipe.Ingredients = &[]model.Ingredient{}
	ironOreRecipe.Products = &[]model.Product{{Item: ironOre, Amount: 60}}
	ironOre.Recipes = &[]model.Recipe{*ironOreRecipe}
	items = append(items, *ironOre)

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
	items = append(items, *ironIngot)

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
	items = append(items, *ironPlate)

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
	items = append(items, *ironRod)

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
	items = append(items, *screw)

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
	items = append(items, *reinforcedIronPlate)

	return &Data{Items: &items}
}
