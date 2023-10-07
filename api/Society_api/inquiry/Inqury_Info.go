package inquiry

import (
	"github.com/gin-gonic/gin"
	"guanxingtuan_bck/global"
	"guanxingtuan_bck/models"
	"guanxingtuan_bck/models/res"
	"strconv"
)

func (InquiryApi) InquiryInfoView(c *gin.Context) {
	type InquiryRecv struct {
		AccessToken string `json:"access_token"`
	}

	var inquiry InquiryRecv
	var list []models.UserSociety
	studentID := c.Query("studentID")
	global.DB.Where("user_student_id = ?", studentID).Find(&list)
	l := len(list)
	if l == 0 {
		res.OK(list, "没有报名社团报名信息", c)
	} else {
		var retuList []string
		i := 0
		for _, each := range list {
			retuList = append(retuList, each.SOCIETY_ID)
			i += 1
		}
		msg := "共有" + strconv.Itoa(l) + "条社团报名信息"
		res.OK(retuList, msg, c)
	}
	global.Log.Infof("request=%#v", inquiry)
}
