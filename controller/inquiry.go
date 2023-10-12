package controller

import (
	"net/http"
	"ytt-societyservice/controller/res"
	"ytt-societyservice/services"

	"github.com/gin-gonic/gin"
)

func Inquiry(c *gin.Context) {
	studentID := c.Query("studentID")

	if ls, err := services.Inquiry(studentID); err != nil {
		c.JSON(http.StatusOK, res.OK(&gin.H{"isSignup": false}, "success"))
	} else {

		c.JSON(http.StatusOK, res.OK(&gin.H{"societyList": ls}, "success"))
	}
}
