package controllers

import (
	"encoding/json"
	"github.com/devkishor8007/word_master/src/database"
	"github.com/devkishor8007/word_master/src/helper"
	"github.com/devkishor8007/word_master/src/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

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

func GetOwnArticles(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	claims, err := helper.JwtParserClaims(request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusUnauthorized)
		return
	}

	var article []models.Article

	if err := database.DB.Where("author_id = ?", claims.UserID).Preload("User", func(db *gorm.DB) *gorm.DB {
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

	claims, err := helper.JwtParserClaims(request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusUnauthorized)
		return
	}

	var article models.Article

	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&article); err != nil {
		http.Error(writer, "failed to decode JSON", http.StatusBadRequest)
		return
	}

	article.AuthorID = claims.UserID
	database.DB.Create(&article)

	responseMessage := map[string]string{"message": "article created successfully"}
	jsonResponse, _ := json.Marshal(responseMessage)

	writer.WriteHeader(http.StatusOK)
	writer.Write(jsonResponse)
}

func DeleteArticle(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(request)
	articleID := vars["article_id"]

	articleIDInt, errArticleID := strconv.Atoi(articleID)

	if errArticleID != nil {
		http.Error(writer, "Invalid article_id", http.StatusBadRequest)
		return
	}

	claims, err := helper.JwtParserClaims(request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusUnauthorized)
		return
	}

	var article models.Article

	result := database.DB.Where("article_id = ? AND author_id = ?", articleIDInt, claims.UserID).First(&article)

	if result.Error != nil {
		http.Error(writer, "Article not found", http.StatusNotFound)
		return
	} else {
		deleteResult := database.DB.Delete(&article)

		if deleteResult.Error != nil {
			http.Error(writer, "Error deleting record", http.StatusInternalServerError)
			return
		}
	}

	responseMessage := map[string]string{"message": "article deleted successfully"}
	jsonResponse, _ := json.Marshal(responseMessage)

	writer.WriteHeader(http.StatusOK)
	writer.Write(jsonResponse)
}
