package routers

import "guanxingtuan_bck/api"

func (router RouterGroup) InquiryRouter() {
	Inquiryapi := api.ApiGroupApp.InquiryApi
	router.GET("inquiry", Inquiryapi.InquiryInfoView)

}
