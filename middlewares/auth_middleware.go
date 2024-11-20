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
		// Extract the token from the Authorization header
		token := r.Header.Get("Authorization")
		if token == "" {
			// Log missing token error
			log.WithFields(logrus.Fields{
				"method": r.Method,
				"uri":    r.RequestURI,
			}).Warn("Missing token in request")
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}

		// Remove "Bearer " prefix if present
		if strings.HasPrefix(token, "Bearer ") {
			token = strings.TrimPrefix(token, "Bearer ")
		} else {
			// Log invalid token format error
			log.WithFields(logrus.Fields{
				"method": r.Method,
				"uri":    r.RequestURI,
				"token":  token,
			}).Warn("Invalid token format")
			http.Error(w, "invalid token format", http.StatusUnauthorized)
			return
		}

		// Validate the token
		_, err := utils.ValidateToken(token)
		if err != nil {
			// Log invalid token error
			log.WithFields(logrus.Fields{
				"method": r.Method,
				"uri":    r.RequestURI,
				"token":  token,
				"error":  err.Error(),
			}).Error("Invalid token")
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		// Log successful token validation
		log.WithFields(logrus.Fields{
			"method": r.Method,
			"uri":    r.RequestURI,
		}).Info("Valid token, request authorized")

		// If the token is valid, pass the request to the next handler
		next.ServeHTTP(w, r)
	})
}
