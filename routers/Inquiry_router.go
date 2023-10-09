package routers

import (
	"guanxingtuan_bck/api"
	"guanxingtuan_bck/middleware"
)

func (router RouterGroup) InquiryRouter() {
	Inquiryapi := api.ApiGroupApp.InquiryApi
	router.GET("inquiry", middleware.Jwt(), Inquiryapi.InquiryInfoView)

}
