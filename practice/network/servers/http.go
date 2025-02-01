package servers

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func HTTPServer() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Println(err)
	}
}
