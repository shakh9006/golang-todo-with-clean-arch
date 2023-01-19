package main

import (
	"example.com/golang-gin-auth/cmd/app"
	"example.com/golang-gin-auth/config"
	"log"
)

var (
	cfg *config.Config
	err error
)

func init() {
	cfg, err = config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %v", err.Error())
	}
}

func main() {
	db, err := app.NewApp(cfg)
	if err != nil {
		log.Fatalf("New app error: %v", err.Error())
	}

	app.Run(db, cfg)
}
