package ret

func NewError(code Code, msg string) *Error {
	return &Error{
		Code: code,
		Des:  msg,
	}
}
