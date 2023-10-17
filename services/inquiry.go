package services

import (
	"ytt-societyservice/dao"
)

func Inquiry(uid string) ([]string, error) {
	if ls, err := dao.Inquiry(uid); err != nil {
		return nil, err
	} else {
		return ls, nil
	}

}
