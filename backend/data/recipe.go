package data

import (
	"github.com/mattzsong/godev-backend/types"
)

type Ingredient struct {
	IngredientName string
	Amount         float32
	Unit           types.MeasuringUnits
}

type Recipe struct {
	RecipeName     string       `bson:"recipeName" json:"recipeName"`
	IngredientList []Ingredient `bson:"ingredientList" json:"ingredientList"`
	Description    string       `bson:"description" json:"description"`
}
