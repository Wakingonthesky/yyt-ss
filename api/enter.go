package api

import (
	"guanxingtuan_bck/api/Society_api/SignUp_api"
	"guanxingtuan_bck/api/Society_api/inquiry"
)

type ApiGROUP struct {
	SignUpApi  SignUp_api.SignUpApi
	InquiryApi inquiry.InquiryApi
}

var ApiGroupApp = new(ApiGROUP)
