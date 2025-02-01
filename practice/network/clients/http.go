package clients

import (
	"bufio"
	"log"
	"net/http"
)

func HTTPClientSimpleGet() {
	resp, err := http.Get("http://localhost:80/hello")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	log.Println("Response status:", resp.Status)

	scaner := bufio.NewScanner(resp.Body)

	for i := 0; scaner.Scan() && i < 5; i++ {
		log.Println(scaner.Text())
	}

	if err := scaner.Err(); err != nil {
		panic(err)
	}
}

func HTTPClientHeadersGet() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:80/headers", nil)
	req.Header.Add("x-filter-group", "1")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	log.Println("Server headers:")
	for name, headers := range resp.Header {
		for _, h := range headers {
			log.Printf("H: %v: %v\n", name, h)
		}
	}
	log.Println("Response status:", resp.Status)

	scaner := bufio.NewScanner(resp.Body)
	for i := 0; scaner.Scan() && i < 5; i++ {
		log.Println(scaner.Text())
	}
	if err := scaner.Err(); err != nil {
		panic(err)
	}
}
