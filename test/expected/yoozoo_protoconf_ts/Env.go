// Code generated by protoapi; DO NOT EDIT.

package yoozoo_protoconf_ts

// Env
type Env struct {
	Env_id int `json:"env_id"`
	Env_name string `json:"env_name"`
}

func (r Env) Validate() *ValidateError {
	errs := []*FieldError{}
	if len(errs) > 0 {
		return &ValidateError{Errors: errs}
	}
	return nil
}
