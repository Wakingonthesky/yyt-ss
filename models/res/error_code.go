package res

type ErrorCode int

const (
	SettingsError    ErrorCode = 1001 //系统错误
	RequestInfoError ErrorCode = 2001 //请求信息错误
	RequestTypeError ErrorCode = 2002 //请求类型错误
	InfoNotFound     ErrorCode = 3001 //目标信息未找到
)

var (
	ErrorMap = map[ErrorCode]string{
		SettingsError:    "系统错误",
		RequestInfoError: "请求信息错误",
		RequestTypeError: "请求类型错误",
		InfoNotFound:     "目标信息未找到",
	}
)
