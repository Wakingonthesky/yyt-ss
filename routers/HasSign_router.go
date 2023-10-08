package routers

import "guanxingtuan_bck/api"

func (router RouterGroup) HasSignRouter() {
	HasSignapi := api.ApiGroupApp.HasSignApi
	router.POST("hassign", HasSignapi.HasSignInfoView)
}
