package controllers

import (
	"encoding/json"
	"github.com/devkishor8007/word_master/src/database"
	"github.com/devkishor8007/word_master/src/middleware"
	"net/http"
)

func ViewProfile(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	claims, err := middleware.JwtParserClaimss(request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusUnauthorized)
		return
	}

	type User struct {
		UserID   uint   `json:"user_id"`
		Email    string `json:"email"`
		Username string `json:"username"`
	}

	var user User

	if err := database.DB.Where("user_id = ?", claims.UserID).First(&user).Error; err != nil {
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
