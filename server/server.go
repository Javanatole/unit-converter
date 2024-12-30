package server

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

// ResultData represents the structure of the API response
type ResultData struct {
	ConvertedValue float64 `json:"converted_value"`
	ConvertedUnit  string  `json:"converted_unit"`
}

// ResultError represents the structure of the API response
type ResultError struct {
	Error string `json:"error"`
}

func ExecuteErrorTemplate(tmpl *template.Template, w http.ResponseWriter, error string) {
	err := tmpl.ExecuteTemplate(w, "error.html", ResultError{
		Error: error,
	})
	if err != nil {
		panic(err)
	}
}

func InitHttpRoutes() {
	// Parse templates
	tmpl := template.Must(template.ParseFiles("static/result.html", "static/error.html"))

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

		convertType := r.URL.Query().Get("type")

		convertedValue := 0.0
		var conversionError error

		switch convertType {
		case "distance":
			convertedValue, conversionError = ConvertDistance(value, unitFrom, unitTo)
		case "weight":
			convertedValue = ConvertWeight(value, unitFrom, unitTo)
		case "temperature":
			convertedValue = ConvertTemperature(value, unitFrom, unitTo)
		}

		if conversionError != nil {
			missing := []string{}
			if unitFrom == "" {
				missing = append(missing, "« unit from »")
			}
			if unitTo == "" {
				missing = append(missing, "« unit to »")
			}
			if len(missing) > 0 {
				missing := "We need: " + strings.Join(missing, " and ") + " for computing the result"
				ExecuteErrorTemplate(tmpl, w, fmt.Sprintf(missing))
				return
			}
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
