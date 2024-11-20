package models

import (
	"errors"

	"github.com/sirupsen/logrus"
)

type User struct {
	Email    string
	Password string
}

var users = map[string]string{} // Email: HashedPassword
var log = logrus.New()

// CreateUser adds a new user to the users map
func CreateUser(email, password string) error {
	// Check if the user already exists
	if _, exists := users[email]; exists {
		// Log that the user already exists
		log.WithFields(logrus.Fields{
			"email": email,
		}).Warn("User already exists")
		return errors.New("user already exists")
	}

	// Add user to the "database" (map)
	users[email] = password

	// Log successful user creation
	log.WithFields(logrus.Fields{
		"email": email,
	}).Info("User created successfully")
	return nil
}

// GetUser retrieves the password of the user
func GetUser(email string) (string, error) {
	// Retrieve password from the "database" (map)
	password, exists := users[email]
	if !exists {
		// Log that the user was not found
		log.WithFields(logrus.Fields{
			"email": email,
		}).Warn("User not found")
		return "", errors.New("user not found")
	}

	// Log successful user retrieval
	log.WithFields(logrus.Fields{
		"email": email,
	}).Info("User found successfully")

	return password, nil
}
