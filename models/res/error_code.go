package res

type Code int

const (
	SwitchingProtocols          Code = 101 //系统错误
	OKCode                      Code = 200 //请求信息错误
	Created                     Code = 201 //请求类型错误
	Accepted                    Code = 202
	NonAuthoritativeInformation Code = 203
	NoContent                   Code = 204
	ResetContent                Code = 205
	PartialContent              Code = 206

	MultipleChoices   Code = 300
	MovedPermanently  Code = 301
	Found             Code = 302
	SeeOther          Code = 303
	NotModified       Code = 304
	UseProxy          Code = 305
	Unused            Code = 306
	TemporaryRedirect Code = 307

	BadRequest                  Code = 400
	Unauthorized                Code = 401
	PaymentRequired             Code = 402
	Forbidden                   Code = 403
	NotFound                    Code = 404
	MethodNotAllowed            Code = 405
	NotAcceptable               Code = 406
	ProxyAuthenticationRequired Code = 407
	RequestTimeout              Code = 408
	Conflict                    Code = 409
	Gone                        Code = 410
	LengthRequired              Code = 411
	PreconditionFailed          Code = 412
	RequestEntityTooLarge       Code = 413
	RequestURITooLarge          Code = 414
	UnsupportedMediaType        Code = 415
	Requestedrangenotsatisfi    Code = 416
	ExpectationFailed           Code = 417

	InternalServerError     Code = 500
	NotImplemented          Code = 501
	BadGateway              Code = 502
	ServiceUnavailable      Code = 503
	GatewayTimeout          Code = 504
	HTTPVersionnotsupported Code = 505
)

var (
	ErrorMap = map[Code]string{
		SwitchingProtocols:          "切换协议。服务器根据客户端的请求切换协议。只能切换到更高级的协议,例如,切换到HTTP的新版本协议",
		OKCode:                      "请求成功。一般用于GET与POST请求",
		Created:                     "已创建。成功请求并创建了新的资源",
		Accepted:                    "已接受。已经接受请求,但未处理完成",
		NonAuthoritativeInformation: "非授权信息。请求成功。但返回的meta信息不在原始的服务器,而是一个副本",
		NoContent:                   "无内容。服务器成功处理,但未返回内容。在未更新网页的情况下,可确保浏览器继续显示当前文档",
		ResetContent:                "重置内容。服务器处理成功,用户终端（例如：浏览器）应重置文档视图。可通过此返回码清除浏览器的表单域",
		PartialContent:              "部分内容。服务器成功处理了部分GET请求",

		MultipleChoices:   "多种选择。请求的资源可包括多个位置,相应可返回一个资源特征与地址的列表用于用户终端（例如：浏览器）选择",
		MovedPermanently:  "永久移动。请求的资源已被永久的移动到新URI,返回信息会包括新的URI,浏览器会自动定向到新URI。今后任何新的请求都应使用新的URI代替",
		Found:             "临时移动。与301类似。但资源只是临时被移动。客户端应继续使用原有URI",
		SeeOther:          "查看其它地址。与301类似。使用GET和POST请求查看",
		NotModified:       "未修改。所请求的资源未修改,服务器返回此状态码时,不会返回任何资源。客户端通常会缓存访问过的资源,通过提供一个头信息指出客户端希望只返回在指定日期之后修改的资源",
		UseProxy:          "使用代理。所请求的资源必须通过代理访问",
		Unused:            "已经被废弃的HTTP状态码",
		TemporaryRedirect: "临时重定向。与302类似。使用GET请求重定向",

		BadRequest:                  "客户端请求的语法错误,服务器无法理解",
		Unauthorized:                "请求要求用户的身份认证",
		PaymentRequired:             "保留,将来使用",
		Forbidden:                   "服务器理解请求客户端的请求,但是拒绝执行此请求",
		NotFound:                    "服务器无法根据客户端的请求找到资源（网页）。通过此代码,网站设计人员可设置'您所请求的资源无法找到'的个性页面",
		MethodNotAllowed:            "客户端请求中的方法被禁止",
		NotAcceptable:               "服务器无法根据客户端请求的内容特性完成请求",
		ProxyAuthenticationRequired: "请求要求代理的身份认证,与401类似,但请求者应当使用代理进行授权",
		RequestTimeout:              "服务器等待客户端发送的请求时间过长,超时",
		Conflict:                    "服务器完成客户端的 PUT 请求时可能返回此代码,服务器处理请求时发生了冲突",
		Gone:                        "客户端请求的资源已经不存在。410不同于404,如果资源以前有现在被永久删除了可使用410代码,网站设计人员可通过301代码指定资源的新位置",
		LengthRequired:              "服务器无法处理客户端发送的不带Content-Length的请求信息",
		PreconditionFailed:          "客户端请求信息的先决条件错误",
		RequestEntityTooLarge:       "由于请求的实体过大,服务器无法处理,因此拒绝请求。为防止客户端的连续请求,服务器可能会关闭连接。如果只是服务器暂时无法处理,则会包含一个Retry-After的响应信息",
		RequestURITooLarge:          "请求的URI过长（URI通常为网址）,服务器无法处理",
		UnsupportedMediaType:        "服务器无法处理请求附带的媒体格式",
		Requestedrangenotsatisfi:    "客户端请求的范围无效",
		ExpectationFailed:           "服务器无法满足Expect的请求头信息",

		InternalServerError:     "服务器内部错误,无法完成请求",
		NotImplemented:          "服务器不支持请求的功能,无法完成请求",
		BadGateway:              "作为网关或者代理工作的服务器尝试执行请求时,从远程服务器接收到了一个无效的响应",
		ServiceUnavailable:      "由于超载或系统维护,服务器暂时的无法处理客户端的请求。延时的长度可包含在服务器的Retry-After头信息中",
		GatewayTimeout:          "充当网关或代理的服务器,未及时从远端服务器获取请求",
		HTTPVersionnotsupported: "服务器不支持请求的HTTP协议的版本,无法完成处理",
	}
)
