package controllers

import (
	"encoding/json"
	"github.com/devkishor8007/word_master/src/database"
	"github.com/devkishor8007/word_master/src/models"
	"net/http"
)

// Get godoc
// @Summary Get comment
// @Description Get comment in articles
// @Tags comment
// @Accept  json
// @Produce  json
// @Router /comment [get]
func GetComments(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var comment []models.Comment

	if err := database.DB.Find(&comment).Error; err != nil {
		http.Error(writer, "Comment not found", http.StatusNotFound)
		return
	}

	responseMessage := map[string]interface{}{
		"status":  200,
		"message": "fetched successfully",
		"data":    comment,
		"count":   len(comment),
	}
	jsonResponse, _ := json.Marshal(responseMessage)

	writer.WriteHeader(http.StatusOK)
	writer.Write(jsonResponse)
}

// Create godoc
// @Summary Create comment
// @Description Create comment in articles
// @Tags comment
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param Authorization header string true "Token value"
// @Router /comment [post]
func CreateComment(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var comment models.Comment

	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&comment); err != nil {
		http.Error(writer, "failed to decode JSON", http.StatusBadRequest)
		return
	}

	database.DB.Create(&comment)

	responseMessage := map[string]string{"message": "comment created successully"}
	jsonResponse, _ := json.Marshal(responseMessage)

	writer.WriteHeader(http.StatusOK)
	writer.Write(jsonResponse)
}
