package middleware

import (
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"github.com/devkishor8007/word_master/src/config"
	"github.com/devkishor8007/word_master/src/utilis"
)

func RequiredAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		
		secret := config.JWTSecret

		tokenString := request.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(writer, "Token not found in the header", http.StatusUnauthorized)
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &utilis.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})
		if err != nil || !token.Valid {
			http.Error(writer, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(writer, request)
	})
}
