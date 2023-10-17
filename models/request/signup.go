package request

type user_info struct {
	User_UID       string `json:"user_UID"`
	User_QQ        string `json:"user_QQ"`
	User_Phone     string `json:"user_Phone"`
	User_StudentID string `json:"user_StudentID"`
	User_Name      string `json:"user_Name"`
	SocietyID      string `json:"societyID"`
}

type Signup struct {
	Access_token string `json:"access_token"`
	User_info    user_info
}
