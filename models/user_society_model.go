package models

import "gorm.io/gorm"

type UserSociety struct {
	gorm.Model
	USER_UID        string `gorm:"size:255"` //用户ID
	USER_REALNAME   string `gorm:"size:255"` //用户姓名
	USER_STUDENT_ID string `gorm:"size:25"`  //用户学号
	USER_PHONE      string `gorm:"size:"255` //用户手机号
	USER_QQ         string `gorm:"size:"255` //用户QQ号
	SOCIETY_ID      string `gorm:"size:32"`  //社团编号
}

func (UserSociety) TableName() string {
	return "USER_SOCIETY"
}
