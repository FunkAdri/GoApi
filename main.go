package main

import (
	"log"
	"net/http"

	"github.com/synbioz/go_api/config"
	"github.com/synbioz/go_api/models"
)

func main() {
	config.DatabaseInit()
	router := InitializeRouter()

	// Populate database
	models.NewUser(&models.User{Username: "FunkAdri", Mail: "FunkAdri@gmail.com", Password: "123456"})

	log.Fatal(http.ListenAndServe(":5432", router))
}
