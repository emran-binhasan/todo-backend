package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

func Load(){
	if err := godotenv.Load(); err != nil {
		log.Println("No, .env file found, using system environment variables")
	}
	setDefault("PORT", "3000")
	setDefault("DB_HOST", "localhost")
	setDefault("DB_PORT", "5432")
	setDefault("DB_SSLMODE", "disable")
}

func setDefault(key, defaultValue string){
	if os.Getenv(key) == "" {
		os.Setenv(key, defaultValue)
	}
}