package sighup_api

import (
	"github.com/gin-gonic/gin"
	"guanxingtuan_bck/global"
	"guanxingtuan_bck/models"
	"guanxingtuan_bck/models/res"
	"time"
)

type Recv struct {
	AccessToken string `json:"access_token"`
	UserInfo    struct {
		UserName      string `json:"user_Name"`
		UserPhone     string `json:"user_Phone"`
		UserStudentID string `json:"user_StudentID"`
		SocietyID     string `json:"societyID"`
	} `json:"user_Info"`
}

func (SignUpApi) SignUpInfoView(c *gin.Context) {
	var (
		body Recv
	)
	if err := c.BindJSON(&body); err != nil {
		global.Log.Infof("request=%#v", body)
		global.Log.Errorf("%s", err)
		res.FailWithCode(400, c)
		return
	}
	
	global.DB.Create(&models.UserInfo{
		USER_REALNAME:  body.UserInfo.UserName,
		USER_PHONE:     body.UserInfo.UserPhone,
		USER_STUDENTID: body.UserInfo.UserStudentID,
		USER_HASSIGN:   1,
		CREATED_TIME:   time.Now(),
		UPDATED_TIME:   time.Now(),
	})

	global.DB.Create(&models.UserSociety{
		SOCIETY_ID:      body.UserInfo.SocietyID,
		USER_REALNAME:   body.UserInfo.UserName,
		USER_PHONE:      body.UserInfo.UserPhone,
		USER_STUDENT_ID: body.UserInfo.UserStudentID,
		CREATED_TIME:    time.Now(),
		UPDATED_TIME:    time.Now(),
	})
	res.OKWithMessage("报名成功", c)
}
