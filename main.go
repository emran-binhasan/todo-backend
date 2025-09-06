package main

import (
	"todo-backend/config"
	"todo-backend/database"
)

func main() {
	config.Load()
	database.Connect()
	database.Migrate()
}