package config

import (
	"fmt"
	"log"
	"ytt-societyservice/models"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	C  *viper.Viper
	DB *gorm.DB
)

func initMysql() {
	userName := C.GetString("mysql.user")
	Password := C.GetString("mysql.password")
	ip := C.GetString("mysql.host")
	port := C.GetString("mysql.port")
	dbName := C.GetString("mysql.db")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", userName, Password, ip, port, dbName)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	if DB == nil {
		// MySQL未初始化
		fmt.Println("MySQL未初始化")
	} else {
		// MySQL已初始化
		fmt.Println("MySQL已初始化")
	}

	sqlDB, err := DB.DB()

	sqlDB.SetMaxOpenConns(10) // 设置连接池中的最大连接数
	sqlDB.SetMaxIdleConns(5)  // 设置连接池中的最大空闲连接数

	DB.AutoMigrate(&models.UserSociety{})
}

func initViper() {
	C = viper.New()
	C.SetConfigName("config")
	C.SetConfigType("yaml")
	C.AddConfigPath("./config")
	err := C.ReadInConfig() // 读取配置文件
	if err != nil {
		log.Fatal("无法读取配置文件:", err)
	}

}

func Init() {
	initViper()
	initMysql()
}
