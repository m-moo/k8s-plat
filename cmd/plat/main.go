package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/m-moo/k8s-plat/datastore"
	router "github.com/m-moo/k8s-plat/router"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample swagger doc for test

//	@host		localhost:5000
//	@BasePath	/api

//	@securityDefinitions.basic	BasicAuth
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	datastore.Init()

	r := router.Init()
	r.Run(":5000")
}
