package main

import (
	"awesomeProject34/internal/adapter"
	booksrepository "awesomeProject34/internal/books/repository"
	booksservice "awesomeProject34/internal/books/service"
	"awesomeProject34/internal/config"
	router "awesomeProject34/internal/transport/http"
	"awesomeProject34/pkg/postgres"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Ошибка открытия файла лога:", err)
	}
	logger := log.New(file, "BOOKSERVICE: ", log.Ldate|log.Ltime|log.Lshortfile)

	cfg := config.NewConfig()

	db, err := postgres.NewPostgres(cfg.DB)
	if err != nil {
		logger.Fatal("Ошибка подключения к базе данных:", err)
	}

	repo := booksrepository.NewRepository(db)

	whetherAdapter := adapter.NewAdapter()

	service := booksservice.NewService(repo, whetherAdapter)

	handler := router.NewHandler(service)

	r := router.NewRouter(cfg.RouterConfig, handler)

	r.Run()
}
