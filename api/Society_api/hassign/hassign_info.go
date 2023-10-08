package hassign

import (
	"github.com/gin-gonic/gin"
	"guanxingtuan_bck/global"
	"guanxingtuan_bck/models"
	"guanxingtuan_bck/models/res"
)

type InquiryRecv struct {
	StudentID string `json:"user_StudentID"`
	SocietyID string `json:"societyID"`
}

func (HasSignApi) HasSignInfoView(c *gin.Context) {
	var (
		body   InquiryRecv
		result []models.UserSociety
	)
	//读取请求
	if err := c.BindJSON(&body); err != nil {
		global.Log.Infof("request=%#v", body)
		global.Log.Errorf("%s", err)
		res.FailWithCode(400, c)
	}
	//在User_Society_Info表中查找
	global.DB.Where("user_student_id=?", body.StudentID).Find(&result)
	flag := 0
	for _, index := range result {
		if index.USER_STUDENT_ID == body.StudentID {
			flag = 1
			res.OKWithMessage("用户已报名", c)
			return
		}
	}
	if flag == 0 {
		res.OKWithMessage("用户未报名", c)
	}
}
