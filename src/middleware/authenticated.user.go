package middleware

import (
	"fmt"
    "github.com/golang-jwt/jwt/v5"
	"net/http"
    "github.com/devkishor8007/word_master/src/utilis"
    "github.com/devkishor8007/word_master/src/config"
)

func JwtParserClaimss(r *http.Request) (*utilis.JWTClaims, error) {
    tokenString := r.Header.Get("Authorization")

    if tokenString == "" {
        return nil, fmt.Errorf("Authorization header is missing")
    }

    secretKey := config.JWTSecret

    token, err := jwt.ParseWithClaims(tokenString, &utilis.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
        return secretKey, nil
    })

    if err != nil {
        return nil, fmt.Errorf("Invalid token: %v", err)
    }

    if token.Valid {
        claims, ok := token.Claims.(*utilis.JWTClaims)
        if !ok {
            return nil, fmt.Errorf("Invalid claims")
        }

        return claims, nil
    }

    return nil, fmt.Errorf("Invalid token")
}
