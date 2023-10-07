package routers

import (
	"guanxingtuan_bck/api"
)

func (router RouterGroup) SignUpRouter() {
	Signapi := api.ApiGroupApp.SignUpApi
	router.POST("hassign", Signapi.InquirySigned)
	router.POST("signup", Signapi.SignUpInfoView)
}
