package routers

import (
	"guanxingtuan_bck/api"
)

func (router RouterGroup) LoginJsonRouter() {
	loginApi := api.ApiGroupApp.LoginJsonApi
	router.POST("Login_json.php", loginApi.LoginJsonInfoView)
}
