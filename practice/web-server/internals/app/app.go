package app

import (
	"context"
	"net/http"
	"time"
	"web-server/api"
	"web-server/api/middleware"
	"web-server/internals/app/handlers"
	"web-server/internals/app/processors"
	"web-server/internals/cfg"

	db3 "web-server/internals/app/db"

	log "github.com/sirupsen/logrus"

	"github.com/jackc/pgx/v4/pgxpool"
)

type AppServer struct {
	config cfg.Cfg
	ctx    context.Context
	srv    *http.Server
	db     *pgxpool.Pool
}

func NewServer(config cfg.Cfg, ctx context.Context) *AppServer {
	server := new(AppServer)
	server.ctx = ctx
	server.config = config
	return server
}

func (server *AppServer) Serve() {
	log.Println("Starting server")
	log.Println(server.config.GetDBString())

	var err error
	server.db, err = pgxpool.Connect(server.ctx, server.config.GetDBString()) //создаем пул соединений с БД и сохраним его для закрытия при остановке приложения
	if err != nil {
		log.Fatalln(err)
	}

	carsStorage := db3.NewCarsStorage(server.db)   //создаем экземпляр storage для работы с бд и всем что связано с машинами
	usersStorage := db3.NewUsersStorage(server.db) //создаем экземпляр storage для работы с бд и всем что связано с пользователями

	carsProcessor := processors.NewCarsProcessor(carsStorage) //инициализируем процессоры соотвествующими storage
	usersProcessor := processors.NewUsersProcessor(usersStorage)

	userHandler := handlers.NewUsersHandler(usersProcessor) //инициализируем handlerы нашими процессорами
	carsHandler := handlers.NewCarsHandler(carsProcessor)

	routes := api.CreateRoutes(userHandler, carsHandler)
	routes.Use(middleware.RequestLog)

	server.srv = &http.Server{ //в отличие от примеров http, здесь мы передаем наш server в поле структуры, для работы в Shutdown
		Addr:    ":" + server.config.Port,
		Handler: routes,
	}

	log.Println("Server started")

	err = server.srv.ListenAndServe() //запускаем сервер

	if err != nil && err != http.ErrServerClosed {
		log.Fatalln(err)
	}

	return
}

func (server *AppServer) Shutdown() {
	log.Println("Server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	server.db.Close() //закрываем соединение с БД
	defer func() {
		cancel()
	}()
	var err error
	if err = server.srv.Shutdown(ctxShutDown); err != nil { //выключаем сервер, с ограниченным по времени контекстом
		log.Fatalf("Server shutdown failed:%s", err)
	}

	log.Println("Server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}
}
