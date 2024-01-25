package router

import (
	"go-project/app/controller"
	"go-project/app/middleware"

	"github.com/gin-gonic/gin"
)

func SetupApiRouter(engine *gin.Engine) {
	apiRouter := engine.Group("api")

	// test
	testApiRouter := apiRouter.Group("/test", middleware.ParamFilter)
	testApiRouter.GET("/hello", controller.SayHello)

	// auth

	// user

	// ...
}
