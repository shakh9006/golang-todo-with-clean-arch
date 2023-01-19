.PHONY: build
build:
		go build -o main ./cmd/main.go

migrate-up:
	migrate -path migrations/ -database 'postgres://postgres:postgres@0.0.0.0:5432/postgres?sslmode=disable' up

migrate-down:
	migrate -path migrations/ -database 'postgres://postgres:postgres@0.0.0.0:5432/postgres?sslmode=disable' down

migrate-drop:
	migrate -path migrations/ -database 'postgres://postgres:postgres@0.0.0.0:5432/postgres?sslmode=disable' drop

.DEFAULT_GOAL := build