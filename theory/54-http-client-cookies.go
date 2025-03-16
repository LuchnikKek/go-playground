package theory

import (
	"log"
	"net/http"
	"net/http/cookiejar"
)

func getRequestCookies(client *http.Client) {
	req, err := http.NewRequest("GET", "http://vk.com", nil)
	if err != nil {
		log.Println(err)
		return
	}
	req.AddCookie(&http.Cookie{
		Name:  "ID",
		Value: "52",
	})
	req.AddCookie(&http.Cookie{
		Name:   "Token",
		Value:  "TEST_TOKEN",
		MaxAge: 360,
		// Path: ,
		// Domain: ,
		// Expires: ,
		// Secure: ,
		// HttpOnly: ,
	})
	_, err = client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
}

func getRequestCookiesJar() {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &http.Client{Jar: jar}
	req, err := http.NewRequest("GET", "http://vk.com", nil)
	if err != nil {
		log.Println(err)
		return
	}
	cookie := &http.Cookie{
		Name:   "Token_2",
		Value:  "TEST_TOKEN_222",
		MaxAge: 360,
	}
	req.AddCookie(cookie)
	_, err = client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
}

func MainHttpClientCookies() {
	client := &http.Client{}
	getRequestCookies(client)
}
