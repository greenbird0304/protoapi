// Code generated by protoapi:go; DO NOT EDIT.

package nested

// BindError
type BindError struct {
	Message string `json:"message"`
}

func (r *BindError) GetMessage() string {
	if r == nil {
		var zeroVal string
		return zeroVal
	}
	return r.Message
}
