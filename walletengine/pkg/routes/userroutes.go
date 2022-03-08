package routes

import (
	"RPay/walletengine/pkg/controllers"

	"github.com/gin-gonic/gin"
)

var RegisterRoutes = func(router *gin.Engine) {
	router.GET("/walletengine/user/details", func(c *gin.Context) {
		controllers.GetUsers(c)
	})

	router.POST("/walletengine/user/add", func(c *gin.Context) {
		controllers.AddUser(c)
	})

	router.GET("/favicon.ico", func(c *gin.Context) {
	})

}
