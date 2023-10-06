package routers

import (
	"github.com/gin-gonic/gin"
	"guanxingtuan_bck/global"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	apiRouterGroup := router.Group("/api/society")
	RouterGroupApp := RouterGroup{apiRouterGroup}
	//系统操作api
	RouterGroupApp.LoginJsonRouter()
	RouterGroupApp.SignUpRouter()
	return router
}
