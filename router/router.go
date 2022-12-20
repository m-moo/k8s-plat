package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	user "github.com/m-moo/k8s-plat/controllers/user"
	swaggerfiles "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
)

func Init() {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/user", user.GetUser)
	}
	router.NoRoute(func(ctx *gin.Context) {
		ctx.AbortWithStatus(http.StatusNotFound)
	})

	router.GET("/swagger/*any", swagger.WrapHandler(swaggerfiles.Handler))
	router.Run(":5000")
}
