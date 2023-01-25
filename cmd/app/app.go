package app

import (
	"database/sql"
	"example.com/golang-gin-auth/config"
	delivery "example.com/golang-gin-auth/internal/app/delivery/http/echo"
	"example.com/golang-gin-auth/internal/app/models"
	"example.com/golang-gin-auth/internal/app/service"
	"example.com/golang-gin-auth/pkg/logger"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

var (
	appLogger *logger.Logger
)

func connectToDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PG_HOST"), fmt.Sprint(os.Getenv("PG_PORT")), os.Getenv("PG_USER"), os.Getenv("PG_PASSWORD"), os.Getenv("PG_DB_NAME"))

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func NewApp(cfg *config.Config) (*sql.DB, error) {
	appLogger = logger.NewLogger(cfg.Log.Level)
	appLogger.Info("Something")

	if err := godotenv.Load(); err != nil {
		appLogger.Fatal("Error loading .env file")
		return nil, err
	}

	db, err := connectToDB()
	if err != nil {
		appLogger.Fatal("Database connection err %v: ", err)
		return nil, err
	}
	defer db.Close()

	return db, nil
}

func Run(db *sql.DB, cfg *config.Config) {
	store := models.NewTodoRepository(db)
	todoService := service.NewTodoService(store)

	todoCtrl := delivery.NewEchoTodoCtrl(todoService)
	router := delivery.NewEchoRouter(todoCtrl, logger.NewLogger(cfg.Level))
	appLogger.Error(router.InitAndServeRoutes(cfg.HTTP.Port))
}
