package theory

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

var carByModel = map[string]string{
	"renault logan":  "any data",
	"renault duster": "any data",
	"bmw x6":         "any data",
	"bmw m5":         "any data",
	"vw passat":      "any data",
	"vw jetta":       "any data",
	"audi a4":        "any data",
	"audi q7":        "any data",
}

func modelFunc(brand, model string) (string, bool) {
	cBrand := strings.ToLower(brand)
	cModel := strings.ToLower(model)

	key := cBrand + " " + cModel
	data, exists := carByModel[key]
	return data, exists
}

// hand snippet
func modelHandle(w http.ResponseWriter, r *http.Request) {
	brand := chi.URLParam(r, "brand")
	model := chi.URLParam(r, "model")

	car, exists := modelFunc(brand, model)

	if !exists {
		http.Error(w, "unknown model: "+model, http.StatusNotFound)
	}
	w.Write([]byte(car))
}

func MainChiServerAdv() {
	r := chi.NewRouter()
	// r.Get("/cars", carsHandle) // GET /cars
	// r.Get("/cars/{brand}", brandHandle) // GET /cars/renault
	// r.Get("/cars/{brand}/{model}", modelHandle) // GET /cars/renault/duster

	// то же самое можно описать, используя Route
	r.Route(`/cars`, func(r chi.Router) {
		// r.Get("/", carsHandle) // GET /cars
		r.Route(`/{brand}`, func(r chi.Router) {
			// r.Get("/", brandHandle)        // GET /cars/renault
			r.Get(`/{model}`, modelHandle) // GET /cars/renault/duster
		})
	})

	http.ListenAndServe(":8080", r)
}
