package handlers

import (
	"auth-service/services"
	"auth-service/utils"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func SignUp(w http.ResponseWriter, r *http.Request) {
	var data map[string]string
	// Decode the request body
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		// Log error if decoding fails
		log.WithFields(logrus.Fields{
			"method": r.Method,
			"uri":    r.RequestURI,
			"error":  err.Error(),
		}).Error("Failed to decode request body")
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Call the SignUp service
	err = services.SignUp(data["email"], data["password"])
	if err != nil {
		// Log the error during user creation
		log.WithFields(logrus.Fields{
			"email": data["email"],
			"error": err.Error(),
		}).Error("User creation failed")
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Log successful user creation
	log.WithFields(logrus.Fields{
		"email": data["email"],
	}).Info("User created successfully")
	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"message": "user created"})
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	var data map[string]string
	// Decode the request body
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		// Log error if decoding fails
		log.WithFields(logrus.Fields{
			"method": r.Method,
			"uri":    r.RequestURI,
			"error":  err.Error(),
		}).Error("Failed to decode request body")
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Call the SignIn service
	token, err := services.SignIn(data["email"], data["password"])
	if err != nil {
		// Log error during sign-in
		log.WithFields(logrus.Fields{
			"email": data["email"],
			"error": err.Error(),
		}).Error("Invalid credentials")
		utils.RespondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	// Log successful sign-in
	log.WithFields(logrus.Fields{
		"email": data["email"],
	}).Info("SignIn successful, token generated")

	// Respond with the generated token
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"token": token})
}
