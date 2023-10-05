package flag

import (
	"guanxingtuan_bck/global"
	"guanxingtuan_bck/models"
)

func Makemigrations() {
	var err error
	//生成表
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.UserInfo{},
		&models.UserSociety{},
		&models.SocietyInfo{},
	)

	if err != nil {
		global.Log.Error("[error]生成数据库失败")
	}
	global.Log.Info("[success]生成数据库成功")
}
