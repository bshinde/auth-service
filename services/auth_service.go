package services

import (
	"auth-service/models"
	"auth-service/utils"
	"errors"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

var log = logrus.New()

func SignUp(email, password string) error {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// Log the error and return it
		log.WithFields(logrus.Fields{
			"email": email,
			"error": err,
		}).Error("Error hashing password")
		return err
	}

	// Create the user in the database
	err = models.CreateUser(email, string(hashedPassword))
	if err != nil {
		// Log any errors from creating the user
		log.WithFields(logrus.Fields{
			"email": email,
			"error": err,
		}).Error("Error creating user in the database")
		return err
	}

	// Log successful sign-up
	log.WithFields(logrus.Fields{
		"email": email,
	}).Info("User signed up successfully")

	return nil
}

func SignIn(email, password string) (string, error) {
	// Get the hashed password for the user
	hashedPassword, err := models.GetUser(email)
	if err != nil {
		// Log the error and return an invalid credentials error
		log.WithFields(logrus.Fields{
			"email": email,
			"error": err,
		}).Error("Error retrieving user from database")
		return "", errors.New("invalid credentials")
	}

	// Compare the hashed password with the provided password
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		// Log the invalid credentials error
		log.WithFields(logrus.Fields{
			"email": email,
		}).Warn("Invalid credentials provided during sign-in")
		return "", errors.New("invalid credentials")
	}

	// Generate a token for the user after successful sign-in
	token, err := utils.GenerateToken(email)
	if err != nil {
		// Log the error in token generation
		log.WithFields(logrus.Fields{
			"email": email,
			"error": err,
		}).Error("Error generating token")
		return "", err
	}

	// Log successful sign-in
	log.WithFields(logrus.Fields{
		"email": email,
	}).Info("User signed in successfully")

	return token, nil
}
