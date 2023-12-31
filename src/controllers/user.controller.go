package controllers

import (
	"encoding/json"
	"github.com/devkishor8007/word_master/src/database"
	"github.com/devkishor8007/word_master/src/helper"
	"github.com/devkishor8007/word_master/src/models"
	"net/http"
)

// Get godoc
// @Summary Get Profile
// @Description Get user profile info
// @Tags users
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param Authorization header string true "Token value"
// @Router /user/profile [get]
func ViewProfile(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	claims, err := helper.JwtParserClaims(request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusUnauthorized)
		return
	}

	var user models.User

	if err := database.DB.Where("user_id = ?", claims.UserID).Preload("Articles", "author_id = ?", claims.UserID).Preload("Articles.User").Preload("Articles.Category").First(&user).Error; err != nil {
		http.Error(writer, "User not found", http.StatusNotFound)
		return
	}

	responseMessage := map[string]interface{}{
		"status":  200,
		"message": "fetched successfully",
		"data":    user,
	}
	jsonResponse, _ := json.Marshal(responseMessage)

	writer.WriteHeader(http.StatusOK)
	writer.Write(jsonResponse)
}
