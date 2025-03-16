package theory

import (
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var cars = map[string]string{
	"1": "Renault Logan",
	"2": "Renault Duster",
	"3": "BMW X6",
	"4": "BMW M5",
	"5": "VW Passat",
	"6": "VW Jetta",
	"7": "Audi A4",
	"8": "Audi Q7",
}

func carFunc(id string) string {
	if c, ok := cars[id]; ok {
		return c
	}
	return "unknown identifier " + id
}

func carHandle(rw http.ResponseWriter, r *http.Request) {
	// при запросе "/car" вернётся ошибка 404 => id всегда не пустой
	rw.Write([]byte(carFunc(chi.URLParam(r, "id"))))
}

func MainChiServer() {
	r := chi.NewRouter()
	r.Get("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("chi")) // chi
	})
	r.Get("/item/{id}", func(rw http.ResponseWriter, r *http.Request) { // localhost:8080/item/lol
		id := chi.URLParam(r, "id")
		io.WriteString(rw, fmt.Sprintf("requested id = %s", id)) // requested id = lol
	})
	r.Get("/car/{id}", carHandle)
	http.ListenAndServe(":8080", r)
}
