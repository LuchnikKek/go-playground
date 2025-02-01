package theory

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type SlicesTags []string

func (tags *SlicesTags) UnmarshalYAML(value *yaml.Node) error {
	if value != nil {
		*tags = strings.Split(value.Value, ",")
	}
	return nil
}

type Messages struct {
	Tags SlicesTags `yaml:"tags"`
}

type Subs struct {
	Messages Messages `yaml:"messages"`
}

func MainYamlFeatures() {
	parseYamlWithCustomStruct()
}

func parseYamlWithCustomStruct() {
	config := &Subs{}

	yamlFile, err := os.ReadFile("theory/42-custom-unmapping.yaml")
	if err != nil {
		log.Printf("Yaml file not found %#v", err)
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.Fatalln("Unmarshalling error:", err)
	}

	fmt.Println(config)
}
