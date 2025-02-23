package theory

import (
	"log"
	"net/http"
	"time"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://yandex.ru/", http.StatusMovedPermanently)
}

func internalErrorPage(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Internal server error", http.StatusInternalServerError)
}

func notFoundPage(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func longLoadedPage(w http.ResponseWriter, r *http.Request) {
	time.Sleep(7 * time.Second)
}

func MainHttpHandleHelpers() {
	// same 301
	http.HandleFunc("/search/", redirect)
	http.Handle("/dummy", http.RedirectHandler("https://google.com", http.StatusMovedPermanently))

	// same 404
	http.HandleFunc("/404", notFoundPage)
	http.Handle("/404-2", http.NotFoundHandler())

	// 500
	http.HandleFunc("/500", internalErrorPage)

	// 503
	http.Handle(`/timed`, http.TimeoutHandler(http.HandlerFunc(longLoadedPage), 2*time.Second, "Timeout exceeded"))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
