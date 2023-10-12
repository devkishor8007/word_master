package middleware

import (
	"github.com/devkishor8007/word_master/src/config"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

func RequiredAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		secretKey := config.JWTSecret

		tokenString := request.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(writer, "Token not found in the header", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(writer, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(writer, request)
	})
}
