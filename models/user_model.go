package models

import "time"

type UserInfo struct {
	MODEL
	TENANT_ID      string    `gorm:"size:32","primarykey"` //租户号
	REVISION       string    `gorm:"size:32"`              //乐观锁
	CREATED_BY     string    `gorm:"size:32"`              //创建人
	CREATED_TIME   time.Time //创建时间
	UPDATED_BY     string    `gorm:"size:32"` //更新人
	UPDATED_TIME   time.Time //更新时间
	USER_UID       string    `gorm:"size:32"`  //用户ID
	USER_REALNAME  string    `gorm:"size:32"`  //用户真实姓名
	USER_STUDNETID int       `gorm:"size:10"`  //用户学号
	USER_PHONE     int       `gorm:"size:11"`  //用户手机号
	USER_QQ        string    `gorm:"size:255"` //用户QQ号

}

func (UserInfo) TableName() string {
	return "USER_INFO"
}
