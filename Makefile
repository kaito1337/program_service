#!/usr/bin/make -f
.DEFAULT_GOAL := help

include .env

help:
	@echo "Program Service"
	@echo "Usage:"
	@echo "> make run - Run Backend Server"
	@echo "> make dev - Run Backend Server in dev mode"
	@echo "> make db-status - Check database status"
	@echo "> make db-up - Create database"
	@echo "> make db-down - Drop database"
	@echo "Dependencies:"
	@echo "Air (go install github.com/cosmtrek/air@latest)"
	@echo "Goose (go install github.com/pressly/goose/v3/cmd/goose@latest)"


run:
	@echo " > Running Backend Server"
	go run .\cmd\app\main.go

dev:
	@echo " > Running Backend Server in dev mode"
			air -c .air.toml

db-status:
	@echo " > Checking database status"
		goose -dir='./internal/database/migrations' postgres 'host=${DB_HOST} port=${DB_PORT} user=${DB_USERNAME} password=${DB_PASSWORD} dbname=${DB_DATABASE} sslmode=disable' status

db-up:
	@echo " > Creating database"
		goose -dir='./internal/database/migrations' postgres 'host=${DB_HOST} port=${DB_PORT} user=${DB_USERNAME} password=${DB_PASSWORD} dbname=${DB_DATABASE} sslmode=disable' up

db-down:
	@echo " > Dropping database"
		goose -dir='./internal/database/migrations' postgres 'host=${DB_HOST} port=${DB_PORT} user=${DB_USERNAME} password=${DB_PASSWORD} dbname=${DB_DATABASE} sslmode=disable' down-to -1
