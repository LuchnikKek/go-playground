package theory

import "net/http"

func EasyFileServer() {
	// простейший сервер, которому доступны все файлы в поддиректории 48-static/
	// GET / => ls ./48-static
	// GET /assets/ => ls ./48-static/assets/
	err := http.ListenAndServe(":8080", http.FileServer(http.Dir("./theory/48-static")))
	if err != nil {
		panic(err)
	}
}

func MappedFileServer() {
	// GET /assets/page.html => ./48-static/page.html
	mux := http.NewServeMux()

	// GET / => 404
	// GET /files/ => ls ./48-static
	fs := http.FileServer(http.Dir("./theory/48-static"))
	mux.Handle("/files/", http.StripPrefix("/files/", fs))

	// GET /page => ./theory/48-static/page.html
	mux.HandleFunc("/page", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./theory/48-static/page.html")
	})
	// same
	// GET /page-copy => ./theory/48-static/page.html
	mux.HandleFunc("/page-copy", staticPage)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}

func staticPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./theory/48-static/page.html")
}
