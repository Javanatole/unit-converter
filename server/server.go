package server

import (
	"html/template"
	"net/http"
	"strconv"
)

// ResultData represents the structure of the API response
type ResultData struct {
	ConvertedValue float64 `json:"converted_value"`
	ConvertedUnit  string  `json:"converted_unit"`
}

func InitHttpRoutes() {
	// Parse templates
	tmpl := template.Must(template.ParseFiles("static/result.html"))

	// handle convert distance
	http.HandleFunc("/convert-distance", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		// Parse form values
		length, err := strconv.ParseFloat(r.FormValue("length"), 64)
		if err != nil {
			http.Error(w, "Invalid length value", http.StatusBadRequest)
			return
		}
		unitFrom := r.FormValue("unit_from")
		unitTo := r.FormValue("unit_to")

		distance, err := ConvertDistance(length, unitFrom, unitTo)

		if err != nil {
			http.Error(w, "Can't convert in another unit", http.StatusBadRequest)
		}

		err = tmpl.ExecuteTemplate(w, "result.html", ResultData{
			ConvertedValue: distance,
			ConvertedUnit:  unitTo,
		})
		if err != nil {
			panic(err)
		}
	})

	// handle convert weight
	http.HandleFunc("/convert-weight", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		// Parse form values
		weight, err := strconv.ParseFloat(r.FormValue("weight"), 64)
		if err != nil {
			http.Error(w, "Invalid weight value", http.StatusBadRequest)
			return
		}
		unitFrom := r.FormValue("unit_from")
		unitTo := r.FormValue("unit_to")

		convertedWeight := ConvertWeight(weight, unitFrom, unitTo)

		if err != nil {
			http.Error(w, "Can't convert in another unit", http.StatusBadRequest)
		}

		err = tmpl.ExecuteTemplate(w, "result.html", ResultData{
			ConvertedValue: convertedWeight,
			ConvertedUnit:  unitTo,
		})

		if err != nil {
			panic(err)
		}
	})

	// handle convert weight
	http.HandleFunc("/convert-temperature", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		// Parse form values
		temperature, err := strconv.ParseFloat(r.FormValue("temperature"), 64)
		if err != nil {
			http.Error(w, "Invalid weight value", http.StatusBadRequest)
			return
		}
		unitFrom := r.FormValue("unit_from")
		unitTo := r.FormValue("unit_to")

		convertedTemperature := ConvertTemperature(temperature, unitFrom, unitTo)

		if err != nil {
			http.Error(w, "Can't convert in another unit", http.StatusBadRequest)
		}

		err = tmpl.ExecuteTemplate(w, "result.html", ResultData{
			ConvertedValue: convertedTemperature,
			ConvertedUnit:  unitTo,
		})

		if err != nil {
			panic(err)
		}
	})
}
