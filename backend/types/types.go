package types

type MeasuringUnits string

type UnitsMetric MeasuringUnits

const (
	Gram      UnitsMetric = "g"
	Kilogram  UnitsMetric = "kg"
	MiliLiter UnitsMetric = "ml"
	Liter     UnitsMetric = "l"
	Celsius   UnitsMetric = "C"
)

type UnitsImperial MeasuringUnits

const (
	Dash       UnitsImperial = "dash"
	Tsp        UnitsImperial = "tsp"
	Tbsp       UnitsImperial = "tbsp"
	Cup        UnitsImperial = "cup"
	Pound      UnitsImperial = "pound"
	FluidOunce UnitsImperial = "floz"
	Fahrenheit UnitsImperial = "F"
)
