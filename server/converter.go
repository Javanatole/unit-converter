package server

import (
	"fmt"
	"strings"
)

// ConvertDistance convert distance from a unit to another unit
func ConvertDistance(distance float64, unitFrom string, unitTo string) (float64, error) {
	conversionFactors := map[string]float64{
		"meter":      1,
		"kilometer":  0.001,
		"centimeter": 100,
	}
	unitFrom = strings.ToLower(unitFrom)
	unitTo = strings.ToLower(unitTo)
	fromFactor, fromExists := conversionFactors[unitFrom]
	toFactor, toExists := conversionFactors[unitTo]
	if !fromExists || !toExists {
		return -1, fmt.Errorf("unknow unit")
	}
	// Convert distance to meters first, then to the target unit
	return (distance / fromFactor) * toFactor, nil
}

// ConvertWeight convert weight from a unit to another unit
func ConvertWeight(weight float64, unitFrom string, unitTo string) float64 {
	conversionFactors := map[string]float64{
		"kilogram": 1,
		"gram":     1000,
		"pound":    2.20462,
	}

	unitFrom = strings.ToLower(unitFrom)
	unitTo = strings.ToLower(unitTo)

	fromFactor, fromExists := conversionFactors[unitFrom]
	toFactor, toExists := conversionFactors[unitTo]

	if !fromExists || !toExists {
		panic(fmt.Sprintf("Invalid units: %s or %s", unitFrom, unitTo))
	}

	// Convert weight to kilograms first, then to the target unit
	return (weight / fromFactor) * toFactor
}

// ConvertTemperature convert temperature from a unit to another unit
func ConvertTemperature(temp float64, unitFrom string, unitTo string) float64 {
	// Conversion formulas between Celsius, Kelvin, and Fahrenheit
	switch strings.ToLower(unitFrom) {
	case "celsius":
		switch strings.ToLower(unitTo) {
		case "kelvin":
			return temp + 273.15
		case "fahrenheit":
			return (temp * 9 / 5) + 32
		}
	case "kelvin":
		switch strings.ToLower(unitTo) {
		case "celsius":
			return temp - 273.15
		case "fahrenheit":
			return ((temp - 273.15) * 9 / 5) + 32
		}
	case "fahrenheit":
		switch strings.ToLower(unitTo) {
		case "celsius":
			return (temp - 32) * 5 / 9
		case "kelvin":
			return ((temp - 32) * 5 / 9) + 273.15
		}
	}

	// If no valid conversion exists, panic with an error
	panic(fmt.Sprintf("Invalid units: %s or %s", unitFrom, unitTo))
}
