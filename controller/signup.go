package controller

import (
	"net/http"
	"ytt-societyservice/controller/res"
	"ytt-societyservice/models/request"

	"ytt-societyservice/services"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Signup(c *gin.Context) {
	var f request.Signup
	c.ShouldBindBodyWith(&f, binding.JSON)

	//services.Signup(&f)

	//c.JSON(http.StatusOK, res.OK(&gin.H{"isSignup": f}, "fail"))

	if err := services.Signup(&f); err != nil {
		c.JSON(http.StatusOK, res.OK(&gin.H{"isSignup": f}, "fail"))

	} else {
		c.JSON(http.StatusOK, res.OK(&gin.H{"isSignup": f}, "success"))

	}

}
