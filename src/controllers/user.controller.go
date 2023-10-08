package controllers

import (
	"net/http"
	"encoding/json"
	"github.com/devkishor8007/word_master/src/database"
)

func ViewProfile(writer http.ResponseWriter, request *http.Request) {
    writer.Header().Set("Content-Type", "application/json")

	var requestBody struct {
		UserID   uint   `json:"user_id"`
    }

	type User struct {
		UserID   uint   `json:"user_id"`
		Email    string `json:"email"`
		Username    string `json:"username"`
    }

	var user User

	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&requestBody); err != nil {
		http.Error(writer, "failed to decode JSON", http.StatusBadRequest)
		return
	}

	if err := database.DB.Where("user_id = ?", requestBody.UserID).First(&user).Error; err != nil {
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
