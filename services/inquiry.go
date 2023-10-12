package services

import (
	"ytt-societyservice/dao"
)

func Inquiry(st string) ([]string, error) {
	if ls, err := dao.Inquiry(st); err != nil {
		return nil, err
	} else {
		return ls, nil
	}

}
