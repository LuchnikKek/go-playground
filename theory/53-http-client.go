package theory

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func getRequest() {
	resp, err := http.Get("https://ya.ru")
	if err != nil {
		log.Println(err)
	}

	log.Printf("Status: %d\r\n", resp.StatusCode)
	for k, v := range resp.Header {
		log.Printf("%s: %v\r\n", k, v[0])
	}
	log.Printf("Content-Type: %v", resp.Header.Get("Content-Type"))    // первое значение
	log.Printf("Content-Type: %v", resp.Header.Values("Content-Type")) // все значения

	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return
	}
	_ = body
}

func postRequest() {
	data := `{"name": "Илья", "age": 52}`
	resp, err := http.Post("https://ya.ru", "application/json", strings.NewReader(data))
	if err != nil {
		log.Println(err)
	}
	_ = resp
}

func getRequestBatch512() {
	resp, err := http.Get("https://mail.ru")
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	batch, err := io.ReadAll(io.LimitReader(resp.Body, 512)) // var 1
	if err != nil {
		log.Println(err)
		return
	}

	// if _, err := io.CopyN(os.Stdout, resp.Body, 512); err != nil { // var 2
	// 	log.Println(err)
	// 	return
	// }

	// batch, err := io.ReadAll(resp.Body) // var 3
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// if len(batch) > 512 {
	// 	batch = batch[:512]
	// }

	// output
	log.Println(len(batch))
}

func showRedirects() {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			log.Printf("Redirected to: %s\n", req.URL) // Redirected to: https://m.vk.com/
			return nil
		},
	}
	client.Get("https://vk.com")
}

func postRequestDiscardBody() {
	// http.Transport - реализация RoundTripper,
	// переиспользует открытые TCP-подключения (keep-alive)
	resp, err := http.Post("http://example.com", "text/html", strings.NewReader("lol"))
	if err != nil {
		log.Println(err)
		return
	}
	_, err = io.Copy(io.Discard, resp.Body) // discard
	resp.Body.Close()
	if err != nil {
		log.Println(err)
		return
	}
}

func getNewRequest() {
	req, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Println(err)
		return
	}
}

func MainHttpClient() {
	getRequest()
	postRequest()
	getRequestBatch512()
	showRedirects()
	postRequestDiscardBody()
	getNewRequest()
}
