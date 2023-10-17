package dao

import (
	"ytt-societyservice/config"
	"ytt-societyservice/models"
)

func IsSign(uid string, so string) (bool, error) {
	var count int64
	DB := config.DB
	res := DB.Where("user_uid = ? AND society_ID = ? ", uid, so).Find(&models.UserSociety{}).Count(&count)

	if res.Error != nil {
		return false, res.Error
	}

	if count >= 1 {
		return false, nil
	} else {
		return true, nil
	}

}
