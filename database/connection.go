package database

import (
	"fmt"
	"log"
	"os"
	"todo-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Connect initializes the database connection
func Connect() {
	var err error
	
	// Build connection string from environment variables
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)
	
	// Connect to database
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Enable SQL logging in development
	})
	
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	
	log.Println("Successfully connected to database")
}

// Migrate runs database migrations
func Migrate() {
	err := DB.AutoMigrate(&models.Todo{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	log.Println("Database migration completed successfully")
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return DB
}