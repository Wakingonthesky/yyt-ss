package dao

import (
	"fmt"
	"ytt-societyservice/config"
	"ytt-societyservice/models"
)

func Signup(uid string, unm string, up string, ustd string, uqq string, sid string) error {

	m := &models.UserSociety{
		USER_UID:        uid,
		USER_REALNAME:   unm,
		USER_PHONE:      up,
		USER_STUDENT_ID: ustd,
		USER_QQ:         uqq,
		SOCIETY_ID:      sid,
	}

	res := config.DB.Create(&m)

	fmt.Println(m.USER_REALNAME)

	if res.Error != nil {
		fmt.Println(res.Error)
		return res.Error
	} else {
		return nil
	}
}
