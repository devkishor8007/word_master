package utilis

import (
    "github.com/dgrijalva/jwt-go"
)

type JWTClaims struct {
    UserID uint `json:"user_id"`
    jwt.StandardClaims
}
