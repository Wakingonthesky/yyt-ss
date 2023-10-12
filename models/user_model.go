package models

type UserInfo struct {
	USER_UID       string `gorm:"size:32"`  //用户ID
	USER_REALNAME  string `gorm:"size:32"`  //用户真实姓名
	USER_STUDENTID string `gorm:"size:10"`  //用户学号
	USER_PHONE     string `gorm:"size:32"`  //用户手机号
	USER_QQ        string `gorm:"size:255"` //用户QQ号
	USER_HASSIGN   int    `gorm:"size:"4"`  //用户是否已经报名

}

func (UserInfo) TableName() string {
	return "USER_INFO"
}
