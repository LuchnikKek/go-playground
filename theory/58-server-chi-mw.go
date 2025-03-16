package theory

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func TimerTrace(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		log.Println("Request duration:", duration)
	})
}

func MainChiServerMw() {
	r := chi.NewRouter()

	// r.Use(middleware.RealIP, middleware.Logger, middleware.Recoverer)
	r.Use(TimerTrace)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("chi")) // chi
	})

	http.ListenAndServe(":8080", r)
}
