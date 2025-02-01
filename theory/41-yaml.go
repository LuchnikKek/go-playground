package theory

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Conf struct {
	Hits int64 `yaml:"hits"`
	Time int64 `yaml:"time"`
}

func MainYaml() {
	getConf()
	createConf()
	readBigConf()
}

func getConf() {
	// анмаршалим YAML
	config := &Conf{}

	yamlFile, err := os.ReadFile("theory/41-example.yaml")
	if err != nil {
		log.Printf("Yaml file not found %#v", err)
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.Fatalln("Unmarshalling error:", err)
	}
	fmt.Println(config) // &{52 171247818427}
}

func createConf() {
	// маршалим YAML
	config := &Conf{
		Hits: 12,
		Time: 174724178277,
	}
	out, _ := yaml.Marshal(config)
	fmt.Println(string(out))
}

type BuildConf struct { // лениво собираем
	Definitions map[string]interface{} `yaml:"definitions"`
	Pipelines   map[string]interface{} `yaml:"pipelines"`
}

func readBigConf() {
	config := &BuildConf{}

	yamlFile, err := os.ReadFile("theory/41-1-anchor.yaml")
	if err != nil {
		log.Printf("Yaml file not found %#v", err)
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.Fatalln("Unmarshalling error:", err)
	}

	fmt.Println(config) // разумеется, это map
	// Но главное - якори отработали
}
