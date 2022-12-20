package main

import router "github.com/m-moo/k8s-plat/router"

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample swagger doc for test

//	@host		localhost:5000
//	@BasePath	/api

//	@securityDefinitions.basic	BasicAuth
func main() {
	router.Init()
}
