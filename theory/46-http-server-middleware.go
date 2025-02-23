package theory

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// Single MW
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // пример
		// Замыкание
		next.ServeHTTP(w, r)
	})
}

// or

// Many MWs
type Middleware func(http.Handler) http.Handler

func Conveyor(h http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}

func logMiddleware(next http.Handler) http.Handler {
	log_format := "[%s] %s: %s %s - %v\n"
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf(log_format, time.Now().Format("2006-01-02 15:04:05.999999"), "Request", r.Method, r.RequestURI, r.ContentLength)
		next.ServeHTTP(w, r)
		time.Sleep(1 * time.Second) // симуляция работы
		fmt.Printf(log_format, time.Now().Format("2006-01-02 15:04:05.999999"), "Response", r.Method, r.RequestURI, r.ContentLength)
	})
}

func outerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Запрос начат")
		next.ServeHTTP(w, r)
		fmt.Println("Ответ отправлен")
	})
}

func infoPage(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "1")
	fmt.Fprint(w, "2")
	w.Write([]byte("3\n"))
	// "123\n"
}

func MainHttpServerMiddleware() {
	mux := http.NewServeMux()

	mux.Handle(`/info`, logMiddleware(http.HandlerFunc(infoPage)))
	mux.Handle(`/stats`, Conveyor(http.HandlerFunc(infoPage), corsMiddleware, logMiddleware, outerMiddleware))
	// Запрос начат
	// [2025-02-23 06:10:28.442506] Request: POST /stats - 7
	// [2025-02-23 06:10:29.443714] Response: POST /stats - 7
	// Ответ отправлен

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
