package theory

import (
	"log"
	"net/http"

	"github.com/go-resty/resty/v2"
)

func getRestyAuth(client *resty.Client) {
	resp, err := client.R().
		SetAuthToken("Bearer <TOKEN>").
		Get("http://example.com")

	log.Println("Error:", err)
	log.Println("Status Code:", resp.StatusCode())
	log.Println("Status:", resp.Status())
	log.Println("Time:", resp.Time())
	log.Println("Received At:", resp.ReceivedAt())
	// log.Println("Body:\n", resp)
}

type RestyUser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func getRestyUsers(client *resty.Client) {
	users := []RestyUser{}
	resp, err := client.R().
		SetResult(&users).
		Get("https://jsonplaceholder.typicode.com/users")

	if err != nil {
		log.Println("Error:", err)
		return
	}
	if resp.StatusCode() != http.StatusOK {
		log.Println("Error: status code is ", resp.StatusCode())
		return
	}

	log.Println(users[0])          // {1 Bret Sincere@april.biz}
	log.Println(users[0].Username) // Bret
	log.Println(len(users))        // 10
}

func MainHttpClientResty() {
	client := resty.New()
	getRestyAuth(client)
	getRestyUsers(client)
}
