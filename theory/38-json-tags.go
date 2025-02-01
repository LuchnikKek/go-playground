package theory

import (
	"encoding/json"
	"fmt"
	"os"
)

type RequestContent struct {
	User    string
	Message string `json:"msg"` // указывает анмаршаллеру, что в json надо собирать из другого поля
}

type Request struct {
	Request RequestContent
	Author  string `json:"user"`
}

func MainJsonTags() {
	jsonData, _ := os.ReadFile("theory/38-example.json")

	LoadAndParseJson(jsonData)
	LoadAndParseRawMsgToMap(jsonData)
	LoadAndParseOnly(jsonData)
}

func LoadAndParseJson(jsonData []byte) {
	// Парсим из файла
	var request Request
	_ = json.Unmarshal(jsonData, &request)
	fmt.Println(request) // {{Admin Sending message} Admin}
}

func LoadAndParseRawMsgToMap(jsonData []byte) {
	// мапа с названиями полей (для динамической обработки, например)
	var objmap map[string]interface{}
	_ = json.Unmarshal(jsonData, &objmap)
	fmt.Println(objmap)
	// map[request:map[msg:Sending message user:Admin] user:Admin]  - мапа строк
}

func LoadAndParseOnly(jsonData []byte) {
	// Парсим только часть мапы
	var objmap map[string]json.RawMessage // json.RawMessage - массив байтов
	_ = json.Unmarshal(jsonData, &objmap)

	var internalMap map[string]string
	_ = json.Unmarshal(objmap["request"], &internalMap)
	fmt.Println(internalMap) // map[msg:Sending message user:Admin]
}
