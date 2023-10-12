package request

type user_info struct {
	User_UID       string
	User_Name      string
	User_Phone     string
	User_StudentID string
	SocietyID      string
}

type Signup struct {
	Access_token string
	User_info    user_info
}
