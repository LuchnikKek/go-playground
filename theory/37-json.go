package theory

import (
	"encoding/json"
	"fmt"
)

type Dimensions struct {
	Height int
	Width  int
}

type Bird struct {
	Species     string
	Description string
	Dimensions  Dimensions
}

func MainJson() {
	MainParseJson()
	MainCreateJson()
}

const (
	birdJson  = `{"Species":"Eagle","Description":"Cool eagle","Dimensions":{"Height":100,"Width":50}}`
	arrayJson = `["one","two","three"]`
)

func MainParseJson() {
	var bird Bird
	parseJson(birdJson, &bird)
	fmt.Println(bird) // {Eagle Cool eagle {100 50}}

	var nums []string
	parseJson(arrayJson, &nums)
	fmt.Println(nums) // [one two three]
}

func parseJson(jsonStr string, ref any) {
	err := json.Unmarshal([]byte(jsonStr), ref)
	if err != nil {
		panic(err)
	}
}

func MainCreateJson() {
	bird := Bird{
		Species:     "Eagle",
		Description: "Cool eagle",
		Dimensions: Dimensions{
			Height: 100,
			Width:  50,
		},
	}
	data, _ := json.Marshal(bird) // маршалит в сплошную строку
	// data, _ := json.MarshalIndent(bird, "", "  ") // маршалит в строку, но с корректными индентами
	fmt.Println(string(data))
}
