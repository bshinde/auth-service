package utils

import (
	"errors"
	"time"

	"auth-service/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(email string) (string, error) {
	expirationTime := time.Now().Add(10 * time.Minute)
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(config.SecretKey))

	if err != nil {
		log.WithFields(logrus.Fields{
			"email": email,
			"error": err,
		}).Error("Error signing the token")
		return "", err
	}

	log.WithFields(logrus.Fields{
		"email": email,
		"exp":   expirationTime,
	}).Info("Token generated successfully")
	return signedToken, nil
}

func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SecretKey), nil
	})

	if err != nil {
		log.WithFields(logrus.Fields{
			"token": tokenString,
			"error": err,
		}).Error("Error parsing token")
		return nil, errors.New("invalid token")
	}

	if !token.Valid {
		log.WithFields(logrus.Fields{
			"token": tokenString,
		}).Warn("Token is not valid")
		return nil, errors.New("invalid token")
	}

	// Check if the token is revoked
	if IsTokenRevoked(tokenString) {
		log.WithFields(logrus.Fields{
			"token": tokenString,
		}).Warn("Token is revoked")
		return nil, errors.New("token has been revoked")
	}

	log.WithFields(logrus.Fields{
		"email": claims.Email,
	}).Info("Token validated successfully")
	return claims, nil
}

// RenewToken generates a new token using the claims of an existing token
func RenewToken(oldToken string) (string, error) {
	claims, err := ValidateToken(oldToken)
	if err != nil {
		// Log the error with details
		log.Error("Error validating token", "error", err)
		return "", err
	}

	// Log the successful token validation with email
	log.Info("Token validated successfully", "email", claims.Email)

	// Generate a new token with the same email
	newToken, err := GenerateToken(claims.Email)
	if err != nil {
		// Log the error during token generation
		log.Error("Error generating new token", "error", err)
		return "", err
	}

	// Log successful generation of the new token
	log.Info("New token generated successfully", "email", claims.Email)

	return newToken, nil
}
