package models

import "time"

type SocietyInfo struct {
	TENANT_ID     string    `gorm:"size:32"` //租户号
	REVISION      string    `gorm:"size:32"` //乐观锁
	CREATED_BY    string    `gorm:"size:32"` //创建人
	CREATED_TIME  time.Time //创建时间
	UPDATED_BY    string    `gorm:"size:32"` //更新人
	UPDATED_TIME  time.Time //更新时间
	SOCIETY_ID    int       `gorm:"size:4"`   //社团编号
	SOCIETY_NAME  string    `gorm:"size:255"` //社团名称
	SOCIETY_TYPE  string    `gorm:"size:255"` //社团类型
	SOCIETY_INDEX string    `gorm:"size:900"` //社团简介
}

func (SocietyInfo) TableName() string {
	return "SOCIETY_INFO"
}
