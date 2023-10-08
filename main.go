package main

import (
	"guanxingtuan_bck/core"
	"guanxingtuan_bck/flag"
	"guanxingtuan_bck/global"
	"guanxingtuan_bck/routers"
)

func main() {
	//读取配置文件
	core.InitConf()
	//初始化日志
	global.Log = core.InitLogger()

	//连接数据库
	global.DB = core.InitGorm()

	//命令行参数配置
	option := flag.Parse()
	global.Log.Infof("option=%v", option)
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}
	//api设置
	router := routers.InitRouter()
	addr := global.Config.System.Addr()

	global.Log.Infof("guanxingtuan_bck运行在server:%s", addr)
	err := router.Run(addr)
	if err != nil {
		global.Log.Fatalf(err.Error())
	}
}
