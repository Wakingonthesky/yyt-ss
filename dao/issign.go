package dao

import (
	"ytt-societyservice/config"
	"ytt-societyservice/models"
)

func IsSign(st string, so string) (bool, error) {
	DB := config.DB
	res := DB.Where("studentID = ? AND societyID = ? ", st, so).Find(&models.UserSociety{})

	if res.RowsAffected >= 1 {
		return false, res.Error
	} else {
		return true, nil
	}

}
