// Code generated by protoapi; DO NOT EDIT.

package yoozoo_protoconf_ts

// EnvListRequest
type EnvListRequest struct {
}

func (r EnvListRequest) Validate() *ValidateError {
	errs := []*FieldError{}
	if len(errs) > 0 {
		return &ValidateError{Errors: errs}
	}
	return nil
}
