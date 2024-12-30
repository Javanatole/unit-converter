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
func ConvertWeight(weight float64, unitFrom string, unitTo string) (float64, error) {
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
		return -1, fmt.Errorf("unknow unit")
	}

	// Convert weight to kilograms first, then to the target unit
	return (weight / fromFactor) * toFactor, nil
}

// ConvertTemperature convert temperature from a unit to another unit
func ConvertTemperature(temp float64, unitFrom string, unitTo string) (float64, error) {
	// Conversion formulas between Celsius, Kelvin, and Fahrenheit
	value := 0.0
	var err error
	err = nil
	switch strings.ToLower(unitFrom) {
	case "celsius":
		switch strings.ToLower(unitTo) {
		case "kelvin":
			value = temp + 273.15
		case "fahrenheit":
			value = (temp * 9 / 5) + 32
		default:
			err = fmt.Errorf("unknow unit")
		}
	case "kelvin":
		switch strings.ToLower(unitTo) {
		case "celsius":
			value = temp - 273.15
		case "fahrenheit":
			value = ((temp - 273.15) * 9 / 5) + 32
		default:
			err = fmt.Errorf("unknow unit")
		}
	case "fahrenheit":
		switch strings.ToLower(unitTo) {
		case "celsius":
			value = (temp - 32) * 5 / 9
		case "kelvin":
			value = ((temp - 32) * 5 / 9) + 273.15
		default:
			err = fmt.Errorf("unknow unit")
		}
	default:
		err = fmt.Errorf("unknow unit")

	}
	return value, err
}
