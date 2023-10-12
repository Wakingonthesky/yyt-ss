package controller

import (
	"net/http"
	"ytt-societyservice/controller/res"
	"ytt-societyservice/models/request"

	"ytt-societyservice/services"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	var f request.Signup
	c.BindJSON(&f)

	services.Signup(&f)

	c.JSON(http.StatusOK, res.OK(&gin.H{"isSignup": f}, "fail"))

	return

	if err := services.Signup(&f); err != nil {
		c.JSON(http.StatusOK, res.OK(&gin.H{"isSignup": false}, "fail"))

	} else {
		c.JSON(http.StatusOK, res.OK(&gin.H{"signup": true}, "success"))

	}

}
