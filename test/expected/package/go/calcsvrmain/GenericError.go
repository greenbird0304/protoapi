// Code generated by protoapi:go; DO NOT EDIT.

package calcsvrmain

// GenericError
type GenericError struct {
	Message string `json:"message"`
}

func (r *GenericError) GetMessage() string {
	if r == nil {
		var zeroVal string
		return zeroVal
	}
	return r.Message
}