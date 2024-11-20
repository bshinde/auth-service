package controllers

import (
	"auth-service/handlers"
	"auth-service/middlewares"

	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/signup", handlers.SignUp).Methods("POST")
	router.HandleFunc("/signin", handlers.SignIn).Methods("POST")

	protected := router.PathPrefix("/api").Subrouter()
	protected.Use(middlewares.AuthMiddleware)
	protected.HandleFunc("/protected", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Access granted"))
	}).Methods("GET")
}
