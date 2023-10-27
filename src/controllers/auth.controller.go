package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/devkishor8007/word_master/src/config"
	"github.com/devkishor8007/word_master/src/database"
	"github.com/devkishor8007/word_master/src/models"
	"github.com/devkishor8007/word_master/src/utilis"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type User struct {
	UserID   uint   `gorm:"primaryKey" json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type BodyLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Create godoc
// @Summary Register user
// @Description Register user api if not exists
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   payload     body    User     true        "User Data"
// @Router /signup [post]
func Register(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	// create a new user instance
	var user User

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

// Login godoc
// @Summary Login user
// @Description Login user api with email and password
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   payload     body    BodyLogin     true        "User Data"
// @Router /signin [post]
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
