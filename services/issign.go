package services

import "ytt-societyservice/dao"

func IsSign(studentID string, societyID string) (bool, error) {
	if flag, err := dao.IsSign(studentID, societyID); err != nil {
		return false, nil
	} else {
		if !flag {
			return false, nil
		} else {
			return true, nil
		}
	}
}
