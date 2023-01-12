.PHONY: build
build:
		go build -v ./cmd/app

migrate-up:
	migrate -path migrations/ -database 'postgres://postgres:postgres@0.0.0.0:5432/postgres?sslmode=disable' up

migrate-down:
	migrate -path migrations/ -database 'postgres://postgres:postgres@0.0.0.0:5432/postgres?sslmode=disable' up

migrate-drop:
	migrate -path migrations/ -database 'postgres://postgres:postgres@0.0.0.0:5432/postgres?sslmode=disable' up

.DEFAULT_GOAL := build