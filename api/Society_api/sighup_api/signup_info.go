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
type SignUpRtrn struct {
	status string
}

type InquiryRecv struct {
	StudentID string `json:"user_StudentID"`
	SocietyID string `json:"societyID"`
}

func (SignUpApi) SignUpInfoView(c *gin.Context) {
	var (
		body   Recv
		result []models.UserSociety
		user   models.UserInfo
	)
	if err := c.BindJSON(&body); err != nil {
		global.Log.Infof("request=%#v", body)
		global.Log.Errorf("%s", err)
		res.FailWithCode(400, c)
		return
	}

	global.DB.Where("user_realname=?", body.UserInfo.UserName).First(&user)
	if user.USER_HASSIGN == 1 { //是否用户已经注册
		res.OKWithMessage("用户已报名", c)
	} else {
		res.OKWithMessage("用户未报名", c)
		global.DB.Create(&models.UserInfo{
			USER_REALNAME:  body.UserInfo.UserName,
			USER_PHONE:     body.UserInfo.UserPhone,
			USER_STUDNETID: body.UserInfo.UserStudentID,
			USER_HASSIGN:   1,
			CREATED_TIME:   time.Now(),
			UPDATED_TIME:   time.Now(),
		})
	}

	global.DB.Where("user_student_id=?", body.UserInfo.UserStudentID).Find(&result)
	flag := 0
	for _, index := range result {
		if index.SOCIETY_ID == body.UserInfo.SocietyID {
			flag = 1
			res.OKWithMessage("社团已报名", c)
		}
	}
	if flag == 0 {
		global.DB.Create(&models.UserSociety{
			SOCIETY_ID:     body.UserInfo.SocietyID,
			USER_REALNAME:  body.UserInfo.UserName,
			USER_PHONE:     body.UserInfo.UserPhone,
			USER_STUDNETID: body.UserInfo.UserStudentID,
			CREATED_TIME:   time.Now(),
			UPDATED_TIME:   time.Now(),
		})
	}

}

func (SignUpApi) InquirySigned(c *gin.Context) {
	var (
		body   InquiryRecv
		result []models.UserSociety
	)
	if err := c.BindJSON(&body); err != nil {
		global.Log.Infof("request=%#v", body)
		global.Log.Errorf("%s", err)
		res.FailWithCode(400, c)
	}

	global.DB.Where("society_id=?", body.SocietyID).Find(&result)
	flag := 0
	for _, index := range result {
		if index.USER_STUDNETID == body.StudentID {
			flag = 1
			res.OKWithMessage("用户已报名", c)
			return
		}
	}
	if flag == 0 {
		res.OKWithMessage("用户未报名", c)
	}
}
