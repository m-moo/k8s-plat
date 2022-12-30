package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m-moo/k8s-plat/controllers/helm"
	"github.com/m-moo/k8s-plat/controllers/user"
	docs "github.com/m-moo/k8s-plat/docs"
	swaggerfiles "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
)

func Init() *gin.Engine {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api"

	api := router.Group("/api")
	{
		api.GET("/user", user.GetUserHandler)
		api.GET("/chart", helm.GetChartsHandler)
	}
	router.NoRoute(func(ctx *gin.Context) {
		ctx.AbortWithStatus(http.StatusNotFound)
	})

	router.GET("/swagger/*any", swagger.WrapHandler(swaggerfiles.Handler))

	return router
}
