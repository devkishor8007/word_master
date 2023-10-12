package controllers

import (
	"github.com/devkishor8007/word_master/src/database"
	"github.com/devkishor8007/word_master/src/models"
	"net/http"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"fmt"
	"time"
	"github.com/golang-jwt/jwt/v5"
	"github.com/devkishor8007/word_master/src/config"
	"github.com/devkishor8007/word_master/src/utilis"
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

func Login(writer http.ResponseWriter, request *http.Request) {
    writer.Header().Set("Content-Type", "application/json")

	secret := config.JWTSecret
    expiry := config.TokenExpiry

	var requestBody struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
	
	var user models.User

	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&requestBody); err != nil {
		http.Error(writer, "failed to decode JSON", http.StatusBadRequest)
		return
	}

	if err := database.DB.Where("email = ?", requestBody.Email).First(&user).Error; err != nil {
        http.Error(writer, "User not found", http.StatusNotFound)
        return
    }

	storedHashedPassword := user.Password 
	plainPassword := requestBody.Password

	// compare the password using bcrypt
	errPassword := bcrypt.CompareHashAndPassword([]byte(storedHashedPassword), []byte(plainPassword))
    if errPassword != nil {
        fmt.Println("Email and Password does not match")
        return
    }

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, utilis.JWTClaims{
		user.UserID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secret)
		if err != nil {
			http.Error(writer, "Failed to generate token", http.StatusInternalServerError)
			return
		}

	responseMessage := map[string]string{"access_token": tokenString}
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
