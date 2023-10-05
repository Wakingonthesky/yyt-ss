package loginJson_api

import (
	"github.com/gin-gonic/gin"
	"guanxingtuan_bck/global"
	"guanxingtuan_bck/models/res"
)

type LoginInfo struct {
	UserName     string `json:"user_name"`
	UserPassword string `json:"user_password"`
}

type Recv struct {
	LoginInfo LoginInfo `json:"login_info"`
	UserType  string    `json:"user_type"`
}

func (LoginJsonApi) LoginJsonInfoView(c *gin.Context) {
	var body Recv
	if err := c.BindJSON(&body); err != nil {
		res.FailWithCode(res.RequestInfoError, c)
		return
	}
	global.Log.Infof("request=%#v", body)
	res.OKWithMessage("成功登录", c)
}
