package api

import (
	"guanxingtuan_bck/api/Society_api/hassign"
	"guanxingtuan_bck/api/Society_api/inquiry"
	"guanxingtuan_bck/api/Society_api/sighup_api"
)

type ApiGROUP struct {
	SignUpApi  sighup_api.SignUpApi
	InquiryApi inquiry.InquiryApi
	HasSignApi hassign.HasSignApi
}

var ApiGroupApp = new(ApiGROUP)
