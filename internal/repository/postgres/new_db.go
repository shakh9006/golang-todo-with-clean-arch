package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

func NewPostgresDB(databaseURL string) (*sql.DB, error) {
	fmt.Println(databaseURL)
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
