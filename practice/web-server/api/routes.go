package api

import (
	"web-server/internals/app/handlers"

	"github.com/gorilla/mux"
)

func CreateRoutes(userHandler *handlers.UsersHandler, carsHandler *handlers.CarsHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users/create", userHandler.Create).Methods("POST") //каждая функция реализует один и тот же интерфейс
	r.HandleFunc("/users/list", userHandler.List).Methods("GET")
	r.HandleFunc("/users/find/{id:[0-9]+}", userHandler.Find).Methods("GET")

	r.HandleFunc("/cars/create", carsHandler.Create).Methods("POST")
	r.HandleFunc("/cars/list", carsHandler.List).Methods("GET")
	r.HandleFunc("/cars/find/{id:[0-9]+}", carsHandler.Find).Methods("GET") //Methods определяют какой глагол можно использовать, если будет другой - вернется 404

	r.NotFoundHandler = r.NewRoute().HandlerFunc(handlers.NotFound).GetHandler() //оборачиваем 404, для обработки NotFound

	return r
}
