package controllers

import (
	"github.com/devkishor8007/word_master/src/database"
	"github.com/devkishor8007/word_master/src/models"
	"net/http"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"fmt"
	"time"
	"github.com/dgrijalva/jwt-go"
)

type JWTClaims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

var (
    jwtSecret   = []byte("your-secret-key") // change the secret-key
    tokenExpiry = 7 * 24 * time.Hour   // Token expiration after 7 days        
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
		UserID: user.UserID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenExpiry).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

		// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtSecret)
		if err != nil {
			http.Error(writer, "Failed to generate token", http.StatusInternalServerError)
			return
		}

	responseMessage := map[string]string{"access_token": tokenString}
	jsonResponse, _ := json.Marshal(responseMessage)
	
	writer.WriteHeader(http.StatusOK)
	writer.Write(jsonResponse)
}

func RequiredAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		tokenString := request.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(writer, "Token not found in the header", http.StatusUnauthorized)
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			http.Error(writer, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Token is valid; proceed to the protected route
		next.ServeHTTP(writer, request)
	})
}

func Home(writer http.ResponseWriter, request *http.Request) {
    writer.Header().Set("Content-Type", "application/json")
	
	writer.WriteHeader(http.StatusOK)
	response := []byte(`{"message": "hello"}`)
	writer.Write(response)
}

