package data

import "github.com/mattzsong/godev-backend/types"

type Ingredient struct {
	ingredientName string,
	amount float32,
	unit MeasuringUnits,
}

type Recipe struct {
	recipe_name string,
	ingredientList []Ingredient,
	description string,
}


