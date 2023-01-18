package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func NewMysqlDB(databaseURL string) (*sql.DB, error) {
	fmt.Println("databaseURL: ", databaseURL)
	db, err := sql.Open("mysql", databaseURL)
	if err != nil {
		return nil, err
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}
