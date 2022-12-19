package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	user "github.com/m-moo/k8s-plat/controllers/user"
)

func init() {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/user", user.GetUser)
	}
	router.NoRoute(func(ctx *gin.Context) {
		ctx.AbortWithStatus(http.StatusNotFound)
	})
	router.Run(":5000")
}
