package servers

import (
	"log"
	"net"
)

func TCPServer() {
	l, err := net.Listen("tcp", "localhost:1234")
	defer func() { _ = l.Close() }()
	if err != nil {
		log.Panicf("Error listening: %v\n", err.Error())
	}

	log.Println("Server started on localhost:1234")

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Panicf("Error accepting: %v\n", err.Error())
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	log.Println("Connection opened")
	buf := make([]byte, 1024)
	reqLen, err := conn.Read(buf)
	if err != nil {
		log.Printf("Error reading: %v\n", err.Error())
	}
	log.Printf("Request len=%d, data=%s", reqLen, string(buf))

	bytesWritten, err := conn.Write([]byte("Message received\n"))
	if err != nil {
		log.Printf("Error writing: %v\n", err.Error())
	}
	log.Printf("Bytes written: %d\n", bytesWritten)
	conn.Close()
	log.Println("Connection closed")
}
