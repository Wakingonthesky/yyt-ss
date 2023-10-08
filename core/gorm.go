package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"guanxingtuan_bck/global"
	"log"
	"time"
)

func InitGorm() *gorm.DB {
	//如果此时的yaml文件没有相关配置，则返回nil
	if global.Config.Mysql.Host == "" {
		log.Println("未配置mysql，请先配置相应参数")
		return nil
	}
	//赋值给dsn
	dsn := global.Config.Mysql.Dsn()

	//创建接口（打印相关日志用）
	var mysqlLogger logger.Interface

	if global.Config.System.Env == "dev" {
		//开发环境显示所有的sql
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Info) //只打印错误的sql
	}
	//连接数据库，第一个参数为Dialector
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	//打印连接结果
	if err != nil {
		global.Log.Fatalf(fmt.Sprintf("[%s] mysql连接失败", dsn))
	}
	global.Log.Infof("已连接到数据库：%s", dsn)
	//其他配置
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)  //最大空闲连接数
	sqlDB.SetMaxOpenConns(100) //最多可容纳
	sqlDB.SetConnMaxLifetime(time.Hour * 4)
	return db
}
