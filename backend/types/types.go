package types

type MeasuringUnits string

type UnitsMetric MeasuringUnits

const (
	Gram UnitsMetric = "g"
	Kilogram UnitsMetric = "kg"
	MiliLiter UnitsMetric = "ml"
	Liter UnitsMetric = "l"
	Celsius UnitsMetric = "C"
)

type UnitsImperial MeasuringUnits

const (

	dash UnitsImperial = "dash"
	tsp UnitsImperial = "tsp"
	tbsp UnitsImperial = "tbsp"
	cup UnitsImperial = "cup"
	pound UnitsImperial = "pound"
	fluidOunce UnitsImperial = "floz"
	Fahrenheit UnitsImperial = "F"
)