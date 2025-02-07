package main

import (
	"context"
	"os"
	"os/signal"
	"web-server/internals/app"
	"web-server/internals/cfg"

	log "github.com/sirupsen/logrus"
)

func main() {
	config := cfg.LoadAndStoreConfig()

	ctx, cancel := context.WithCancel(context.Background())

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	server := app.NewServer(config, ctx)

	go func() { // горутина, ожидающая сообщения системы
		oscall := <-c //как только чет пришло
		log.Printf("system call:%+v", oscall)
		server.Shutdown() //выключаем сервер
		cancel()          //отменяем контекст
	}()

	server.Serve()
}
