// Code generated by protoapi; DO NOT EDIT.

package yoozoo_protoconf_ts

// SearchKeyValueListRequest
type SearchKeyValueListRequest struct {
	Key string `json:"key"`
	Service_id int `json:"service_id"`
	Env_id int `json:"env_id"`
}

func (r SearchKeyValueListRequest) Validate() *ValidateError {
	errs := []*FieldError{}
	if len(errs) > 0 {
		return &ValidateError{Errors: errs}
	}
	return nil
}
