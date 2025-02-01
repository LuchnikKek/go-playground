package theory

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func MainJsonCustomUnmapping() {
	// На этапе анмаршаллинга решаем проблему поля tags
	jsonData, _ := os.ReadFile("theory/39-example.json")

	var request RequestTagged
	_ = json.Unmarshal(jsonData, &request)
	fmt.Println(request) // {{Admin Sending message} Admin}
}

type RequestTagged struct {
	Request RequestContentTag
	Author  string `json:"user"`
}

type RequestContentTag struct {
	User    string
	Message string   `json:"msg"`
	Tags    strslice `json:"tags"`
}

type strslice []string // маппинг слайса над строкой

func (ss *strslice) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	*ss = strings.Split(s, ",") // дереференсируем его и присваиваем ему сплит
	return nil
}
