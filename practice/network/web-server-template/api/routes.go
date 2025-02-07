package api

import (
	"web-server/internals/app/handlers"

	"github.com/gorilla/mux"
)

func CreateRoutes() *mux.Router {
	r := mux.NewRouter()

	r.NotFoundHandler = r.NewRoute().HandlerFunc(handlers.NotFound).GetHandler() //оборачиваем 404, для обработки NotFound
	return r
}
