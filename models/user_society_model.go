package models

import "time"

type UserSociety struct {
	TENANT_ID       string    `gorm:"size:32"` //租户号
	REVISION        string    `gorm:"size:32"` //乐观锁
	CREATED_BY      string    `gorm:"size:32"` //创建人
	CREATED_TIME    time.Time //创建时间
	UPDATED_BY      string    `gorm:"size:32"` //更新人
	UPDATED_TIME    time.Time //更新时间
	USER_SOCIETY_ID string    `gorm:"size:32"`  //报名ID
	USER_UID        string    `gorm:"size:255"` //用户ID
	USER_REALNAME   string    `gorm:"size:255"` //用户姓名
	USER_STUDENT_ID string    `gorm:"size:25"`  //用户学号
	USER_PHONE      string    `gorm:"size:"255` //用户手机号
	SOCIETY_ID      string    `gorm:"size:32"`  //社团编号
	SOCIETY_NAME    string    `gorm:"size:255"` //社团名称
}

func (UserSociety) TableName() string {
	return "USER_SOCIETY"
}
