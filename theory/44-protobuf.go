package theory

import (
	"log"
	"playground/theory/gen"

	"google.golang.org/protobuf/proto"
)

// protolint lint -fix .		- linter
// protoc theory/44_example.proto --go_out=./theory      - генерирует ./theory/44_example.pb.go (мапперы и ремапперы)

func MainProtobuf() {
	protoUnmarshal(protoMarshal())
}

func protoMarshal() []byte {
	mike := gen.Person{
		Name: "Mike",
		Age:  52,
	}

	data, err := proto.Marshal(&mike)
	if err != nil {
		log.Fatal("proto marshalling error:", err)
	}

	// raw protobuff obj
	log.Println(data) // [10 4 77 105 107 101 16 52]
	return data
}

func protoUnmarshal(input []byte) {
	decoded := &gen.Person{}

	if err := proto.Unmarshal(input, decoded); err != nil {
		log.Fatal("proto unmarshalling error:", err)
	}

	log.Println(decoded.GetAge())  // 52
	log.Println(decoded.GetName()) // Mike
}
