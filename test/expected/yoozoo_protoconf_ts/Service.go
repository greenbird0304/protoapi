// Code generated by protoapi; DO NOT EDIT.

package yoozoo_protoconf_ts

// Service
type Service struct {
	Service_id int `json:"service_id"`
	Service_name string `json:"service_name"`
	Product_id string `json:"product_id"`
	Product_name string `json:"product_name"`
	Desc string `json:"desc"`
	Tags []*Tag `json:"tags"`
}

func (r Service) Validate() *ValidateError {
	errs := []*FieldError{}
	if len(errs) > 0 {
		return &ValidateError{Errors: errs}
	}
	return nil
}
