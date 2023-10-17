package services

import (
	"log"
	"ytt-societyservice/dao"
	"ytt-societyservice/models/request"
)

func Signup(f *request.Signup) error {

	if err := dao.Signup(f.User_info.User_UID, f.User_info.User_Name, f.User_info.User_Phone, f.User_info.User_StudentID, f.User_info.User_QQ, f.User_info.SocietyID); err != nil {
		log.Panic(err)
		return err
	} else {
		Code(f.User_info.User_Phone, f.User_info.User_Name, f.User_info.SocietyID)
		return nil
	}

}
