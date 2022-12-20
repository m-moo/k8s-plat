package controllers

import "github.com/gin-gonic/gin"

//	@BasePath	/api

// PingExample godoc
//	@Summary	get user example
//	@Schemes
//	@Description	get user
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	getuser
//	@Router			/user [get]
func GetUser(c *gin.Context) {
	println("getuser")
}
