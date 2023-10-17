package dao

import (
	"fmt"
	"ytt-societyservice/config"
	"ytt-societyservice/models"
)

func Inquiry(uid string) ([]string, error) {
	DB := config.DB

	var ls []string

	if result := DB.Model(&models.UserSociety{}).Where("user_UID = ?", uid).Select("society_id").Find(&ls); result.Error != nil {
		return nil, result.Error
	} else {
		fmt.Printf(ls[0])
		return ls, nil
	}
}
