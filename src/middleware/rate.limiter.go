package middleware

import (
	"golang.org/x/time/rate"
	"net/http"
)

const burst = 5

var limiter = rate.NewLimiter(10, burst)

func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if limiter.Allow() {
			next.ServeHTTP(writer, request)
		} else {
			http.Error(writer, "Rate limit exceeded", http.StatusTooManyRequests)
		}
	})
}
