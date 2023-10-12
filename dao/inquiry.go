package dao

import (
	"fmt"
	"ytt-societyservice/config"
	"ytt-societyservice/models"
)

func Inquiry(st string) ([]string, error) {
	DB := config.DB

	var ls []string

	if result := DB.Model(&models.UserSociety{}).Where("studentID = ?", st).Find(ls); result.Error != nil {
		return nil, result.Error
	} else {
		fmt.Printf(ls[0])
		return ls, nil
	}
}
