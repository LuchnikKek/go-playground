package theory

import (
	"encoding/xml"
	"fmt"
)

func MainXml() {
	DecodeXml(CreateXml())
	NestedXml()
}

type Plant struct {
	XMLName xml.Name `xml:"plant"` // это название тега
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant(Id=%v, Name=%v, Origin=%v)", p.Id, p.Name, p.Origin)
}

func CreateXml() string {
	// XML from Struct
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	out, _ := xml.MarshalIndent(coffee, "", "  ")
	return string(out)
}

func DecodeXml(input string) {
	// Struct from XML
	var p Plant
	if err := xml.Unmarshal([]byte(input), &p); err != nil {
		panic(err)
	}
	fmt.Println("decoded:", p)
}

func NestedXml() {
	// XML nesting +quick Unmarshal
	coffee := &Plant{Id: 11, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	peas := &Plant{Id: 52, Name: "Peas"}
	peas.Origin = []string{"India", "China"}

	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, peas}

	out, _ := xml.MarshalIndent(nesting, "", "  ")
	fmt.Println(string(out))

	res := &Nesting{}
	_ = xml.Unmarshal(out, res)

	fmt.Println(res)
}
