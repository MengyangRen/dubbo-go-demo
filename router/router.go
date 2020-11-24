package router

import (
	"gin-dubbogo-consumer/controllers"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	Router = gin.New()

	//Dobubo-Demo-API接口
	var uctl controllers.UserController
	apiV1Group := Router.Group("/v1/dobubo-demo")
	{
		apiV1Group.GET("/user/*id", uctl.User)
		apiV1Group.GET("/users", uctl.Users)
		apiV1Group.POST("/user", uctl.Store)
		apiV1Group.PUT("/user", uctl.Update)
		apiV1Group.DELETE("/user", uctl.Destroy)
	}
}
