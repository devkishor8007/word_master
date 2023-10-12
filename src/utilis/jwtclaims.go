package utilis

import "github.com/golang-jwt/jwt/v5"

type JWTClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}
