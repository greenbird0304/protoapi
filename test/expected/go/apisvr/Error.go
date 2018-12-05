// Code generated by protoapi:go; DO NOT EDIT.

package apisvr

// Error
type Error struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}

func (r *Error) GetCode() ErrorCode {
	if r == nil {
		var zeroVal ErrorCode
		return zeroVal
	}
	return r.Code
}

func (r *Error) GetMessage() string {
	if r == nil {
		var zeroVal string
		return zeroVal
	}
	return r.Message
}
