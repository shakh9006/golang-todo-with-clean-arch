package app

import (
	"example.com/golang-gin-auth/config"
	controller "example.com/golang-gin-auth/internal/controller/http/v1/echo"
	router "example.com/golang-gin-auth/internal/delivery/http/v1/echo"
	repository "example.com/golang-gin-auth/internal/repository/postgres"
	"example.com/golang-gin-auth/internal/service"
	"example.com/golang-gin-auth/pkg/logger"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func Run(cfg *config.Config) {
	l := logger.NewLogger(cfg.Log.Level)
	l.Info("Something")

	err := godotenv.Load()
	if err != nil {
		l.Fatal("Error loading .env file")
	}

	databaseURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("PG_HOST"), os.Getenv("PG_PORT"), os.Getenv("PG_USER"), os.Getenv("PG_DB_NAME"), os.Getenv("PG_PASSWORD"), "false",
	)

	db, err := repository.NewPostgresDB(databaseURL)
	if err != nil {
		l.Fatal("Database connection err %v: ", err)
	}

	defer db.Close()
	store := repository.NewPostgresStore(db)
	todoService := service.NewTodoService(store)

	todoCtrl := controller.NewEchoTodoCtrl(todoService)
	delivery := router.NewEchoRouter(todoCtrl, l)
	l.Error(delivery.InitAndServeRoutes(cfg.HTTP.Port))
}
