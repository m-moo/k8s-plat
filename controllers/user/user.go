package user

import "github.com/gin-gonic/gin"

//	@BasePath	/api

// user handler example
//	@Summary	get user example
//	@Schemes
//	@Description	get user
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	getuser
//	@Router			/user [get]
func GetUserHandler(c *gin.Context) {
	println("getuser")
}
