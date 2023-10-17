package routers

import (
	c "ytt-societyservice/controller"
	"ytt-societyservice/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("v1")
	{
		v1.POST("/signup", middleware.JwtPost(), c.Signup)
		v1.GET("/issign", middleware.JwtGet(), c.Issign)
		v1.GET("/inquiry", middleware.JwtGet(), c.Inquiry)
	}

	return router
}
