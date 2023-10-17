package controller

import (
	"net/http"
	"ytt-societyservice/controller/res"
	"ytt-societyservice/services"

	"github.com/gin-gonic/gin"
)

func Issign(c *gin.Context) {
	UID := c.Query("uid")
	societyID := c.Query("societyID")

	if flag, err := services.IsSign(UID, societyID); err != nil {
		c.JSON(http.StatusOK, res.OK(&gin.H{"isSignup": false}, "success"))
	} else {

		if flag == false {
			c.JSON(http.StatusOK, res.OK(&gin.H{"isSignup": false}, "success"))
		} else {
			c.JSON(http.StatusOK, res.OK(&gin.H{"isSignup": true}, "success"))
		}
	}
}
