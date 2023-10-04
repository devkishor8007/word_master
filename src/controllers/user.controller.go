package controllers

import (
	"github.com/devkishor8007/word_master/src/database"
	"github.com/devkishor8007/word_master/src/models"
	"net/http"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

func Register(writer http.ResponseWriter, request *http.Request) {
    writer.Header().Set("Content-Type", "application/json")
	
	// create a new user instance
	var user models.User

	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(writer, "failed to decode JSON", http.StatusBadRequest)
		return
	}

	// extract the password from the user struct
	password := user.Password 

	// hash the password using bcrypt
	passwordBytes := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// store the hashed password in the user struct
	user.Password = string(hashedPassword)

	// store in the database 'user'
	database.DB.Create(&user)

	responseMessage := map[string]string{"message": "user created successully"}
	jsonResponse, _ := json.Marshal(responseMessage)
	
	writer.WriteHeader(http.StatusOK)
	writer.Write(jsonResponse)
}

func Home(writer http.ResponseWriter, request *http.Request) {
    writer.Header().Set("Content-Type", "application/json")
	
	writer.WriteHeader(http.StatusOK)
	response := []byte(`{"message": "hello"}`)
	writer.Write(response)
}

