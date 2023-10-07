package inquiry

import (
	"github.com/gin-gonic/gin"
	"guanxingtuan_bck/global"
	"guanxingtuan_bck/models"
	"guanxingtuan_bck/models/res"
)

func (InquiryApi) InquiryInfoView(c *gin.Context) {
	type InquiryRecv struct {
		AccessToken string `json:"access_token"`
	}

	var inquiry InquiryRecv
	var list []models.UserSociety
	studentID := c.Query("studentID")
	err := c.BindJSON(&inquiry)
	if err != nil {
		res.FailWithCode(res.RequestInfoError, c)
		return
	}
	global.DB.Where("user_student_id = ?", studentID).Find(&list)
	res.OKWithData(list, c)
	global.Log.Infof("request=%#v", inquiry)
}
