package routers

import (
	"guanxingtuan_bck/api"
	"guanxingtuan_bck/middleware"
)

func (router RouterGroup) SignUpRouter() {
	Signapi := api.ApiGroupApp.SignUpApi
	router.POST("signup", middleware.Jwt(), Signapi.SignUpInfoView)
}
