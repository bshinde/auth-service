package middlewares

import (
	"auth-service/utils"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the incoming request details
		log.Info("Incoming request", "method", r.Method, "url", r.URL.Path)

		// Extract the token from the Authorization header
		token := r.Header.Get("Authorization")
		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			log.Error("Missing or invalid token", "method", r.Method, "url", r.URL.Path)
			http.Error(w, "missing or invalid token", http.StatusUnauthorized)
			return
		}
		token = strings.TrimPrefix(token, "Bearer ")

		// Validate the token
		_, err := utils.ValidateToken(token)
		if err != nil {
			log.Error("Invalid or revoked token", "method", r.Method, "url", r.URL.Path, "error", err.Error())
			http.Error(w, "invalid or revoked token", http.StatusUnauthorized)
			return
		}

		// Log success after validation
		log.Info("Token validated successfully", "method", r.Method, "url", r.URL.Path)

		// Call the next handler if token is valid
		next.ServeHTTP(w, r)
	})
}
