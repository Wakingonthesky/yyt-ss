package res

type ErrorCode int

const (
	SettingsError    ErrorCode = 1001 //系统错误
	RequestInfoError ErrorCode = 2001 //请求信息错误
	RequestTypeError ErrorCode = 2002 //请求类型错误
	InfoNotFound     ErrorCode = 3001 //目标信息未找到
	ExistError       ErrorCode = 3002 //目标信息已存在

)

var (
	ErrorMap = map[ErrorCode]string{
		SettingsError:    "系统错误",
		RequestInfoError: "请求信息错误",
		RequestTypeError: "请求类型错误",
		InfoNotFound:     "目标信息未找到",
		ExistError:       "目标信息已存在",
	}
)
