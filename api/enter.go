package api

import (
	"guanxingtuan_bck/api/Society_api/SignUp_api"
	"guanxingtuan_bck/api/User_api/acount_api/loginJson_api"
)

type ApiGROUP struct {
	LoginJsonApi loginJson_api.LoginJsonApi
	SignUpApi    SignUp_api.SignUpApi
}

var ApiGroupApp = new(ApiGROUP)
