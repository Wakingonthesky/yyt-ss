package SignUp_api

import (
	"github.com/gin-gonic/gin"
	"guanxingtuan_bck/global"
	"guanxingtuan_bck/models/res"
)

type SignUpRecv struct {
	AccessToken string `json:"access_token"`
	UserInfo    struct {
		UserName      string `json:"user_Name"`
		UserPhone     string `json:"user_Phone"`
		UserStudentID struct {
		} `json:"user_StudentID"`
		SocietyID string `json:"societyID"`
	} `json:"user_Info"`
}

type SignUpRtrn struct {
	status string
}

func (SignUpApi) SignUpInfoView(c *gin.Context) {
	var body SignUpRecv
	if err := c.BindJSON(&body); err != nil {
		res.FailWithCode(res.RequestInfoError, c)
		return
	}
	global.Log.Infof("request=%#v", body)

}
