package servers

import (
	"log"
	"net"
)

func UDPServer() {
	p := make([]byte, 2048)
	addr := net.UDPAddr{
		Port: 1234,
		IP:   net.ParseIP("localhost"),
	}
	ser, err := net.ListenUDP("udp", &addr) // можно через Listen(), как у TCP
	defer func() { _ = ser.Close() }()
	if err != nil {
		log.Panicf("Error listening: %v\n", err.Error())
	}

	log.Println("Listening on 127.0.0.1:1234")
	for {
		_, remoteaddr, err := ser.ReadFromUDP(p)
		log.Printf("Read message from %v: %s\n", remoteaddr, string(p))

		if err != nil {
			log.Println("An error occured while reading:", err.Error())
			continue
		}

		go sendResponse(ser, remoteaddr)
	}
}

func sendResponse(ser *net.UDPConn, remoteaddr *net.UDPAddr) {
	_, err := ser.WriteToUDP([]byte("Hello, Client. Your IP: "+remoteaddr.IP.String()), remoteaddr)
	if err != nil {
		log.Printf("Error writing: %v\n", err.Error())
	}
}
