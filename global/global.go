package global

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"guanxingtuan_bck/config"
)

var (
	//全局变量，用于保存配置文件
	Config *config.Config
	DB     *gorm.DB
	Log    *logrus.Logger
)
