package clients

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func TCPClient() {
	p := make([]byte, 1024)
	conn, err := net.Dial("tcp", "localhost:1234")
	defer func() { _ = conn.Close() }()

	if err != nil {
		log.Fatalln("Cannot set Dial connection:", err.Error())
	}

	fmt.Fprintf(conn, "Any important client data! 52") // конвертация в байты, т.к. conn - Writer
	_, err = bufio.NewReader(conn).Read(p)

	if err != nil {
		log.Println("Unexpected error:", err.Error())
	} else {
		log.Println(string(p))
	}
}
