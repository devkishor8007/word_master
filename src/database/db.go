package database

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
    "github.com/devkishor8007/word_master/src/models"
    "fmt"
)

var DB *gorm.DB

func InitDB() {
    dsn := "host=localhost user=postgres password=postba123@ dbname=word_master sslmode=disable"
    
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
