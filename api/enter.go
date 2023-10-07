package api

import (
	"guanxingtuan_bck/api/Society_api/sighup_api"
)

type ApiGROUP struct {
	SignUpApi sighup_api.SignUpApi
}

var ApiGroupApp = new(ApiGROUP)
