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

	// handle all type of conversion
	http.HandleFunc("/convert", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		// Parse form values
		value, err := strconv.ParseFloat(r.FormValue("value"), 64)
		if err != nil {
			http.Error(w, "Invalid value", http.StatusBadRequest)
			return
		}
		// get unit from and to
		unitFrom := r.FormValue("unit_from")
		unitTo := r.FormValue("unit_to")

		// check if unit from or to are missing
		if unitFrom == "" || unitTo == "" {
			http.Error(w, "Missing unit", http.StatusBadRequest)
			return
		}

		convertedValue := 0.0
		var conversionError error

		convertType := r.URL.Query().Get("type")

		switch convertType {
		case "distance":
			convertedValue, conversionError = ConvertDistance(value, unitFrom, unitTo)
		case "weight":
			convertedValue, conversionError = ConvertWeight(value, unitFrom, unitTo)
		case "temperature":
			convertedValue, conversionError = ConvertTemperature(value, unitFrom, unitTo)
		default:
			http.Error(w, "Invalid type", http.StatusBadRequest)
		}

		if conversionError != nil {
			http.Error(w, "Invalid unit", http.StatusBadRequest)
			return
		}

		err = tmpl.ExecuteTemplate(w, "result.html", ResultData{
			ConvertedValue: convertedValue,
			ConvertedUnit:  unitTo,
		})

		if err != nil {
			panic(err)
		}
	})
}
