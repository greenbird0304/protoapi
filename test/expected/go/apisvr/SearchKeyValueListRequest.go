// Code generated by protoapi:go; DO NOT EDIT.

package apisvr

// SearchKeyValueListRequest
type SearchKeyValueListRequest struct {
	Key        string `json:"key"`
	Service_id int    `json:"service_id"`
	Env_id     int    `json:"env_id"`
}

func (r *SearchKeyValueListRequest) GetKey() string {
	if r == nil {
		var zeroVal string
		return zeroVal
	}
	return r.Key
}

func (r *SearchKeyValueListRequest) GetService_id() int {
	if r == nil {
		var zeroVal int
		return zeroVal
	}
	return r.Service_id
}

func (r *SearchKeyValueListRequest) GetEnv_id() int {
	if r == nil {
		var zeroVal int
		return zeroVal
	}
	return r.Env_id
}
