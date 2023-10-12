package models

type SocietyInfo struct {
	SOCIETY_ID    int    `gorm:"size:4"`   //社团编号
	SOCIETY_NAME  string `gorm:"size:255"` //社团名称
	SOCIETY_TYPE  string `gorm:"size:255"` //社团类型
	SOCIETY_INDEX string `gorm:"size:900"` //社团简介
}

func (SocietyInfo) TableName() string {
	return "SOCIETY_INFO"
}
