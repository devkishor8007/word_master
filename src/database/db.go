package database

import (
	"fmt"
	"github.com/devkishor8007/word_master/src/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
)

var DB *gorm.DB

func InitDB() {
	env_err := godotenv.Load()
	if env_err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbUser := os.Getenv("DB_USER")
	dbPort := os.Getenv("DB_PORT")

	port, err_port := strconv.Atoi(dbPort)
	if err_port != nil {
		log.Fatal("Invalid port:", err_port)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", dbHost, dbUser, dbPassword, dbName, port)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = DB.AutoMigrate(&models.User{}, &models.Article{}, &models.Comment{}, &models.Category{})

	if err != nil {
		log.Printf("Error performing auto-migration: %v", err)
		return
	}

	fmt.Println("Database connection along with migration successfully")
}
