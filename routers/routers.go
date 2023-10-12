package routers

import (
	c "ytt-societyservice/controller"
	"ytt-societyservice/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("v1/society")
	{
		v1.POST("/signup", c.Signup)
		v1.POST("/hassign", middleware.Jwt(), c.Issign)
		v1.GET("/inquiry", middleware.Jwt(), c.Inquiry)
	}

	return router
}
