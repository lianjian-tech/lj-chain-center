package errno

var (
	OK          = &Errno{Code: 0, Msg: "ok"}
	HandleError = &Errno{Code: 1, Msg: "handle error"}
	ParamError  = &Errno{Code: 2, Msg: "request param error"}
	AuthError   = &Errno{Code: 3, Msg: "auth error"}
)
