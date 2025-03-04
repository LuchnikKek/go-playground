package theory

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type HelloHandler struct{}

func (h HelloHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	data := []byte("Привет!")
	res.Write(data)
}

func mainPage(res http.ResponseWriter, req *http.Request) {
	body := fmt.Sprintf("Method: %s\r\n", req.Method)

	body += "Header ===============\r\n"
	for k, v := range req.Header {
		body += fmt.Sprintf("%s: %v\r\n", k, v)
	}

	body += "Query parameters ===============\r\n"
	for k, v := range req.URL.Query() {
		body += fmt.Sprintf("%s: %v\r\n", k, v)
	}

	body += "Query parameters (Form) ===============\r\n"
	if err := req.ParseForm(); err != nil {
		// при ошибке можно так
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	for k, v := range req.Form {
		body += fmt.Sprintf("%s: %v\r\n", k, v)
	}

	if req.Method == http.MethodPost {
		body += "Body ===============\r\n"
		req_body, err := io.ReadAll(req.Body)
		if err != nil {
			// а можно так (200)
			res.Write([]byte(err.Error()))
			return
		}
		body += fmt.Sprintf("%s\r\n", req_body)
	}

	res.Header().Set("lol", "1")
	res.Header().Add("kek", "3")

	res.WriteHeader(http.StatusOK)

	res.Write([]byte(body))
}

type Subj struct {
	Product string `json:"name"`
	Price   int    `json:"price"`
}

func apiPage(w http.ResponseWriter, req *http.Request) {
	subj := Subj{"Milk", 50}

	resp, err := json.Marshal(subj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(resp)
}

const form = `<html>
    <head>
    <title></title>
    </head>
    <body>
        <form action="/" method="post">
            <label>Логин <input type="text" name="login"></label>
            <label>Пароль <input type="password" name="password"></label>
            <input type="submit" value="Login">
        </form>
    </body>
</html>`

func Auth(login, password string) bool {
	return login == `guest` && password == `demo`
}

func authPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		login := r.FormValue("login")
		password := r.FormValue("password")
		if Auth(login, password) {
			io.WriteString(w, "Добро пожаловать!")
		} else {
			http.Error(w, "Неверный логин или пароль", http.StatusUnauthorized)
		}
		return
	} else {
		io.WriteString(w, form)
	}
}

func MainHttpServer() {
	// curl -X POST http://localhost:8080/\?id\=12345\&name\=John%20Doe\&filter\=town\&filter\=country -H "Content-Type: text/plain" --data "Hello" -i
	var helloPage HelloHandler

	mux := http.NewServeMux()

	mux.Handle(`/hello`, helloPage)
	mux.HandleFunc(`/`, mainPage)
	mux.HandleFunc(`/api/`, apiPage)
	mux.HandleFunc(`/auth`, authPage)
	// `/api/` - всё + редирект /api->/api/
	// `/api` - только /api

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
