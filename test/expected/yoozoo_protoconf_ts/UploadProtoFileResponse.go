// Code generated by protoapi; DO NOT EDIT.

package yoozoo_protoconf_ts

// UploadProtoFileResponse
type UploadProtoFileResponse struct {
	Service_id int `json:"service_id"`
	Env_id int `json:"env_id"`
	Key_count int `json:"key_count"`
}

func (r UploadProtoFileResponse) Validate() *ValidateError {
	errs := []*FieldError{}
	if len(errs) > 0 {
		return &ValidateError{Errors: errs}
	}
	return nil
}
