package app

import (
	"database/sql"
	"example.com/golang-gin-auth/config"
	controller "example.com/golang-gin-auth/internal/controller/http/v1/echo"
	router "example.com/golang-gin-auth/internal/delivery/http/v1/echo"
	repository "example.com/golang-gin-auth/internal/repository/mysql"
	"example.com/golang-gin-auth/internal/service"
	"example.com/golang-gin-auth/pkg/logger"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

var (
	pgDatabaseURL    string
	mysqlDatabaseURL string
)

func ConnectPgDB() (*sql.DB, error) {
	pgDatabaseURL = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("PG_HOST"), os.Getenv("PG_PORT"), os.Getenv("PG_USER"), os.Getenv("PG_DB_NAME"), os.Getenv("PG_PASSWORD"), "false",
	)

	return nil, nil // repository.NewPostgresDB(pgDatabaseURL)
}

func ConnectMysqlDB() (*sql.DB, error) {
	mysqlDatabaseURL = fmt.Sprintf("%s:%s@/%s", os.Getenv("MYSQL_DB_USER"),
		os.Getenv("MYSQL_DB_PASSWORD"), os.Getenv("MYSQL_DB_NAME"),
	)
	return repository.NewMysqlDB(mysqlDatabaseURL)
}

func Run(cfg *config.Config) {
	l := logger.NewLogger(cfg.Log.Level)
	l.Info("Something")

	err := godotenv.Load()
	if err != nil {
		l.Fatal("Error loading .env file")
	}

	db, err := ConnectMysqlDB()
	if err != nil {
		l.Fatal("Database connection err %v: ", err)
	}

	defer db.Close()

	store := repository.NewMysqlStore(db)
	todoService := service.NewTodoService(store)

	todoCtrl := controller.NewEchoTodoCtrl(todoService)
	delivery := router.NewEchoRouter(todoCtrl, l)
	l.Error(delivery.InitAndServeRoutes(cfg.HTTP.Port))
}
