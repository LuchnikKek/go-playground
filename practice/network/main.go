package main

import (
	"playground/practice/network/clients"
	"playground/practice/network/servers"
	"time"
)

func main() {
	go servers.HTTPServer()
	time.Sleep(4 * time.Second)

	go clients.HTTPClientHeadersGet()
	time.Sleep(5 * time.Second)
}
