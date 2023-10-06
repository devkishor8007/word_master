package config

import (
    "time"
    "os"
)

var (
    JWTSecret   = []byte(os.Getenv("JWT_SECRET")) // Change this to your secret key
    TokenExpiry = 7 * 24 * time.Hour        // Token expiration after 7 days
)
