package atom

type Error interface {
	Code() int
	Error() string
}

type MyError struct {
	errMsg     string
	errCode    int         // 小于-10不打日志的错误，大于-10的需要打印日志
	httpStatus int         // http的状态码，为0或200表示输出均为200，其他值表示需要返回该http状态码
	data       interface{} // 用于需要传数据的情况
}

func NewMyErrorByCode(code int) Error {
	e := new(MyError)
	e.errMsg = GetMsgByCode(code)
	e.errCode = code
	return e
}

func NewMyError(code int, msg string) Error {
	e := new(MyError)
	e.errMsg = msg
	e.errCode = code
	return e
}

func (e MyError) Code() int {
	return e.errCode
}

func (e MyError) Error() string {
	return e.errMsg
}

func (e MyError) HttpStatus() int {
	return e.httpStatus
}
