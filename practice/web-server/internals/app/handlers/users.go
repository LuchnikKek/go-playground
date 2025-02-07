package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"web-server/internals/app/models"
	"web-server/internals/app/processors"

	"github.com/gorilla/mux"
)

type UsersHandler struct {
	processor *processors.UsersProcessor
}

func NewUsersHandler(processor *processors.UsersProcessor) *UsersHandler {
	handler := new(UsersHandler)
	handler.processor = processor
	return handler
}

func (handler *UsersHandler) Create(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		WrapError(w, err)
		return
	}

	if err := handler.processor.CreateUser(r.Context(), newUser); err != nil {
		WrapError(w, err)
		return
	}

	var m = map[string]interface{}{
		"result": "ok",
		"data":   "user created",
	}

	WrapOK(w, m)
}

func (handler *UsersHandler) List(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	// параметры будут обёрнуты в кавычки, надо развернуть
	list, err := handler.processor.ListUsers(r.Context(), strings.Trim(vars.Get("name"), "\""))

	if err != nil {
		WrapError(w, err)
		return
	}

	var m = map[string]interface{}{
		"result": "ok",
		"data":   list,
	}

	WrapOK(w, m)
}

func (handler *UsersHandler) Find(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //переменные, обьявленные в ресурсах попадают в Vars и могут быть адресованы
	if vars["id"] == "" {
		WrapError(w, errors.New("missing id"))
		return
	}

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		WrapError(w, err)
		return
	}

	user, err := handler.processor.FindUser(r.Context(), id)
	if err != nil {
		WrapError(w, err)
		return
	}

	var m = map[string]interface{}{
		"result": "OK",
		"data":   user,
	}

	WrapOK(w, m)
}
