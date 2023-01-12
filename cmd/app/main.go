package main

import (
	"example.com/golang-gin-auth/config"
	"example.com/golang-gin-auth/internal/app"
	"log"
)

func init() {

}

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
