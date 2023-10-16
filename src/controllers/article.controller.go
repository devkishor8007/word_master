package controllers

import (
	"encoding/json"
	"github.com/devkishor8007/word_master/src/database"
	"github.com/devkishor8007/word_master/src/middleware"
	"github.com/devkishor8007/word_master/src/models"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type Article struct {
	ArticleID       uint      `gorm:"primaryKey" json:"article_id"`
	Title           string    `json:"title"`
	Content         string    `json:"content"`
	PublicationDate time.Time `json:"publication_date"`
	AuthorID        uint      `json:"author_id"`
	CategoryID      uint      `json:"category_id"`
}

func GetArticles(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var article []models.Article

	if err := database.DB.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("user_id", "username", "email")
	}).Preload("Category").Find(&article).Error; err != nil {
		http.Error(writer, "Article not found", http.StatusNotFound)
		return
	}

	responseMessage := map[string]interface{}{
		"status":  200,
		"message": "fetched successfully",
		"data":    article,
		"count":   len(article),
	}
	jsonResponse, _ := json.Marshal(responseMessage)

	writer.WriteHeader(http.StatusOK)
	writer.Write(jsonResponse)
}

func CreateArticle(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	claims, err := middleware.JwtParserClaimss(request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusUnauthorized)
		return
	}

	var article Article

	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&article); err != nil {
		http.Error(writer, "failed to decode JSON", http.StatusBadRequest)
		return
	}

	article.AuthorID = claims.UserID
	database.DB.Create(&article)

	responseMessage := map[string]string{"message": "article created successully"}
	jsonResponse, _ := json.Marshal(responseMessage)

	writer.WriteHeader(http.StatusOK)
	writer.Write(jsonResponse)
}
